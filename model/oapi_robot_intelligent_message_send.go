package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotIntelligentMessageSendRequest() *OapiRobotIntelligentMessageSendRequest {
	return &OapiRobotIntelligentMessageSendRequest{}
}

type OapiRobotIntelligentMessageSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiRobotIntelligentMessageSendResponse
	AtUnionIds         string
	MsgKey             string
	MsgParam           string
	OpenConversationId string
	ReceiverUnionIds   string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiRobotIntelligentMessageSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotIntelligentMessageSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotIntelligentMessageSendRequest) SetAtUnionIds(atUnionIds2 string) {
	this.AtUnionIds = atUnionIds2
}
func (this *OapiRobotIntelligentMessageSendRequest) GetAtUnionIds() string {
	return this.AtUnionIds
}
func (this *OapiRobotIntelligentMessageSendRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiRobotIntelligentMessageSendRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiRobotIntelligentMessageSendRequest) SetMsgParam(msgParam2 string) {
	this.MsgParam = msgParam2
}
func (this *OapiRobotIntelligentMessageSendRequest) GetMsgParam() string {
	return this.MsgParam
}
func (this *OapiRobotIntelligentMessageSendRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiRobotIntelligentMessageSendRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiRobotIntelligentMessageSendRequest) SetReceiverUnionIds(receiverUnionIds2 string) {
	this.ReceiverUnionIds = receiverUnionIds2
}
func (this *OapiRobotIntelligentMessageSendRequest) GetReceiverUnionIds() string {
	return this.ReceiverUnionIds
}
func (this *OapiRobotIntelligentMessageSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.intelligent.message.send"
}
func (this *OapiRobotIntelligentMessageSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotIntelligentMessageSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotIntelligentMessageSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotIntelligentMessageSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotIntelligentMessageSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["at_union_ids"] = this.AtUnionIds
	txtParams["msg_key"] = this.MsgKey
	txtParams["msg_param"] = this.MsgParam
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["receiver_union_ids"] = this.ReceiverUnionIds
	return txtParams
}
func (this *OapiRobotIntelligentMessageSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.AtUnionIds, 999, "atUnionIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotIntelligentMessageSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotIntelligentMessageSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotIntelligentMessageSendResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
