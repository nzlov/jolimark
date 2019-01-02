package jolimark

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetRequest(url string) (Resp, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Resp{}, err
	}
	defer resp.Body.Close()

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Resp{}, err
	}

	return newResp(respbody)

}
func PostRequest(url string, bodyType string, body io.Reader) (Resp, error) {
	resp, err := http.Post(url, bodyType, body)
	if err != nil {
		return Resp{}, err
	}
	defer resp.Body.Close()

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Resp{}, err
	}
	return newResp(respbody)
}

//func PostJSONRequest(url string, data interface{}) (Resp, error) {
//	bytesData, err := json.Marshal(data)
//	if err != nil {
//		return Resp{}, err
//	}
//	return PostRequest(url, "application/json;charset=UTF-8", bytes.NewReader(bytesData))
//}

func PostFormRequest(urls string, data map[string]interface{}) (Resp, error) {
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
	return newResp(respbody)
}
