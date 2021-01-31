package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartworkHrmSmsSendforpayslipRequest() *OapiSmartworkHrmSmsSendforpayslipRequest {
	return &OapiSmartworkHrmSmsSendforpayslipRequest{}
}

type OapiSmartworkHrmSmsSendforpayslipRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmSmsSendforpayslipResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmSmsSendforpayslipRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.sms.sendforpayslip"
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmSmsSendforpayslipRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PayslipSmsMessageParam struct {
	Month      int64    `json:"month,omitempty"`
	UseridList []string `json:"userid_list,omitempty"`
}
type OapiSmartworkHrmSmsSendforpayslipResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  SendSmsMessageResult `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type SendSmsMessageResult struct {
	FailedUseridList  []string `json:"failed_userid_list,omitempty"`
	InvalidUseridList []string `json:"invalid_userid_list,omitempty"`
	SuccessUseridList []string `json:"success_userid_list,omitempty"`
}
