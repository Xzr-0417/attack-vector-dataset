import sys

def unicode_escape(text):
    """将字符串转换为Unicode转义序列"""
    escaped = []
    for char in text:
        code_point = ord(char)
        if code_point <= 0xFFFF:
            escaped.append(f"\\u{code_point:04X}")
        else:
            escaped.append(f"\\U{code_point:08X}")
    return "".join(escaped)

def process_file(input_filename, output_filename):
    try:
        with open(input_filename, "rb") as src_file, \
                open(output_filename, "w", encoding="utf-8") as dest_file:

            for line_num, byte_line in enumerate(src_file, 1):
                try:
                    # 解码并清理行内容
                    decoded_line = byte_line.decode('utf-8').rstrip('\n')
                    # 生成Unicode转义序列
                    encoded_line = unicode_escape(decoded_line)
                    dest_file.write(f"{encoded_line}\n")
                except UnicodeDecodeError:
                    print(f"跳过第 {line_num} 行：无效的UTF-8编码字节")
                except Exception as e:
                    print(f"跳过第 {line_num} 行：处理异常 - {str(e)}")

        print(f"处理完成！有效内容已保存到 {output_filename}")

    except FileNotFoundError:
        print(f"错误：找不到 {input_filename} 文件")
    except PermissionError:
        print("错误：没有文件写入权限")
    except Exception as e:
        print(f"程序异常终止：{str(e)}")

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("用法: python3 script.py <输入文件> <输出文件>")
        sys.exit(1)
    
    input_file = sys.argv[1]
    output_file = sys.argv[2]
    process_file(input_file, output_file)
