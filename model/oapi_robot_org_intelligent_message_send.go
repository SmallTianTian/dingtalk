package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotOrgIntelligentMessageSendRequest() *OapiRobotOrgIntelligentMessageSendRequest {
	return &OapiRobotOrgIntelligentMessageSendRequest{}
}

type OapiRobotOrgIntelligentMessageSendRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiRobotOrgIntelligentMessageSendResponse
	AtUserIds          string
	MsgKey             string
	MsgParam           string
	OpenConversationId string
	ReceiverUserIds    string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiRobotOrgIntelligentMessageSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) SetAtUserIds(atUserIds2 string) {
	this.AtUserIds = atUserIds2
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetAtUserIds() string {
	return this.AtUserIds
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) SetMsgParam(msgParam2 string) {
	this.MsgParam = msgParam2
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetMsgParam() string {
	return this.MsgParam
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) SetReceiverUserIds(receiverUserIds2 string) {
	this.ReceiverUserIds = receiverUserIds2
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetReceiverUserIds() string {
	return this.ReceiverUserIds
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.org.intelligent.message.send"
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["at_user_ids"] = this.AtUserIds
	txtParams["msg_key"] = this.MsgKey
	txtParams["msg_param"] = this.MsgParam
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["receiver_user_ids"] = this.ReceiverUserIds
	return txtParams
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.AtUserIds, 999, "atUserIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotOrgIntelligentMessageSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotOrgIntelligentMessageSendResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
