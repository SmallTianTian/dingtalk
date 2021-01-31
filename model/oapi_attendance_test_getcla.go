package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceTestGetclassRequest() *OapiAttendanceTestGetclassRequest {
	return &OapiAttendanceTestGetclassRequest{}
}

type OapiAttendanceTestGetclassRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceTestGetclassResponse
	ClassId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceTestGetclassRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceTestGetclassRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceTestGetclassRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiAttendanceTestGetclassRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiAttendanceTestGetclassRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.test.getclass"
}
func (this *OapiAttendanceTestGetclassRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceTestGetclassRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceTestGetclassRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceTestGetclassRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceTestGetclassRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["classId"] = this.ClassId
	return txtParams
}
func (this *OapiAttendanceTestGetclassRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceTestGetclassRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceTestGetclassRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceTestGetclassResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
