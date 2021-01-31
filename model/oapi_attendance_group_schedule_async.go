package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupScheduleAsyncRequest() *OapiAttendanceGroupScheduleAsyncRequest {
	return &OapiAttendanceGroupScheduleAsyncRequest{}
}

type OapiAttendanceGroupScheduleAsyncRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupScheduleAsyncResponse
	GroupId         int64
	OpUserId        string
	Schedules       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupScheduleAsyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) SetSchedules(schedules2 string) {
	this.Schedules = schedules2
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) GetSchedules() string {
	return this.Schedules
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.schedule.async"
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_id"] = this.GroupId
	txtParams["op_user_id"] = this.OpUserId
	txtParams["schedules"] = this.Schedules
	return txtParams
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupId, "groupId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupScheduleAsyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopScheduleParam struct {
	IsRest   bool   `json:"is_rest,omitempty"`
	ShiftId  int64  `json:"shift_id,omitempty"`
	Userid   string `json:"userid,omitempty"`
	WorkDate int64  `json:"work_date,omitempty"`
}
type OapiAttendanceGroupScheduleAsyncResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
