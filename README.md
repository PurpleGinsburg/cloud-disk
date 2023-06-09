# CloudDisk

>轻量级云盘系统，基于go-zero、zorm实现

集成go-zero和xorm
xorm内置了三种IMapper实现：core.SnakeMapper,core.SameMapper和core.GonicMapper,默认的名称映射规则为SnakeMapper(支持struct为驼峰式命名，表结构中为下划线命名之间的转换。)
corelogic.go中的Core写具体业务逻辑

使用到的命令
```text
# 创建API服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yaml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero

```

```
用户模块
代码逻辑：req传json数据，在logic中通过req在mysql中取对应值

密码登陆接口
利用 Name 和 Password 返回 Token
基于 Token 的身份验证方法:
1、客户端使用用户名跟密码请求登录
2、服务端收到请求，去验证用户名与密码
3、验证成功后，服务端会签发一个 Token，再把这个 Token 发送给客户端
4、客户端收到 Token 以后可以把它存储起来，比如放在 Cookie 里或者 Local Storage 里
   客户端每次向服务端请求资源的时候需要带着服务端签发的 Token
5、服务端收到请求，然后去验证客户端请求里面带着的 Token，如果验证成功，就向客户端返回请求的数据
主要用于前后端分离项目

JWT（JSON Web Token）主要由三部分组成：Header（包含 Token 的类型（即 JWT）和所使用的算法，通常采用的算法是 HMAC SHA256 或 RSA），Payload(包含了要传递的信息) 和Signature（使用私钥对 Header 和 Payload 进行签名生成的，用来验证消息确实是由发送方发出的，以及在传输过程中没有被篡改过）
go 中 jwt 的使用方法
创建一个自定义的 Claims（ go.jwt包中使用jwt.StandardClaims 加 自定义类型），使用 JWT 签名算法生成 token

邮箱验证码登陆接口
想在编程环境中发送邮件，首先需要到自己所使用的邮箱中开启POP3/SMTP服务后获取授权密码

抽取MySQL和redis配置
go-zero配置文件的加载是通过将Yaml中的信息 映射到config.go中
go-zero主要是通过 conf.MustLoad(*configFile, &c) 这个方法去做映射，可以不用默认的etc/core-api.yaml 通过-f参数你可以用任何想映射的配置文件
在MustLoad-loaders中定义了三种格式并且指定了将[]byte映射到struct的方法；可以看到目前仅支持json,yaml,yml三种文件格式

文件上传
用到了腾讯云存储COS，独立于自己服务器的文件存储，可以分担服务器存储压力，分担服务器带宽成本
存储桶域名为：	tanyuyan-19990720-1318633606.cos.ap-nanjing.myqcloud.com

TODO:Q：twelve-factor应用提倡将配置存储在环境变量中。任何从开发环境切换到生产环境时需要修改的东西都从代码抽取到环境变量里，本项目我使用了os.Setenv配置环境变量。
解决方案1：go get github.com/joho/godotenv可以在可执行程序相同目录下，添加一个.env文件，从.env文件中读取配置， 然后存储到程序的环境变量中。
但是go run生成的 exe 为 temp 文件，这种情况如何配置环境变量？

```

```
存储池模块
代码逻辑：req打json标签定义数据，方便我们在handler中处理文件，把相关信息后丢入req,再在logic中处理
TODO:将返回文件和用户关联存储
```
