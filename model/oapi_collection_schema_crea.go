package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCollectionSchemaCreateRequest() *OapiCollectionSchemaCreateRequest {
	return &OapiCollectionSchemaCreateRequest{}
}

type OapiCollectionSchemaCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCollectionSchemaCreateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCollectionSchemaCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCollectionSchemaCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCollectionSchemaCreateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiCollectionSchemaCreateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiCollectionSchemaCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.collection.schema.create"
}
func (this *OapiCollectionSchemaCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCollectionSchemaCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCollectionSchemaCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCollectionSchemaCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCollectionSchemaCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiCollectionSchemaCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCollectionSchemaCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCollectionSchemaCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCollectionSchemaCreateResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
