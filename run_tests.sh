#!/bin/bash

# 运行所有测试用例的脚本

echo "===== 运行 go-commons 所有测试用例 ====="
echo ""

# 设置颜色
GREEN="\033[0;32m"
RED="\033[0;31m"
YELLOW="\033[0;33m"
NC="\033[0m" # No Color

# 检查是否安装了Go
if ! command -v go &> /dev/null; then
    echo -e "${RED}错误: Go未安装，请先安装Go。${NC}"
    exit 1
fi

# 运行所有测试
echo -e "${YELLOW}运行所有测试...${NC}"
go test -v ./...
TEST_RESULT=$?

echo ""

# 检查测试结果
if [ $TEST_RESULT -eq 0 ]; then
    echo -e "${GREEN}所有测试通过！${NC}"
else
    echo -e "${RED}测试失败，请检查上面的错误信息。${NC}"
fi

echo ""
echo "===== 测试完成 ====="