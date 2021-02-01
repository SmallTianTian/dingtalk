package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiReportGetunreadcountRequest() *OapiReportGetunreadcountRequest {
	return &OapiReportGetunreadcountRequest{}
}

type OapiReportGetunreadcountRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiReportGetunreadcountResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiReportGetunreadcountRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportGetunreadcountRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportGetunreadcountRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiReportGetunreadcountRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiReportGetunreadcountRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.getunreadcount"
}
func (this *OapiReportGetunreadcountRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportGetunreadcountRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportGetunreadcountRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportGetunreadcountRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportGetunreadcountRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiReportGetunreadcountRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiReportGetunreadcountRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportGetunreadcountRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportGetunreadcountResponse struct {
	taobao.TaobaoResponse
	Count int64 `json:"count,omitempty"`
}
