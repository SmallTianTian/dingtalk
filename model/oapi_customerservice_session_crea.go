package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceSessionCreateRequest() *OapiCustomerserviceSessionCreateRequest {
	return &OapiCustomerserviceSessionCreateRequest{}
}

type OapiCustomerserviceSessionCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomerserviceSessionCreateResponse
	CreateSession   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCustomerserviceSessionCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceSessionCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceSessionCreateRequest) SetCreateSession(createSession2 string) {
	this.CreateSession = createSession2
}
func (this *OapiCustomerserviceSessionCreateRequest) GetCreateSession() string {
	return this.CreateSession
}
func (this *OapiCustomerserviceSessionCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.session.create"
}
func (this *OapiCustomerserviceSessionCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceSessionCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceSessionCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceSessionCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceSessionCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["create_session"] = this.CreateSession
	return txtParams
}
func (this *OapiCustomerserviceSessionCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceSessionCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceSessionCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SessionSourceDTO struct {
	Channel           string `json:"channel,omitempty"`
	ChannelAvatarUrl  string `json:"channel_avatar_url,omitempty"`
	ChannelNick       string `json:"channel_nick,omitempty"`
	ChannelUid        string `json:"channel_uid,omitempty"`
	ChannelUserSource string `json:"channel_user_source,omitempty"`
	CmsId             string `json:"cms_id,omitempty"`
	SessionEndTime    int64  `json:"session_end_time,omitempty"`
	SessionSource     string `json:"session_source,omitempty"`
	SessionStartTime  int64  `json:"session_start_time,omitempty"`
	SessionStatus     int64  `json:"session_status,omitempty"`
	Sid               string `json:"sid,omitempty"`
	Summary           string `json:"summary,omitempty"`
}
type SessionTargetDTO struct {
	BuId          string `json:"bu_id,omitempty"`
	ServiceId     string `json:"service_id,omitempty"`
	SessionSource string `json:"session_source,omitempty"`
}
type CreateSessionDTO struct {
	Source SessionSourceDTO `json:"source,omitempty"`
	Target SessionTargetDTO `json:"target,omitempty"`
}
type OapiCustomerserviceSessionCreateResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
