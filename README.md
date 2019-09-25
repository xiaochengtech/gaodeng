# gaodeng

[![Latest Tag](https://img.shields.io/badge/tag-v0.0.1-blue.svg)](https://gitee.com/cuckoopark/gaodeng/releases)

用Golang封装了高灯电子发票SDK。

文档地址：[http://developer-doc.fapiaoer.cn/](http://developer-doc.fapiaoer.cn/)

### 简介

* 支持全局配置`AppKey`和`AppSecret`。
* 全部参数和返回值均使用`struct`类型传递，而不是`map`类型。
* 调用时，按`struct`传入参数，内部自动将其转换为原生SDK支持的`map`类型。
* 返回时，先按通用类型解析：
  * 异常时，记录日志；
  * 正常后，仅返回业务部分数据(即`data`部分)的字节流，需要在具体的业务内转为对应的`struct`类型。

### 术语清单

| 术语 | 对应 | 使用的系统 |
| ---- | ---- | ---- |
| 服务商 | 布谷 | 服务商平台 |
| 销货方 | 车场 | 商家平台　|
| 购买方 | 车主 | 微信 |

### 代码简介

* `constant.go`：常数定义。
* `model.go`：通用返回类型定义。
* `config.go`：HTTP请求客户端的环境配置。
* `client.go`：HTTP请求客户端。
* `gd_*.go`：相关的业务接口。
* `gd_*_test.go`：业务接口对应的单元测试类。

### 接口

* `gd_invoice_blue.go`：发票开具接口。
* `gd_invoice_red.go`：发票冲红接口(批量查询，但一次最多20张)。
* `gd_invoice_status.go`：查询发票信息接口(单张发票的查询)。
* `gd_invoice_amount.go`：查询发票余量接口(`todo`)。
* `gd_invoice_verify.go`：发票查验接口。

> 开票结果异步通知接口，将写在云端服务中
