package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappCreateRequest() *OapiMicroappCreateRequest {
	return &OapiMicroappCreateRequest{}
}

type OapiMicroappCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappCreateResponse
	AppDesc         string
	AppIcon         string
	AppName         string
	HomepageUrl     string
	OmpLink         string
	PcHomepageUrl   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMicroappCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappCreateRequest) SetAppDesc(appDesc2 string) {
	this.AppDesc = appDesc2
}
func (this *OapiMicroappCreateRequest) GetAppDesc() string {
	return this.AppDesc
}
func (this *OapiMicroappCreateRequest) SetAppIcon(appIcon2 string) {
	this.AppIcon = appIcon2
}
func (this *OapiMicroappCreateRequest) GetAppIcon() string {
	return this.AppIcon
}
func (this *OapiMicroappCreateRequest) SetAppName(appName2 string) {
	this.AppName = appName2
}
func (this *OapiMicroappCreateRequest) GetAppName() string {
	return this.AppName
}
func (this *OapiMicroappCreateRequest) SetHomepageUrl(homepageUrl2 string) {
	this.HomepageUrl = homepageUrl2
}
func (this *OapiMicroappCreateRequest) GetHomepageUrl() string {
	return this.HomepageUrl
}
func (this *OapiMicroappCreateRequest) SetOmpLink(ompLink2 string) {
	this.OmpLink = ompLink2
}
func (this *OapiMicroappCreateRequest) GetOmpLink() string {
	return this.OmpLink
}
func (this *OapiMicroappCreateRequest) SetPcHomepageUrl(pcHomepageUrl2 string) {
	this.PcHomepageUrl = pcHomepageUrl2
}
func (this *OapiMicroappCreateRequest) GetPcHomepageUrl() string {
	return this.PcHomepageUrl
}
func (this *OapiMicroappCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.create"
}
func (this *OapiMicroappCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["appDesc"] = this.AppDesc
	txtParams["appIcon"] = this.AppIcon
	txtParams["appName"] = this.AppName
	txtParams["homepageUrl"] = this.HomepageUrl
	txtParams["ompLink"] = this.OmpLink
	txtParams["pcHomepageUrl"] = this.PcHomepageUrl
	return txtParams
}
func (this *OapiMicroappCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappCreateResponse struct {
	taobao.TaobaoResponse
	Agentid int64 `json:"agentid,omitempty"`
}
