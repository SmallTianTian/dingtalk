package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripMonthbillUrlGetRequest() *OapiAlitripBtripMonthbillUrlGetRequest {
	return &OapiAlitripBtripMonthbillUrlGetRequest{}
}

type OapiAlitripBtripMonthbillUrlGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripMonthbillUrlGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripMonthbillUrlGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.monthbill.url.get"
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripMonthbillUrlGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenAccountRq struct {
	BillMonth string `json:"bill_month,omitempty"`
	Corpid    string `json:"corpid,omitempty"`
}
type OapiAlitripBtripMonthbillUrlGetResponse struct {
	taobao.TaobaoResponse
	Module  []OpenAccountRs `json:"module,omitempty"`
	Success bool            `json:"success,omitempty"`
}
type OpenAccountRs struct {
	EndDate   string `json:"end_date,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	Url       string `json:"url,omitempty"`
}
