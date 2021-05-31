FROM golang:alpine
#还需要加上一个mysql的环境部署，再把mysql的账号密码整成gorm 123456
#参考db.json里面的账号密码，这俩环境部署了应该就可以运行啦！！！

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
	
# 移动到工作目录：/home/www/goWebBlog 这个目录 是你项目代码 放在linux上  
# 这是我的代码跟目录 
# 你们得修改成自己的
WORKDIR /home/www/goWebBlog

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件  可执行文件名为 app
RUN go build -o app .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /home/www/goWebBlog 目录复制到这里
RUN cp /home/www/goWebBlog/app .
# 在容器目录 /dist 创建一个目录 为src
RUN mkdir src .
# 在容器目录 把宿主机的静态资源文件 拷贝到 容器/dist/src目录下
# 这个步骤可以略  因为项目是引用到了 外部静态资源
RUN cp -r /home/www/goWebBlog/src/static ./src/
# 声明服务端口
EXPOSE 9090

# 启动容器时运行的命令
CMD ["/dist/app"]