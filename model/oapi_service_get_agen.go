package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiServiceGetAgentRequest() *OapiServiceGetAgentRequest {
	return &OapiServiceGetAgentRequest{}
}

type OapiServiceGetAgentRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceGetAgentResponse
	Agentid         string
	AuthCorpid      string
	PermanentCode   string
	SuiteKey        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceGetAgentRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceGetAgentRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceGetAgentRequest) SetAgentid(agentid2 string) {
	this.Agentid = agentid2
}
func (this *OapiServiceGetAgentRequest) GetAgentid() string {
	return this.Agentid
}
func (this *OapiServiceGetAgentRequest) SetAuthCorpid(authCorpid2 string) {
	this.AuthCorpid = authCorpid2
}
func (this *OapiServiceGetAgentRequest) GetAuthCorpid() string {
	return this.AuthCorpid
}
func (this *OapiServiceGetAgentRequest) SetPermanentCode(permanentCode2 string) {
	this.PermanentCode = permanentCode2
}
func (this *OapiServiceGetAgentRequest) GetPermanentCode() string {
	return this.PermanentCode
}
func (this *OapiServiceGetAgentRequest) SetSuiteKey(suiteKey2 string) {
	this.SuiteKey = suiteKey2
}
func (this *OapiServiceGetAgentRequest) GetSuiteKey() string {
	return this.SuiteKey
}
func (this *OapiServiceGetAgentRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.get_agent"
}
func (this *OapiServiceGetAgentRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceGetAgentRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceGetAgentRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceGetAgentRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceGetAgentRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["auth_corpid"] = this.AuthCorpid
	txtParams["permanent_code"] = this.PermanentCode
	txtParams["suite_key"] = this.SuiteKey
	return txtParams
}
func (this *OapiServiceGetAgentRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiServiceGetAgentRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceGetAgentRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceGetAgentResponse struct {
	taobao.TaobaoResponse
	Agentid     int64  `json:"agentid,omitempty"`
	Close       int64  `json:"close,omitempty"`
	Description string `json:"description,omitempty"`

	LogoUrl string `json:"logo_url,omitempty"`
	Name    string `json:"name,omitempty"`
}
