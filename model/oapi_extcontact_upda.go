package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiExtcontactUpdateRequest() *OapiExtcontactUpdateRequest {
	return &OapiExtcontactUpdateRequest{}
}

type OapiExtcontactUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiExtcontactUpdateResponse
	Contact         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiExtcontactUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiExtcontactUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiExtcontactUpdateRequest) SetContact(contact2 string) {
	this.Contact = contact2
}
func (this *OapiExtcontactUpdateRequest) GetContact() string {
	return this.Contact
}
func (this *OapiExtcontactUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.extcontact.update"
}
func (this *OapiExtcontactUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiExtcontactUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiExtcontactUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiExtcontactUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiExtcontactUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact"] = this.Contact
	return txtParams
}
func (this *OapiExtcontactUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiExtcontactUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiExtcontactUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiExtcontactUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
