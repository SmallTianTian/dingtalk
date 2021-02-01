package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewIsvCallGetuserlistRequest() *IsvCallGetuserlistRequest {
	return &IsvCallGetuserlistRequest{}
}

type IsvCallGetuserlistRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            IsvCallGetuserlistResponse
	Offset          int64
	Start           int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *IsvCallGetuserlistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *IsvCallGetuserlistRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *IsvCallGetuserlistRequest) GetOffset() int64 {
	return this.Offset
}
func (this *IsvCallGetuserlistRequest) SetStart(start2 int64) {
	this.Start = start2
}
func (this *IsvCallGetuserlistRequest) GetStart() int64 {
	return this.Start
}
func (this *IsvCallGetuserlistRequest) GetApiMethodName() string {
	return "dingtalk.isv.call.getuserlist"
}
func (this *IsvCallGetuserlistRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *IsvCallGetuserlistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *IsvCallGetuserlistRequest) GetTopApiCallType() string {
	return "top"
}
func (this *IsvCallGetuserlistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *IsvCallGetuserlistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *IsvCallGetuserlistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["start"] = this.Start
	return txtParams
}
func (this *IsvCallGetuserlistRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *IsvCallGetuserlistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *IsvCallGetuserlistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type IsvCallGetuserlistResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
