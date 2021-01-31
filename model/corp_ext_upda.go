package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpExtUpdateRequest() *CorpExtUpdateRequest {
	return &CorpExtUpdateRequest{}
}

type CorpExtUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtUpdateResponse
	Contact         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpExtUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtUpdateRequest) SetContact(contact2 string) {
	this.Contact = contact2
}
func (this *CorpExtUpdateRequest) GetContact() string {
	return this.Contact
}
func (this *CorpExtUpdateRequest) GetApiMethodName() string {
	return "dingtalk.corp.ext.update"
}
func (this *CorpExtUpdateRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtUpdateRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact"] = this.Contact
	return txtParams
}
func (this *CorpExtUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpExtUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpExtUpdateResponse struct {
	taobao.TaobaoResponse
}
