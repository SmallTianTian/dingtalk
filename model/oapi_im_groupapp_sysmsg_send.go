package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImGroupappSysmsgSendRequest() *OapiImGroupappSysmsgSendRequest {
	return &OapiImGroupappSysmsgSendRequest{}
}

type OapiImGroupappSysmsgSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImGroupappSysmsgSendResponse
	MsgKey             string
	MsgParam           string
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiImGroupappSysmsgSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImGroupappSysmsgSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImGroupappSysmsgSendRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiImGroupappSysmsgSendRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiImGroupappSysmsgSendRequest) SetMsgParam(msgParam2 string) {
	this.MsgParam = msgParam2
}
func (this *OapiImGroupappSysmsgSendRequest) GetMsgParam() string {
	return this.MsgParam
}
func (this *OapiImGroupappSysmsgSendRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImGroupappSysmsgSendRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImGroupappSysmsgSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.groupapp.sysmsg.send"
}
func (this *OapiImGroupappSysmsgSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImGroupappSysmsgSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImGroupappSysmsgSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImGroupappSysmsgSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImGroupappSysmsgSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["msg_key"] = this.MsgKey
	txtParams["msg_param"] = this.MsgParam
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *OapiImGroupappSysmsgSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MsgKey, "msgKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImGroupappSysmsgSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImGroupappSysmsgSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImGroupappSysmsgSendResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
