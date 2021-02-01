package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScencegroupMessageSendRequest() *OapiImChatScencegroupMessageSendRequest {
	return &OapiImChatScencegroupMessageSendRequest{}
}

type OapiImChatScencegroupMessageSendRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                     OapiImChatScencegroupMessageSendResponse
	AtMobiles                string
	IsAtAll                  bool
	MsgMediaIdParamMap       string
	MsgParamMap              string
	MsgTemplateId            string
	ReceiverMobiles          string
	ReceiverUnionIds         string
	ReceiverUserIds          string
	RobotCode                string
	TargetOpenConversationId string
	TopHttpMethod            string
	TopResponseType          string
}

func (this *OapiImChatScencegroupMessageSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScencegroupMessageSendRequest) SetAtMobiles(atMobiles2 string) {
	this.AtMobiles = atMobiles2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetAtMobiles() string {
	return this.AtMobiles
}
func (this *OapiImChatScencegroupMessageSendRequest) SetIsAtAll(isAtAll2 bool) {
	this.IsAtAll = isAtAll2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetIsAtAll() bool {
	return this.IsAtAll
}
func (this *OapiImChatScencegroupMessageSendRequest) SetMsgMediaIdParamMap(msgMediaIdParamMap2 string) {
	this.MsgMediaIdParamMap = msgMediaIdParamMap2
}
func (this *OapiImChatScencegroupMessageSendRequest) SetMsgMediaIdParamMapString(msgMediaIdParamMap2 string) {
	this.MsgMediaIdParamMap = msgMediaIdParamMap2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetMsgMediaIdParamMap() string {
	return this.MsgMediaIdParamMap
}
func (this *OapiImChatScencegroupMessageSendRequest) SetMsgParamMap(msgParamMap2 string) {
	this.MsgParamMap = msgParamMap2
}
func (this *OapiImChatScencegroupMessageSendRequest) SetMsgParamMapString(msgParamMap2 string) {
	this.MsgParamMap = msgParamMap2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetMsgParamMap() string {
	return this.MsgParamMap
}
func (this *OapiImChatScencegroupMessageSendRequest) SetMsgTemplateId(msgTemplateId2 string) {
	this.MsgTemplateId = msgTemplateId2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetMsgTemplateId() string {
	return this.MsgTemplateId
}
func (this *OapiImChatScencegroupMessageSendRequest) SetReceiverMobiles(receiverMobiles2 string) {
	this.ReceiverMobiles = receiverMobiles2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetReceiverMobiles() string {
	return this.ReceiverMobiles
}
func (this *OapiImChatScencegroupMessageSendRequest) SetReceiverUnionIds(receiverUnionIds2 string) {
	this.ReceiverUnionIds = receiverUnionIds2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetReceiverUnionIds() string {
	return this.ReceiverUnionIds
}
func (this *OapiImChatScencegroupMessageSendRequest) SetReceiverUserIds(receiverUserIds2 string) {
	this.ReceiverUserIds = receiverUserIds2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetReceiverUserIds() string {
	return this.ReceiverUserIds
}
func (this *OapiImChatScencegroupMessageSendRequest) SetRobotCode(robotCode2 string) {
	this.RobotCode = robotCode2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetRobotCode() string {
	return this.RobotCode
}
func (this *OapiImChatScencegroupMessageSendRequest) SetTargetOpenConversationId(targetOpenConversationId2 string) {
	this.TargetOpenConversationId = targetOpenConversationId2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetTargetOpenConversationId() string {
	return this.TargetOpenConversationId
}
func (this *OapiImChatScencegroupMessageSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scencegroup.message.send"
}
func (this *OapiImChatScencegroupMessageSendRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *OapiImChatScencegroupMessageSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScencegroupMessageSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScencegroupMessageSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScencegroupMessageSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScencegroupMessageSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["at_mobiles"] = this.AtMobiles
	txtParams["is_at_all"] = this.IsAtAll
	txtParams["msg_media_id_param_map"] = this.MsgMediaIdParamMap
	txtParams["msg_param_map"] = this.MsgParamMap
	txtParams["msg_template_id"] = this.MsgTemplateId
	txtParams["receiver_mobiles"] = this.ReceiverMobiles
	txtParams["receiver_union_ids"] = this.ReceiverUnionIds
	txtParams["receiver_user_ids"] = this.ReceiverUserIds
	txtParams["robot_code"] = this.RobotCode
	txtParams["target_open_conversation_id"] = this.TargetOpenConversationId
	return txtParams
}
func (this *OapiImChatScencegroupMessageSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.AtMobiles, 999, "atMobiles"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScencegroupMessageSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScencegroupMessageSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScencegroupMessageSendResponse struct {
	taobao.TaobaoResponse
	Succ bool `json:"succ,omitempty"`
}
