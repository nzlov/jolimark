package jolimark

import (
	"fmt"
	"net/http"
)

type TaskResult struct {
	// 客户系统订单流水号
	BillNo string `json:"bill_no"`
	// 打印机编码
	PrinterCode string `json:"printer_code"`
	// 结果状态码，10000-打印失败，打印机下载数据失败，10001-打印成功，10002-打印失败,打印过程开盖，10003-打印失败,打印过程缺纸，10004-通讯超时，10005-打印失败，推送时打印机掉线，10006-保留值，10007-保留值，10008-超时未打印，10009-打印失败,解绑清理，10010-打印失败，推送时打印机开盖，10011-打印失败，推送时打印机缺纸，10012-打印失败，推送时打印机未注册，10013-打印失败，推送时未找到打印机可用状态
	ResultCode string `json:"result_code"`
	// 时间戳
	Timestamp string `json:"time_stamp"`
	// 随机字符串
	NonceStr string `json:"nonce_str"`
	// 签名
	Sign string `json:"sign"`
}

func (t TaskResult) check(key string) bool {
	return t.Sign == MD5(fmt.Sprintf("bill_no=%s&nonce_str=%s&printer_code=%s&result_code=%s&sign_type=MD5&time_stamp=%s&key=%s", t.BillNo, t.NonceStr, t.PrinterCode, t.ResultCode, t.Timestamp, key))
}

//TaskCallbackVerify
func (c *Client) TaskCallbackVerify(result TaskResult) bool {
	return result.check(c.callbackkey)
}

// TaskCallback 打印任务结果回调更新
func (c *Client) TaskCallback(callback func(TaskResult) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(fmt.Sprintf(`{"status":0,"data":"%s"}`, err.Error())))
			return
		}

		result := TaskResult{
			BillNo:      r.FormValue("bill_no"),
			PrinterCode: r.FormValue("printer_code"),
			ResultCode:  r.FormValue("result_code"),
			Timestamp:   r.FormValue("time_stamp"),
			NonceStr:    r.FormValue("nonce_str"),
			Sign:        r.FormValue("sign"),
		}

		if !result.check(c.callbackkey) {
			w.Write([]byte(`{"status":0,"data":"sign"}`))
			return
		}
		err = callback(result)
		if err != nil {
			w.Write([]byte(fmt.Sprintf(`{"status":0,"data":"%s"}`, err.Error())))
			return
		}
		w.Write([]byte(`{"status":1,"data":""}`))
	}
}

type DeviceStatus struct {
	// 状态改变时间，时间格式形如：yyyy-MM-dd HH:mm:ss格式
	FaultTime string `json:"fault_time"`
	// 打印机编码
	PrinterCode string `json:"printer_code"`
	// 状态码，1-正常， 2-缺纸，3-开盖，4-网络故障，99-未知异常
	ResultCode string `json:"result_code"`
	// 时间戳
	Timestamp string `json:"time_stamp"`
	// 随机字符串
	NonceStr string `json:"nonce_str"`
	// 签名
	Sign string `json:"sign"`
}

func (t DeviceStatus) check(key string) bool {
	return t.Sign == MD5(fmt.Sprintf("fault_time=%s&nonce_str=%s&printer_code=%s&result_code=%s&sign_type=MD5&time_stamp=%s&key=%s", t.FaultTime, t.NonceStr, t.PrinterCode, t.ResultCode, t.Timestamp, key))
}

//DeviceStatusCallbackVerify
func (c *Client) DeviceStatusCallbackVerify(device DeviceStatus) bool {
	return device.check(c.callbackkey)
}

// DeviceStatusCallback 打印设备状态回调更新
func (c *Client) DeviceStatusCallback(callback func(DeviceStatus) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(fmt.Sprintf(`{"status":0,"data":"%s"}`, err.Error())))
			return
		}

		result := DeviceStatus{
			FaultTime:   r.FormValue("fault_time"),
			PrinterCode: r.FormValue("printer_code"),
			ResultCode:  r.FormValue("result_code"),
			Timestamp:   r.FormValue("time_stamp"),
			NonceStr:    r.FormValue("nonce_str"),
			Sign:        r.FormValue("sign"),
		}

		if !result.check(c.callbackkey) {
			w.Write([]byte(`{"status":0,"data":"sign"}`))
			return
		}
		err = callback(result)
		if err != nil {
			w.Write([]byte(fmt.Sprintf(`{"status":0,"data":"%s"}`, err.Error())))
			return
		}
		w.Write([]byte(`{"status":1,"data":""}`))
	}
}
