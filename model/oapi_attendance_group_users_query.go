package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupUsersQueryRequest() *OapiAttendanceGroupUsersQueryRequest {
	return &OapiAttendanceGroupUsersQueryRequest{}
}

type OapiAttendanceGroupUsersQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupUsersQueryResponse
	Cursor          string
	GroupKey        string
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupUsersQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupUsersQueryRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiAttendanceGroupUsersQueryRequest) SetGroupKey(groupKey2 string) {
	this.GroupKey = groupKey2
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetGroupKey() string {
	return this.GroupKey
}
func (this *OapiAttendanceGroupUsersQueryRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupUsersQueryRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.users.query"
}
func (this *OapiAttendanceGroupUsersQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupUsersQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupUsersQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["group_key"] = this.GroupKey
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAttendanceGroupUsersQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupKey, "groupKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupUsersQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupUsersQueryResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
