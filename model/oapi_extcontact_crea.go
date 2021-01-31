package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiExtcontactCreateRequest() *OapiExtcontactCreateRequest {
	return &OapiExtcontactCreateRequest{}
}

type OapiExtcontactCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiExtcontactCreateResponse
	Contact         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiExtcontactCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiExtcontactCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiExtcontactCreateRequest) SetContact(contact2 string) {
	this.Contact = contact2
}
func (this *OapiExtcontactCreateRequest) GetContact() string {
	return this.Contact
}
func (this *OapiExtcontactCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.extcontact.create"
}
func (this *OapiExtcontactCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiExtcontactCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiExtcontactCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiExtcontactCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiExtcontactCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact"] = this.Contact
	return txtParams
}
func (this *OapiExtcontactCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiExtcontactCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiExtcontactCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiExtcontactCreateResponse struct {
	taobao.TaobaoResponse
	Userid string `json:"userid,omitempty"`
}
