# Changelog

所有项目的显著变更都将记录在此文件中。

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
并且本项目遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

## [未发布]

### 新增
- 完善了README和README-zh文档
- 添加了GitHub Actions自动化工作流
- 添加了代码覆盖率报告
- 添加了netutils网络工具包（IP验证、域名验证、端口检查、URL可达性检查等）
- 添加了cryptutils加密解密工具包（哈希计算、Base64编解码、AES加密解密、UUID生成等）
- 添加了validationutils验证工具包（邮箱、手机号、URL、IP地址验证、密码强度检查等）

### 修复
- 修复了测试文件中的格式问题
- 修复了cmd/apidocs/main.go中未检查的错误返回值
- 修复了cmd/apidocs/main.go中缺少log包导入的问题

### 变更
- 无

## [0.1.0] - 2025-09-08

### 新增
- 初始版本发布
- 添加了stringutils包的核心功能
- 实现了systemutils包的基础功能（CPU、内存、磁盘监控）
- 添加了跨平台支持（Linux、macOS、Windows）
- 创建了示例和文档