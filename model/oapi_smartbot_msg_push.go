package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartbotMsgPushRequest() *OapiSmartbotMsgPushRequest {
	return &OapiSmartbotMsgPushRequest{}
}

type OapiSmartbotMsgPushRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartbotMsgPushResponse
	ChatIdList      string
	Msg             string
	ToAllUser       bool
	TopHttpMethod   string
	TopResponseType string
	UserIdList      string
}

func (this *OapiSmartbotMsgPushRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartbotMsgPushRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartbotMsgPushRequest) SetChatIdList(chatIdList2 string) {
	this.ChatIdList = chatIdList2
}
func (this *OapiSmartbotMsgPushRequest) GetChatIdList() string {
	return this.ChatIdList
}
func (this *OapiSmartbotMsgPushRequest) SetMsg(msg2 string) {
	this.Msg = msg2
}
func (this *OapiSmartbotMsgPushRequest) GetMsg() string {
	return this.Msg
}
func (this *OapiSmartbotMsgPushRequest) SetToAllUser(toAllUser2 bool) {
	this.ToAllUser = toAllUser2
}
func (this *OapiSmartbotMsgPushRequest) GetToAllUser() bool {
	return this.ToAllUser
}
func (this *OapiSmartbotMsgPushRequest) SetUserIdList(userIdList2 string) {
	this.UserIdList = userIdList2
}
func (this *OapiSmartbotMsgPushRequest) GetUserIdList() string {
	return this.UserIdList
}
func (this *OapiSmartbotMsgPushRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartbot.msg.push"
}
func (this *OapiSmartbotMsgPushRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartbotMsgPushRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartbotMsgPushRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartbotMsgPushRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartbotMsgPushRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id_list"] = this.ChatIdList
	txtParams["msg"] = this.Msg
	txtParams["to_all_user"] = this.ToAllUser
	txtParams["user_id_list"] = this.UserIdList
	return txtParams
}
func (this *OapiSmartbotMsgPushRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.ChatIdList, 500, "chatIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartbotMsgPushRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartbotMsgPushRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartbotMsgPushResponse struct {
	taobao.TaobaoResponse
	TaskId string `json:"task_id,omitempty"`
}
