# 目录结构

1. bin:可执行路径
2. etc:配置文件路径
3. pkg（可忽略):编译文件
4. src:源文件,其中just.com为自己编写的代码
5. log:日志文件
6. src/just.com/main.go:主程序,可以`go run main.go dev(production)`运行程序

# 说明
### 1. 软件版本
1. go 1.4.2
2. PostgreSQL 9.3

### 2. 环境变量
1. 安装目录的bin放到PATH中
2. GOROOT:安装目录(golang安装路径),GOROOT=/usr/local/go
3. GOPATH:项目目录(server文件夹路径),GOPATH=~/project/just_bs/server
4. JUST_PATH:也是项目目录(server文件夹路径),这个是我自己定义的,为的是找到项目的配置文件和资源文件,JUST_PATH=~/project/just_bs/server

# get started
增加环境变量 JUST_PATH 为server文件夹路径
### 1. 修改配置
修改`etc`目录下的`config.json`,如端口号,数据库配置,七牛配置,send_cloud配置文件

### 2. 运行命令
    ```
    cd src/just.com
    go run main.go dev(production),默认启动方式为dev
    ```
    或者
    ```
    ./start.bash dev(production),默认启动方式为dev
    ```

### 3. nginx配置文件
如果需要访问前端Web工程,还需要通过`nginx`代理,这边其实可以优化一下,当时没想到更好的解决方案
安装`nginx`后用`res`目录下的`nginx.conf`替换原来的即可
此时,后端的端口最好别改(default 9090,如果改的话修改`nginx.conf`下http->server->location),`nginx`的端口可以随便改(我默认8086)
通过修改后,`nginx`中的虚拟服务器目录如下:
1. / 根目录为前端目录,映射`JUST_PATH`的上一层路径
2. /api 后端api路径,如`/api/v1/courses`
3. /res 后端测试demo,映射`JUST_PATH/res`目录

### 4. 数据库导入
 数据库备份文件在`res`文件夹中 mlearning.dump,包括数据和数据都在里面
 自己用PostreSQL的命令恢复一下
 
### 5. API page:
API利用Swagger UI,修改 `res/dist`目录下的`swagger.yaml`,具体yaml的数据结构可以看Swagerr UI的Github文档
本地访问地址 http://localhost:8086/res/dist/ (如果通过`nginx`代理的话,否则端口为9090)
为了防止API暴露,需要在界面左上角`api_key`中增加数据库中有效的邮箱

# 响应状态码说明
1. 200 | ok:成功状态,对应GET/PUT/DELETE/PATCH,不能用于post
2. 400 | Bad Request:请求格式错误
3. 401 | Unauthorized:未授权,在本系统中就是没有携带有效的token
4. 403 | Forbidden:授权成功,但是用户没有权限
5. 500 | 服务器内部错误


# banner 在线生成工具
http://patorjk.com/software/taag/#p=display&h=0&f=3D-ASCII&t=Hello%20World

# 更新记录(以下内容可以忽略)

## 4.30
1. user的view去掉role,改成单独的role_name,因为role的id传递过去也没什么意义,而且最好不要传递,设计到的接口有 `POST /tokens`和 `GET /courses`
2. 加入`Authorization`控制,在`Header`的`Authorization`里面增加`token`(不知道前端是不是要序列化,反正我的后端是反序列化得到的`token`),白名单在`etc/config.json`的`while_list`控制.
3. 在swagger里面测试API的时候需要在`api_key`那边加上已经注册的邮箱,我加了一层中间件会自动给这些request加header`Authorization`
4. 接口 : 用户关注课程 `users/courses`

## 5.1
1. 还是改成`/users/:user__id/courses`(统一成复数会好一点)干脆把获取其他人关注的课程也一起暴露出来
2. 增加部署时的环境变量,如则启动命令改成 `go run main.go dev(production)`,默认启动方式为`dev`
3. 用户更新接口
4. 根据id获取用户信息(貌似感觉这个接口没什么用)
5. files/tokens,token有限时间为20min,key由本地客户端生成(callback还没来得及写)

### 5.2
1. 更新课程 patch/courses

### 5.3 
1. 中间件修改
2. bug,关注课程列表,修改用户数据
3. 管理员接口:禁用账户/重置密码

### 5.6 
1. chapter的增加/列表和修改,其中order字段只是简单的起排序的作用
2. 课程的返回值也做了一定的修改
> 课程的增加和修改需要改表结构,我这边估计还有点问题

### 5.7
1. user list(`get /users`接口),其中返回的对象为UserDetailView,比普通的UserView增加了2个对象
2. 判断icon,空则返回default.png

3. chapter delete 接口
4. 新增课程接口 POST COURSE,其中 course的view有更新,多了 create_time,update_time和 attachment_list
5. 用户登出 `DELETE TOKENS`

### 5.8
1. course view增加 `mark_status`字段
2. 更新课程接口
3. 教师获取自己创建的课程

### 5.25
1. 课程章节接口修改,增加video
2. 增加脚本启动`start.bash`
dev环境 `./start.bash`
production环境 `./start.bash production`