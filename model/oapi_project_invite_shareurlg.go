package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProjectInviteShareurlGetRequest() *OapiProjectInviteShareurlGetRequest {
	return &OapiProjectInviteShareurlGetRequest{}
}

type OapiProjectInviteShareurlGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProjectInviteShareurlGetResponse
	InviteInfo      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProjectInviteShareurlGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProjectInviteShareurlGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProjectInviteShareurlGetRequest) SetInviteInfo(inviteInfo2 string) {
	this.InviteInfo = inviteInfo2
}
func (this *OapiProjectInviteShareurlGetRequest) GetInviteInfo() string {
	return this.InviteInfo
}
func (this *OapiProjectInviteShareurlGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.project.invite.shareurl.get"
}
func (this *OapiProjectInviteShareurlGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProjectInviteShareurlGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProjectInviteShareurlGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProjectInviteShareurlGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProjectInviteShareurlGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["invite_info"] = this.InviteInfo
	return txtParams
}
func (this *OapiProjectInviteShareurlGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProjectInviteShareurlGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProjectInviteShareurlGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type InviteUrlQuery struct {
	ChatId        string `json:"chat_id,omitempty"`
	ExpireSeconds int64  `json:"expire_seconds,omitempty"`
	Redirect      string `json:"redirect,omitempty"`
	SceneId       string `json:"scene_id,omitempty"`
	TermType      string `json:"term_type,omitempty"`
	Userid        string `json:"userid,omitempty"`
}
type OapiProjectInviteShareurlGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
