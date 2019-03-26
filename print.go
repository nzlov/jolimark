package jolimark

import "github.com/json-iterator/go"

// PrintHtmlUrl 打印简单HTML页面-传URL地址。
// 适用于票据页面简单的餐饮、超市等小票打印。
// 客户提供一个html页面的url地址，由打印机直接下载页面解析打印，此页面内的标签必须符合映美定义的html标签规范。标签规范详情请详见《HTML模板设置说明文档》。
func (c *Client) PrintHtmlUrl(accesstoken, device_ids, cus_orderid, bill_content string, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintHtmlUrl", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"bill_content": bill_content,
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(bill_content),
	})
}

// PrintHtmlCode 打印简单html页面-传HTML代码。
// 在一些特定应用中，客户无法生成html页面url，可以传html源代码，由映美服务器代为生成打印。
// 打印应用提供简单标签的html源代码，由映美服务器转为html文件，打印机再从映美服务器直接下载解析打印，此打印方式跟简单标签的html页面一致，标签必须符合映美定义的html标签规范，标签规范详情请 详见《HTML模板设置说明文档》。
func (c *Client) PrintHtmlCode(accesstoken, device_ids, cus_orderid, bill_content string, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintHtmlCode", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"bill_content": bill_content,
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(bill_content),
	})
}

// PrintRichHtmlCode 打印复杂HTML页面-转图片。
// html页面有复杂表格和复杂排版样式的情况，以黑白双色打印。
// 打印客户提供的html页面，支持复杂的html标签，样式只能支持到IE11内核，可以打印表格数据以及复杂排版页面，html页面的body宽度不大于打印机的纸张宽度值，Body宽度的像素最大值等于纸张宽度*打印机dpi/25.4，打印机的dpi值，详见打印机默认纸张宽度和DPI值。
// ps：页面包含表格的情况下，表格线条需要设置成黑色。
func (c *Client) PrintRichHtmlCode(accesstoken, device_ids, cus_orderid, bill_content string, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintRichHtmlCode", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"bill_content": bill_content,
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(bill_content),
	})
}

// PrintHtmlToPic 打印复杂html页面-传url转图片。
// 适用于html页面含有复杂排版样式与表格的情况，以黑白双色打印。
// html可以包含复杂的标签，用于呈现打印页面上的表格数据以及复杂排版，最高支持IE11内核。html页面的body宽度不大于打印机的纸张宽度值，body宽度的像素最大值等于纸张宽度*打印机dpi/25.4，打印机的dpi值详见附录1.3。
// 注：在页面包含表格的情况下，表格线条需要设置为黑色。将要打印的内容宽度设置成打印机可以支持的最大宽度，打印效果会比较好。
func (c *Client) PrintHtmlToPic(accesstoken, device_ids, cus_orderid, bill_content string, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintHtmlToPic", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"bill_content": bill_content,
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(bill_content),
	})
}

// PrintHtmlToGrayPic 打印简单html页面-转灰度图。
// 此接口适用于页面中含有图片需要打印的情况，以多层级灰度打印，仅适用于热敏打印机。
// 打印客户提供的html页面，支持复杂的html标签，样式只能支持到IE11内核，可以打印灰度图片数据，html页面的body宽度不大于打印机的纸张宽度值，Body宽度的像素最大值等于纸张宽度*打印机dpi/25.4，打印机的dpi值，详见打印机默认纸张宽度和DPI值。
func (c *Client) PrintHtmlToGrayPic(accesstoken, device_ids, cus_orderid, bill_content string, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintHtmlToGrayPic", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"bill_content": bill_content,
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(bill_content),
	})
}

// PrintEsc 打印ESC指令
// 适用于开发者懂Esc指令排版布局，详见《打印机编程手册esc指令》
// 将ESC标准指令代码以文本格式传输到接口进行打印
func (c *Client) PrintEsc(accesstoken, device_ids, cus_orderid, bill_content string, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintEsc", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"bill_content": bill_content,
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(bill_content),
	})
}

type Point struct {
	// 文本相对于left位置的横坐标
	X float64 `json:"x"`
	// 文本相对于top位置的纵坐标
	Y float64 `json:"y"`
	// 打印的文本内容
	V string `json:"v"`
	// 文本说明，此属性可以不填写，为了方便客户检查数据用
	Dsp string `json:"dsp"`
}

type PointText struct {
	High    float64 `json:"high"`
	Left    float64 `json:"left"`
	Top     float64 `json:"top"`
	Content []Point `json:"content"`
}

// PrintPointText 打印定点坐标文本
// 此接口应用于客户需要自己设定文本打印位置的票据打印，比如增值税发票，门票等票据。
// 将文本打印到客户指定的坐标位置。
func (c *Client) PrintPointText(accesstoken, device_ids, cus_orderid string, bill_content PointText, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	content, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(bill_content)
	if err != nil {
		return resp, err
	}
	return c.postFormRequest(HOST+"/sys/PrintPointText", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"bill_content": string(content),
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(string(content)),
	})
}

// PrintExpress 打印快递面单
// 直接对接快递面单打印
// 打印标准的快递面单，如顺丰，圆通，EMS等
func (c *Client) PrintExpress(accesstoken, device_ids, cus_orderid, template_id, jj_dwmc, jj_jjr, jj_lxdh, jj_dz, sj_dwmc, sj_sjr, sj_lxdh, sj_dz, wp_jtw, wp_smjz string, copies, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintExpress", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"template_id":  template_id,
		"jj_dwmc":      jj_dwmc,
		"jj_jjr":       jj_jjr,
		"jj_lxdh":      jj_lxdh,
		"jj_dz":        jj_dz,
		"sj_dwmc":      sj_dwmc,
		"sj_sjr":       sj_sjr,
		"sj_lxdh":      sj_lxdh,
		"sj_dz":        sj_dz,
		"wp_jtw":       wp_jtw,
		"wp_smjz":      wp_smjz,
		"time_out":     time_out,
	})
}

// PrintHtmlTemplate 打印云模板
// 用户创建打印模板
// 打印用户在映美云模板上的排版设计页面。映美云模板地址： http://form.jolimark.com
func (c *Client) PrintHtmlTemplate(accesstoken, device_ids, cus_orderid, template_id, bill_content string, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintHtmlTemplate", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"template_id":  template_id,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"bill_content": bill_content,
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(bill_content),
	})
}

// PrintFile 打印本地文档, bill_content base64格式的文件二进制
// 适用于调用方不方便提供文档的url址址，如C/S程序对接。适用机型为CFP-535。
// 将pdf、图片(jpg、png )转换成base64格式的数据后，调用此接口，打印机便可按要求进行打印。其中pdf数据大小不超过过4M，图片数据大小不超过2M。
func (c *Client) PrintFile(accesstoken, device_ids, cus_orderid, file_type, bill_content string, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintFile", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"file_type":    file_type,
		"bill_content": bill_content,
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(bill_content),
	})
}

// PrintFileByUrl 打印远程文档
// 适用于调用方不方便提供文档的url址址，如C/S程序对接。适用机型为CFP-535。
// 调用时传pdf、图片(jpg、png ) 对应的 url地址，打印机便可进行打印。
func (c *Client) PrintFileByUrl(accesstoken, device_ids, cus_orderid, file_type, bill_content string, copies, paper_width, paper_height, paper_type, time_out int) (resp Resp, err error) {
	return c.postFormRequest(HOST+"/sys/PrintFileByUrl", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_ids":   device_ids,
		"copies":       copies,
		"cus_orderid":  cus_orderid,
		"file_type":    file_type,
		"bill_content": bill_content,
		"paper_width":  paper_width,
		"paper_height": paper_height,
		"paper_type":   paper_type,
		"time_out":     time_out,
		//"sign_type":    "MD5",
		//"sign":         MD5(bill_content),
	})
}
