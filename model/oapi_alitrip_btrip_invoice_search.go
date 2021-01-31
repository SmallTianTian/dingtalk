package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripInvoiceSearchRequest() *OapiAlitripBtripInvoiceSearchRequest {
	return &OapiAlitripBtripInvoiceSearchRequest{}
}

type OapiAlitripBtripInvoiceSearchRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripInvoiceSearchResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripInvoiceSearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripInvoiceSearchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripInvoiceSearchRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripInvoiceSearchRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripInvoiceSearchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.invoice.search"
}
func (this *OapiAlitripBtripInvoiceSearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripInvoiceSearchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripInvoiceSearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripInvoiceSearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripInvoiceSearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripInvoiceSearchRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripInvoiceSearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripInvoiceSearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenInvoiceRq struct {
	Corpid string `json:"corpid,omitempty"`
	Title  string `json:"title,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type OapiAlitripBtripInvoiceSearchResponse struct {
	taobao.TaobaoResponse
	Invoice []OpenInvoiceDo `json:"invoice,omitempty"`
	Success bool            `json:"success,omitempty"`
}
