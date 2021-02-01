package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiStatisticsDetailsRequest() *OapiStatisticsDetailsRequest {
	return &OapiStatisticsDetailsRequest{}
}

type OapiStatisticsDetailsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiStatisticsDetailsResponse
	TopHttpMethod   string
	TopResponseType string
	Type            string
}

func (this *OapiStatisticsDetailsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiStatisticsDetailsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiStatisticsDetailsRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiStatisticsDetailsRequest) GetType() string {
	return this.Type
}
func (this *OapiStatisticsDetailsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.statistics.details"
}
func (this *OapiStatisticsDetailsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiStatisticsDetailsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiStatisticsDetailsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiStatisticsDetailsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiStatisticsDetailsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiStatisticsDetailsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiStatisticsDetailsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiStatisticsDetailsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiStatisticsDetailsResponse struct {
	taobao.TaobaoResponse
	Result []EnterpriseDataVo `json:"result,omitempty"`
}
type EnterpriseDataVo struct {
	Date         string `json:"date,omitempty"`
	ReturnFields string `json:"return_fields,omitempty"`
	Type         string `json:"type,omitempty"`
	Url          string `json:"url,omitempty"`
}
