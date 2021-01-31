package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatUpdategroupnickRequest() *OapiChatUpdategroupnickRequest {
	return &OapiChatUpdategroupnickRequest{}
}

type OapiChatUpdategroupnickRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatUpdategroupnickResponse
	Chatid          string
	GroupNick       string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiChatUpdategroupnickRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatUpdategroupnickRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatUpdategroupnickRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatUpdategroupnickRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatUpdategroupnickRequest) SetGroupNick(groupNick2 string) {
	this.GroupNick = groupNick2
}
func (this *OapiChatUpdategroupnickRequest) GetGroupNick() string {
	return this.GroupNick
}
func (this *OapiChatUpdategroupnickRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiChatUpdategroupnickRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiChatUpdategroupnickRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.updategroupnick"
}
func (this *OapiChatUpdategroupnickRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatUpdategroupnickRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatUpdategroupnickRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatUpdategroupnickRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatUpdategroupnickRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams["group_nick"] = this.GroupNick
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiChatUpdategroupnickRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatUpdategroupnickRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatUpdategroupnickRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatUpdategroupnickResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
