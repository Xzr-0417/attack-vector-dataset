import html
import sys

def html_encode_file(input_filename, output_filename):
    try:
        with open(input_filename, "rb") as src_file, \
                open(output_filename, "w", encoding="utf-8") as dest_file:

            for line_num, byte_line in enumerate(src_file, 1):
                # 尝试解码和编码处理
                try:
                    # 解码并去除行尾换行符
                    decoded_line = byte_line.decode('utf-8').rstrip('\n')
                    # HTML实体编码（同时转义引号）
                    encoded_line = html.escape(decoded_line, quote=True)
                    dest_file.write(f"{encoded_line}\n")
                except UnicodeDecodeError:
                    print(f"跳过第 {line_num} 行：包含无效UTF-8编码内容")
                except Exception as e:
                    print(f"跳过第 {line_num} 行：处理失败 - {str(e)}")

        print(f"处理完成！有效内容已保存到 {output_filename}")

    except FileNotFoundError:
        print(f"错误：{input_filename} 文件不存在")
    except PermissionError:
        print(f"错误：没有 {output_filename} 的写入权限")
    except Exception as e:
        print(f"程序异常终止：{str(e)}")

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("用法: python3 html_encoder.py <输入文件> <输出文件>")
        sys.exit(1)
    
    input_file = sys.argv[1]
    output_file = sys.argv[2]
    html_encode_file(input_file, output_file)
