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

## nginx配置文件
安装`nginx`后用`res`目录下的`nginx.conf`替换原来的即可
此时,后端的端口最好别改(default 9090,如果改的话修改`nginx.conf`下http->server->location)
nginx的端口可以随便改(我默认8086)
1. / 根目录为前端目录,映射`JUST_PATH`的上一层路径
2. /api 后端api路径,如`/api/v1/courses`
3. /res 后端测试demo,映射`JUST_PATH/res`目录

## API page:
http://localhost:8086/res/dist/

# 更新记录

## 4.30
1. user的view去掉role,改成单独的role_name,因为role的id传递过去也没什么意义,而且最好不要传递,设计到的接口有 `POST /tokens`和 `GET /courses`
2. 加入`Authorization`控制,在`Header`的`Authorization`里面增加`token`(不知道前端是不是要序列化,反正我的后端是反序列化得到的`token`),白名单在`etc/config.json`的`while_list`控制.
3. 在swagger里面测试API的时候需要在`api_key`那边加上已经注册的邮箱,我加了一层中间件会自动给这些request加header`Authorization`
4. 接口 : 用户关注课程 `users/courses`
