package jolimark // import "github.com/nzlov/jolimark"
import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var HOST = "https://mcp.jolimark.com/mcp/v2"

type Client struct {
	appid       string
	key         string
	callbackkey string
	log         *logrus.Logger
}

func NewClient(appid, key, callbackkey string) *Client {
	return &Client{
		appid:       appid,
		key:         key,
		callbackkey: callbackkey,
	}
}

func (c *Client) SetLog(l *logrus.Logger) {
	c.log = l
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	// 有效时间，单位为秒
	Expire float64 `json:"expires_in"`
	// 生成时间
	Create string `json:"create_time"`
}

func MD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}
func accesstokenSign(appid, key string, timestamp int64) string {
	return MD5(fmt.Sprintf("app_id=%s&sign_type=MD5&time_stamp=%d&key=%s", appid, timestamp, key))
}

// AccessToken 获取AccessToken
func (c *Client) AccessToken() (token AccessToken, err error) {
	t := time.Now().Unix()
	resp, err := c.getRequest(fmt.Sprintf("%s/sys/GetAccessToken?app_id=%s&time_stamp=%d&sign=%s&sign_type=MD5", HOST, c.appid, t, accesstokenSign(c.appid, c.key, t)))
	if err != nil {
		return token, err
	}
	err = resp.Decode(&token)
	return
}
