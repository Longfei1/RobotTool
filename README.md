# RobotTool
用于测试服务器协议的机器人工具，使用golang、vue3开发。golang用于实现网络通信、行为树等逻辑，vue3用于编写网页界面

## 帮助
运行本项目，需要提前安装以下 [依赖](#依赖)，初次使用以及协议变更时，需要通过构建流程生成可执行文件。   
构建相关脚本：
- syncproto.bat: 同步项目的proto文件，需修改脚本中的项目路径
- makeproto.bat: 处理.proto，生成.pb.go以及对应的json schema文件
- build.bat: 编译项目，生成可执行文件（包含前两个过程）

## 项目结构
```
RobotTool # 项目根目录
├── behavior3editor # 行为树编辑器
├── cache # 缓存目录，chrome用户缓存所在目录
├── logs # 日志目录
├── src # golang源码目录
│   ├── app #
│   ├── common # 公共模块
│   ├── jsdef # 与js交互的定义
│   │   ├── config # 配置
│   │   └── jsmsg # 通信结构定义
│   ├── network # 通用网络相关
│   ├── robot # 机器人逻辑
│   │   ├── dyj # 大赢家机器人实现
│   │   │   ├── pb # protobuf文件目录
│   │   │   ├── pbgo # go语言protobuf文件目录
│   │   │   ├── schema # json schema文件目录
│   │   │   └── *.go # 源码文件
│   │   └── robot.go # 机器人接口文件
│   └── main.go # 主入口
├── ui # vue源码目录（界面）
│   ├── dist # 构建产出目录
│   ├── public # 公共资源目录
│   └── src # 源码目录
├── build.bat # 构建脚本
├── makeproto.bat
├── README.md
└── syncproto.bat
```

## 开发环境
### 1.golang
- go 1.20

### 2.vue3
- node.js

- vue.js 3.4.36

- vite 

创建vue工程：
```
npm create vue@latest
```

## 依赖
### 1.protoc

### 2.protoc-gen-jsonschema
go install github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema@latest