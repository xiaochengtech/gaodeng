package gaodeng

type InvoiceBlueCallbackBody struct {
	AppKey                  string `json:"appkey"`                      // 服务商appkey
	Message                 string `json:"message"`                     // 返回结果详情
	NotifyType              string `json:"notify_type"`                 // 通知类型，参见constant.go
	NotifyTime              string `json:"notify_time"`                 // 通知时间 格式"2017-09-09 10:20:12"
	OrderId                 string `json:"order_id"`                    // 商户交易流水号(由商户维护唯一性)
	GlobalUniqueId          string `json:"g_unique_id"`                 // 平台交易流水号(由平台维护唯一性,同order_sn，兼容老平台用户)
	OrderSn                 string `json:"order_sn"`                    // 平台唯一交易订单号(由平台维护唯一性)
	TicketSn                string `json:"ticket_sn"`                   // 发票号码 开票失败时值为空
	TicketCode              string `json:"ticket_code"`                 // 发票代码 开票失败时值为空
	TicketDate              string `json:"ticket_date"`                 // 开票时间
	TicketStatus            uint8  `json:"ticket_status"`               // 发票状态，参见constant.go
	TicketTotalAmountHasTax string `json:"ticket_total_amount_has_tax"` // 发票含税总价 开票失败时值为空
	TicketTotalAmountNoTax  string `json:"ticket_total_amount_no_tax"`  // 发票去税总价 开票失败时值为空
	TicketTaxAmount         string `json:"ticket_tax_amount"`           // 发票税额 开票失败时值为空
	QrCode                  string `json:"qrcode"`                      // 发票二维码base64内容，最大10Kb。渠道不同可能返回为空
	CheckCode               string `json:"check_code"`                  // 发票校验码 开票失败时值为空
	CipherText              string `json:"cipher_text"`                 // 发票密文
	PdfUrl                  string `json:"pdf_url"`                     // 发票url地址 开票失败时值为空(纸票没有pdf)
}
