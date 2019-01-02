package jolimark

import (
	"encoding/json"
	"github.com/json-iterator/go"
)

type PrinterState struct {
	// 打印机编码
	PrintCode string `json:"print_code"`
	// 状态编码；0-不存在，1-正常在线，2-缺纸，3-故障或开盖，4-离线，5-选择纸张错误，6-不属于此应用id，7-未绑定，8-已绑定到其他商户，99-其他异常
	StatusCode int `json:"status_code"`
	// 状态信息；0-不存在，1-正常在线，2-缺纸，3-故障或开盖，4-离线，5-选择纸张错误，6-不属于此应用id，7-未绑定，8-已绑定到其他商户，99-其他异常
	StatusMsg string `json:"status_msg"`
}

type Resp struct {
	Code         int             `json:"return_code"`
	Data         json.RawMessage `json:"return_data"`
	Msg          string          `json:"return_msg"`
	PrinterState []PrinterState  `json:"printer_state"`
}

func newResp(data []byte) (Resp, error) {
	r := Resp{}
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, &r)
	return r, err
}

func (r Resp) Decode(obj interface{}) error {
	if r.Code != 0 {
		return Error{r.Code}
	}
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(r.Data, obj)
}
