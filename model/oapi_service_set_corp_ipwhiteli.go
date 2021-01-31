package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiServiceSetCorpIpwhitelistRequest() *OapiServiceSetCorpIpwhitelistRequest {
	return &OapiServiceSetCorpIpwhitelistRequest{}
}

type OapiServiceSetCorpIpwhitelistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceSetCorpIpwhitelistResponse
	AuthCorpid      string
	IpWhitelist     []string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceSetCorpIpwhitelistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceSetCorpIpwhitelistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceSetCorpIpwhitelistRequest) SetAuthCorpid(authCorpid2 string) {
	this.AuthCorpid = authCorpid2
}
func (this *OapiServiceSetCorpIpwhitelistRequest) GetAuthCorpid() string {
	return this.AuthCorpid
}
func (this *OapiServiceSetCorpIpwhitelistRequest) SetIpWhitelist(ipWhitelist2 []string) {
	this.IpWhitelist = ipWhitelist2
}
func (this *OapiServiceSetCorpIpwhitelistRequest) GetIpWhitelist() []string {
	return this.IpWhitelist
}
func (this *OapiServiceSetCorpIpwhitelistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.set_corp_ipwhitelist"
}
func (this *OapiServiceSetCorpIpwhitelistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceSetCorpIpwhitelistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceSetCorpIpwhitelistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceSetCorpIpwhitelistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceSetCorpIpwhitelistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["auth_corpid"] = this.AuthCorpid
	txtParams["ip_whitelist"] = this.IpWhitelist
	return txtParams
}
func (this *OapiServiceSetCorpIpwhitelistRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiServiceSetCorpIpwhitelistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceSetCorpIpwhitelistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceSetCorpIpwhitelistResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
