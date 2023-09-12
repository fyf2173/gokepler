package jcq

import (
	"encoding/json"
	"fmt"

	"github.com/fyf2173/ysdk-go/xhttp"
)

// Request 消费/确认
func (jc *Client) Request(method string, path string, params interface{}, out interface{}) error {
	jh := jc.NewHeader()
	jh.Signature = jc.GetSignature(jh.GetSignSource(params))

	opts := []xhttp.Option{
		xhttp.SetRequestHeader("accessKey", jh.AccessKey),
		xhttp.SetRequestHeader("dateTime", jh.DateTime),
		xhttp.SetRequestHeader("signature", jh.Signature),
	}
	var resp CommonConsumerResp
	if err := xhttp.Request(method, fmt.Sprintf("%s%s", endPoint, path), params, &resp, opts...); err != nil {
		return err
	}
	if resp.Error != nil && resp.Error.Code != 0 {
		return fmt.Errorf("code=%d,message=%s,status=%s", resp.Error.Code, resp.Error.Message, resp.Error.Status)
	}
	if out == nil {
		return nil
	}
	return json.Unmarshal(resp.Result, &out)
}
