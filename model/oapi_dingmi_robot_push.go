package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingmiRobotPushRequest() *OapiDingmiRobotPushRequest {
	return &OapiDingmiRobotPushRequest{}
}

type OapiDingmiRobotPushRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingmiRobotPushResponse
	ConversationId  string
	MsgKey          string
	MsgParam        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingmiRobotPushRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingmiRobotPushRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingmiRobotPushRequest) SetConversationId(conversationId2 string) {
	this.ConversationId = conversationId2
}
func (this *OapiDingmiRobotPushRequest) GetConversationId() string {
	return this.ConversationId
}
func (this *OapiDingmiRobotPushRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiDingmiRobotPushRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiDingmiRobotPushRequest) SetMsgParam(msgParam2 string) {
	this.MsgParam = msgParam2
}
func (this *OapiDingmiRobotPushRequest) GetMsgParam() string {
	return this.MsgParam
}
func (this *OapiDingmiRobotPushRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingmi.robot.push"
}
func (this *OapiDingmiRobotPushRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingmiRobotPushRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingmiRobotPushRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingmiRobotPushRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingmiRobotPushRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conversation_id"] = this.ConversationId
	txtParams["msg_key"] = this.MsgKey
	txtParams["msg_param"] = this.MsgParam
	return txtParams
}
func (this *OapiDingmiRobotPushRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingmiRobotPushRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingmiRobotPushRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingmiRobotPushResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
