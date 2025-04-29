import base64
import sys

def process_lines(input_filename, output_filename):
    try:
        with open(input_filename, "r", encoding="utf-8") as input_file, \
                open(output_filename, "w", encoding="utf-8") as output_file:

            for line_number, line in enumerate(input_file, 1):
                original_line = line.rstrip('\n')

                try:
                    # 使用正确的base64编码方式
                    encoded_bytes = base64.b64encode(original_line.encode("utf-8"))
                    encoded_str = encoded_bytes.decode("utf-8")
                    output_file.write(f"{encoded_str}\n")
                except UnicodeEncodeError as e:
                    print(f"跳过第 {line_number} 行：包含非法字符 - {str(e)}")
                except Exception as e:
                    print(f"跳过第 {line_number} 行：处理失败 - {str(e)}")

        print(f"处理完成！结果已保存到 {output_filename}")

    except FileNotFoundError:
        print(f"错误：{input_filename} 文件不存在")
    except PermissionError:
        print(f"错误：没有 {output_filename} 的写入权限")
    except Exception as e:
        print(f"程序异常终止：{str(e)}")

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("用法: python3 base64_encoder.py <输入文件> <输出文件>")
        print("示例: python3 base64_encoder.py input.txt encoded.txt")
        sys.exit(1)
    
    process_lines(sys.argv[1], sys.argv[2])
