package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripInvoiceSettingDeleteRequest() *OapiAlitripBtripInvoiceSettingDeleteRequest {
	return &OapiAlitripBtripInvoiceSettingDeleteRequest{}
}

type OapiAlitripBtripInvoiceSettingDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripInvoiceSettingDeleteResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) GetRequest() string {
	return this.Request
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.invoice.setting.delete"
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripInvoiceSettingDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenInvoiceDeleteRq struct {
	Corpid      string `json:"corpid,omitempty"`
	ThirdPartId string `json:"third_part_id,omitempty"`
}
type OapiAlitripBtripInvoiceSettingDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Module  bool   `json:"module,omitempty"`
	Success bool   `json:"success,omitempty"`
}
