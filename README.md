# Go版本agenda开发项目
对于该项目具体的实现以及相关包的使用见博客：https://blog.csdn.net/C486C/article/details/83663022
## 安装使用
获取agenda项目，并安装
```
go get github.com/cyulei/agenda 
go install github.com/cyulei/agenda
```
安装可能会出现`cannot find package "golang.org/x/..."`，这时候需要安装相应的Go工具，见[博客](https://blog.csdn.net/C486C/article/details/82871020)中安装Go工具部分。

## 目录结构
```
│  .gitignore                 git配置文件
│  .travis.yml                Travis CI需要的配置文件  
│  LICENSE                    cobra License 
│  main.go                    cobra生成的main文件
│  README.md                  README
│  
├─cmd                         命令具体实现
│      cancelmeeting.go       取消会议
│      changeparticipator.go  增删参与人员
│      clearmeeting.go        清除所有会议
│      createmeeting.go       创建会议
│      deleteuser.go          删除用户
│      exitmeeting.go         退出会议
│      login.go               登录
│      logout.go              登出
│      querymeeting.go        查询会议
│      queryuser.go           查询用户
│      register.go            用户注册
│      root.go                cobra生成的root文件
│      
├─datarw                      数据存取
│      Agenda.log             log日志
│      CurUser.json           保存当前登录的用户
│      meetingrw.go           会议的Json文件存取接口
│      Meetings.json          会议的Json文件
│      userrw.go              用户的Json文件存取接口
│      Users.json             用户的Json文件
│      
├─entity                      实体
│      Date.go                日期实体
│      EntityFunc.go          实体和命令的接口
│      EntityFunc_test.go     test函数
│      Meeting.go             会议实体
│      User.go                用户实体
│      
└─test_pic                    测试样例中的图片
```
## 命令介绍
以下是agenda中的所有命令介绍：
### agenda help

显示所有可用命令。

可在help后添加相应的命令名字，如`agenda help login`即可显示`login`命令的相关帮助

![help](https://github.com/kevinli36/agenda/raw/master/test_pic/help.JPG)

### agenda login
用户登录。用户登陆需要提供用户名和密码，只有在没有登陆的状态下，用户名注册过且密码正确则登陆成功。

使用范例：(already log out) 
```
agenda login -n username1 -p password
```

可用参数列表
```
  -h, --help              help for login
  -p, --password string   user password
  -n, --username string   user name
```
![login](https://github.com/kevinli36/agenda/raw/master/test_pic/login.JPG)

### agenda logout
用户登出。用户登出不用使用任何参数，在已经登陆的情况下可以退出成功。

使用范例：(already log in) 
```
agenda logout
```

可用参数列表
```
  -h, --help   help for logout
```
![logout](https://github.com/kevinli36/agenda/raw/master/test_pic/logout.JPG)

### agenda register
用户注册。注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。

使用范例：
```
agenda register -n username1 -p password
```

可用参数列表
```
  -h, --help              help for register
  -n, --name string       user's name
  -p, --password string   user's password
```
![register](https://github.com/kevinli36/agenda/raw/master/test_pic/register.JPG)

### agenda queryuser
用户查询。
已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。
使用参数-n，可以查询单个用户

使用范例：（已登录）
```
agenda queryuser                //查询所有用户
agenda queryuser -n username1   //查询单个用户
```

可用参数列表
```
  -h, --help          help for queryuser
  -n, --name string   user's name
```
![queryuser](https://github.com/kevinli36/agenda/raw/master/test_pic/queryuser.JPG)

### agenda deleteuser
用户删除
已登录的用户可以删除本用户账户（即销号）。
以该用户为 发起者 的会议将被删除
以该用户为 参与者 的会议将从 参与者 列表中移除该用户。若因此造成会议 参与者 人数为0，则会议也将被删除。

使用范例：（已登录）
```
agenda deleteuser
```

可用参数列表
```
  -h, --help   help for deleteuser
```
![deleteuser](https://github.com/kevinli36/agenda/raw/master/test_pic/deleteuser.JPG)

### agenda createmeeting

已登录的用户可以使用createmeeting创建一个新的会议。创建新的会议需要提供会议名称，会议开始/结束时间和会议的参加者。如果会议名称与已有会议相同，或者会议开始时间/结束时间无效，或者没有有效的会议参与者，创建会议都会失败。

使用范例：(already log in) 
```
agenda createmeeting -t m1 -s 2008-11-11-11-11 -e 2008-11-11-12-12 -p p1-p2-p3
```

可用参数列表
```
  -e, --end string     meeting end date
  -h, --help           help for createmeeting
  -p, --part string    meeting participators
  -s, --start string   meeting start date
  -t, --title string   meeting title
```
![createmeeting](https://github.com/kevinli36/agenda/raw/master/test_pic/createmeeting.JPG)

### agenda changeparticipator

已登录的用户可以更改自己作为sponsor的会议的参与者。更改会议参与者需要提供会议的名字，会议需要做的操作（增加/删除会议参与者）以及涉及的参与者名字。如果不存在此会议，或者说存在但会议sponsor不是当前用户，更改会议参与者的操作都会失败。在增加和删除会议参与者的过程中，该命令只会增加/删除有效的参与者，即：删除命令只会删除原本在会议中的参与者，增加命令只会增加原本不再会议中且空闲的user。

使用范例：(already log in) 
```
agenda changeparticipator -t m1 (-a) -d -p p1-p2
注：如果不写操作方式，默认为增加参与者，如果写了两个操作方式参数，视作进行删除操作
```

可用参数列表
```
  -a, --add            add participator (default true)
  -d, --delete         delete participator
  -h, --help           help for changeparticipator
  -p, --name string    participator's name
  -t, --title string   meeting title
```
![changeparticipator](https://github.com/kevinli36/agenda/raw/master/test_pic/changeparticipator.JPG)

### agenda clearmeeting

已登录的用户可以删除所有自己作为sponsor的会议。如果想要查看被删除的会议，可以在命令后添加-i

使用范例：(already log in) 
```
agenda clearmeeting (-i)
```

可用参数列表
```
  -h, --help   help for clearmeeting
  -i, --info   show meetings cleared
```
![clearmeeting](https://github.com/kevinli36/agenda/raw/master/test_pic/clearmeeting.JPG)

### agenda cancelmeeting

  已登录用户取消自己创建的会议记录，参数：会议标题（-t title）
    
   使用范例
```
   agenda cancelmeeting -t health
```
![cancelmeeting](https://github.com/kevinli36/agenda/raw/master/test_pic/cancelmeeting.JPG)

### agenda querymeeting

  查询会议并打印，参数：-a,-s 开始日期, -e 结束日期,
  
  范例：
  
```
  agenda querymeeting
```

  命令翻译：查询已登录用户的所有会议，包括作为发起者和参与者的会议，必须先登录

```
   agenda querymeeting -a
```

  命令翻译：查询所有用户的所有会议，允许未登录状态下查询
    
```
    agenda querymeeting -s 2018-1-1-1-1
```

  命令翻译：查询开始时间晚于2018年1月1日1时1分的会议
    
```
    agenda querymeeting -e 2018-1-2-1-1
```

  命令翻译： 查询结束时间早于2018年1月2日1时1分的会议
  
  注意：  -a,-s,-t可以随意套用

![querymeeting](https://github.com/kevinli36/agenda/raw/master/test_pic/querymeeting.JPG)

### agenda exitmeeting

  退出作为参与者的会议，如果是会议发起者，那么删除,必须先登录
  
  范例：
```
    agenda exitmeeting -t title
```
![exitmeeting](https://github.com/kevinli36/agenda/raw/master/test_pic/exitmeeting.JPG)
