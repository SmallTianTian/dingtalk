package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatTagDeleteRequest() *OapiChatTagDeleteRequest {
	return &OapiChatTagDeleteRequest{}
}

type OapiChatTagDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatTagDeleteResponse
	Chatid          string
	GroupTag        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatTagDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatTagDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatTagDeleteRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatTagDeleteRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatTagDeleteRequest) SetGroupTag(groupTag2 int64) {
	this.GroupTag = groupTag2
}
func (this *OapiChatTagDeleteRequest) GetGroupTag() int64 {
	return this.GroupTag
}
func (this *OapiChatTagDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.tag.delete"
}
func (this *OapiChatTagDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatTagDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatTagDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatTagDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatTagDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams["group_tag"] = this.GroupTag
	return txtParams
}
func (this *OapiChatTagDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatTagDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatTagDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatTagDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
