# gaodeng

用Go封装了高灯电子发票SDK。

### 简介

* 支持全局配置`AppKey`和`AppSecret`。
* 全部参数和返回值均使用`struct`类型传递，而不是`map`类型。
* 调用时，按`struct`传入参数，内部自动将其转换为原生SDK支持的`map`类型。
* 返回时，先按通用类型解析：
  * 异常时，记录日志；
  * 正常后，仅返回业务部分数据(即`data`部分)的字节流，需要在具体的业务内转为对应的`struct`类型。

### 代码简介

* `constant.go`：常数定义。
* `model.go`：通用返回类型定义。
* `config.go`：HTTP请求客户端的环境配置。
* `client.go`：HTTP请求客户端。
* `gd_*.go`：相关的业务接口。
* `gd_*_test.go`：业务接口对应的单元测试类。

### 接口

* `gd_invoice_blue.go`：发票开具接口。
* `gd_invoice_blue_callback.go`：发票开具的异步通知信息。
* `gd_invoice_red.go`：发票冲红接口(批量查询，但一次最多20张)。
* `gd_invoice_status.go`：查询发票信息接口(单张发票的查询)。
* `gd_invoice_amount.go`：查询发票余量接口。
* `gd_invoice_verify.go`：发票查验接口。

### 测试

需要跟高灯的商务人员，获取测试账号，然后在环境变量中，添加`client_test.go`中需要的测试变量：

```shell
export GDTestAppKey=xxxxxxxx
export GDTestAppSecret=xxxxxxxx
export GDTestTaxPayerNumber=xxxxxxxx
export GDTestSellerName=xxxxxxxx
export GDTestSellerAddress=xxxxxxxx
```

然后直接运行：

```go
go test
```

### 资料

* [云票儿开发文档](http://developer-doc.fapiaoer.cn/)
* [Github开源API库](https://github.com/golden-corp/openplatform-sdk-go)
