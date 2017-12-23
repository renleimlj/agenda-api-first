
## 我的工作

* 编写dockerfile
* 调试程序，修复bug，生成镜像
* 将镜像push到docker hub

## 工作工程及重要截图

### docker安装
过程略。

在测试能否正常使用的时候，遇到的问题：

 ![](http://img.blog.csdn.net/20171219224315099)

解决：
使用国内的Docker仓库daocloud。

然后就可以跑hello-world的demo镜像了。

 ![](http://img.blog.csdn.net/20171219224244119)


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

![](http://img.blog.csdn.net/20171219224058253)


为啥要用`golang：1.8`呢。想试试看别的版本会怎样，于是看了一下我的`docker version`，发现时1.8.3版本的，就想用`golang:1.8.3`建个镜像。

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

![](http://img.blog.csdn.net/20171219224213662 )




### 上传镜像

利用dockerhub的自动生成功能再次创建该镜像，并成功上传。

![](http://img.blog.csdn.net/20171223140230902 )

```
Building in Docker Cloud's infrastructure...
Cloning into '.'...

KernelVersion: 4.4.0-93-generic
Arch: amd64
BuildTime: 2017-08-17T22:50:04.828747906+00:00
ApiVersion: 1.30
Version: 17.06.1-ce
MinAPIVersion: 1.12
GitCommit: 874a737
Os: linux
GoVersion: go1.8.3
Starting build of index.docker.io/raincreek/agenda-api-first3:latest...
Step 1/6 : FROM golang:1.8

 ---> ba52c9ef0f5c

Step 2/6 : WORKDIR $GOPATH/src/agenda-api-first/cli

 ---> 63dec7122cf4

Removing intermediate container 406e761c5cd7

Step 3/6 : ADD . $GOPATH/src/agenda-api-first

 ---> 680537a2dcf9

Removing intermediate container ee6e41534ed0

Step 4/6 : EXPOSE 8080

 ---> Running in a01cb9394b32

 ---> dcd17279a060

Removing intermediate container a01cb9394b32

Step 5/6 : CMD agenda-api-first/main

 ---> Running in 903bca4ad90b

 ---> 888bea1928b6

Removing intermediate container 903bca4ad90b

Step 6/6 : VOLUME /data

 ---> Running in 590e5ad35228

 ---> 368b87105303

Removing intermediate container 590e5ad35228

Successfully built 368b87105303

Successfully tagged raincreek/agenda-api-first3:latest

Pushing index.docker.io/raincreek/agenda-api-first3:latest...
Done!
Build finished
```


[【镜像地址】](https://hub.docker.com/r/raincreek/agenda-api-first3/builds/)




