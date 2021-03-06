package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmObjectdataCustomobjectUpdateRequest() *OapiCrmObjectdataCustomobjectUpdateRequest {
	return &OapiCrmObjectdataCustomobjectUpdateRequest{}
}

type OapiCrmObjectdataCustomobjectUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectdataCustomobjectUpdateResponse
	Instance        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectdataCustomobjectUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) SetInstance(instance2 string) {
	this.Instance = instance2
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) GetInstance() string {
	return this.Instance
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.customobject.update"
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["instance"] = this.Instance
	return txtParams
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataCustomobjectUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataCustomobjectUpdateResponse struct {
	taobao.TaobaoResponse
	Result  ObjectDataCreateDto `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
