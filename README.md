# RobotTool
用于测试服务器协议的机器人工具，使用golang、vue3开发。golang用于实现网络通信、行为树等逻辑，vue3用于编写网页界面

## 帮助
运行本项目，需要提前安装以下 [依赖](#依赖)，初次使用以及协议变更时，需要通过构建流程生成可执行文件。   
构建相关脚本：
- syncproto.bat: 同步项目的proto文件，需修改脚本中的项目路径
- makeproto.bat: 处理.proto，生成.pb.go以及对应的json schema文件
- build.bat: 编译项目，生成可执行文件（包含前两个过程）

## 开发环境
### 1.golang
- go 1.20

### 2.vue3
- node.js

- vue.js 3.4.36

- vite
```
npm create vue@latest
```

## 依赖
### 1.protoc

### 2.protoc-gen-jsonschema
go install github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema@latest