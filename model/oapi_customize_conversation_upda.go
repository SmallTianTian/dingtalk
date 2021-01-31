package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCustomizeConversationUpdateRequest() *OapiCustomizeConversationUpdateRequest {
	return &OapiCustomizeConversationUpdateRequest{}
}

type OapiCustomizeConversationUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomizeConversationUpdateResponse
	ChatId          string
	ExtensionKey    string
	ExtensionValue  string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCustomizeConversationUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomizeConversationUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomizeConversationUpdateRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiCustomizeConversationUpdateRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiCustomizeConversationUpdateRequest) SetExtensionKey(extensionKey2 string) {
	this.ExtensionKey = extensionKey2
}
func (this *OapiCustomizeConversationUpdateRequest) GetExtensionKey() string {
	return this.ExtensionKey
}
func (this *OapiCustomizeConversationUpdateRequest) SetExtensionValue(extensionValue2 string) {
	this.ExtensionValue = extensionValue2
}
func (this *OapiCustomizeConversationUpdateRequest) GetExtensionValue() string {
	return this.ExtensionValue
}
func (this *OapiCustomizeConversationUpdateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCustomizeConversationUpdateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCustomizeConversationUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customize.conversation.update"
}
func (this *OapiCustomizeConversationUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomizeConversationUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomizeConversationUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomizeConversationUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomizeConversationUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id"] = this.ChatId
	txtParams["extension_key"] = this.ExtensionKey
	txtParams["extension_value"] = this.ExtensionValue
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCustomizeConversationUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatId, "chatId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCustomizeConversationUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomizeConversationUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCustomizeConversationUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
