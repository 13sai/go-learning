## 开发工具


### [json2go](https://mholt.github.io/json-to-go/)
用于将 json 文本转换为 struct.

### [gormt](https://github.com/xxjwxc/gormt)
mysql数据库转 struct 工具,可以将mysql数据库自动生成golang sturct结构，带大驼峰命名规则。带json标签

### [sql2go](http://stming.cn/tool/sql2go.html)
用于将 sql 语句转换为 golang 的 struct. 使用 ddl 语句即可。
例如对于创建表的语句: show create table xxx. 将输出的语句，直接粘贴进去就行。

### [toml2go](https://xuri.me/toml-to-go/)
用于将编码后的 toml 文本转换问 golang 的 struct.

### [curl2go](https://mholt.github.io/curl-to-go/)
用来将 curl 命令转化为具体的 golang 代码.

### [mysql 转 ES 工具](http://www.ischoolbar.com/EsParser/)

### go 内置命令
go list 可以查看某一个包的依赖关系.
go vet 可以检查代码不符合 golang 规范的地方。

### [热编译工具](https://github.com/silenceper/gowatch)


### [revive](https://github.com/mgechev/revive)
golang 代码质量检测工具


### [Go Callvis](https://github.com/TrueFurby/go-callvis)
golang 的代码调用链图工具


### [Realize](https://github.com/oxequa/realize)
开发流程改进工具


### [Gotests](https://github.com/cweill/gotests)
自动生成测试用例工具


## 调试工具
### perf
代理工具，支持内存，cpu，堆栈查看，并支持火焰图.
perf 工具和 go-torch 工具，快捷定位程序问题.
https://github.com/uber-archive/go-torch
https://github.com/google/gops

### dlv 远程调试
基于 goland+dlv 可以实现远程调式的能力.
[delve](https://github.com/go-delve/delve)
提供了对 golang 原生的支持，相比 gdb 调试，简单太多。

### 网络代理工具
goproxy 代理，支持多种协议，支持 ssh 穿透和 kcp 协议.
[goproxy](https://github.com/snail007/goproxy)

### 抓包工具
go-sniffer 工具，可扩展的抓包工具，可以开发自定义协议的工具包. 现在只支持了 http，mysql，redis，mongodb.
基于这个工具，我们开发了 qapp 协议的抓包。
[go-sniffer](https://github.com/40t/go-sniffer)

### 反向代理工具，快捷开放内网端口供外部使用

ngrok 可以让内网服务外部调用
https://ngrok.com/
https://github.com/inconshreveable/ngrok

### 配置化生成证书
从根证书，到业务侧证书一键生成.
https://github.com/cloudflare/cfssl

### 免费的证书获取工具
基于 acme 协议，从 letsencrypt 生成免费的证书，有效期 1 年，可自动续期。
[acme](https://github.com/Neilpang/acme.sh)

### 开发环境管理工具，单机搭建可移植工具的利器。支持多种虚拟机后端。
vagrant常被拿来同 docker 相比，值得拥有。
https://github.com/hashicorp/vagrant

### 轻量级容器调度工具
nomad 可以非常方便的管理容器和传统应用，相比 k8s 来说，简单不要太多.
https://github.com/hashicorp/nomad

### 敏感信息和密钥管理工具
https://github.com/hashicorp/vault

### 高度可配置化的 http 转发工具，基于 etcd 配置。
https://github.com/gojek/weaver

### 进程监控工具 supervisor

### 基于procFile进程管理工具. 相比 supervisor 更加简单。
[procFile](https://github.com/ddollar/foreman)

### 基于 http，https，websocket 的调试代理工具，配置功能丰富。在线教育的 nohost web 调试工具，基于此开发

https://github.com/avwo/whistle

### 分布式调度工具

https://github.com/shunfei/cronsun/blob/master/README_ZH.md
https://github.com/ouqiang/gocron

### 自动化运维平台 Gaia

[Gaia](https://github.com/gaia-pipeline/gaia)