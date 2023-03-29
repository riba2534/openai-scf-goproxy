# 腾讯云函数1分钟搭建 OpenAI 国内代理

> 项目地址: https://github.com/riba2534/openai-scf-goproxy

最近有一个好消息，OpenAI 开放了自己的 API，开发者可以很方便的调用各种语言模型来完成自己的创意，但是由于众所周知的原因国内访问 OpenAI 时接口可能大概率超时或者调不通，那解决无非是通过 proxy 的方式：

- 直接在境外服务器运行自己的服务，缺点是国内访问可能比较慢
- 国内服务器运行服务，把 OpenAPI 的相关请求用境外服务器做一层转发

本文介绍一种对于国内相对而言比较方便的办法，使用腾讯云函数来完成一个指向 OpenAI 的反向代理服务搭建，完成后开发者开发时直接把请求 OpenAPI 的接口直接指向腾讯云函数的地址即可。

直接开始正题


## 第一步：新建云函数

1. 打开腾讯云函数控制台: https://console.cloud.tencent.com/scf/list?rid=5&ns=default
2. 页面左边「函数服务」中，点击「新建」，然后照着下面图填：

- 点「从头开始」
- 函数类型选 web函数
- 名称自己随便填
- 地域选择一个境外的，推荐新加坡(香港好像不在openai支持地区内)
- 运行环境选 Go1
- 时区选上海
- 提交方法：本地上传zip包
- 日志投递也推荐选上，方便看日志
- 触发器配置照着图看

![云函数.png](https://image-1252109614.cos.ap-beijing.myqcloud.com/2023/03/09/64096590d8255.png)

3. 注意，上传的 zip 包可以在本项目 [releases](https://github.com/riba2534/openai-scf-goproxy/releases) 中下载到，最新的包地址是： [main.zip](https://github.com/riba2534/openai-scf-goproxy/releases/download/V2.0/main.zip)


## 第二步：查看部署信息

新建好之后，在腾讯云函数列表中找到你刚创建的，从左边 「函数管理」-> 「函数代码」，找到你的访问路径

![1678337783998.png](https://image-1252109614.cos.ap-beijing.myqcloud.com/2023/03/09/640966f88f891.png)


这个访问路径就是你之后请求 OpenAPI 的访问路径，访问路径的格式是 `https://service-xxxxxx.hk.apigw.tencentcs.com/release/`

注意: 这里的访问路径后面有个 `/release/` 你在用的时候把这个去掉，即: `https://service-xxxxxx.hk.apigw.tencentcs.com`



**重要提示**：云函数默认访问的超时时间较短，而调用 openai 的时间可能很长，所以我们需要改一下云函数配置，把超时时间调大，在左边「函数管理」-> 「函数配置」 里面，把访问的超时时间和并发度调大，如下图：

[![超时时间.png](https://image-1252109614.cos.ap-beijing.myqcloud.com/2023/03/13/640ece0cc7848.png)](https://image-1252109614.cos.ap-beijing.myqcloud.com/2023/03/13/640ece0cc7848.png)

[![并发配置.png](https://image-1252109614.cos.ap-beijing.myqcloud.com/2023/03/13/640ece0c4dd58.png)](https://image-1252109614.cos.ap-beijing.myqcloud.com/2023/03/13/640ece0c4dd58.png)

### 大功告成

至此，一个指向 openAPI 的反向代理就搭好了，你在开发的时候使用国内服务器，只需要把 `api.openapi.com` 换成这个新的地址就可以了.

我们可以通过类似 postman 这种工具来测试一下是否可用，查询一个完成模型试试，可以看到，成功的返回了信息!

![1678338111283.png](https://image-1252109614.cos.ap-beijing.myqcloud.com/2023/03/09/6409683fa484c.png)


# 玩耍

接下来就需要去看看 OpenAI 的接口文档了: https://platform.openai.com/docs/introduction
