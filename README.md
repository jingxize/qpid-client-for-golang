use golang to send and receive message to  message queue which's type is qpid

QPID 是 Apache Foundation 的一个开源项目，是一个 AMQP 实现。关于AMPQ和QPID的详细资料，可以查看apache的官方网站：https://qpid.apache.org/ 这里是用golang语言实现的一个QPID的客户端，可用于消息的收发。

本项目的环境：

操作系统：OS X El Capitan 10.11.6

system :OS X El Capitan 10.11.6

golang版本：go1.9.2 darwin/amd64

golang version:go1.9.2 darwin/amd64

注意：如果大家在其它环境下测试这个项目，需要重新编译pack.ag下的amqp包。因为go 对不同文件在不同环境下编译的结果是不一样的。

NOTE:if your system is not OS X El Capitan,you must recompile the package /src/pack.ag/amqp.Because
     one package at different system will get different binary file.


1、部署

   因为这个项目主要是一个golang连接qpid的演示实例,不包括qpid服务端的实现，所以部署十分简单。首先，当然是把项目下载到本地了。

   git clone https://github.com/jingxize/qpid-client-for-golang.git
1、install

   Because the project is just a demo for use golang to connect qpid, don’t include the service side of the qpid,so its install is easy,you just need download the project from the github.

   git clone https://github.com/jingxize/qpid-client-for-golang.git
2、测试

   比如我们把项目代码下载到的 /usr/local/golangqpid/。
   首先，我们需要把这个路径加到GOPATH,不然，编译代码会提示找不到对应的包文件。
   export GOPATH=/usr/local/golangqpid
   接着，我们编译go文件。
   进入到测试的脚本文件目录：cd src
   编译测试的脚本文件：go build qpid.go
   执行编译成功的二进行文件： ./qpid (如果当前目前找不到生成的二进制文件，请用go env命令看自己的二进制文件的目录设置)


   这时，我们会看到输出结果：Message received: Hello!


2、runing

   now ,we have download the project to the folder, /usr/local/golangqpid
   then,we need add this path to the GOPATH.  export GOPATH=/usr/local/golangqpid
   then,we can compile the go file.   cd src & go build qpdi.go
   then,we can run the binary file.   ./qpid

   now ,we can see the out put :Message received: Hello!




声明：项目中用到的pack包（amqp）来自https://github.com/vcabbage/amqp

NOTE: the pack amqp  is from https://github.com/vcabbage/amqp
