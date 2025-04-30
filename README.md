# 题目二:全栈开发挑战(计算器)
任务要求:
- 后端:使用 Go+ConnectRPC 实现一个服务，支持基础的加减乘除运算。
- 前端:使用 Next.js 构建U，包含输入框、操作符选择和结果展示。
- 交互:通过ConnectRPC调用后端，并显示运算结果。需要使用基于Connect协议的调用而不是HTTP POST和JSON的调用


附加题:
- 将代码提交到公开Git仓库(GitHub/GitLab等均可)
- 为前后端代码编写单元测试。


全栈计算器应用，基于以下技术栈：
- **前端**: Next.js + Connect-Web
  - nodejs v20.11.0
- **后端**: Go + ConnectRPC
  - go v1.19

## 功能特性
✅ 基本四则运算
✅ 实时错误处理
✅ 响应式设计
✅ 单元测试覆盖

## 项目部署
将项目拉取到GOROOT下的src文件夹下，如我的本地GOROOT是F:\golang\go1.19则拉取到F:\golang\go1.19\src下进入项目文件夹F:\golang\go1.19\src\calculator-backend

## 快速开始
```bash
# 后端准备工作（当前路径calculator-backend文件夹下）
cd backend
go clean -modcache
go mod tidy
rmdir /s /q gen
buf generate proto

# 启动后端
go run main.go

# 后端测试
go test


# 前端准备工作（当前路径calculator-backend文件夹下）
cd frontend
npm install

# 启动前端
npm run dev

# 前端测试（还存在bug，暂未完成）
npm run test:watch
npm run test:coverage
