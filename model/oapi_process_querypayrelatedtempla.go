package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessQuerypayrelatedtemplateRequest() *OapiProcessQuerypayrelatedtemplateRequest {
	return &OapiProcessQuerypayrelatedtemplateRequest{}
}

type OapiProcessQuerypayrelatedtemplateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessQuerypayrelatedtemplateResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessQuerypayrelatedtemplateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.querypayrelatedtemplate"
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessQuerypayrelatedtemplateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessQuerypayrelatedtemplateResponse struct {
	taobao.TaobaoResponse
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
