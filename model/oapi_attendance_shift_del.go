package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceShiftDeleteRequest() *OapiAttendanceShiftDeleteRequest {
	return &OapiAttendanceShiftDeleteRequest{}
}

type OapiAttendanceShiftDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceShiftDeleteResponse
	OpUserId        string
	ShiftId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceShiftDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceShiftDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceShiftDeleteRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceShiftDeleteRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceShiftDeleteRequest) SetShiftId(shiftId2 int64) {
	this.ShiftId = shiftId2
}
func (this *OapiAttendanceShiftDeleteRequest) GetShiftId() int64 {
	return this.ShiftId
}
func (this *OapiAttendanceShiftDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.shift.delete"
}
func (this *OapiAttendanceShiftDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceShiftDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceShiftDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceShiftDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceShiftDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_user_id"] = this.OpUserId
	txtParams["shift_id"] = this.ShiftId
	return txtParams
}
func (this *OapiAttendanceShiftDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceShiftDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceShiftDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceShiftDeleteResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
