package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceApproveFinishRequest() *OapiAttendanceApproveFinishRequest {
	return &OapiAttendanceApproveFinishRequest{}
}

type OapiAttendanceApproveFinishRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp             OapiAttendanceApproveFinishResponse
	ApproveId        string
	BizType          int64
	CalculateModel   int64
	DurationUnit     string
	FromTime         string
	JumpUrl          string
	OvertimeDuration string
	OvertimeToMore   int64
	SubType          string
	TagName          string
	ToTime           string
	TopHttpMethod    string
	TopResponseType  string
	Userid           string
}

func (this *OapiAttendanceApproveFinishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceApproveFinishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceApproveFinishRequest) SetApproveId(approveId2 string) {
	this.ApproveId = approveId2
}
func (this *OapiAttendanceApproveFinishRequest) GetApproveId() string {
	return this.ApproveId
}
func (this *OapiAttendanceApproveFinishRequest) SetBizType(bizType2 int64) {
	this.BizType = bizType2
}
func (this *OapiAttendanceApproveFinishRequest) GetBizType() int64 {
	return this.BizType
}
func (this *OapiAttendanceApproveFinishRequest) SetCalculateModel(calculateModel2 int64) {
	this.CalculateModel = calculateModel2
}
func (this *OapiAttendanceApproveFinishRequest) GetCalculateModel() int64 {
	return this.CalculateModel
}
func (this *OapiAttendanceApproveFinishRequest) SetDurationUnit(durationUnit2 string) {
	this.DurationUnit = durationUnit2
}
func (this *OapiAttendanceApproveFinishRequest) GetDurationUnit() string {
	return this.DurationUnit
}
func (this *OapiAttendanceApproveFinishRequest) SetFromTime(fromTime2 string) {
	this.FromTime = fromTime2
}
func (this *OapiAttendanceApproveFinishRequest) GetFromTime() string {
	return this.FromTime
}
func (this *OapiAttendanceApproveFinishRequest) SetJumpUrl(jumpUrl2 string) {
	this.JumpUrl = jumpUrl2
}
func (this *OapiAttendanceApproveFinishRequest) GetJumpUrl() string {
	return this.JumpUrl
}
func (this *OapiAttendanceApproveFinishRequest) SetOvertimeDuration(overtimeDuration2 string) {
	this.OvertimeDuration = overtimeDuration2
}
func (this *OapiAttendanceApproveFinishRequest) GetOvertimeDuration() string {
	return this.OvertimeDuration
}
func (this *OapiAttendanceApproveFinishRequest) SetOvertimeToMore(overtimeToMore2 int64) {
	this.OvertimeToMore = overtimeToMore2
}
func (this *OapiAttendanceApproveFinishRequest) GetOvertimeToMore() int64 {
	return this.OvertimeToMore
}
func (this *OapiAttendanceApproveFinishRequest) SetSubType(subType2 string) {
	this.SubType = subType2
}
func (this *OapiAttendanceApproveFinishRequest) GetSubType() string {
	return this.SubType
}
func (this *OapiAttendanceApproveFinishRequest) SetTagName(tagName2 string) {
	this.TagName = tagName2
}
func (this *OapiAttendanceApproveFinishRequest) GetTagName() string {
	return this.TagName
}
func (this *OapiAttendanceApproveFinishRequest) SetToTime(toTime2 string) {
	this.ToTime = toTime2
}
func (this *OapiAttendanceApproveFinishRequest) GetToTime() string {
	return this.ToTime
}
func (this *OapiAttendanceApproveFinishRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceApproveFinishRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceApproveFinishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.approve.finish"
}
func (this *OapiAttendanceApproveFinishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceApproveFinishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceApproveFinishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceApproveFinishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceApproveFinishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["approve_id"] = this.ApproveId
	txtParams["biz_type"] = this.BizType
	txtParams["calculate_model"] = this.CalculateModel
	txtParams["duration_unit"] = this.DurationUnit
	txtParams["from_time"] = this.FromTime
	txtParams["jump_url"] = this.JumpUrl
	txtParams["overtime_duration"] = this.OvertimeDuration
	txtParams["overtime_to_more"] = this.OvertimeToMore
	txtParams["sub_type"] = this.SubType
	txtParams["tag_name"] = this.TagName
	txtParams["to_time"] = this.ToTime
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceApproveFinishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxLength(this.ApproveId, 100, "approveId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceApproveFinishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceApproveFinishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceApproveFinishResponse struct {
	taobao.TaobaoResponse
	Result TopDurationVo `json:"result,omitempty"`
}
