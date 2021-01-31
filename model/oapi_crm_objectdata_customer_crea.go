package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmObjectdataCustomerCreateRequest() *OapiCrmObjectdataCustomerCreateRequest {
	return &OapiCrmObjectdataCustomerCreateRequest{}
}

type OapiCrmObjectdataCustomerCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectdataCustomerCreateResponse
	Instance        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectdataCustomerCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataCustomerCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataCustomerCreateRequest) SetInstance(instance2 string) {
	this.Instance = instance2
}
func (this *OapiCrmObjectdataCustomerCreateRequest) GetInstance() string {
	return this.Instance
}
func (this *OapiCrmObjectdataCustomerCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.customer.create"
}
func (this *OapiCrmObjectdataCustomerCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataCustomerCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataCustomerCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataCustomerCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataCustomerCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["instance"] = this.Instance
	return txtParams
}
func (this *OapiCrmObjectdataCustomerCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmObjectdataCustomerCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataCustomerCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataCustomerCreateResponse struct {
	taobao.TaobaoResponse
	Result  ObjectDataCreateDto `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
