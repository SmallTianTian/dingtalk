package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupSearchRequest() *OapiAttendanceGroupSearchRequest {
	return &OapiAttendanceGroupSearchRequest{}
}

type OapiAttendanceGroupSearchRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupSearchResponse
	GroupName       string
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupSearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupSearchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupSearchRequest) SetGroupName(groupName2 string) {
	this.GroupName = groupName2
}
func (this *OapiAttendanceGroupSearchRequest) GetGroupName() string {
	return this.GroupName
}
func (this *OapiAttendanceGroupSearchRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupSearchRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupSearchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.search"
}
func (this *OapiAttendanceGroupSearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupSearchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupSearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupSearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupSearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_name"] = this.GroupName
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiAttendanceGroupSearchRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupName, "groupName"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupSearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupSearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupSearchResponse struct {
	taobao.TaobaoResponse
	Errcode int64                  `json:"errcode,omitempty"`
	Errmsg  string                 `json:"errmsg,omitempty"`
	Result  []TopMinimalismGroupVO `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
type TopMinimalismGroupVO struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
