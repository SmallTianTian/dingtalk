package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupUsersAddRequest() *OapiAttendanceGroupUsersAddRequest {
	return &OapiAttendanceGroupUsersAddRequest{}
}

type OapiAttendanceGroupUsersAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupUsersAddResponse
	GroupKey        string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
	UserIdList      string
}

func (this *OapiAttendanceGroupUsersAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupUsersAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupUsersAddRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupUsersAddRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupUsersAddRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupUsersAddRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupUsersAddRequest) SetUserIdList(userIdList2 string) {
	this.UserIdList = userIdList2
}
func (this *OapiAttendanceGroupUsersAddRequest) GetUserIdList() string {
	return this.UserIdList
}
func (this *OapiAttendanceGroupUsersAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.users.add"
}
func (this *OapiAttendanceGroupUsersAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupUsersAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupUsersAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupUsersAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupUsersAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	txtParams["user_id_list"] = this.UserIdList
	return txtParams
}
func (this *OapiAttendanceGroupUsersAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupUsersAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupUsersAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupUsersAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
