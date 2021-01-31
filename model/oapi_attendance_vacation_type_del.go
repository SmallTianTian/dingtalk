package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceVacationTypeDeleteRequest() *OapiAttendanceVacationTypeDeleteRequest {
	return &OapiAttendanceVacationTypeDeleteRequest{}
}

type OapiAttendanceVacationTypeDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceVacationTypeDeleteResponse
	LeaveCode       string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceVacationTypeDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceVacationTypeDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceVacationTypeDeleteRequest) SetLeaveCode(leaveCode2 string) {
	this.LeaveCode = leaveCode2
}
func (this *OapiAttendanceVacationTypeDeleteRequest) GetLeaveCode() string {
	return this.LeaveCode
}
func (this *OapiAttendanceVacationTypeDeleteRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceVacationTypeDeleteRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceVacationTypeDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.vacation.type.delete"
}
func (this *OapiAttendanceVacationTypeDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceVacationTypeDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceVacationTypeDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceVacationTypeDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceVacationTypeDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["leave_code"] = this.LeaveCode
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiAttendanceVacationTypeDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.LeaveCode, "leaveCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceVacationTypeDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceVacationTypeDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceVacationTypeDeleteResponse struct {
	taobao.TaobaoResponse
	Result  LeaveTypeVo `json:"result,omitempty"`
	Success bool        `json:"success,omitempty"`
}
