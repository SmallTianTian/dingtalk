package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceListRecordRequest() *OapiAttendanceListRecordRequest {
	return &OapiAttendanceListRecordRequest{}
}

type OapiAttendanceListRecordRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceListRecordResponse
	CheckDateFrom   string
	CheckDateTo     string
	IsI18n          bool
	TopHttpMethod   string
	TopResponseType string
	UserIds         []string
}

func (this *OapiAttendanceListRecordRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceListRecordRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceListRecordRequest) SetCheckDateFrom(checkDateFrom2 string) {
	this.CheckDateFrom = checkDateFrom2
}
func (this *OapiAttendanceListRecordRequest) GetCheckDateFrom() string {
	return this.CheckDateFrom
}
func (this *OapiAttendanceListRecordRequest) SetCheckDateTo(checkDateTo2 string) {
	this.CheckDateTo = checkDateTo2
}
func (this *OapiAttendanceListRecordRequest) GetCheckDateTo() string {
	return this.CheckDateTo
}
func (this *OapiAttendanceListRecordRequest) SetIsI18n(isI18n2 bool) {
	this.IsI18n = isI18n2
}
func (this *OapiAttendanceListRecordRequest) GetIsI18n() bool {
	return this.IsI18n
}
func (this *OapiAttendanceListRecordRequest) SetUserIds(userIds2 []string) {
	this.UserIds = userIds2
}
func (this *OapiAttendanceListRecordRequest) GetUserIds() []string {
	return this.UserIds
}
func (this *OapiAttendanceListRecordRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.listRecord"
}
func (this *OapiAttendanceListRecordRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceListRecordRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceListRecordRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceListRecordRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceListRecordRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["checkDateFrom"] = this.CheckDateFrom
	txtParams["checkDateTo"] = this.CheckDateTo
	txtParams["isI18n"] = this.IsI18n
	txtParams["userIds"] = this.UserIds
	return txtParams
}
func (this *OapiAttendanceListRecordRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CheckDateFrom, "checkDateFrom"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceListRecordRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceListRecordRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceListRecordResponse struct {
	taobao.TaobaoResponse

	Recordresult []Recordresult `json:"recordresult,omitempty"`
}
type Recordresult struct {
	ApproveId         int64     `json:"approveId,omitempty"`
	BaseAccuracy      string    `json:"baseAccuracy,omitempty"`
	BaseAddress       string    `json:"baseAddress,omitempty"`
	BaseCheckTime     time.Time `json:"baseCheckTime,omitempty"`
	BaseLatitude      string    `json:"baseLatitude,omitempty"`
	BaseLongitude     string    `json:"baseLongitude,omitempty"`
	BaseMacAddr       string    `json:"baseMacAddr,omitempty"`
	BaseSsid          string    `json:"baseSsid,omitempty"`
	BizId             string    `json:"bizId,omitempty"`
	CheckType         string    `json:"checkType,omitempty"`
	ClassId           int64     `json:"classId,omitempty"`
	DeviceId          string    `json:"deviceId,omitempty"`
	DeviceSN          string    `json:"deviceSN,omitempty"`
	GmtCreate         time.Time `json:"gmtCreate,omitempty"`
	GmtModified       time.Time `json:"gmtModified,omitempty"`
	GroupId           int64     `json:"groupId,omitempty"`
	Id                int64     `json:"id,omitempty"`
	InvalidRecordMsg  string    `json:"invalidRecordMsg,omitempty"`
	InvalidRecordType string    `json:"invalidRecordType,omitempty"`
	IsLegal           string    `json:"isLegal,omitempty"`
	LocationMethod    string    `json:"locationMethod,omitempty"`
	LocationResult    string    `json:"locationResult,omitempty"`
	OutsideRemark     string    `json:"outsideRemark,omitempty"`
	PlanCheckTime     time.Time `json:"planCheckTime,omitempty"`
	PlanId            int64     `json:"planId,omitempty"`
	ProcInstId        string    `json:"procInstId,omitempty"`
	SourceType        string    `json:"sourceType,omitempty"`
	TimeResult        string    `json:"timeResult,omitempty"`
	UserAccuracy      string    `json:"userAccuracy,omitempty"`
	UserAddress       string    `json:"userAddress,omitempty"`
	UserCheckTime     time.Time `json:"userCheckTime,omitempty"`
	UserId            string    `json:"userId,omitempty"`
	UserLatitude      string    `json:"userLatitude,omitempty"`
	UserLongitude     string    `json:"userLongitude,omitempty"`
	UserMacAddr       string    `json:"userMacAddr,omitempty"`
	UserSsid          string    `json:"userSsid,omitempty"`
	WorkDate          time.Time `json:"workDate,omitempty"`
}
