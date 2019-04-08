package jolimark

import "fmt"

type PrintTaskStatus struct{
	// 客户系统的订单号
	CusOrderID string `json:"cus_orderid"`
	// 映美云打印平台的任务id
	OrderID string `json:"order_id"`
	// 任务状态码；状态码解释：1-打印成功，2-打印失败，打印机下载数据失败，3-超时未打印，4-打印失败,解绑清理；5-保留值，6-打印失败,打印过程开盖，7-打印失败,打印过程缺纸，8-打印失败，推送时打印机掉线，9-打印失败，推送时打印机开盖，10-打印失败，推送时打印机缺纸，11-打印失败，推送时打印机未注册，12-打印失败，推送时未找到打印机可用状态，13-用户取消的任务
	OrderStatus string `json:"order_status"`
}

// QueryPrintTaskStatus 根据客户订单号来查询打印任务状态的接口。
func (c *Client) QueryPrintTaskStatus(accesstoken, cus_orderid string) (p PrintTaskStatus, err error) {
	resp, err := c.getRequest(fmt.Sprintf("%s/sys/QueryPrintTaskStatus?app_id=%s&access_token=%s&cus_orderid=%s", HOST, c.appid, accesstoken, cus_orderid))
	if err != nil {
		return p, err
	}

	if resp.Code != 0 {
		return p, Error{Code: resp.Code}
	}

	err = resp.Decode(&p)
	return
}

type NotPrintTaskStatus struct{
	// 未打印的任务数
	Count int `json:"count"`
	// 未打印的任务集合
	Result interface{} `json:"result"`
	// 客户系统的订单号
	CusOrderID string `json:"cus_orderid"`
	// 映美云打印平台的任务id
	OrderID string `json:"order_id"`
	// 打印的任务内容
	OrderContent string `json:"order_content"`
	// 任务创建的时间
	OrderDate string `json:"order_date"`
}

// QueryNotPrintTask 查询指定打印机未打印任务。
func (c *Client) QueryNotPrintTask(accesstoken, device_id string) (p PrintTaskStatus, err error) {
	resp, err := c.getRequest(fmt.Sprintf("%s/sys/QueryNotPrintTask?app_id=%s&access_token=%s&device_id=%s", HOST, c.appid, accesstoken, device_id))
	if err != nil {
		return p, err
	}

	if resp.Code != 0 {
		return p, Error{Code: resp.Code}
	}

	err = resp.Decode(&p)
	return
}

// CancelNotPrintTask 取消指定打印机当前未完成的打印任务。
func (c *Client) CancelNotPrintTask(accesstoken, device_id string) (p NotPrintTaskStatus, err error) {
	resp,err:= c.postFormRequest(HOST+"/sys/CancelNotPrintTask", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_id":    device_id,
	})
	if err!=nil{
		return p ,err
	}

	if resp.Code != 0 {
		return p, Error{Code: resp.Code}
	}

	err = resp.Decode(&p)
	return
}