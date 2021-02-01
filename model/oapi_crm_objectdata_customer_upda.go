package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmObjectdataCustomerUpdateRequest() *OapiCrmObjectdataCustomerUpdateRequest {
	return &OapiCrmObjectdataCustomerUpdateRequest{}
}

type OapiCrmObjectdataCustomerUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectdataCustomerUpdateResponse
	Instance        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectdataCustomerUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) SetInstance(instance2 string) {
	this.Instance = instance2
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) GetInstance() string {
	return this.Instance
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.customer.update"
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["instance"] = this.Instance
	return txtParams
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataCustomerUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataCustomerUpdateResponse struct {
	taobao.TaobaoResponse
	Result  ObjectDataCreateDto `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
