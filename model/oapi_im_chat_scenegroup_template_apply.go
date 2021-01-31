package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScenegroupTemplateApplyRequest() *OapiImChatScenegroupTemplateApplyRequest {
	return &OapiImChatScenegroupTemplateApplyRequest{}
}

type OapiImChatScenegroupTemplateApplyRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImChatScenegroupTemplateApplyResponse
	ApplyMode          int64
	OpenConversationId string
	OwnerUserId        string
	TemplateId         string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiImChatScenegroupTemplateApplyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScenegroupTemplateApplyRequest) SetApplyMode(applyMode2 int64) {
	this.ApplyMode = applyMode2
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetApplyMode() int64 {
	return this.ApplyMode
}
func (this *OapiImChatScenegroupTemplateApplyRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImChatScenegroupTemplateApplyRequest) SetOwnerUserId(ownerUserId2 string) {
	this.OwnerUserId = ownerUserId2
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetOwnerUserId() string {
	return this.OwnerUserId
}
func (this *OapiImChatScenegroupTemplateApplyRequest) SetTemplateId(templateId2 string) {
	this.TemplateId = templateId2
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetTemplateId() string {
	return this.TemplateId
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scenegroup.template.apply"
}
func (this *OapiImChatScenegroupTemplateApplyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScenegroupTemplateApplyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScenegroupTemplateApplyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["apply_mode"] = this.ApplyMode
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["owner_user_id"] = this.OwnerUserId
	txtParams["template_id"] = this.TemplateId
	return txtParams
}
func (this *OapiImChatScenegroupTemplateApplyRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationId, "openConversationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScenegroupTemplateApplyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScenegroupTemplateApplyResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
