package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceStatusGetRequest() *OapiCustomerserviceStatusGetRequest {
	return &OapiCustomerserviceStatusGetRequest{}
}

type OapiCustomerserviceStatusGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomerserviceStatusGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCustomerserviceStatusGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceStatusGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceStatusGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.status.get"
}
func (this *OapiCustomerserviceStatusGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceStatusGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceStatusGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceStatusGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceStatusGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiCustomerserviceStatusGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceStatusGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceStatusGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCustomerserviceStatusGetResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
