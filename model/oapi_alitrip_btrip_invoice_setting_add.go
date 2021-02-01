package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripInvoiceSettingAddRequest() *OapiAlitripBtripInvoiceSettingAddRequest {
	return &OapiAlitripBtripInvoiceSettingAddRequest{}
}

type OapiAlitripBtripInvoiceSettingAddRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripInvoiceSettingAddResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripInvoiceSettingAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.invoice.setting.add"
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripInvoiceSettingAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenInvoiceModifyAndNewRq struct {
	Address     string `json:"address,omitempty"`
	BankName    string `json:"bank_name,omitempty"`
	BankNo      string `json:"bank_no,omitempty"`
	Corpid      string `json:"corpid,omitempty"`
	TaxNo       string `json:"tax_no,omitempty"`
	Tel         string `json:"tel,omitempty"`
	ThirdPartId string `json:"third_part_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        int64  `json:"type,omitempty"`
}
type OapiAlitripBtripInvoiceSettingAddResponse struct {
	taobao.TaobaoResponse
	Module  int64 `json:"module,omitempty"`
	Success bool  `json:"success,omitempty"`
}
