import os
import sys
from typing import List, Dict

def merge_txt_files() -> List[str]:
    """合并当前目录及其子目录下的所有txt文件内容"""
    output = []
    for root, dirs, files in os.walk("."):
        for file in files:
            if file.endswith(".txt"):
                file_path = os.path.join(root, file)
                try:
                    with open(file_path, 'r', encoding='utf-8') as f:
                        output.append(f.read())
                except UnicodeDecodeError:
                    print(f"跳过无法解码的文件: {file_path}")
    return output

def remove_empty_lines(lines: List[str]) -> List[str]:
    """从列表中删除所有空行"""
    return [line for line in lines if line.strip() != '']

def remove_duplicate_lines(lines: List[str]) -> List[str]:
    """从列表中删除所有重复行，同时保持顺序"""
    seen = {}
    unique_lines = []
    for line in lines:
        if line not in seen:
            seen[line] = True
            unique_lines.append(line)
    return unique_lines

def save_to_file(lines: List[str], output_path: str) -> None:
    """将行列表保存到文件"""
    with open(output_path, 'w', encoding='utf-8') as f:
        f.writelines(lines)
    print(f"结果已保存到 {output_path}")

def main():
    try:
        if len(sys.argv) != 2:
            print("用法：python script.py <输出文件>")
            sys.exit(1)
    
        output_file = sys.argv[1]

        # 第一步：合并所有txt文件
        merged_content = merge_txt_files()
        if not merged_content:
            print("未找到任何txt文件")
            return
        
        # 第二步：将合并内容按换行分割成列表
        all_lines = "\n".join(merged_content).splitlines(True)
        
        # 第三步：移除空行
        non_empty_lines = remove_empty_lines(all_lines)
        
        # 第四步：移除重复行
        unique_lines = remove_duplicate_lines(non_empty_lines)
        
        # 第五步：保存最终结果
        save_to_file(unique_lines, output_file)
        
    except Exception as e:
        print(f"处理过程中发生错误：{e}")

if __name__ == "__main__":
    main()
