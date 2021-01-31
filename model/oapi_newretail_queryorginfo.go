package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiNewretailQueryorginfoRequest() *OapiNewretailQueryorginfoRequest {
	return &OapiNewretailQueryorginfoRequest{}
}

type OapiNewretailQueryorginfoRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiNewretailQueryorginfoResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiNewretailQueryorginfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiNewretailQueryorginfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiNewretailQueryorginfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.newretail.queryorginfo"
}
func (this *OapiNewretailQueryorginfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiNewretailQueryorginfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiNewretailQueryorginfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiNewretailQueryorginfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiNewretailQueryorginfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiNewretailQueryorginfoRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiNewretailQueryorginfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiNewretailQueryorginfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiNewretailQueryorginfoResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Org    `json:"result,omitempty"`
}
type Org struct {
	Licensemediaid          string `json:"licensemediaid,omitempty"`
	Orgname                 string `json:"orgname,omitempty"`
	Registnum               string `json:"registnum,omitempty"`
	Unifiedsocialcredit     string `json:"unifiedsocialcredit,omitempty"`
	Unnameorganizationcoded string `json:"unnameorganizationcoded,omitempty"`
}
