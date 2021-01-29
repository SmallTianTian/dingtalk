package taobao

import (
	"encoding/json"
	"strconv"
)

type SetResponse interface {
	SetResponse(req TaobaoRequest, resp *HttpResponse) error
}

type SimpleResponse struct{}

func (ssr *SimpleResponse) SetResponse(req TaobaoRequest, resp *HttpResponse) error {
	if err := json.Unmarshal(resp.Body, req.GetRespInstance()); err != nil {
		return err
	}
	tr := req.GetTaobaoResp()
	tr.Code = strconv.Itoa(resp.Code)
	tr.HeaderContent = resp.Header
	tr.Body = string(resp.Body)
	return nil
}
