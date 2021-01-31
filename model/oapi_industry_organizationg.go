package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiIndustryOrganizationGetRequest() *OapiIndustryOrganizationGetRequest {
	return &OapiIndustryOrganizationGetRequest{}
}

type OapiIndustryOrganizationGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiIndustryOrganizationGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiIndustryOrganizationGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIndustryOrganizationGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIndustryOrganizationGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.industry.organization.get"
}
func (this *OapiIndustryOrganizationGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIndustryOrganizationGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIndustryOrganizationGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIndustryOrganizationGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIndustryOrganizationGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiIndustryOrganizationGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiIndustryOrganizationGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIndustryOrganizationGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiIndustryOrganizationGetResponse struct {
	taobao.TaobaoResponse
	Result  OpenIndustryOrgInfo `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
type OpenIndustryOrgInfo struct {
	RegionId       string `json:"region_id,omitempty"`
	RegionLocation string `json:"region_location,omitempty"`
	RegionType     string `json:"region_type,omitempty"`
}
