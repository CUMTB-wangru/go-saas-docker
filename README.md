# go-saas-docker
docker部署golang前后端分离项目

# [Docker部署Go+Mysql+Redis ](https://www.cnblogs.com/wylshkjj/p/16694290.html)

两种方式Docker和Docker Compose部署web项目，相对于Go语言来说，不管是使用docker部署还是直接服务器部署都相当方便，比python要简单很多。

## 1、Dockerfile结构解析

**From**

我们正在使用基础镜像`golang:alpine`来创建我们的镜像。这和我们要创建的镜像一样是一个我们能够访问的存储在Docker仓库的基础镜像。这个镜像运行的是alpine Linux发行版，该发行版的大小很小并且内置了Go，非常适合我们的用例。有大量公开可用的Docker镜像，请查看https://hub.docker.com/_/golang

**Env**

用来设置我们编译阶段需要用的环境变量。

**WORKDIR，COPY，RUN**

WORKDIR：切换工作目录

COPY：把文件夹/文件复制到容器中

RUN：用于执行后面跟着的命令行命令，对于Go是把代码编译成可执行的二进制文件

**注**：Dockerfile 的指令(**RUN**)每执行一次都会在 docker 上新建一层。所以过多无意义的层，会造成镜像膨胀过大。

例如：

```dockerfile
FROM centos
RUN yum -y install wget
RUN wget -O redis.tar.gz "http://download.redis.io/releases/redis-6.2.6.tar.gz"
RUN tar -xvf redis.tar.gz
```

以上执行会创建 3 层镜像。可简化为以下格式：

```dockerfile
FROM centos
RUN yum -y install wget \
    && wget -O redis.tar.gz "http://download.redis.io/releases/redis-6.2.6.tar.gz" \
    && tar -xvf redis.tar.gz
```

如上，以 **&&** 符号连接命令，这样执行后，只会创建 1 层镜像。

**EXPORT，CMD**

声明服务端口，因为应用程序监听的是这个端口并通过这个端口对外提供服务。并且还定义了在运行镜像的时候默认执行的命令`CMD ["/dist/app"]`。CMD指令指定的程序可被**docker run**命令行参数中指定要运行的程序覆盖掉。

**注**：如果 Dockerfile 中如果存在多个 CMD 指令，仅最后一个生效。

## 2、Docker的Golang Web项目部署方式

### [示例项目的代码地址](https://gitee.com/wylshkjj/go-saas/tree/master/tracer)

项目中的tracer就是go部署的示例文件，需要在根目录tracer目录下创建配置文件夹config，在congif下创建配置文件config.ini，里面缺少的配置文件详细我粘贴在下面，将下面的配置粘贴到文件config.ini中：

注：host.docker.internal这个是用在本地Docker容器部署后访问本机的地址，也就是说使用Docker跑起项目后，需要关联mysql数据库和redis，而mysql和redis是安装在本的地并没有创建对应的Docker容器，所以需要使用host.docker.internal来访问，127.0.0.1是不能访问到的，host.docker.internal这个方法在Go中有版本限制，某些较低版本不可用。

```bash
[server]
AppMode = debug
HttpPort = :8000
JwtKey = 123qwe456asd789zxc

[database]
Db = mysql
DbHost = host.docker.internal
; DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassWord = 1234
DbName = tracer

[redis]
Rdb = redis
RdbHost = host.docker.internal
; RdbHost = 127.0.0.1
RdbPort = 6379
RdbUser =
DdbPassWord =
RdbName = 0
RdbPoolSize = 100
```

**目录结构：**(注：Dockerfile、wait-for-it.sh、docker-compose.yml文件自己创建到根目录tracer下)

```bash
.
├── Dockerfile
├── app
│   ├── blog
│   │   ├── blog.go
│   │   └── routers.go
│   └── user
│       ├── mode.go
│       ├── router.go
│       └── views.go
├── build
├── config
│   └── config.ini
├── dao
│   ├── db
│   │   └── user.go
│   └── redis
│       └── rdb.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── init.sql
├── log
│   ├── log
├── main.go
├── middleware
│   ├── cors.go
│   ├── jwt.go
│   └── logger.go
├── model
│   ├── db.go
│   ├── price_policy.go
│   ├── rdb.go
│   └── user.go
├── routers.go
├── tmp
│   ├── build-errors.log
│   └── main
├── utils
│   ├── errmsg
│   │   └── errmsg.go
│   ├── settings.go
│   ├── sms
│   │   ├── Tcaptcha.go
│   │   └── sms.go
│   └── tool.go
└── wait-for-it.sh
```

### **Dockerfile配置，分阶段构建**

Go程序编译后得到的是一个可执行的二进制文件，所以在最终的镜像中是不需要Go编译器的，也就是说只需要运行一个二进制文件即可。所以可以通过仅保留二进制文件来减小镜像的大小，而实现这种方式的技术称为**多阶段构建技术**，这就意味着可以通过多个步骤构建镜像。

把COPY静态文件的步骤放在上层，把COPY二进制可执行文件放在下层，争取多使用缓存。

```dockerfile
FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 tracer
RUN go build -o tracer .

# 创建一个小镜像
FROM scratch

# 静态文件
COPY ./log /log
COPY ./config /config
# 当出现static文件或者template文件的时候需要配置
# COPY ./static /static
# COPY ./template /template

# 从builder镜像中把/dist/tracer 拷贝到当前目录
COPY --from=builder /build/tracer /

# 需要运行的命令
ENTRYPOINT ["/tracer", "config/config.ini"]
```

## 3、关联其他容器

上面示例使用的方式是Docker+本地的mysql和本地的redis环境，关联其他容器后使用的是Docker的镜像环境

容器别名`mysql8019`；`root`用户密码为`1234`；挂载容器中的`/var/lib/mysql`到本地的`/Users/user/docker/mysql`目录；内部服务端口为3306，映射到外部的13306端口。

### 1.创建容器映射本地并运行

> **关联mysql**

```bash
# mysql
docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=1234 -v /Users/user/docker/mysql:/var/lib/mysql -d mysql:8.0.19
```

创建好后使用命令测试进入Docker容器中的Mysql：`mysql -uroot -p1234 -h 127.0.0.1 -P13306`

**示例：**

```bash
user@C02FP58GML7H ~ % mysql -uroot -p1234 -h 127.0.0.1 -P13306 
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.19 MySQL Community Server - GPL

Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```

> **关联redis**

docker运行的每个容器都是隔离的，redis默认不允许外部连接，因此想要部署在docker容器内的应用连接上redis，需要修改redis默认配置，这里我们以配置文件运行redis即可，也就是挂载本地的配置文件到docker容器。

配置文件需要去下载一份redis的压缩包解压后提取或复制redis.conf文件或者内容放到一个位置，此位置是用来专门挂载配置文件的。[redis下载地址](https://github.com/antirez/redis/releases)下载redis发行版，版本选择与容器内的相匹配即可。

博主本人是在mac本地是安装了redis的，所以偷懒直接引用了本地redis里面的配置文件进行挂载。这里默认使用root为用户名，密码为空，不输入即可。

```bash
# redis
docker run --name redis626 -p 16379:6379 -v /Users/user/redis-6.2.6/redis.conf:/etc/redis/redis.conf -v /Users/user/docker/data:/data --restart=always -d redis:6.2.6 redis-server --appendonly yes
```

创建好后使用命令测试进入Docker容器中的Redis：`./redis-cli -c -h 127.0.0.1 -p 16379`

**示例：**（注：我没有配置环境变量，所以需要先切换到redis的目录下去使用本地的redis-cli）

```bash
user@C02FP58GML7H src % cd /Users/user/redis-6.2.6/src
user@C02FP58GML7H src % ./redis-cli -c -h 127.0.0.1 -p 16379
127.0.0.1:16379> 
```

### 2.修改配置中的host为镜像别名，重新构建镜像

**配置信息：**

```bash
[server]
AppMode = debug
HttpPort = :8000
JwtKey = 123qwe456asd789zxc

[database]
Db = mysql
; DbHost = host.docker.internal
; DbHost = 127.0.0.1
DbHost = mysql8019
DbPort = 3306
DbUser = root
DbPassWord = 1234
DbName = tracer

[redis]
Rdb = redis
; RdbHost = host.docker.internal
; RdbHost = 127.0.0.1
RdbHost = redis626
RdbPort = 6379
RdbUser =
DdbPassWord =
RdbName = 0
RdbPoolSize = 100
```

> **构建镜像命令：**docker build . -t tracer_containers

> **运行容器命令：**docker run --link=mysql8019:mysql8019 --link=redis626:redis626 -p 8888:8000 tracer_containers

**注：**运行`tracer_containers`容器的时候需要使用`--link`的方式与上面的`mysql8019`容器和`redis626`容器关联起来，还有若显示redis无法连接修改映射的配置文件中的bind=0.0.0.0即可

## 4、Docker Compose模式

除了使用**--link**的方式来关联两个容器之外，还可以使用**Docker Compose**定义和运行多个容器。

**Compose：**是用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，可以使用 YML 文件来配置应用程序需要的所有服务。然后只要使用一个命令，就可以从 YML 文件配置中创建并启动所有服务。

**使用Compose的三步过程：**

1. 使用**Dockerfile**定义你的应用环境以便可以在任何地方复制。
2. 定义组成应用程序的服务，**docker-compose.yml**以便它们可以在隔离的环境中一起运行。
3. 执行**docker-compose up**命令来启动并运行整个应用程序。

**示例项目：**需要三个容器分别运行`mysql`、`redis`和`tracer_containers`，这里需要修改redis映射的配置，`bind 127.0.0.1 -::1`修改为：`bind 0.0.0.0 -::1`，由于我这个是redis6.2.6版本，其他版本就把`127.0.0.1`换成`0.0.0.0`即可，否则无法访问redis。

**注：**depends_on字段仅能保证web服务启动时，mysql服务处于Running状态而不是Ready状态，所以tracer_containers需要等待mysql启动后再启动，因此需要添加一个`wait-for-it.sh`脚本文件，检测mysql服务是否处于Ready状态也就是mysql是否已经启动完毕。[Docker官网给出的wait-for-it.sh文件地址](https://github.com/vishnubob/wait-for-it) 下载后放到项目的根目录下，配置文件中的mysql8019挂载了一个init.sql文件，是用来判断创建数据库的。

> **docker-compose.yml文件配置**

```yaml
# yaml 配置
version: "3.7"
services:
  mysql8019:
    image: mysql:8.0.19
    restart: "always"
    ports:
      - 13306:3306
    command: "--default-authentication-plugin=mysql_native_password --init-file /data/application/init.sql"
    # 账号密码
    environment:
      MYSQL_ROOT_PASSWORD: "1234"
      MYSQL_DATABASE: "tracer"
      MYSQL_PASSWORD: "1234"
    # 文件夹以及文件映射
    volumes:
      - ./init.sql:/data/application/init.sql
  redis626:
    # 镜像版本号
    image: redis:6.2.6
    # 端口号
    ports:
      - 16379:6379
    # 失败后总是重启
    restart: "always"
    # 以配置文件的方式启动 redis.conf
    command: "redis-server /etc/redis/redis.conf --appendonly yes"
    # 文件夹以及文件映射
    volumes:
      - /Users/user/docker/data:/data
      - /Users/user/redis-6.2.6/redis.conf:/etc/redis/redis.conf
  tracer_containers:
    # 相对当前 docker-compose.yml 文件所在目录，基于名称为 Dockerfile 的文件构建镜像
    build: .
    restart: "always"
    # command: sh -c  "./wait-for-it.sh mysql8019:3306 -- ./tracer ./config/config.ini"  # shell脚本方式启动
    command: ["/wait-for-it.sh", "mysql8019:3306", "--", "/tracer", "config/config.ini"]
    # 依赖启动项
    depends_on:
      - mysql8019
      - redis626
    ports:
      - 8888:8000
```

> **DockerFile文件修改**

```dockerfile
FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 tracer
RUN go build -o tracer .

# 创建一个小镜像
FROM debian:stretch-slim

# 静态文件
COPY ./wait-for-it.sh /
COPY ./log /log
COPY ./config /config

# 从builder镜像中把/dist/tracer 拷贝到当前目录
COPY --from=builder /build/tracer /

RUN chmod 755 wait-for-it.sh

# 需要运行的命令（注释掉下面这一行）
# ENTRYPOINT ["/tracer", "config/config.ini"]
```

> **执行命令：docker-compose up（docker-compose up -d是后台运行）**

运行成功后的部分结果（docker-compose up）：

```bash
tracer_containers_1  | wait-for-it.sh: mysql8019:3306 is available after 10 seconds
tracer_containers_1  | redis626:6379  0 100
tracer_containers_1  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
tracer_containers_1  |  - using env:    export GIN_MODE=release
tracer_containers_1  |  - using code:   gin.SetMode(gin.ReleaseMode)
tracer_containers_1  | 
tracer_containers_1  | [GIN-debug] POST   /api/user/sms             --> tracer/app/user.sendSms (3 handlers)
tracer_containers_1  | [GIN-debug] POST   /api/user/register        --> tracer/app/user.registerHandler (3 handlers)
tracer_containers_1  | [GIN-debug] POST   /api/user/login           --> tracer/app/user.loginHandler (3 handlers)
tracer_containers_1  | [GIN-debug] POST   /api/user/login_sms       --> tracer/app/user.loginSmsHandler (3 handlers)
tracer_containers_1  | [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
tracer_containers_1  | Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
tracer_containers_1  | [GIN-debug] Listening and serving HTTP on :8000
```

后台运行启动结果（docker-compose up -d）：

```bash
user@C02FP58GML7H tracer % docker-compose up -d
Creating network "tracer_default" with the default driver
Creating tracer_redis626_1  ... done
Creating tracer_mysql8019_1 ... done
Creating tracer_tracer_containers_1 ... done
```

