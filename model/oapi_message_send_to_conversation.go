package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMessageSendToConversationRequest() *OapiMessageSendToConversationRequest {
	return &OapiMessageSendToConversationRequest{}
}

type OapiMessageSendToConversationRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageSendToConversationResponse
	ActionCard      string
	Cid             string
	File            string
	Image           string
	Link            string
	Markdown        string
	Msg             string
	Msgtype         string
	Oa              string
	Sender          string
	Text            string
	TopHttpMethod   string
	TopResponseType string
	Voice           string
}

func (this *OapiMessageSendToConversationRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageSendToConversationRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageSendToConversationRequest) SetActionCard(actionCard2 string) {
	this.ActionCard = actionCard2
}
func (this *OapiMessageSendToConversationRequest) GetActionCard() string {
	return this.ActionCard
}
func (this *OapiMessageSendToConversationRequest) SetCid(cid2 string) {
	this.Cid = cid2
}
func (this *OapiMessageSendToConversationRequest) GetCid() string {
	return this.Cid
}
func (this *OapiMessageSendToConversationRequest) SetFile(file2 string) {
	this.File = file2
}
func (this *OapiMessageSendToConversationRequest) GetFile() string {
	return this.File
}
func (this *OapiMessageSendToConversationRequest) SetImage(image2 string) {
	this.Image = image2
}
func (this *OapiMessageSendToConversationRequest) GetImage() string {
	return this.Image
}
func (this *OapiMessageSendToConversationRequest) SetLink(link2 string) {
	this.Link = link2
}
func (this *OapiMessageSendToConversationRequest) GetLink() string {
	return this.Link
}
func (this *OapiMessageSendToConversationRequest) SetMarkdown(markdown2 string) {
	this.Markdown = markdown2
}
func (this *OapiMessageSendToConversationRequest) GetMarkdown() string {
	return this.Markdown
}
func (this *OapiMessageSendToConversationRequest) SetMsg(msg2 string) {
	this.Msg = msg2
}
func (this *OapiMessageSendToConversationRequest) GetMsg() string {
	return this.Msg
}
func (this *OapiMessageSendToConversationRequest) SetMsgtype(msgtype2 string) {
	this.Msgtype = msgtype2
}
func (this *OapiMessageSendToConversationRequest) GetMsgtype() string {
	return this.Msgtype
}
func (this *OapiMessageSendToConversationRequest) SetOa(oa2 string) {
	this.Oa = oa2
}
func (this *OapiMessageSendToConversationRequest) GetOa() string {
	return this.Oa
}
func (this *OapiMessageSendToConversationRequest) SetSender(sender2 string) {
	this.Sender = sender2
}
func (this *OapiMessageSendToConversationRequest) GetSender() string {
	return this.Sender
}
func (this *OapiMessageSendToConversationRequest) SetText(text2 string) {
	this.Text = text2
}
func (this *OapiMessageSendToConversationRequest) GetText() string {
	return this.Text
}
func (this *OapiMessageSendToConversationRequest) SetVoice(voice2 string) {
	this.Voice = voice2
}
func (this *OapiMessageSendToConversationRequest) GetVoice() string {
	return this.Voice
}
func (this *OapiMessageSendToConversationRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.send_to_conversation"
}
func (this *OapiMessageSendToConversationRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageSendToConversationRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageSendToConversationRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageSendToConversationRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageSendToConversationRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["action_card"] = this.ActionCard
	txtParams["cid"] = this.Cid
	txtParams["file"] = this.File
	txtParams["image"] = this.Image
	txtParams["link"] = this.Link
	txtParams["markdown"] = this.Markdown
	txtParams["msg"] = this.Msg
	txtParams["msgtype"] = this.Msgtype
	txtParams["oa"] = this.Oa
	txtParams["sender"] = this.Sender
	txtParams["text"] = this.Text
	txtParams["voice"] = this.Voice
	return txtParams
}
func (this *OapiMessageSendToConversationRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMessageSendToConversationRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageSendToConversationRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageSendToConversationResponse struct {
	taobao.TaobaoResponse
	Errcode  int64  `json:"errcode,omitempty"`
	Errmsg   string `json:"errmsg,omitempty"`
	Receiver string `json:"receiver,omitempty"`
}
