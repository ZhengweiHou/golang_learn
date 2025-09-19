
## Cobra简明教程
### 创建新的CLI项目
#### 1. 安装Cobra工具
```
go install github.com/spf13/cobra-cli@latest
```

#### 2. 初始化新的CLI项目
- 创建项目及结构
```
mkdir cobra_cmd
cd cobra_cmd
go mod init cobra_cmd
cobra-cli init # cobra-cli会寻找当前项目mod name，作为import
```
这会生成基本的项目结构，包括：
- `cmd/` 目录：包含所有的命令代码
- `main.go` 文件：项目的入口文件
- `go.mod` 文件：Go模块依赖文件
