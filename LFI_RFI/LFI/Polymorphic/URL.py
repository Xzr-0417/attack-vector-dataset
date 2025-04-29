import urllib.parse
import sys

def url_encode_file(input_filename, output_filename):
    try:
        with open(input_filename, "rb") as src_file, \
                open(output_filename, "w", encoding="utf-8") as dest_file:

            for line_num, byte_line in enumerate(src_file, 1):
                # 尝试解码行内容
                try:
                    str_line = byte_line.decode('utf-8').rstrip('\n')
                except UnicodeDecodeError:
                    print(f"跳过第 {line_num} 行：无法解码的字节内容")
                    continue

                # 尝试URL编码
                try:
                    encoded = urllib.parse.quote(str_line)
                    dest_file.write(f"{encoded}\n")
                except Exception as e:
                    print(f"跳过第 {line_num} 行：编码失败 - {str(e)}")

        print(f"处理完成！有效行已保存到 {output_filename}")

    except FileNotFoundError:
        print(f"错误：{input_filename} 文件不存在")
    except PermissionError:
        print(f"错误：没有 {output_filename} 的写入权限")
    except Exception as e:
        print(f"程序异常终止：{str(e)}")


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("用法: python3 url_encoder.py <输入文件> <输出文件>")
        print("示例: python3 url_encoder.py input.txt encoded.txt")
        sys.exit(1)
    
    input_file = sys.argv[1]
    output_file = sys.argv[2]
    url_encode_file(input_file, output_file)
