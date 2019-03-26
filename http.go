package jolimark

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Client) getRequest(url string) (Resp, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Resp{}, err
	}
	defer resp.Body.Close()

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Resp{}, err
	}

	if c.log != nil {
		c.log.Debugln("[Jolimark] Get:", url, "Resp:", string(respbody))
	}

	return newResp(respbody)

}

func (c *Client) postFormRequest(urls string, data map[string]interface{}) (Resp, error) {
	vs := url.Values{}
	for k, v := range data {
		vs[k] = []string{fmt.Sprint(v)}
	}
	resp, err := http.PostForm(urls, vs)
	if err != nil {
		return Resp{}, err
	}
	defer resp.Body.Close()

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Resp{}, err
	}
	if c.log != nil {
		c.log.Debugln("[Jolimark] PostForm:", urls, "Form:", data, "Resp:", string(respbody))
	}
	return newResp(respbody)
}
