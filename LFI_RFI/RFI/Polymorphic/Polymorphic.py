import urllib.parse
import sys
import base64
import html

#======================= 核心编码函数 ========================
def url_encode(line):
    return urllib.parse.quote(line)

def case_convert(line):
    converted = []
    letter_count = 0
    for char in line:
        if char.isalpha():
            letter_count += 1
            converted.append(char.swapcase() if letter_count % 2 == 1 else char)
        else:
            converted.append(char)
    return ''.join(converted)

def base64_encode(line):
    return base64.b64encode(line.encode()).decode()

def unicode_escape(line):
    escaped = []
    for char in line:
        cp = ord(char)
        if cp <= 0xFFFF:
            escaped.append(f"\\u{cp:04X}")
        else:
            escaped.append(f"\\U{cp:08X}")
    return ''.join(escaped)

def html_encode(line):
    return html.escape(line, quote=True)

#======================= 编码器列表 ========================
ENCODERS = [
    ("url", url_encode),
    ("case", case_convert),
    ("base64", base64_encode),
    ("unicode", unicode_escape),
    ("html", html_encode)
]

#======================= 主处理逻辑 ========================
def process_file(input_file, output_file, encode_mode):
    try:
        with open(input_file, "rb") as src, \
             open(output_file, "w", encoding="utf-8") as dest:
            
            for line_num, byte_line in enumerate(src, 1):
                try:
                    raw_line = byte_line.decode('utf-8').rstrip('\n')
                except UnicodeDecodeError:
                    print(f"跳过第 {line_num} 行：无效的UTF-8编码")
                    continue
                
                # 默认模式：全编码处理
                if encode_mode == "all":
                    for enc_name, encoder in ENCODERS:
                        try:
                            encoded = encoder(raw_line)
                            dest.write(f"{encoded}\n")
                        except Exception as e:
                            print(f"第 {line_num} 行 [{enc_name}] 编码失败 - {str(e)}")
                # 单编码模式
                else:
                    encoder = dict(ENCODERS).get(encode_mode)
                    if not encoder:
                        raise ValueError(f"无效编码类型：{encode_mode}")
                    encoded = encoder(raw_line)
                    dest.write(f"{encoded}\n")

        print(f"处理完成！结果已保存到 {output_file}")

    except FileNotFoundError:
        print(f"错误：输入文件 {input_file} 不存在")
    except PermissionError:
        print(f"错误：没有 {output_file} 的写入权限")
    except Exception as e:
        print(f"程序异常终止：{str(e)}")

#======================= 命令行参数处理 ========================
if __name__ == "__main__":
    if len(sys.argv) not in [3, 4]:
        print("用法: python3 unified_encoder.py <输入文件> <输出文件> [编码类型]")
        print("编码类型（可选）: url, case, base64, unicode, html")
        print("示例1（全编码）: python3 unified_encoder.py input.txt output.txt")
        print("示例2（单编码）: python3 unified_encoder.py input.txt output.txt url")
        sys.exit(1)

    # 参数解析
    if len(sys.argv) == 3:
        _, input_file, output_file = sys.argv
        encode_mode = "all"
    else:
        _, input_file, output_file, encode_type = sys.argv
        encode_mode = encode_type.lower()

    try:
        process_file(input_file, output_file, encode_mode)
    except ValueError as ve:
        print(ve)
        sys.exit(1)
