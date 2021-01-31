package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiAttendanceScheduleResultListbyidsRequest() *OapiAttendanceScheduleResultListbyidsRequest {
	return &OapiAttendanceScheduleResultListbyidsRequest{}
}

type OapiAttendanceScheduleResultListbyidsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceScheduleResultListbyidsResponse
	OpUserId        string
	ScheduleIds     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceScheduleResultListbyidsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) SetScheduleIds(scheduleIds2 string) {
	this.ScheduleIds = scheduleIds2
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) GetScheduleIds() string {
	return this.ScheduleIds
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.schedule.result.listbyids"
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_user_id"] = this.OpUserId
	txtParams["schedule_ids"] = this.ScheduleIds
	return txtParams
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceScheduleResultListbyidsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceScheduleResultListbyidsResponse struct {
	taobao.TaobaoResponse
	Errcode int64                 `json:"errcode,omitempty"`
	Errmsg  string                `json:"errmsg,omitempty"`
	Result  []TopScheduleResultVo `json:"result,omitempty"`
	Success bool                  `json:"success,omitempty"`
}
type TopScheduleResultVo struct {
	BaseCheckTime  time.Time `json:"base_check_time,omitempty"`
	CheckType      string    `json:"check_type,omitempty"`
	CorpId         string    `json:"corp_id,omitempty"`
	GmtCreate      time.Time `json:"gmt_create,omitempty"`
	GmtModified    time.Time `json:"gmt_modified,omitempty"`
	GroupId        int64     `json:"group_id,omitempty"`
	Id             int64     `json:"id,omitempty"`
	IsLegal        string    `json:"is_legal,omitempty"`
	LocationResult string    `json:"location_result,omitempty"`
	PlanCheckTime  time.Time `json:"plan_check_time,omitempty"`
	RecordId       int64     `json:"record_id,omitempty"`
	ScheduleId     int64     `json:"schedule_id,omitempty"`
	TimeResult     string    `json:"time_result,omitempty"`
	UserCheckTime  time.Time `json:"user_check_time,omitempty"`
	UserId         string    `json:"user_id,omitempty"`
	WorkDate       time.Time `json:"work_date,omitempty"`
}
