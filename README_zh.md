# webdog - 轻量自定义 web 服务

## 什么是 webdog？
webdog 是一个自定义的轻量 web 服务程序，他允许我们通过 yaml 配置路由、外部实现控制器，你可以用 bash, lua, python 等任何语言去实现外部控制器。当然，他也可以是一个简单的静态资源 web 服务。

## 安装
### 执行文件下载
去 [releases](https://github.com/edboffical/webdog/releases/) 下载对应操作系统的可执行文件
### 源码编译
```bash
git clone https://github.com/edboffical/webdog
cd webdog
# for mac
make build-mac
# for linux
make build-linux
# for windows
make build-win
```

## 文档
[使用文档](./DOCUMENT_zh.md)

## 依赖项目
- [httprouter](https://github.com/julienschmidt/httprouter)
- [fsnotify](https://github.com/fsnotify/fsnotify)