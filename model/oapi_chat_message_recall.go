package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatMessageRecallRequest() *OapiChatMessageRecallRequest {
	return &OapiChatMessageRecallRequest{}
}

type OapiChatMessageRecallRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatMessageRecallResponse
	Chatid          string
	Msgid           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatMessageRecallRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatMessageRecallRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatMessageRecallRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatMessageRecallRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatMessageRecallRequest) SetMsgid(msgid2 string) {
	this.Msgid = msgid2
}
func (this *OapiChatMessageRecallRequest) GetMsgid() string {
	return this.Msgid
}
func (this *OapiChatMessageRecallRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.message.recall"
}
func (this *OapiChatMessageRecallRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatMessageRecallRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatMessageRecallRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatMessageRecallRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatMessageRecallRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams["msgid"] = this.Msgid
	return txtParams
}
func (this *OapiChatMessageRecallRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatMessageRecallRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatMessageRecallRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatMessageRecallResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
