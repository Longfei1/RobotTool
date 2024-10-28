# -- coding: utf-8 -

import os
import re

def replace_json_tags(file_path, old_pattern, new_pattern):
    with open(file_path, 'r', encoding='utf-8') as file:
        contents = file.read()

    # 使用正则表达式替换 JSON tags
    contents, count = re.subn(old_pattern, new_pattern, contents)

    if count > 0:
        with open(file_path, 'w', encoding='utf-8') as file:
            file.write(contents)
        print(f"Updated {count} occurrences in {file_path}")
    else:
        print(f"No occurrences found in {file_path}")

if __name__ == '__main__':
    # 设置要搜索的目录
    directory = "./src/robot/dyj/pbgo"

    # 设置旧的 JSON tag 模式和新的模式
    old_pattern = r'json:\"([^\"]+?),omitempty\"'
    new_pattern = r'json:"\1"'  # 如果你想删除标签，使用空字符串

    # 递归地查找所有 .pb.go 文件
    for root, dirs, files in os.walk(directory):
        for file in files:
            if file.endswith(".pb.go"):
                file_path = os.path.join(root, file)
                replace_json_tags(file_path, old_pattern, new_pattern)