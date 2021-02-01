package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScenegroupTemplateCloseRequest() *OapiImChatScenegroupTemplateCloseRequest {
	return &OapiImChatScenegroupTemplateCloseRequest{}
}

type OapiImChatScenegroupTemplateCloseRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImChatScenegroupTemplateCloseResponse
	ApplyMode          int64
	OpenConversationId string
	OwnerUserId        string
	TemplateId         string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiImChatScenegroupTemplateCloseRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScenegroupTemplateCloseRequest) SetApplyMode(applyMode2 int64) {
	this.ApplyMode = applyMode2
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetApplyMode() int64 {
	return this.ApplyMode
}
func (this *OapiImChatScenegroupTemplateCloseRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImChatScenegroupTemplateCloseRequest) SetOwnerUserId(ownerUserId2 string) {
	this.OwnerUserId = ownerUserId2
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetOwnerUserId() string {
	return this.OwnerUserId
}
func (this *OapiImChatScenegroupTemplateCloseRequest) SetTemplateId(templateId2 string) {
	this.TemplateId = templateId2
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetTemplateId() string {
	return this.TemplateId
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scenegroup.template.close"
}
func (this *OapiImChatScenegroupTemplateCloseRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScenegroupTemplateCloseRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScenegroupTemplateCloseRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["apply_mode"] = this.ApplyMode
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["owner_user_id"] = this.OwnerUserId
	txtParams["template_id"] = this.TemplateId
	return txtParams
}
func (this *OapiImChatScenegroupTemplateCloseRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationId, "openConversationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScenegroupTemplateCloseRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScenegroupTemplateCloseResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
