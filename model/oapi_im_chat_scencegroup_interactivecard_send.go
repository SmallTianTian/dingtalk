package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScencegroupInteractivecardSendRequest() *OapiImChatScencegroupInteractivecardSendRequest {
	return &OapiImChatScencegroupInteractivecardSendRequest{}
}

type OapiImChatScencegroupInteractivecardSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                     OapiImChatScencegroupInteractivecardSendResponse
	CardMediaidParamMap      string
	CardParamMap             string
	CardTemplateId           string
	OutTrackId               string
	ReceiverUseridList       string
	RobotCode                string
	TargetOpenConversationId string
	TopHttpMethod            string
	TopResponseType          string
}

func (this *OapiImChatScencegroupInteractivecardSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetCardMediaidParamMap(cardMediaidParamMap2 string) {
	this.CardMediaidParamMap = cardMediaidParamMap2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetCardMediaidParamMapString(cardMediaidParamMap2 string) {
	this.CardMediaidParamMap = cardMediaidParamMap2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetCardMediaidParamMap() string {
	return this.CardMediaidParamMap
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetCardParamMap(cardParamMap2 string) {
	this.CardParamMap = cardParamMap2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetCardParamMapString(cardParamMap2 string) {
	this.CardParamMap = cardParamMap2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetCardParamMap() string {
	return this.CardParamMap
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetCardTemplateId(cardTemplateId2 string) {
	this.CardTemplateId = cardTemplateId2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetCardTemplateId() string {
	return this.CardTemplateId
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetOutTrackId(outTrackId2 string) {
	this.OutTrackId = outTrackId2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetOutTrackId() string {
	return this.OutTrackId
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetReceiverUseridList(receiverUseridList2 string) {
	this.ReceiverUseridList = receiverUseridList2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetReceiverUseridList() string {
	return this.ReceiverUseridList
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetRobotCode(robotCode2 string) {
	this.RobotCode = robotCode2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetRobotCode() string {
	return this.RobotCode
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetTargetOpenConversationId(targetOpenConversationId2 string) {
	this.TargetOpenConversationId = targetOpenConversationId2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetTargetOpenConversationId() string {
	return this.TargetOpenConversationId
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scencegroup.interactivecard.send"
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["card_mediaid_param_map"] = this.CardMediaidParamMap
	txtParams["card_param_map"] = this.CardParamMap
	txtParams["card_template_id"] = this.CardTemplateId
	txtParams["out_track_id"] = this.OutTrackId
	txtParams["receiver_userid_list"] = this.ReceiverUseridList
	txtParams["robot_code"] = this.RobotCode
	txtParams["target_open_conversation_id"] = this.TargetOpenConversationId
	return txtParams
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CardTemplateId, "cardTemplateId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScencegroupInteractivecardSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScencegroupInteractivecardSendResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
