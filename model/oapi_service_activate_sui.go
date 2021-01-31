package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiServiceActivateSuiteRequest() *OapiServiceActivateSuiteRequest {
	return &OapiServiceActivateSuiteRequest{}
}

type OapiServiceActivateSuiteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceActivateSuiteResponse
	AuthCorpid      string
	PermanentCode   string
	SuiteKey        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceActivateSuiteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceActivateSuiteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceActivateSuiteRequest) SetAuthCorpid(authCorpid2 string) {
	this.AuthCorpid = authCorpid2
}
func (this *OapiServiceActivateSuiteRequest) GetAuthCorpid() string {
	return this.AuthCorpid
}
func (this *OapiServiceActivateSuiteRequest) SetPermanentCode(permanentCode2 string) {
	this.PermanentCode = permanentCode2
}
func (this *OapiServiceActivateSuiteRequest) GetPermanentCode() string {
	return this.PermanentCode
}
func (this *OapiServiceActivateSuiteRequest) SetSuiteKey(suiteKey2 string) {
	this.SuiteKey = suiteKey2
}
func (this *OapiServiceActivateSuiteRequest) GetSuiteKey() string {
	return this.SuiteKey
}
func (this *OapiServiceActivateSuiteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.activate_suite"
}
func (this *OapiServiceActivateSuiteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceActivateSuiteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceActivateSuiteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceActivateSuiteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceActivateSuiteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["auth_corpid"] = this.AuthCorpid
	txtParams["permanent_code"] = this.PermanentCode
	txtParams["suite_key"] = this.SuiteKey
	return txtParams
}
func (this *OapiServiceActivateSuiteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiServiceActivateSuiteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceActivateSuiteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceActivateSuiteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
