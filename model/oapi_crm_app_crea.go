package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmAppCreateRequest() *OapiCrmAppCreateRequest {
	return &OapiCrmAppCreateRequest{}
}

type OapiCrmAppCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmAppCreateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmAppCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmAppCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmAppCreateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiCrmAppCreateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiCrmAppCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.app.create"
}
func (this *OapiCrmAppCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmAppCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmAppCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmAppCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmAppCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiCrmAppCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmAppCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmAppCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CreateCrmMicroAppRequest struct {
	AppDesc     string `json:"app_desc,omitempty"`
	AppIcon     string `json:"app_icon,omitempty"`
	BizKey      string `json:"biz_key,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
	IpWhiteList string `json:"ip_white_list,omitempty"`
	Name        string `json:"name,omitempty"`
	PcHomepage  string `json:"pc_homepage,omitempty"`
}
type OapiCrmAppCreateResponse struct {
	taobao.TaobaoResponse
	Result GetCrmMicroAppResponse `json:"result,omitempty"`
}
type GetCrmMicroAppResponse struct {
	Agentid     int64  `json:"agentid,omitempty"`
	AppDesc     string `json:"app_desc,omitempty"`
	AppIcon     string `json:"app_icon,omitempty"`
	AppKey      string `json:"app_key,omitempty"`
	AppSecret   string `json:"app_secret,omitempty"`
	BizKey      string `json:"biz_key,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
	IpWhiteList string `json:"ip_white_list,omitempty"`
	Name        string `json:"name,omitempty"`
	PcHomepage  string `json:"pc_homepage,omitempty"`
}
