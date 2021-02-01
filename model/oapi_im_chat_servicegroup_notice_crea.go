package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatServicegroupNoticeCreateRequest() *OapiImChatServicegroupNoticeCreateRequest {
	return &OapiImChatServicegroupNoticeCreateRequest{}
}

type OapiImChatServicegroupNoticeCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImChatServicegroupNoticeCreateResponse
	ChatId          string
	SendDing        bool
	Sticky          bool
	TextContent     string
	Title           string
	TopHttpMethod   string
	TopResponseType string
	UniqueKey       string
	Userid          string
}

func (this *OapiImChatServicegroupNoticeCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetSendDing(sendDing2 bool) {
	this.SendDing = sendDing2
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetSendDing() bool {
	return this.SendDing
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetSticky(sticky2 bool) {
	this.Sticky = sticky2
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetSticky() bool {
	return this.Sticky
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetTextContent(textContent2 string) {
	this.TextContent = textContent2
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetTextContent() string {
	return this.TextContent
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetUniqueKey(uniqueKey2 string) {
	this.UniqueKey = uniqueKey2
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetUniqueKey() string {
	return this.UniqueKey
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.servicegroup.notice.create"
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatServicegroupNoticeCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id"] = this.ChatId
	txtParams["send_ding"] = this.SendDing
	txtParams["sticky"] = this.Sticky
	txtParams["text_content"] = this.TextContent
	txtParams["title"] = this.Title
	txtParams["unique_key"] = this.UniqueKey
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiImChatServicegroupNoticeCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatId, "chatId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatServicegroupNoticeCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatServicegroupNoticeCreateResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
