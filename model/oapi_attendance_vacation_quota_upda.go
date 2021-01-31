package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceVacationQuotaUpdateRequest() *OapiAttendanceVacationQuotaUpdateRequest {
	return &OapiAttendanceVacationQuotaUpdateRequest{}
}

type OapiAttendanceVacationQuotaUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceVacationQuotaUpdateResponse
	LeaveQuotas     string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceVacationQuotaUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) SetLeaveQuotas(leaveQuotas2 string) {
	this.LeaveQuotas = leaveQuotas2
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) GetLeaveQuotas() string {
	return this.LeaveQuotas
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.vacation.quota.update"
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["leave_quotas"] = this.LeaveQuotas
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.LeaveQuotas, 100, "leaveQuotas"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceVacationQuotaUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceVacationQuotaUpdateResponse struct {
	taobao.TaobaoResponse
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
