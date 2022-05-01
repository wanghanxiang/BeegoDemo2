###### 一、项目说明：

新增包后：
go mod tidy

起项目：
bee run


在beego中，目前支持三种数据库驱动，分别是：

MySQL： github.com/go-sql-driver/mysql

PostgreSQL：github.com/lib/pq

Sqlite3：github.com/mattn/go-sqlite3



**二、编译相关**

**编译镜像命令：**

docker build -t beegodemo:v1 .



docker-compose相关

**编译命令：**

docker-compose build

**启动服务命令：**

docker-compose up -d



##### 三、介绍beego

安装beego

```
go get github.com/astaxie/beego
```

安装bee工具

```
go get github.com/beego/bee
```

创建新项目

```
bee new ProjectName
```

启动项目

```
bee run 
```



**beego中文文档：**

https://beego.gocn.vip/beego/zh/developing/#%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B



###### docker端口映射

https://cloud.tencent.com/developer/article/1193005