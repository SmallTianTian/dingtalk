package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupUsersRemoveRequest() *OapiAttendanceGroupUsersRemoveRequest {
	return &OapiAttendanceGroupUsersRemoveRequest{}
}

type OapiAttendanceGroupUsersRemoveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupUsersRemoveResponse
	GroupKey        string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
	UserIdList      string
}

func (this *OapiAttendanceGroupUsersRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupUsersRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupUsersRemoveRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupUsersRemoveRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupUsersRemoveRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupUsersRemoveRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupUsersRemoveRequest) SetUserIdList(userIdList2 string) {
	this.UserIdList = userIdList2
}
func (this *OapiAttendanceGroupUsersRemoveRequest) GetUserIdList() string {
	return this.UserIdList
}
func (this *OapiAttendanceGroupUsersRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.users.remove"
}
func (this *OapiAttendanceGroupUsersRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupUsersRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupUsersRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupUsersRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupUsersRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	txtParams["user_id_list"] = this.UserIdList
	return txtParams
}
func (this *OapiAttendanceGroupUsersRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupUsersRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupUsersRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupUsersRemoveResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
