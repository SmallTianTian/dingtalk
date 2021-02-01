package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatTagSetRequest() *OapiChatTagSetRequest {
	return &OapiChatTagSetRequest{}
}

type OapiChatTagSetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatTagSetResponse
	Chatid          string
	GroupTag        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatTagSetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatTagSetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatTagSetRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatTagSetRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatTagSetRequest) SetGroupTag(groupTag2 int64) {
	this.GroupTag = groupTag2
}
func (this *OapiChatTagSetRequest) GetGroupTag() int64 {
	return this.GroupTag
}
func (this *OapiChatTagSetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.tag.set"
}
func (this *OapiChatTagSetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatTagSetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatTagSetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatTagSetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatTagSetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams["group_tag"] = this.GroupTag
	return txtParams
}
func (this *OapiChatTagSetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatTagSetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatTagSetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatTagSetResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
