package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiChatSendRequest() *OapiChatSendRequest {
	return &OapiChatSendRequest{}
}

type OapiChatSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatSendResponse
	ActionCard      string
	Chatid          string
	File            string
	Image           string
	Link            string
	Markdown        string
	Msg             string
	Msgtype         string
	Oa              string
	Text            string
	TopHttpMethod   string
	TopResponseType string
	Voice           string
}

func (this *OapiChatSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatSendRequest) SetActionCard(actionCard2 string) {
	this.ActionCard = actionCard2
}
func (this *OapiChatSendRequest) GetActionCard() string {
	return this.ActionCard
}
func (this *OapiChatSendRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatSendRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatSendRequest) SetFile(file2 string) {
	this.File = file2
}
func (this *OapiChatSendRequest) GetFile() string {
	return this.File
}
func (this *OapiChatSendRequest) SetImage(image2 string) {
	this.Image = image2
}
func (this *OapiChatSendRequest) GetImage() string {
	return this.Image
}
func (this *OapiChatSendRequest) SetLink(link2 string) {
	this.Link = link2
}
func (this *OapiChatSendRequest) GetLink() string {
	return this.Link
}
func (this *OapiChatSendRequest) SetMarkdown(markdown2 string) {
	this.Markdown = markdown2
}
func (this *OapiChatSendRequest) GetMarkdown() string {
	return this.Markdown
}
func (this *OapiChatSendRequest) SetMsg(msg2 string) {
	this.Msg = msg2
}
func (this *OapiChatSendRequest) GetMsg() string {
	return this.Msg
}
func (this *OapiChatSendRequest) SetMsgtype(msgtype2 string) {
	this.Msgtype = msgtype2
}
func (this *OapiChatSendRequest) GetMsgtype() string {
	return this.Msgtype
}
func (this *OapiChatSendRequest) SetOa(oa2 string) {
	this.Oa = oa2
}
func (this *OapiChatSendRequest) GetOa() string {
	return this.Oa
}
func (this *OapiChatSendRequest) SetText(text2 string) {
	this.Text = text2
}
func (this *OapiChatSendRequest) GetText() string {
	return this.Text
}
func (this *OapiChatSendRequest) SetVoice(voice2 string) {
	this.Voice = voice2
}
func (this *OapiChatSendRequest) GetVoice() string {
	return this.Voice
}
func (this *OapiChatSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.send"
}
func (this *OapiChatSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["action_card"] = this.ActionCard
	txtParams["chatid"] = this.Chatid
	txtParams["file"] = this.File
	txtParams["image"] = this.Image
	txtParams["link"] = this.Link
	txtParams["markdown"] = this.Markdown
	txtParams["msg"] = this.Msg
	txtParams["msgtype"] = this.Msgtype
	txtParams["oa"] = this.Oa
	txtParams["text"] = this.Text
	txtParams["voice"] = this.Voice
	return txtParams
}
func (this *OapiChatSendRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiChatSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BtnJson struct {
	ActionUrl string `json:"action_url,omitempty"`
	Title     string `json:"title,omitempty"`
}
type ActionCard struct {
	Agentid        string    `json:"agentid,omitempty"`
	BtnJsonList    []BtnJson `json:"btn_json_list,omitempty"`
	BtnOrientation string    `json:"btn_orientation,omitempty"`
	HideAvatar     bool      `json:"hide_avatar,omitempty"`
	Markdown       string    `json:"markdown,omitempty"`
	SingleTitle    string    `json:"single_title,omitempty"`
	SingleUrl      string    `json:"single_url,omitempty"`
	Title          string    `json:"title,omitempty"`
}
type Head struct {
	Bgcolor string `json:"bgcolor,omitempty"`
	Text    string `json:"text,omitempty"`
}
type Rich struct {
	Num  string `json:"num,omitempty"`
	Unit string `json:"unit,omitempty"`
}
type Form struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
type Body struct {
	Author    string `json:"author,omitempty"`
	Content   string `json:"content,omitempty"`
	FileCount string `json:"file_count,omitempty"`
	Form      []Form `json:taobao.CONTENT_TYPE_FOR,omitemptyM`
	Image     string `json:"image,omitempty"`
	Rich      Rich   `json:"rich,omitempty"`
	Title     string `json:"title,omitempty"`
}
type Oa struct {
	Body         Body   `json:"body,omitempty"`
	Head         Head   `json:"head,omitempty"`
	MessageUrl   string `json:"message_url,omitempty"`
	PcMessageUrl string `json:"pc_message_url,omitempty"`
}
type Voice struct {
	Duration int64  `json:"duration,omitempty"`
	MediaId  string `json:"media_id,omitempty"`
}
type File struct {
	MediaId string `json:"media_id,omitempty"`
}
type Image struct {
	MediaId string `json:"media_id,omitempty"`
}
type Link struct {
	MessageUrl string `json:"messageUrl,omitempty"`
	PicUrl     string `json:"picUrl,omitempty"`
	Text       string `json:"text,omitempty"`
	Title      string `json:"title,omitempty"`
}
type Text struct {
	Content string `json:"content,omitempty"`
}
type Markdown struct {
	Text  string `json:"text,omitempty"`
	Title string `json:"title,omitempty"`
}
type BtnJsonList struct {
	ActionUrl string `json:"action_url,omitempty"`
	Title     string `json:"title,omitempty"`
}
type Msg struct {
	ActionCard ActionCard `json:"action_card,omitempty"`
	File       File       `json:"file,omitempty"`
	Image      Image      `json:"image,omitempty"`
	Link       Link       `json:"link,omitempty"`
	Markdown   Markdown   `json:"markdown,omitempty"`
	Msgtype    string     `json:"msgtype,omitempty"`
	Oa         Oa         `json:"oa,omitempty"`
	Text       Text       `json:"text,omitempty"`
	Voice      Voice      `json:"voice,omitempty"`
}
type OapiChatSendResponse struct {
	taobao.TaobaoResponse
	Errcode   int64  `json:"errcode,omitempty"`
	Errmsg    string `json:"errmsg,omitempty"`
	MessageId string `json:"messageId,omitempty"`
}
