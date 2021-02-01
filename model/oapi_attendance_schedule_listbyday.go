package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceScheduleListbydayRequest() *OapiAttendanceScheduleListbydayRequest {
	return &OapiAttendanceScheduleListbydayRequest{}
}

type OapiAttendanceScheduleListbydayRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceScheduleListbydayResponse
	DateTime        int64
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
	UserId          string
}

func (this *OapiAttendanceScheduleListbydayRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceScheduleListbydayRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceScheduleListbydayRequest) SetDateTime(dateTime2 int64) {
	this.DateTime = dateTime2
}
func (this *OapiAttendanceScheduleListbydayRequest) GetDateTime() int64 {
	return this.DateTime
}
func (this *OapiAttendanceScheduleListbydayRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceScheduleListbydayRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceScheduleListbydayRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *OapiAttendanceScheduleListbydayRequest) GetUserId() string {
	return this.UserId
}
func (this *OapiAttendanceScheduleListbydayRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.schedule.listbyday"
}
func (this *OapiAttendanceScheduleListbydayRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceScheduleListbydayRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceScheduleListbydayRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceScheduleListbydayRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceScheduleListbydayRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["date_time"] = this.DateTime
	txtParams["op_user_id"] = this.OpUserId
	txtParams["user_id"] = this.UserId
	return txtParams
}
func (this *OapiAttendanceScheduleListbydayRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DateTime, "dateTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceScheduleListbydayRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceScheduleListbydayRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceScheduleListbydayResponse struct {
	taobao.TaobaoResponse
	Result  []TopScheduleVo `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
type TopScheduleVo struct {
	ApproveBizType int64     `json:"approve_biz_type,omitempty"`
	ApproveId      int64     `json:"approve_id,omitempty"`
	ApproveTagName string    `json:"approve_tag_name,omitempty"`
	ApproveType    string    `json:"approve_type,omitempty"`
	BaseCheckTime  time.Time `json:"base_check_time,omitempty"`
	CheckBeginTime time.Time `json:"check_begin_time,omitempty"`
	CheckDateTime  time.Time `json:"check_date_time,omitempty"`
	CheckEndTime   time.Time `json:"check_end_time,omitempty"`
	CheckStatus    string    `json:"check_status,omitempty"`
	CheckType      string    `json:"check_type,omitempty"`
	ClassId        int64     `json:"class_id,omitempty"`
	ClassName      string    `json:"class_name,omitempty"`
	ClassSettingId int64     `json:"class_setting_id,omitempty"`
	CorpId         string    `json:"corp_id,omitempty"`
	Features       string    `json:"features,omitempty"`
	GmtCreate      time.Time `json:"gmt_create,omitempty"`
	GmtModified    time.Time `json:"gmt_modified,omitempty"`
	GroupId        int64     `json:"group_id,omitempty"`
	Id             int64     `json:"id,omitempty"`
	IsRest         string    `json:"is_rest,omitempty"`
	PlanCheckTime  time.Time `json:"plan_check_time,omitempty"`
	RealPlanTime   time.Time `json:"real_plan_time,omitempty"`
	UserId         string    `json:"user_id,omitempty"`
	WorkDate       time.Time `json:"work_date,omitempty"`
}
