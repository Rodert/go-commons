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

# 生成并运行API文档
.PHONY: apidocs
apidocs:
	@echo "生成并运行API文档..."
	@if command -v $(shell go env GOPATH)/bin/swag > /dev/null; then \
		echo "生成Swagger文档..."; \
		$(shell go env GOPATH)/bin/swag init -g cmd/apidocs/main.go -o docs; \
		echo "启动API文档服务器在 http://localhost:8080"; \
		go run cmd/apidocs/main.go; \
	else \
		echo "swag 未安装，请先安装: go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

# 格式化代码
.PHONY: fmt
fmt:
	@echo "格式化代码..."
	@$(GOFMT) -w -s .

# 安装git钩子
.PHONY: hooks
hooks:
	@echo "安装git钩子..."
	@cp -f .git/hooks/pre-commit.sample .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo '#!/bin/bash\n\n# 自动格式化Go代码的pre-commit钩子\necho "Running pre-commit hook: Auto-formatting Go code..."\n\n# 获取所有暂存的Go文件\nSTAGED_GO_FILES=$$(git diff --cached --name-only --diff-filter=ACM | grep "\\.go$$")\n\n# 如果没有Go文件被修改，则退出\nif [[ "$$STAGED_GO_FILES" = "" ]]; then\n  echo "No Go files staged for commit. Skipping formatting."\n  exit 0\nfi\n\n# 格式化所有暂存的Go文件\necho "$$STAGED_GO_FILES" | xargs gofmt -l -w\n\n# 重新添加格式化后的文件到暂存区\necho "$$STAGED_GO_FILES" | xargs git add\n\necho "Go code formatting completed successfully."\n\nexit 0' > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "Git钩子安装完成，现在每次提交前会自动格式化Go代码。"

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