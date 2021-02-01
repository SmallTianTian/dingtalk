package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceShiftQueryRequest() *OapiAttendanceShiftQueryRequest {
	return &OapiAttendanceShiftQueryRequest{}
}

type OapiAttendanceShiftQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceShiftQueryResponse
	OpUserId        string
	ShiftId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceShiftQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceShiftQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceShiftQueryRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceShiftQueryRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceShiftQueryRequest) SetShiftId(shiftId2 int64) {
	this.ShiftId = shiftId2
}
func (this *OapiAttendanceShiftQueryRequest) GetShiftId() int64 {
	return this.ShiftId
}
func (this *OapiAttendanceShiftQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.shift.query"
}
func (this *OapiAttendanceShiftQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceShiftQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceShiftQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceShiftQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceShiftQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_user_id"] = this.OpUserId
	txtParams["shift_id"] = this.ShiftId
	return txtParams
}
func (this *OapiAttendanceShiftQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceShiftQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceShiftQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceShiftQueryResponse struct {
	taobao.TaobaoResponse
	Result  TopShiftVo `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
