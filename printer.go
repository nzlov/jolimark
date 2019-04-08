package jolimark

import (
	"fmt"
)

// BindPrinter 将打印机绑定到当前的应用id。
func (c *Client) BindPrinter(accesstoken, device_codes string) (p Resp, err error) {
	return c.postFormRequest(HOST+"/sys/BindPrinter", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_codes": device_codes,
	})
}

// BindPrinter 从当前应用id中解除绑定的打印机。
func (c *Client) UnBindPrinter(accesstoken, device_id string) (p Resp, err error) {
	return c.postFormRequest(HOST+"/sys/UnBindPrinter", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_id":    device_id,
	})
}

// CheckPrinterEnableBind 校验打印机是否可以绑定。
func (c *Client) CheckPrinterEnableBind(accesstoken, device_id string) (p Resp, err error) {
	return c.postFormRequest(HOST+"/sys/CheckPrinterEnableBind", map[string]interface{}{
		"app_id":       c.appid,
		"access_token": accesstoken,
		"device_id":    device_id,
	})
}

type Printer struct {
	// 打印机编码
	DeviceID string `json:"device_id"`
	// 打印机型号id；1：mcp-58，2：mcp-350，3：mcp-360，4：fp-570，5：mcp-610，6：mcp-230，7：mcp-330，8：clp-180，9：cfp-535；
	DeviceTypeID int `json:"devicetype _id"`
	// 打印机型号名称
	DeviceTypeName string `json:"devicetype _name"`
	// 打印机自身当前软件版本
	VarVersion string `json:"var_version"`
	// 打印机自身当前固件版本
	FirmwareVersion string `json:"firmware_version"`
	// 此型号打印机最新的软件版本
	LastVarVersion string `json:"irmware_version"`
	// 此型号打印机最新的固件版本
	LastFirmwareVersion string `json:"last_firmwareversion"`
}

// QueryPrinterInfo 查询打印机的基础信息，包括当前的软件、固件版本以及打印机型号。
func (c *Client) QueryPrinterInfo(accesstoken, device_id string) (p Printer, err error) {
	resp, err := c.getRequest(fmt.Sprintf("%s/sys/QueryPrinterInfo?app_id=%s&access_token=%s&device_id=%s", HOST, c.appid, accesstoken, device_id))
	if err != nil {
		return p, err
	}

	if resp.Code != 0 {
		return p, Error{Code: resp.Code}
	}

	err = resp.Decode(&p)
	return
}

type PrinterStatus struct {
	// 打印机编码
	DeviceID string `json:"device_id"`
	// 状态编码；0-不存在，1-正常在线，2-缺纸，3-故障或开盖，4-离线，5-选择纸张错误，6-不属于此应用id，7-未绑定，8-已绑定到其他商户，99-其他异常
	StatusCode int `json:"status_code"`
	// 状态信息；0-不存在，1-正常在线，2-缺纸，3-故障或开盖，4-离线，5-选择纸张错误，6-不属于此应用id，7-未绑定，8-已绑定到其他商户，99-其他异常
	StatusMsg string `json:"status_msg"`
}

// QueryPrinterStatus 查询打印机的当前状态，包括打印机是否在线，是否缺纸，是否开盖等。
func (c *Client) QueryPrinterStatus(accesstoken, device_id string) (p PrinterStatus, err error) {
	resp, err := c.getRequest(fmt.Sprintf("%s/sys/QueryPrinterStatus?app_id=%s&access_token=%s&device_id=%s", HOST, c.appid, accesstoken, device_id))
	if err != nil {
		return p, err
	}

	if resp.Code != 0 {
		return p, Error{Code: resp.Code}
	}

	err = resp.Decode(&p)
	return
}
