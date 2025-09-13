# Go Commons Makefile

# 变量定义
GO=go
GOFMT=gofmt
GOLINT=golint
GOVET=go vet
PACKAGES=$(shell $(GO) list ./...)
COVER_PROFILE=coverage.out

# 默认目标：运行所有测试
.PHONY: all
all: fmt test

# 格式化代码
.PHONY: fmt
fmt:
	@echo "格式化代码..."
	@$(GOFMT) -w -s .

# 运行所有测试
.PHONY: test
test:
	@echo "运行所有测试..."
	@$(GO) test -v ./...

# 运行指定包的测试
.PHONY: test-pkg
test-pkg:
	@if [ "$(PKG)" = "" ]; then \
		echo "请指定要测试的包，例如：make test-pkg PKG=./stringutils"; \
		exit 1; \
	fi
	@echo "运行 $(PKG) 的测试..."
	@$(GO) test -v $(PKG)

# 运行测试并生成覆盖率报告
.PHONY: cover
cover:
	@echo "运行测试并生成覆盖率报告..."
	@$(GO) test -coverprofile=$(COVER_PROFILE) ./...
	@$(GO) tool cover -html=$(COVER_PROFILE)
	@rm $(COVER_PROFILE)

# 运行基准测试
.PHONY: bench
bench:
	@echo "运行基准测试..."
	@$(GO) test -bench=. -benchmem ./...

# 代码检查
.PHONY: lint
lint:
	@echo "运行代码检查..."
	@$(GOVET) ./...
	@if command -v $(GOLINT) > /dev/null; then \
		$(GOLINT) ./...; \
	else \
		echo "golint 未安装，跳过 golint 检查"; \
	fi

# 清理生成的文件
.PHONY: clean
clean:
	@echo "清理生成的文件..."
	@$(GO) clean
	@rm -f $(COVER_PROFILE)

# 帮助信息
.PHONY: help
help:
	@echo "Go Commons Makefile 帮助"
	@echo ""
	@echo "可用命令："
	@echo "  make          - 格式化代码并运行所有测试"
	@echo "  make fmt      - 格式化代码"
	@echo "  make test     - 运行所有测试"
	@echo "  make test-pkg PKG=./path/to/package - 运行指定包的测试"
	@echo "  make cover    - 运行测试并生成覆盖率报告"
	@echo "  make bench    - 运行基准测试"
	@echo "  make lint     - 运行代码检查"
	@echo "  make clean    - 清理生成的文件"
	@echo "  make help     - 显示此帮助信息"