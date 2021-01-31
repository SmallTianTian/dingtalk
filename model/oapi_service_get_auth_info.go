package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiServiceGetAuthInfoRequest() *OapiServiceGetAuthInfoRequest {
	return &OapiServiceGetAuthInfoRequest{}
}

type OapiServiceGetAuthInfoRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceGetAuthInfoResponse
	AuthCorpid      string
	SuiteKey        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceGetAuthInfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceGetAuthInfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceGetAuthInfoRequest) SetAuthCorpid(authCorpid2 string) {
	this.AuthCorpid = authCorpid2
}
func (this *OapiServiceGetAuthInfoRequest) GetAuthCorpid() string {
	return this.AuthCorpid
}
func (this *OapiServiceGetAuthInfoRequest) SetSuiteKey(suiteKey2 string) {
	this.SuiteKey = suiteKey2
}
func (this *OapiServiceGetAuthInfoRequest) GetSuiteKey() string {
	return this.SuiteKey
}
func (this *OapiServiceGetAuthInfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.get_auth_info"
}
func (this *OapiServiceGetAuthInfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceGetAuthInfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceGetAuthInfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceGetAuthInfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceGetAuthInfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["auth_corpid"] = this.AuthCorpid
	txtParams["suite_key"] = this.SuiteKey
	return txtParams
}
func (this *OapiServiceGetAuthInfoRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiServiceGetAuthInfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceGetAuthInfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceGetAuthInfoResponse struct {
	taobao.TaobaoResponse
	AuthCorpInfo    AuthCorpInfo    `json:"auth_corp_info,omitempty"`
	AuthInfo        AuthInfo        `json:"auth_info,omitempty"`
	AuthUserInfo    AuthUserInfo    `json:"auth_user_info,omitempty"`
	ChannelAuthInfo ChannelAuthInfo `json:"channel_auth_info,omitempty"`
	Errcode         int64           `json:"errcode,omitempty"`
	Errmsg          string          `json:"errmsg,omitempty"`
}
type Agent struct {
	AdminList []string `json:"admin_list,omitempty"`
	AgentName string   `json:"agent_name,omitempty"`
	Agentid   int64    `json:"agentid,omitempty"`
	Appid     int64    `json:"appid,omitempty"`
	LogoUrl   string   `json:"logo_url,omitempty"`
}
type AuthInfo struct {
	Agent []Agent `json:"agent,omitempty"`
}
type AuthUserInfo struct {
	UserId string `json:"userId,omitempty"`
}
type AuthCorpInfo struct {
	AuthChannel         string `json:"auth_channel,omitempty"`
	AuthChannelType     string `json:"auth_channel_type,omitempty"`
	AuthLevel           int64  `json:"auth_level,omitempty"`
	BelongCorpId        string `json:"belong_corp_id,omitempty"`
	CorpLogoUrl         string `json:"corp_logo_url,omitempty"`
	CorpName            string `json:"corp_name,omitempty"`
	Corpid              string `json:"corpid,omitempty"`
	Industry            string `json:"industry,omitempty"`
	InviteCode          string `json:"invite_code,omitempty"`
	InviteUrl           string `json:"invite_url,omitempty"`
	IsAuthenticated     bool   `json:"is_authenticated,omitempty"`
	LicenseCode         string `json:"license_code,omitempty"`
	UnifiedSocialCredit string `json:"unifiedSocialCredit,omitempty"`
}
type Channelagent struct {
	AgentName string `json:"agent_name,omitempty"`
	Agentid   int64  `json:"agentid,omitempty"`
	Appid     int64  `json:"appid,omitempty"`
	LogoUrl   string `json:"logo_url,omitempty"`
}
type ChannelAuthInfo struct {
	ChannelAgent []Channelagent `json:"channelAgent,omitempty"`
}
