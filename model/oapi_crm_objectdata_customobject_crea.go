package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmObjectdataCustomobjectCreateRequest() *OapiCrmObjectdataCustomobjectCreateRequest {
	return &OapiCrmObjectdataCustomobjectCreateRequest{}
}

type OapiCrmObjectdataCustomobjectCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectdataCustomobjectCreateResponse
	Instance        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectdataCustomobjectCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) SetInstance(instance2 string) {
	this.Instance = instance2
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) GetInstance() string {
	return this.Instance
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.customobject.create"
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["instance"] = this.Instance
	return txtParams
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataCustomobjectCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataCustomobjectCreateResponse struct {
	taobao.TaobaoResponse
	Result  ObjectDataCreateDto `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
