golang 原生服务端开发模板(2017-7-24)
==================
   用原生http简单地模仿nd项目 。尽量实现一些框架的功能,如gin的logger中间件(middleware.go)。
   开发模板有利于帮助理解gin和echo和项目流程,以及net/http包的用法
   缺陷很多，比如模板缓存未实现。未采用上下文context机制 。参考:http://www.tuicool.com/articles/RNvMRbm
   
   作用:说明原生的go  net/http也可以搭建服务。如现在的app 服务，newding 是原生,
   没用什么框架。（何工的高科技框架由于离职被废弃了。。现在好像还在n3d,读不懂。
   已实现用户登陆鉴权流程 auth.go login.go

## 1.配置文件etc
   配置文件是json格式。真正项目应当采用ini格式。ini格式解析包:https://github.com/Unknwon/goconfig
    
   - db.cfg 数据库(在这里修改数据库配置文件)
   - server.cfg 服务器配置(在这里修改端口号)

## 2.框架:
### 推荐（看api就会的，好用）
   - [gin](https://github.com/gin-gonic/gin)
   - [echo](http://go-echo.org/)

### 其他

   - [uweb(柳丁,如后台 )](https://github.com/ot24net/uweb)
   - [xp(xjp 师兄)](https://git.oschina.net/JinpengXu/xp.git)
   - [fasthttp(何工采用过的 )](http://www.qingpingshan.com/jb/go/148471.html)
   - beego 

## 3.newding服务端其他技术（我还没搞懂）：
### (1)反向代理:nginx， haproxy
   现在用haproxy

   使用haproxy：  
        sudo aptitude install haproxy
        sudo  vim  /etc/haproxy/haproxy.cfg
        编辑 ...
        sudo /etc/init.d/haproxy reload 

### (2)守护进程:supervisor
   跑服务器控制终端不能关闭，怎么样关闭控制终端也能跑服务呢？
   见sup  /golibs/src/git.ot24.net/go/sup

## 4:golang后端如何和前端交互
### 数据渲染：
   - (1)go模板  （template/html包），将数据加载到html上，发送给客户端浏览器
   - (2)vue.js: 先将html发送给客户端浏览器。html上的vue会再次请求数据（这个具体怎么做到的我不知道）。服务端收到请求后,将 json格式的数据发送过去。vue.js将json数据渲染到html上的vue模板上。
    
### 按钮:
   按钮触发一些事件。涉及到和后端交互的是ajax异步请求。

### 安卓:
   /jtw/src/app
   安卓重复传输的数据,比如手机设备信息,newding 称它为 ctx(公共头)，又称协议。
   它以json格式放到 request.Body resp.Body中:

    Body: ctx{
        协议数据
        data:[真正的数据]
    }

   数据有newding自己的加密方法。
   这意味着安卓的请求全是post 请求。为了减轻压力，请求还有缓存。
   详见：/jtw/src/app
   开发app服务端是近一年前的事情，还有些细节不得而知


## 5:其他
### 微信

   用的最多的：见微信开发者文档，
 [微信网页授权](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842)  
   jssdk做分享
   流程见各个项目的wxauth.go
    
### 支付宝

 [zfb网页授权](https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.S9FnTv&treeId=193&articleId=105193&docType=1)
   流程见各个项目的aliauth.go
   自己开发要申请测试号(我没弄过)。服务号要钱要域名,还要备案,很麻烦

## 6.还有许多东西，如rpc，甚至不知道的，正在挖掘中
