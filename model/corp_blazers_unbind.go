package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpBlazersUnbindRequest() *CorpBlazersUnbindRequest {
	return &CorpBlazersUnbindRequest{}
}

type CorpBlazersUnbindRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpBlazersUnbindResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpBlazersUnbindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpBlazersUnbindRequest) GetApiMethodName() string {
	return "dingtalk.corp.blazers.unbind"
}
func (this *CorpBlazersUnbindRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpBlazersUnbindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpBlazersUnbindRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpBlazersUnbindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpBlazersUnbindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpBlazersUnbindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *CorpBlazersUnbindRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpBlazersUnbindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpBlazersUnbindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpBlazersUnbindResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
