package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupQueryRequest() *OapiAttendanceGroupQueryRequest {
	return &OapiAttendanceGroupQueryRequest{}
}

type OapiAttendanceGroupQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupQueryResponse
	GroupId         int64
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupQueryRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *OapiAttendanceGroupQueryRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *OapiAttendanceGroupQueryRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupQueryRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.query"
}
func (this *OapiAttendanceGroupQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_id"] = this.GroupId
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiAttendanceGroupQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupId, "groupId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupQueryResponse struct {
	taobao.TaobaoResponse
	Result  TopSimpleGroupVO `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type TopCycleScheduleItemVO struct {
	ClassId   int64  `json:"class_id,omitempty"`
	ClassName string `json:"class_name,omitempty"`
	IsValid   string `json:"is_valid,omitempty"`
}
type TopCycleScheduleVO struct {
	CycleName string                   `json:"cycle_name,omitempty"`
	IsDeleted string                   `json:"is_deleted,omitempty"`
	IsValid   string                   `json:"is_valid,omitempty"`
	ItemList  []TopCycleScheduleItemVO `json:"item_list,omitempty"`
}
type TopSimpleGroupVO struct {
	AddressList    []string             `json:"address_list,omitempty"`
	CycleSchedules []TopCycleScheduleVO `json:"cycle_schedules,omitempty"`
	Id             int64                `json:"id,omitempty"`
	ManagerList    string               `json:"manager_list,omitempty"`
	MemberCount    int64                `json:"member_count,omitempty"`
	Name           string               `json:"name,omitempty"`
	OwnerUserId    string               `json:"owner_user_id,omitempty"`
	ShiftIds       []int64              `json:"shift_ids,omitempty"`
	Type           string               `json:"type,omitempty"`
	Url            string               `json:"url,omitempty"`
	Wifis          []string             `json:"wifis,omitempty"`
	WorkDayList    []int64              `json:"work_day_list,omitempty"`
}
