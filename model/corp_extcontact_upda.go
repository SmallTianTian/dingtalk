package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpExtcontactUpdateRequest() *CorpExtcontactUpdateRequest {
	return &CorpExtcontactUpdateRequest{}
}

type CorpExtcontactUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtcontactUpdateResponse
	Contact         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpExtcontactUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtcontactUpdateRequest) SetContact(contact2 string) {
	this.Contact = contact2
}
func (this *CorpExtcontactUpdateRequest) GetContact() string {
	return this.Contact
}
func (this *CorpExtcontactUpdateRequest) GetApiMethodName() string {
	return "dingtalk.corp.extcontact.update"
}
func (this *CorpExtcontactUpdateRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtcontactUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtcontactUpdateRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtcontactUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtcontactUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtcontactUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact"] = this.Contact
	return txtParams
}
func (this *CorpExtcontactUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpExtcontactUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtcontactUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpExtcontactUpdateResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
