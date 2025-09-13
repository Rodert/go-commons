# 贡献指南

感谢您考虑为go-commons项目做出贡献！以下是一些指导原则，帮助您参与项目开发。

## 开发流程

1. Fork本仓库
2. 创建您的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建一个Pull Request

## 代码规范

- 遵循Go的官方代码规范
- 使用`gofmt`格式化代码
- 添加适当的注释和文档
- 为新功能编写测试用例

## 提交PR前的检查清单

- [ ] 代码已经通过`make test`测试
- [ ] 代码已经通过`make lint`检查
- [ ] 新功能已添加相应的测试用例
- [ ] 文档已更新（如果需要）
- [ ] 更新了CHANGELOG.md（如果适用）

## 报告Bug

报告Bug时，请包含以下信息：

- 问题的简要描述
- 复现步骤
- 预期行为与实际行为
- 环境信息（Go版本、操作系统等）

## 功能请求

如果您有新功能的想法，请先创建一个Issue讨论该功能的必要性和实现方式，然后再开始编码。

## 许可证

通过提交代码，您同意您的贡献将在项目的许可证下发布。