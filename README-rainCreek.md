
## 我的工作

* 编写dockerfile
* 安装图形化管理界面portainer
* 将镜像push到docker hub

## 工作工程及重要截图

### docker安装
过程略。

在测试能否正常使用的时候，遇到的问题：

![]（http://img.blog.csdn.net/20171219213422152）

解决：
使用国内的Docker仓库daocloud。

然后就可以跑hello-world的demo镜像了。

![]（http://img.blog.csdn.net/20171219213550870）


### 编写dockerfile

就按教程里的来。

### 生成镜像

又遇到问题。

问题1：
```
rao@rao-HP-Pavilion-Notebook:~/Desktop/serviceComputing/goworkplace/src/github.com/agenda-api-first$ sudo docker build -t agenda .
Sending build context to Docker daemon  1.355MB
Step 1/8 : FROM golang:1.8
Get https://golang:1/v2/: dial tcp 221.179.46.190:1: getsockopt: connection refused
```

用curl测试`https://golang:1/v2/`
```
rao@rao-HP-Pavilion-Notebook:~/Desktop/serviceComputing/goworkplace/src/github.com/agenda-api-first$ curl https://golang:1/v2/

curl: (7) Failed to connect to golang port 1: Connection refused
```

很难过了。
和同学交流后，同学认为是被墙了。国内仓库不管用了，建议我改hosts。


然后出现了问题2：

```
rao@rao-HP-Pavilion-Notebook:~/Desktop/serviceComputing/goworkplace/src/github.com/agenda-api-first$ sudo docker build -t agenda .

Sending build context to Docker daemon  1.355MB
Step 1/8 : FROM golang:1.8
Get https://golang:1/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
```

又超时，似乎是熟悉的问题，于是照上面的办法又echo了一次，就可以了。

```
docker build -t agenda .
```

![]（http://img.blog.csdn.net/20171219215149376）


为啥要用1golang：1.81呢。想试试看别的版本会怎样，于是看了一下我的`docker version`，发现时1.8.3版本的，就想用`golang:1.8.3`建个镜像。

```
rao@rao-HP-Pavilion-Notebook:~/Desktop/serviceComputing/goworkplace/src/github.com/agenda-api-first$ sudo docker version
Client:
 Version:      17.06.2-ce
 API version:  1.30
 Go version:   go1.8.3
 Git commit:   cec0b72
 Built:        Tue Sep  5 20:00:33 2017
 OS/Arch:      linux/amd64

Server:
 Version:      17.06.2-ce
 API version:  1.30 (minimum version 1.12)
 Go version:   go1.8.3
 Git commit:   cec0b72
 Built:        Tue Sep  5 19:59:26 2017
 OS/Arch:      linux/amd64
 Experimental: false
```

尝试使用`golang:1.8.3`，发现竟然也可以成功build。
又尝试使用`golang:latest`。

得到3个images，show一下，来个大合照。

![]（http://img.blog.csdn.net/20171219215109344）


### 上传镜像



