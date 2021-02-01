package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmAppUpdateRequest() *OapiCrmAppUpdateRequest {
	return &OapiCrmAppUpdateRequest{}
}

type OapiCrmAppUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmAppUpdateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmAppUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmAppUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmAppUpdateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiCrmAppUpdateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiCrmAppUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.app.update"
}
func (this *OapiCrmAppUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmAppUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmAppUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmAppUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmAppUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiCrmAppUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmAppUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmAppUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UpdateCrmMicroAppRequest struct {
	AppDesc     string `json:"app_desc,omitempty"`
	AppIcon     string `json:"app_icon,omitempty"`
	BizKey      string `json:"biz_key,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
	IpWhiteList string `json:"ip_white_list,omitempty"`
	Name        string `json:"name,omitempty"`
	PcHomepage  string `json:"pc_homepage,omitempty"`
}
type OapiCrmAppUpdateResponse struct {
	taobao.TaobaoResponse
	Result GetCrmMicroAppResponse `json:"result,omitempty"`
}
