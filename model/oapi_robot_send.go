package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotSendRequest() *OapiRobotSendRequest {
	return &OapiRobotSendRequest{}
}

type OapiRobotSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRobotSendResponse
	ActionCard      string
	At              string
	FeedCard        string
	Link            string
	Markdown        string
	Msgtype         string
	Text            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRobotSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotSendRequest) SetActionCard(actionCard2 string) {
	this.ActionCard = actionCard2
}
func (this *OapiRobotSendRequest) GetActionCard() string {
	return this.ActionCard
}
func (this *OapiRobotSendRequest) SetAt(at2 string) {
	this.At = at2
}
func (this *OapiRobotSendRequest) GetAt() string {
	return this.At
}
func (this *OapiRobotSendRequest) SetFeedCard(feedCard2 string) {
	this.FeedCard = feedCard2
}
func (this *OapiRobotSendRequest) GetFeedCard() string {
	return this.FeedCard
}
func (this *OapiRobotSendRequest) SetLink(link2 string) {
	this.Link = link2
}
func (this *OapiRobotSendRequest) GetLink() string {
	return this.Link
}
func (this *OapiRobotSendRequest) SetMarkdown(markdown2 string) {
	this.Markdown = markdown2
}
func (this *OapiRobotSendRequest) GetMarkdown() string {
	return this.Markdown
}
func (this *OapiRobotSendRequest) SetMsgtype(msgtype2 string) {
	this.Msgtype = msgtype2
}
func (this *OapiRobotSendRequest) GetMsgtype() string {
	return this.Msgtype
}
func (this *OapiRobotSendRequest) SetText(text2 string) {
	this.Text = text2
}
func (this *OapiRobotSendRequest) GetText() string {
	return this.Text
}
func (this *OapiRobotSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.send"
}
func (this *OapiRobotSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["actionCard"] = this.ActionCard
	txtParams["at"] = this.At
	txtParams["feedCard"] = this.FeedCard
	txtParams["link"] = this.Link
	txtParams["markdown"] = this.Markdown
	txtParams["msgtype"] = this.Msgtype
	txtParams["text"] = this.Text
	return txtParams
}
func (this *OapiRobotSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Msgtype, "msgtype"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}
type Btns struct {
	ActionURL string `json:"actionURL,omitempty"`
	Title     string `json:"title,omitempty"`
}
type Actioncard struct {
	BtnOrientation string `json:"btnOrientation,omitempty"`
	Btns           []Btns `json:"btns,omitempty"`
	HideAvatar     string `json:"hideAvatar,omitempty"`
	SingleTitle    string `json:"singleTitle,omitempty"`
	SingleURL      string `json:"singleURL,omitempty"`
	Text           string `json:"text,omitempty"`
	Title          string `json:"title,omitempty"`
}
type Links struct {
	MessageURL string `json:"messageURL,omitempty"`
	PicURL     string `json:"picURL,omitempty"`
	Title      string `json:"title,omitempty"`
}
type Feedcard struct {
	Links []Links `json:"links,omitempty"`
}
type OapiRobotSendResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
