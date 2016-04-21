## bin
可执行路径
## etc
配置文件路径
## pkg（可忽略）
编译文件
## src
源文件

## get started
增加环境变量 JUST_PATH 为server文件夹路径

## 环境变量
1. 安装目录的bin放到PATH中
2. GOROOT:安装目录(golang安装路径),GOROOT=/usr/local/go
3. GOPATH:项目目录(server文件夹路径),GOPATH=~/project/just_bs/server
4. JUST_PATH:也是项目目录(server文件夹路径),这个是我自己定义的,为的是找到项目的配置文件和资源文件,JUST_PATH=~/project/just_bs/server

## 修改配置
修改`etc`目录下的`config.json`,能改的也只有端口

## 运行命令
```
cd src/just.com
go run main.go
```