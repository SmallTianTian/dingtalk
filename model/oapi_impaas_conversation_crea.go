package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImpaasConversationCreateRequest() *OapiImpaasConversationCreateRequest {
	return &OapiImpaasConversationCreateRequest{}
}

type OapiImpaasConversationCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasConversationCreateResponse
	Channel         string
	Name            string
	OwnerUserid     string
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiImpaasConversationCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasConversationCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasConversationCreateRequest) SetChannel(channel2 string) {
	this.Channel = channel2
}
func (this *OapiImpaasConversationCreateRequest) GetChannel() string {
	return this.Channel
}
func (this *OapiImpaasConversationCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiImpaasConversationCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiImpaasConversationCreateRequest) SetOwnerUserid(ownerUserid2 string) {
	this.OwnerUserid = ownerUserid2
}
func (this *OapiImpaasConversationCreateRequest) GetOwnerUserid() string {
	return this.OwnerUserid
}
func (this *OapiImpaasConversationCreateRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiImpaasConversationCreateRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiImpaasConversationCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.conversation.create"
}
func (this *OapiImpaasConversationCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasConversationCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasConversationCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasConversationCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasConversationCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["channel"] = this.Channel
	txtParams["name"] = this.Name
	txtParams["owner_userid"] = this.OwnerUserid
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiImpaasConversationCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Channel, "channel"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImpaasConversationCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasConversationCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImpaasConversationCreateResponse struct {
	taobao.TaobaoResponse
	Chatid  string `json:"chatid,omitempty"`
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
