use golang to send and receive message to one message queue which's type is qpid

QPID 是 Apache Foundation 的一个开源项目，是一个 AMQP 实现。关于AMPQ和QPID的详细资料，可以查看apache的官方网站：https://qpid.apache.org/ 这里是用golang语言实现的一个QPID的客户端，可用于消息的收发。

本项目的环境：

操作系统：OS X El Capitan 10.11.6

golang版本：go1.9.2 darwin/amd64

1、部署
   因为这个项目主要是一个golang连接qpid的演示实例,不包括qpid服务端的实现，所以部署十分简单。首先，当然是把项目下载到本地了。

1、install
   Because the project is just a demo for use golang to connect qpid, don’t include the service side of the qpid,so its install is easy,you just need download the project from the github.

2、测试
   比如我们把项目代码下载到的 /usr/local/golangqpid/。首先，我们需要把这个路径加到GOPATH,不然，编译代码
   会提示找不到对应的包文件。
   export GOPATH=/usr/local/golangqpid

   
