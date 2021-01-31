package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceIsopensmartreportRequest() *OapiAttendanceIsopensmartreportRequest {
	return &OapiAttendanceIsopensmartreportRequest{}
}

type OapiAttendanceIsopensmartreportRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceIsopensmartreportResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceIsopensmartreportRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceIsopensmartreportRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceIsopensmartreportRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.isopensmartreport"
}
func (this *OapiAttendanceIsopensmartreportRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceIsopensmartreportRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceIsopensmartreportRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceIsopensmartreportRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceIsopensmartreportRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiAttendanceIsopensmartreportRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceIsopensmartreportRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceIsopensmartreportRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceIsopensmartreportResponse struct {
	taobao.TaobaoResponse
	Result IsOpenSmartReportForTopVo `json:"result,omitempty"`
}
type IsOpenSmartReportForTopVo struct {
	SmartReport bool `json:"smart_report,omitempty"`
}
