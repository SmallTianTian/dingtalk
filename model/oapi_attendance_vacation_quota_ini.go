package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceVacationQuotaInitRequest() *OapiAttendanceVacationQuotaInitRequest {
	return &OapiAttendanceVacationQuotaInitRequest{}
}

type OapiAttendanceVacationQuotaInitRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceVacationQuotaInitResponse
	LeaveQuotas     string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceVacationQuotaInitRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceVacationQuotaInitRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceVacationQuotaInitRequest) SetLeaveQuotas(leaveQuotas2 string) {
	this.LeaveQuotas = leaveQuotas2
}
func (this *OapiAttendanceVacationQuotaInitRequest) GetLeaveQuotas() string {
	return this.LeaveQuotas
}
func (this *OapiAttendanceVacationQuotaInitRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceVacationQuotaInitRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceVacationQuotaInitRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.vacation.quota.init"
}
func (this *OapiAttendanceVacationQuotaInitRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceVacationQuotaInitRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceVacationQuotaInitRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceVacationQuotaInitRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceVacationQuotaInitRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["leave_quotas"] = this.LeaveQuotas
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiAttendanceVacationQuotaInitRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.LeaveQuotas, 100, "leaveQuotas"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceVacationQuotaInitRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceVacationQuotaInitRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type LeaveQuotas struct {
	EndTime         int64  `json:"end_time,omitempty"`
	LeaveCode       string `json:"leave_code,omitempty"`
	QuotaCycle      string `json:"quota_cycle,omitempty"`
	QuotaNumPerDay  int64  `json:"quota_num_per_day,omitempty"`
	QuotaNumPerHour int64  `json:"quota_num_per_hour,omitempty"`
	Reason          string `json:"reason,omitempty"`
	StartTime       int64  `json:"start_time,omitempty"`
	Userid          string `json:"userid,omitempty"`
}
type OapiAttendanceVacationQuotaInitResponse struct {
	taobao.TaobaoResponse
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
type Quota struct {
	LeaveCode  string `json:"leave_code,omitempty"`
	QuotaCycle string `json:"quota_cycle,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
