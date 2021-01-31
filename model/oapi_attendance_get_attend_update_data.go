package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGetAttendUpdateDataRequest() *OapiAttendanceGetAttendUpdateDataRequest {
	return &OapiAttendanceGetAttendUpdateDataRequest{}
}

type OapiAttendanceGetAttendUpdateDataRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGetAttendUpdateDataResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
	WorkDate        time.Time
}

func (this *OapiAttendanceGetAttendUpdateDataRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) SetWorkDate(workDate2 time.Time) {
	this.WorkDate = workDate2
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) GetWorkDate() time.Time {
	return this.WorkDate
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.getAttendUpdateData"
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["work_date"] = this.WorkDate
	return txtParams
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGetAttendUpdateDataRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGetAttendUpdateDataResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  AtCheckInfoForOpenVo `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type AtAttendanceResultForOpenVo struct {
	CheckType      string    `json:"check_type,omitempty"`
	ClassId        int64     `json:"class_id,omitempty"`
	GroupId        int64     `json:"group_id,omitempty"`
	LocationMethod string    `json:"location_method,omitempty"`
	LocationResult string    `json:"location_result,omitempty"`
	OutsideRemark  string    `json:"outside_remark,omitempty"`
	PlanCheckTime  time.Time `json:"plan_check_time,omitempty"`
	PlanId         int64     `json:"plan_id,omitempty"`
	ProcInstId     string    `json:"procInst_id,omitempty"`
	RecordId       int64     `json:"record_id,omitempty"`
	SourceType     string    `json:"source_type,omitempty"`
	TimeResult     string    `json:"time_result,omitempty"`
	UserAddress    string    `json:"user_address,omitempty"`
	UserCheckTime  time.Time `json:"user_check_time,omitempty"`
}
type AtApproveForOpenVo struct {
	BeginTime    time.Time `json:"begin_time,omitempty"`
	BizType      int64     `json:"biz_type,omitempty"`
	Duration     string    `json:"duration,omitempty"`
	DurationUnit string    `json:"duration_unit,omitempty"`
	EndTime      time.Time `json:"end_time,omitempty"`
	GmtFinished  time.Time `json:"gmt_finished,omitempty"`
	ProcInstId   string    `json:"procInst_id,omitempty"`
	SubType      string    `json:"sub_type,omitempty"`
	TagName      string    `json:"tag_name,omitempty"`
}
type AtAttendanceRecordForOpenVo struct {
	BaseAccuracy  string    `json:"base_accuracy,omitempty"`
	BaseAddress   string    `json:"base_address,omitempty"`
	RecordId      int64     `json:"record_id,omitempty"`
	SourceType    string    `json:"source_type,omitempty"`
	UserAccuracy  string    `json:"user_accuracy,omitempty"`
	UserCheckTime time.Time `json:"user_check_time,omitempty"`
	UserLatitude  string    `json:"user_latitude,omitempty"`
	UserLongitude string    `json:"user_longitude,omitempty"`
	UserMacAddr   string    `json:"user_mac_addr,omitempty"`
	UserSsid      string    `json:"user_ssid,omitempty"`
	ValidMatched  bool      `json:"valid_matched,omitempty"`
}
type AtCheckInfoForOpenVo struct {
	ApproveList          AtApproveForOpenVo          `json:"approve_list,omitempty"`
	AttendanceResultList AtAttendanceResultForOpenVo `json:"attendance_result_list,omitempty"`
	CheckRecordList      AtAttendanceRecordForOpenVo `json:"check_record_list,omitempty"`
	CorpId               string                      `json:"corpId,omitempty"`
	Userid               string                      `json:"userid,omitempty"`
	WorkDate             time.Time                   `json:"work_date,omitempty"`
}
