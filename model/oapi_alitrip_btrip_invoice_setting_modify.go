package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripInvoiceSettingModifyRequest() *OapiAlitripBtripInvoiceSettingModifyRequest {
	return &OapiAlitripBtripInvoiceSettingModifyRequest{}
}

type OapiAlitripBtripInvoiceSettingModifyRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripInvoiceSettingModifyResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripInvoiceSettingModifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) GetRequest() string {
	return this.Request
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.invoice.setting.modify"
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripInvoiceSettingModifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripInvoiceSettingModifyResponse struct {
	taobao.TaobaoResponse
	Module  int64 `json:"module,omitempty"`
	Success bool  `json:"success,omitempty"`
}
