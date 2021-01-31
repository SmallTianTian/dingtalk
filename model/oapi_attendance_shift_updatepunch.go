package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceShiftUpdatepunchesRequest() *OapiAttendanceShiftUpdatepunchesRequest {
	return &OapiAttendanceShiftUpdatepunchesRequest{}
}

type OapiAttendanceShiftUpdatepunchesRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceShiftUpdatepunchesResponse
	OpUserId        string
	Punches         string
	ShiftId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceShiftUpdatepunchesRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) SetPunches(punches2 string) {
	this.Punches = punches2
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) GetPunches() string {
	return this.Punches
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) SetShiftId(shiftId2 int64) {
	this.ShiftId = shiftId2
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) GetShiftId() int64 {
	return this.ShiftId
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.shift.updatepunches"
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_user_id"] = this.OpUserId
	txtParams["punches"] = this.Punches
	txtParams["shift_id"] = this.ShiftId
	return txtParams
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceShiftUpdatepunchesRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopPunchVO struct {
	FreeCheck bool  `json:"free_check,omitempty"`
	Id        int64 `json:"id,omitempty"`
}
type OapiAttendanceShiftUpdatepunchesResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
