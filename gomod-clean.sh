#!/bin/bash
# 1. 检查参数
if [ -z "$1" ]; then
    echo "用法: $0 <module-name>"
    echo "示例: $0 golang.org/x/crypto"
    exit 1
fi

MODULE="$1"
# 将模块名中的 '.' 转义为 '\.'，防止正则表达式误匹配
REGEX_MODULE="${MODULE//./\\.}"

# 2. 检查 gsed 是否可用
if ! command -v gsed &> /dev/null; then
    echo "❌ 错误: 未找到 gsed。请先安装 GNU sed (macOS: brew install gnu-sed)"
    exit 1
fi

echo "🚀 正在从所有 go.mod 中移除: $MODULE"
echo "--------------------------------------------------"

# 3. 查找并处理所有 go.mod 文件
# 使用 -print0 和 read -d '' 确保能正确处理路径中包含空格的情况
find . -name "go.mod" -type f -print0 | while IFS= read -r -d '' file; do
    echo "✂️  清理: $file"
    
    # 4. 核心操作：使用 gsed 直接删除匹配行
    # -i: 原地修改文件
    # -E: 启用扩展正则表达式
    # \#...\#: 使用 # 作为定界符，避免与模块名中的 / 冲突
    gsed -i -E "\#^[[:space:]]*(require[[:space:]]+)?${REGEX_MODULE}([[:space:]]|$)#d" "$file"
done
