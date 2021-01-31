package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpExtcontactCreateRequest() *CorpExtcontactCreateRequest {
	return &CorpExtcontactCreateRequest{}
}

type CorpExtcontactCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtcontactCreateResponse
	Contact         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpExtcontactCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtcontactCreateRequest) SetContact(contact2 string) {
	this.Contact = contact2
}
func (this *CorpExtcontactCreateRequest) GetContact() string {
	return this.Contact
}
func (this *CorpExtcontactCreateRequest) GetApiMethodName() string {
	return "dingtalk.corp.extcontact.create"
}
func (this *CorpExtcontactCreateRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtcontactCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtcontactCreateRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtcontactCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtcontactCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtcontactCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact"] = this.Contact
	return txtParams
}
func (this *CorpExtcontactCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpExtcontactCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtcontactCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpExtcontactCreateResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
