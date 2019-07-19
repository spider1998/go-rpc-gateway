# go-rpc-gateway
基于Go语言的RPC+HTTP+Gateway+Swagger

##实现功能：通过grpc+gateway+swagger搭建微服务（GO+rpc+gateway+swagger+mysql+log+redis...），示例对外同时兼容HTTP与RPC接口(以  创建人员为例)，服务可在此之上扩展开发##
--

![](http://pic.w-blog.cn/800FD94A-504B-4AA1-9992-C98962E84287.png)

grpc-gateway是protoc的一个插件 。它读取Grpc服务定义，并生成反向代理服务器，将RESTful JSON API请求转换为Grpc的方式调用。主要是根据 google.api.http定义中思想完成的，以下为grpc-gateway结构图： ￼ 

![](https://github.com/spider1998/go-rpc-gateway/blob/master/img/grpc1.png)

一、环境准备
--

·安装ProtocolBuffers 3.0及以上版本:
```shell
{
  mkdir tmp
  cd tmp
  git clone https://github.com/google/protobuf
  cd protobuf
  ./autogen.sh
  ./configure
  make
  make check
  sudo make install
}
```

·go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

·go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

·go get -u github.com/golang/protobuf/protoc-gen-go

·go get -u google.golang.org/grpc

·go get github.com/elazarl/go-bindata-assetfs/...   （使用go-bindata所生成Swagger UI的Go代码，结合net/http对外提供服务）
```
swagger 配置：

·下载Swagger UI文件
  https://github.com/swagger-api/swagger-ui  将其源码下载下来，并将其dist目录下的所有文件拷贝到我们项目中的third_party/swagger-   ui去
·将Swagger UI转换为Go源代码【转换工具是go-bindata】
  它支持将任何文件转换为可管理的Go源代码。用于将二进制数据嵌入到Go程序中。并且在将文件数据转换为原始字节片之前，可以选择压缩文件数据

 安装
    go get -u github.com/jteeuwen/go-bindata/...
    完成后，将$GOPATH/bin下的go-bindata移动到$GOBIN下

转换
  在项目下新建pkg/ui/data/swagger目录，回到third_party/swagger-ui下，执行Makefile   make
检查
  回到pkg/ui/data/swagger目录，检查是否存在datafile.go文件
  
```



二、代码运行
--

--github上拉取Google官方api放入person/pkg/proto/目录  https://github.com/googleapis/googleapis 将google文件夹中需要的文件拷贝至     proto目录

--执行person/pkg/proto/makefile生成相关pb.go及gw.pb.go文件

--go run cmd/server/main.go
 
--访问路径https://127.0.0.1:50052/swagger/hello.swagger.json， 查看输出内容是否为person.swagger.json的内容

--访问路径https://127.0.0.1:50052/api/ 查看内容

   ![](https://github.com/spider1998/go-rpc-gateway/blob/master/img/grpc2.png)


--postman 测试http服务是否正常

   ![](https://github.com/spider1998/go-rpc-gateway/blob/master/img/grpc3.png)
   
--运行person/client 下的main.go测试grpc之间调用

-------------OVER-----------------



三、性能
--

在GO的场景下本身在同等HTTP的情况下经过grpc-gateway和不经过直接到API差距大概在20~30%左右，这样的性能消耗带来的是兼容HTTP并且还可以自动生成swagger（还可以作为调试工具），避免业务代码重写繁琐的过程。

四、参考文档
--

https://my.oschina.net/wenzhenxi/blog/3023874     Grpc-Gateway - Grpc兼容HTTP协议文档自动生成网关
https://segmentfault.com/a/1190000013513469       Grpc+Grpc Gateway实践



