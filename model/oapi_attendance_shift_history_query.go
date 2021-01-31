package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiAttendanceShiftHistoryQueryRequest() *OapiAttendanceShiftHistoryQueryRequest {
	return &OapiAttendanceShiftHistoryQueryRequest{}
}

type OapiAttendanceShiftHistoryQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceShiftHistoryQueryResponse
	OpUserId        string
	ShiftId         int64
	TopHttpMethod   string
	TopResponseType string
	Version         int64
}

func (this *OapiAttendanceShiftHistoryQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceShiftHistoryQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceShiftHistoryQueryRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceShiftHistoryQueryRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceShiftHistoryQueryRequest) SetShiftId(shiftId2 int64) {
	this.ShiftId = shiftId2
}
func (this *OapiAttendanceShiftHistoryQueryRequest) GetShiftId() int64 {
	return this.ShiftId
}
func (this *OapiAttendanceShiftHistoryQueryRequest) SetVersion(version2 int64) {
	this.Version = version2
}
func (this *OapiAttendanceShiftHistoryQueryRequest) GetVersion() int64 {
	return this.Version
}
func (this *OapiAttendanceShiftHistoryQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.shift.history.query"
}
func (this *OapiAttendanceShiftHistoryQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceShiftHistoryQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceShiftHistoryQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceShiftHistoryQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceShiftHistoryQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_user_id"] = this.OpUserId
	txtParams["shift_id"] = this.ShiftId
	txtParams["version"] = this.Version
	return txtParams
}
func (this *OapiAttendanceShiftHistoryQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceShiftHistoryQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceShiftHistoryQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceShiftHistoryQueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  TopShiftVo `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
type TopShiftSettingVo struct {
	AttendDays      string    `json:"attend_days,omitempty"`
	CorpId          string    `json:"corp_id,omitempty"`
	GmtCreate       time.Time `json:"gmt_create,omitempty"`
	GmtModified     time.Time `json:"gmt_modified,omitempty"`
	Id              int64     `json:"id,omitempty"`
	IsDeleted       string    `json:"is_deleted,omitempty"`
	ShiftId         int64     `json:"shift_id,omitempty"`
	WorkTimeMinutes int64     `json:"work_time_minutes,omitempty"`
}
type TopPunchVo struct {
	Across        int64     `json:"across,omitempty"`
	BeginMin      int64     `json:"begin_min,omitempty"`
	CheckTime     time.Time `json:"check_time,omitempty"`
	CheckType     string    `json:"check_type,omitempty"`
	EndMin        int64     `json:"end_min,omitempty"`
	FreeCheck     bool      `json:"free_check,omitempty"`
	Id            int64     `json:"id,omitempty"`
	PermitMinutes int64     `json:"permit_minutes,omitempty"`
}
type TopRestVo struct {
	Across    int64     `json:"across,omitempty"`
	CheckTime time.Time `json:"check_time,omitempty"`
	CheckType string    `json:"check_type,omitempty"`
	Id        int64     `json:"id,omitempty"`
}
type TopSectionVo struct {
	Id              int64        `json:"id,omitempty"`
	Punches         []TopPunchVo `json:"punches,omitempty"`
	Rests           []TopRestVo  `json:"rests,omitempty"`
	WorkTimeMinutes int64        `json:"work_time_minutes,omitempty"`
}
