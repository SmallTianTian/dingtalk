package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpInvoiceGettitleRequest() *CorpInvoiceGettitleRequest {
	return &CorpInvoiceGettitleRequest{}
}

type CorpInvoiceGettitleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpInvoiceGettitleResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpInvoiceGettitleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpInvoiceGettitleRequest) GetApiMethodName() string {
	return "dingtalk.corp.invoice.gettitle"
}
func (this *CorpInvoiceGettitleRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpInvoiceGettitleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpInvoiceGettitleRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpInvoiceGettitleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpInvoiceGettitleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpInvoiceGettitleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *CorpInvoiceGettitleRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpInvoiceGettitleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpInvoiceGettitleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpInvoiceGettitleResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type InvoiceVo struct {
	Account  string `json:"account,omitempty"`
	Address  string `json:"address,omitempty"`
	BankName string `json:"bank_name,omitempty"`
	Contact  string `json:"contact,omitempty"`
	DutyPara string `json:"duty_para,omitempty"`
	Remark   string `json:"remark,omitempty"`
	Title    string `json:"title,omitempty"`
}
