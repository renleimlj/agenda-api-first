
# 服务程序开发实战 - Agenda

## 概述

利用命令行 或 web 客户端调用远端服务是服务开发的重要内容。其中，要点是如何实现 API First 开发，使得团队协作变得更有效率。

### 任务目标

1. 熟悉 API 设计工具，实现从资源（领域）建模，到 API 设计的过程

1. 使用 Github ，通过 API 文档，实现 agenda 命令行项目 与 RESTful 服务项目同步开发

1. 使用 API 设计工具提供 Mock 服务，两个团队独立测试 API

1. 使用 travis 测试相关模块

1. 利用 dockerfile 在 docker hub 上构建一个镜像，同时包含 agenda cli 和 agenda service， 如果 mysql 包含 服务器 和 客户端一样

## 镜像截图

使用国内的Docker仓库daocloud。

跑hello-world的demo镜像。

 ![](http://img.blog.csdn.net/20171219224244119)

使用`golang:1.8`生成镜像
`docker build -t agenda`

![](http://img.blog.csdn.net/20171219224058253)

`docker images`

![](http://img.blog.csdn.net/20171219224213662)

成功上传镜像到dockerhub。

![](http://img.blog.csdn.net/20171223140230902)


## 服务器测试截图

### gotest

![](https://s1.ax1x.com/2017/12/16/L3wcT.png)

### UserLogin_success

![](https://s1.ax1x.com/2017/12/16/L30jU.png)

### UserLogin_fail

![](https://s1.ax1x.com/2017/12/16/L3DuF.png)

### Logout_succeed

![](https://s1.ax1x.com/2017/12/16/L3rB4.png)

### Logout_fail

![](https://s1.ax1x.com/2017/12/16/L3sHJ.png)

### CreateMeeting_succeed

![](https://s1.ax1x.com/2017/12/16/L36E9.png)

### CreateMeeting_fail

![](https://s1.ax1x.com/2017/12/16/L3cNR.png)

### DeleteMeeting_succeed

![](https://s1.ax1x.com/2017/12/16/L3R9x.png)

### DeleteMeeting_fail

![](https://s1.ax1x.com/2017/12/16/L3g41.png)

### QueryMeeting_succeed

![](https://s1.ax1x.com/2017/12/16/L3W36.png)

## 命令行运行截图

### 用户注册

 1. 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。

 1. 如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。

    `main register --un=UserName --pw=password --email=a@xxx.com --phone=xxxxxx`

正确输入：
![此处输入图片的描述][1]

此时User中记录了该条注册信息：
![此处输入图片的描述][2]

缺失信息：
![此处输入图片的描述][3]

用户名已占用：
![此处输入图片的描述][4]

### 用户登录

 1. 用户使用用户名和密码登录 Agenda 系统。

 1. 用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。

    `main login --un=UserName --pw=password`

输入错误的账号密码，登录失败：
![此处输入图片的描述][5]

输入错误的账号密码，登录成功：

![此处输入图片的描述][6]

此时，CurUser文件中写着当前登录的账户名：
![此处输入图片的描述][7]

### 用户查询

 1. 已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。

    `main query --un=UserName`

新注册一个账号，并查询这个账号：
![此处输入图片的描述][8]

在没有登录的情况下，不能使用查询功能：

![此处输入图片的描述][9]

### 用户登出

 1. 已登录的用户登出系统后，只能使用用户注册和用户登录功能。
 
     `main logout`

登出成功，此时CurUser被销毁：

![此处输入图片的描述][10]

若没有登录就运行登出，则失败：

![此处输入图片的描述][11]

### 创建会议

###### 1.会议名是否已经存在 
###### 2.所有参与者是否存在
###### 3.会议的开始时间是否在结束时间之前
###### 4.新会议是否和发起者参与者其他会议有时间冲突

###### 当前用户是has,现有会议：
![此处输入图片的描述][12]
###### 测试：
![此处输入图片的描述][13]

## 取消会议

###### 用户可以取消自己作为sponser的会议
![此处输入图片的描述][14]




## 查询会议


用户输入时间段，agenda查询这个时间段内的所有会议
 1. 若该时间段内存在会议，agenda列出时间段内的所有会议。
 2. 若该时间段内不存在会议，agenda发出提示。
 
 以下为json文件中以创建的会议：
 
 ![此处输入图片的描述](http://img.blog.csdn.net/20171102095131425)
 
 
 `agenda qm --st=2017-08-01/00:00 --et=2017-12-01/00:00`
  
 
 `agenda qm --st=1999-01-01/00:00 --et=2000-01-01/00:00`

 1. 若该时间段内存在会议，agenda列出时间段内的所有会议。
 
 2. 若该时间段内不存在会议，agenda发出提示。
 
![此处输入图片的描述](http://img.blog.csdn.net/20171102095150701)


## 增删会议参与者


根据用户输入的会议名和用户名，agenda增删该会议参与者
 1. 若该会议存在，agenda增删该会议参与者
 （1）若删除参与者后，会议参与者人数为零，agenda自动删除该会议
 2. 若该会议不存在，agenda发出提示。



`addPr --title=happy --pr=newpPrhhhhh`

`delPr --title=new --pr=alice`

该会议存在，agenda增删该会议参与者

 ![此处输入图片的描述](http://img.blog.csdn.net/20171102095201796)

以下为增删会议参与者后的json文件：

 ![此处输入图片的描述](http://img.blog.csdn.net/20171102095139174)


`addPr --title=fruit --pr=orange`

该会议不存在，agenda发出提示。

 ![此处输入图片的描述](http://img.blog.csdn.net/20171102095206601)
 
 

`agenda delPr --title=new --pr=bob`

`agenda delPr --title=new --pr=daisy`

删除参与者后，会议参与者人数为零，agenda自动删除该会议

![此处输入图片的描述](http://img.blog.csdn.net/20171102095211205)

以下为agenda自动删除该会议后的json文件：

![此处输入图片的描述](http://img.blog.csdn.net/20171102095144806)


  [1]: https://s1.ax2x.com/2017/10/31/BCH5S.png
  [2]: https://s1.ax2x.com/2017/10/31/BCVQu.png
  [3]: https://s1.ax2x.com/2017/10/31/BCXUH.png
  [4]: https://s1.ax2x.com/2017/10/31/BCmgN.png
  [5]: https://s1.ax2x.com/2017/10/31/BCYc9.png
  [6]: https://s1.ax2x.com/2017/10/31/BC0fA.png
  [7]: https://s1.ax2x.com/2017/10/31/BCjIe.png
  [8]: https://s1.ax2x.com/2017/10/31/BCWXO.png
  [9]: https://s1.ax2x.com/2017/10/31/BCM9d.png
  [10]: https://s1.ax2x.com/2017/10/31/BCOOR.png
  [11]: https://s1.ax2x.com/2017/10/31/BCbUr.png
  [12]: http://ww1.sinaimg.cn/large/0060lm7Tly1fl2khgvzq2j315m040gma.jpg
  [13]: http://ww2.sinaimg.cn/large/0060lm7Tly1fl2pu2nh3fj31b407u421.jpg
  [14]: http://ww2.sinaimg.cn/large/0060lm7Tly1fl2pubbe9lj30oq04mq49.jpg
