mail-provider
=============

项目地址为:https://github.com/open-falcon/mail-provider.git
这里只是新增了smtp SSL协议,用于云服务器不能连接对外25号端口

## 安装方法

1.二进制安装(推荐)
下载编译好的二进制
```bash
wget https://raw.githubusercontent.com/linyy2233/http-mail/master/falcon-http-mail
```
修改cfg.json文件相关信息，使用
```bash
./control start
```
即可启动客户端

2.源码编译（如无科学上网方法，请勿尝试）
下载之后为源码，安装golang环境，环境配置参考[golang环境配置](http://book.open-falcon.org/zh/quick_install/prepare.html)
编译方法
```bash
cd $GOPATH/src
mkdir github.com/linyy2233/ -p
cd github.com/linyy2233/
git clone https://github.com/linyy2233/http-mail.git
cd http-mail
go get ./...
./control build
```
编译成功之后，修改cfg.json文件相关信息，使用
```bash
./control start
```
即可启动


## 使用方法
下载之后为源码，需要编译

```
curl http://$ip:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"
```

## FAQ

1.对于126.163等邮箱请控制发信频率以免被封
