package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripInvoiceSettingRuleRequest() *OapiAlitripBtripInvoiceSettingRuleRequest {
	return &OapiAlitripBtripInvoiceSettingRuleRequest{}
}

type OapiAlitripBtripInvoiceSettingRuleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripInvoiceSettingRuleResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripInvoiceSettingRuleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) GetRequest() string {
	return this.Request
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.invoice.setting.rule"
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripInvoiceSettingRuleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Entity struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type int64  `json:"type,omitempty"`
}
type OpenInvoiceRuleRq struct {
	AllEmploye  bool     `json:"all_employe,omitempty"`
	Corpid      string   `json:"corpid,omitempty"`
	Entities    []Entity `json:"entities,omitempty"`
	ThirdPartId string   `json:"third_part_id,omitempty"`
}
type OapiAlitripBtripInvoiceSettingRuleResponse struct {
	taobao.TaobaoResponse
	Module  OpenInvoiceRuleRS `json:"module,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type OpenInvoiceRuleRS struct {
	AddNum    int64 `json:"add_num,omitempty"`
	RemoveNum int64 `json:"remove_num,omitempty"`
}
