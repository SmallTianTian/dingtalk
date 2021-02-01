package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmObjectdataContactUpdateRequest() *OapiCrmObjectdataContactUpdateRequest {
	return &OapiCrmObjectdataContactUpdateRequest{}
}

type OapiCrmObjectdataContactUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectdataContactUpdateResponse
	Instance        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectdataContactUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataContactUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataContactUpdateRequest) SetInstance(instance2 string) {
	this.Instance = instance2
}
func (this *OapiCrmObjectdataContactUpdateRequest) GetInstance() string {
	return this.Instance
}
func (this *OapiCrmObjectdataContactUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.contact.update"
}
func (this *OapiCrmObjectdataContactUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataContactUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataContactUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataContactUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataContactUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["instance"] = this.Instance
	return txtParams
}
func (this *OapiCrmObjectdataContactUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmObjectdataContactUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataContactUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataContactUpdateResponse struct {
	taobao.TaobaoResponse
	Result  ObjectDataCreateDto `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
