package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiIndustryPackGetRequest() *OapiIndustryPackGetRequest {
	return &OapiIndustryPackGetRequest{}
}

type OapiIndustryPackGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiIndustryPackGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiIndustryPackGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIndustryPackGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIndustryPackGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.industry.pack.get"
}
func (this *OapiIndustryPackGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIndustryPackGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIndustryPackGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIndustryPackGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIndustryPackGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiIndustryPackGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiIndustryPackGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIndustryPackGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiIndustryPackGetResponse struct {
	taobao.TaobaoResponse
	Result  []PackageDto `json:"result,omitempty"`
	Success bool         `json:"success,omitempty"`
}
type PackageDto struct {
	Id          int64  `json:"id,omitempty"`
	PackageName string `json:"package_name,omitempty"`
}
