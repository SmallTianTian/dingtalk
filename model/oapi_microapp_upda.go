package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappUpdateRequest() *OapiMicroappUpdateRequest {
	return &OapiMicroappUpdateRequest{}
}

type OapiMicroappUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappUpdateResponse
	AgentId         int64
	AppDesc         string
	AppIcon         string
	AppName         string
	HomepageUrl     string
	OmpLink         string
	PcHomepageUrl   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMicroappUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappUpdateRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappUpdateRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappUpdateRequest) SetAppDesc(appDesc2 string) {
	this.AppDesc = appDesc2
}
func (this *OapiMicroappUpdateRequest) GetAppDesc() string {
	return this.AppDesc
}
func (this *OapiMicroappUpdateRequest) SetAppIcon(appIcon2 string) {
	this.AppIcon = appIcon2
}
func (this *OapiMicroappUpdateRequest) GetAppIcon() string {
	return this.AppIcon
}
func (this *OapiMicroappUpdateRequest) SetAppName(appName2 string) {
	this.AppName = appName2
}
func (this *OapiMicroappUpdateRequest) GetAppName() string {
	return this.AppName
}
func (this *OapiMicroappUpdateRequest) SetHomepageUrl(homepageUrl2 string) {
	this.HomepageUrl = homepageUrl2
}
func (this *OapiMicroappUpdateRequest) GetHomepageUrl() string {
	return this.HomepageUrl
}
func (this *OapiMicroappUpdateRequest) SetOmpLink(ompLink2 string) {
	this.OmpLink = ompLink2
}
func (this *OapiMicroappUpdateRequest) GetOmpLink() string {
	return this.OmpLink
}
func (this *OapiMicroappUpdateRequest) SetPcHomepageUrl(pcHomepageUrl2 string) {
	this.PcHomepageUrl = pcHomepageUrl2
}
func (this *OapiMicroappUpdateRequest) GetPcHomepageUrl() string {
	return this.PcHomepageUrl
}
func (this *OapiMicroappUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.update"
}
func (this *OapiMicroappUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentId"] = this.AgentId
	txtParams["appDesc"] = this.AppDesc
	txtParams["appIcon"] = this.AppIcon
	txtParams["appName"] = this.AppName
	txtParams["homepageUrl"] = this.HomepageUrl
	txtParams["ompLink"] = this.OmpLink
	txtParams["pcHomepageUrl"] = this.PcHomepageUrl
	return txtParams
}
func (this *OapiMicroappUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappUpdateResponse struct {
	taobao.TaobaoResponse
	Agentid int64  `json:"agentid,omitempty"`
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
