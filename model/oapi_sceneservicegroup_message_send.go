package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSceneservicegroupMessageSendRequest() *OapiSceneservicegroupMessageSendRequest {
	return &OapiSceneservicegroupMessageSendRequest{}
}

type OapiSceneservicegroupMessageSendRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                OapiSceneservicegroupMessageSendResponse
	AtDingtalkids       string
	AtMobiles           string
	AtUnionids          string
	Bizid               string
	BtnOrientation      string
	Btns                string
	Content             string
	IsAtAll             bool
	MessageType         string
	OpenConversationid  string
	ReceiverDingtalkids string
	ReceiverMobiles     string
	ReceiverUnionids    string
	Title               string
	TopHttpMethod       string
	TopResponseType     string
}

func (this *OapiSceneservicegroupMessageSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSceneservicegroupMessageSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSceneservicegroupMessageSendRequest) SetAtDingtalkids(atDingtalkids2 string) {
	this.AtDingtalkids = atDingtalkids2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetAtDingtalkids() string {
	return this.AtDingtalkids
}
func (this *OapiSceneservicegroupMessageSendRequest) SetAtMobiles(atMobiles2 string) {
	this.AtMobiles = atMobiles2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetAtMobiles() string {
	return this.AtMobiles
}
func (this *OapiSceneservicegroupMessageSendRequest) SetAtUnionids(atUnionids2 string) {
	this.AtUnionids = atUnionids2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetAtUnionids() string {
	return this.AtUnionids
}
func (this *OapiSceneservicegroupMessageSendRequest) SetBizid(bizid2 string) {
	this.Bizid = bizid2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetBizid() string {
	return this.Bizid
}
func (this *OapiSceneservicegroupMessageSendRequest) SetBtnOrientation(btnOrientation2 string) {
	this.BtnOrientation = btnOrientation2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetBtnOrientation() string {
	return this.BtnOrientation
}
func (this *OapiSceneservicegroupMessageSendRequest) SetBtns(btns2 string) {
	this.Btns = btns2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetBtns() string {
	return this.Btns
}
func (this *OapiSceneservicegroupMessageSendRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetContent() string {
	return this.Content
}
func (this *OapiSceneservicegroupMessageSendRequest) SetIsAtAll(isAtAll2 bool) {
	this.IsAtAll = isAtAll2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetIsAtAll() bool {
	return this.IsAtAll
}
func (this *OapiSceneservicegroupMessageSendRequest) SetMessageType(messageType2 string) {
	this.MessageType = messageType2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetMessageType() string {
	return this.MessageType
}
func (this *OapiSceneservicegroupMessageSendRequest) SetOpenConversationid(openConversationid2 string) {
	this.OpenConversationid = openConversationid2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetOpenConversationid() string {
	return this.OpenConversationid
}
func (this *OapiSceneservicegroupMessageSendRequest) SetReceiverDingtalkids(receiverDingtalkids2 string) {
	this.ReceiverDingtalkids = receiverDingtalkids2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetReceiverDingtalkids() string {
	return this.ReceiverDingtalkids
}
func (this *OapiSceneservicegroupMessageSendRequest) SetReceiverMobiles(receiverMobiles2 string) {
	this.ReceiverMobiles = receiverMobiles2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetReceiverMobiles() string {
	return this.ReceiverMobiles
}
func (this *OapiSceneservicegroupMessageSendRequest) SetReceiverUnionids(receiverUnionids2 string) {
	this.ReceiverUnionids = receiverUnionids2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetReceiverUnionids() string {
	return this.ReceiverUnionids
}
func (this *OapiSceneservicegroupMessageSendRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetTitle() string {
	return this.Title
}
func (this *OapiSceneservicegroupMessageSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sceneservicegroup.message.send"
}
func (this *OapiSceneservicegroupMessageSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSceneservicegroupMessageSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSceneservicegroupMessageSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSceneservicegroupMessageSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSceneservicegroupMessageSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["at_dingtalkids"] = this.AtDingtalkids
	txtParams["at_mobiles"] = this.AtMobiles
	txtParams["at_unionids"] = this.AtUnionids
	txtParams["bizid"] = this.Bizid
	txtParams["btn_orientation"] = this.BtnOrientation
	txtParams["btns"] = this.Btns
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["is_at_all"] = this.IsAtAll
	txtParams["message_type"] = this.MessageType
	txtParams["open_conversationid"] = this.OpenConversationid
	txtParams["receiver_dingtalkids"] = this.ReceiverDingtalkids
	txtParams["receiver_mobiles"] = this.ReceiverMobiles
	txtParams["receiver_unionids"] = this.ReceiverUnionids
	txtParams["title"] = this.Title
	return txtParams
}
func (this *OapiSceneservicegroupMessageSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.AtDingtalkids, 999, "atDingtalkids"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSceneservicegroupMessageSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSceneservicegroupMessageSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSceneservicegroupMessageSendResponse struct {
	taobao.TaobaoResponse
}
