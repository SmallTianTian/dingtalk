package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmObjectmetaCustomerDescribeRequest() *OapiCrmObjectmetaCustomerDescribeRequest {
	return &OapiCrmObjectmetaCustomerDescribeRequest{}
}

type OapiCrmObjectmetaCustomerDescribeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectmetaCustomerDescribeResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectmetaCustomerDescribeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectmeta.customer.describe"
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectmetaCustomerDescribeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectmetaCustomerDescribeResponse struct {
	taobao.TaobaoResponse
	Result DObject `json:"result,omitempty"`
}
