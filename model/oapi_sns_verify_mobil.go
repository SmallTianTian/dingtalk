package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSnsVerifyMobileRequest() *OapiSnsVerifyMobileRequest {
	return &OapiSnsVerifyMobileRequest{}
}

type OapiSnsVerifyMobileRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSnsVerifyMobileResponse
	Mobile          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSnsVerifyMobileRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSnsVerifyMobileRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSnsVerifyMobileRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiSnsVerifyMobileRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiSnsVerifyMobileRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sns.verify_mobile"
}
func (this *OapiSnsVerifyMobileRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSnsVerifyMobileRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSnsVerifyMobileRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSnsVerifyMobileRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSnsVerifyMobileRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["mobile"] = this.Mobile
	return txtParams
}
func (this *OapiSnsVerifyMobileRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSnsVerifyMobileRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSnsVerifyMobileRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSnsVerifyMobileResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
