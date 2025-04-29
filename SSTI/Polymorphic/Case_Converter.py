import sys

def convert_line(line):
    converted = []
    letter_count = 0  # 字母计数器（只统计字母）
    for char in line:
        if char.isalpha():
            letter_count += 1
            # 奇数位字母转换大小写，偶数位保持原样
            converted.append(char.swapcase() if letter_count % 2 == 1 else char)
        else:
            converted.append(char)
    return ''.join(converted)


def process_file(input_filename, output_filename):
    try:
        with open(input_filename, "rb") as src_file, \
                open(output_filename, "w", encoding="utf-8") as dest_file:

            for line_num, byte_line in enumerate(src_file, 1):
                try:
                    # 解码并保留原始换行符
                    raw_line = byte_line.decode("utf-8")
                    # 执行转换
                    converted_line = convert_line(raw_line)
                    dest_file.write(converted_line)
                except UnicodeDecodeError:
                    print(f"跳过第 {line_num} 行：包含非法编码内容")
                except Exception as e:
                    print(f"跳过第 {line_num} 行：处理失败 - {str(e)}")

        print(f"转换完成！结果已保存到 {output_filename}")

    except FileNotFoundError:
        print(f"错误：{input_filename} 文件不存在")
    except PermissionError:
        print("错误：没有文件写入权限")
    except Exception as e:
        print(f"程序异常终止：{str(e)}")


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("用法: python3 Case_Converter.py <输入文件名> <输出文件名>")
        sys.exit(1)
    input_filename = sys.argv[1]
    output_filename = sys.argv[2]
    process_file(input_filename, output_filename)
