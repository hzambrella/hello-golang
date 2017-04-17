#必须先在linux下安装graphviz：
##sudo apititude graphviz
#环境变量
##source env.bash

##ctrl:控制
##engine:数据。在init.go修改测试数据。
##lib:自己的包
##vender:GitHub等外来包

#运行方法：
##在ctrl目录下go test即可。
##图片在ctrl/png目录下

##功能
##	将图分为各个连通子图。若子图的链的长度超过significant，就将子图用使用graphviz+DOT绘制出。
