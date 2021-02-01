package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceApproveScheduleSwitchRequest() *OapiAttendanceApproveScheduleSwitchRequest {
	return &OapiAttendanceApproveScheduleSwitchRequest{}
}

type OapiAttendanceApproveScheduleSwitchRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                OapiAttendanceApproveScheduleSwitchResponse
	ApplyShiftId        int64
	ApplyUserid         string
	ApproveId           string
	RebackApplyShiftId  int64
	RebackDate          string
	RebackTargetShiftId int64
	SwitchDate          string
	TargetShiftId       int64
	TargetUserid        string
	TopHttpMethod       string
	TopResponseType     string
	Userid              string
}

func (this *OapiAttendanceApproveScheduleSwitchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetApplyShiftId(applyShiftId2 int64) {
	this.ApplyShiftId = applyShiftId2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetApplyShiftId() int64 {
	return this.ApplyShiftId
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetApplyUserid(applyUserid2 string) {
	this.ApplyUserid = applyUserid2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetApplyUserid() string {
	return this.ApplyUserid
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetApproveId(approveId2 string) {
	this.ApproveId = approveId2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetApproveId() string {
	return this.ApproveId
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetRebackApplyShiftId(rebackApplyShiftId2 int64) {
	this.RebackApplyShiftId = rebackApplyShiftId2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetRebackApplyShiftId() int64 {
	return this.RebackApplyShiftId
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetRebackDate(rebackDate2 string) {
	this.RebackDate = rebackDate2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetRebackDate() string {
	return this.RebackDate
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetRebackTargetShiftId(rebackTargetShiftId2 int64) {
	this.RebackTargetShiftId = rebackTargetShiftId2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetRebackTargetShiftId() int64 {
	return this.RebackTargetShiftId
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetSwitchDate(switchDate2 string) {
	this.SwitchDate = switchDate2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetSwitchDate() string {
	return this.SwitchDate
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetTargetShiftId(targetShiftId2 int64) {
	this.TargetShiftId = targetShiftId2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetTargetShiftId() int64 {
	return this.TargetShiftId
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetTargetUserid(targetUserid2 string) {
	this.TargetUserid = targetUserid2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetTargetUserid() string {
	return this.TargetUserid
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.approve.schedule.switch"
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["apply_shift_id"] = this.ApplyShiftId
	txtParams["apply_userid"] = this.ApplyUserid
	txtParams["approve_id"] = this.ApproveId
	txtParams["reback_apply_shift_id"] = this.RebackApplyShiftId
	txtParams["reback_date"] = this.RebackDate
	txtParams["reback_target_shift_id"] = this.RebackTargetShiftId
	txtParams["switch_date"] = this.SwitchDate
	txtParams["target_shift_id"] = this.TargetShiftId
	txtParams["target_userid"] = this.TargetUserid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ApplyUserid, "applyUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceApproveScheduleSwitchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceApproveScheduleSwitchResponse struct {
	taobao.TaobaoResponse
}
