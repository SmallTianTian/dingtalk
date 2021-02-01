package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSceneservicegroupGroupQueryRequest() *OapiSceneservicegroupGroupQueryRequest {
	return &OapiSceneservicegroupGroupQueryRequest{}
}

type OapiSceneservicegroupGroupQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiSceneservicegroupGroupQueryResponse
	Cursor             int64
	GroupName          string
	OpenConversationid string
	OpenGroupsetid     string
	OpenTeamid         string
	Size               int64
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiSceneservicegroupGroupQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSceneservicegroupGroupQueryRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiSceneservicegroupGroupQueryRequest) SetGroupName(groupName2 string) {
	this.GroupName = groupName2
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetGroupName() string {
	return this.GroupName
}
func (this *OapiSceneservicegroupGroupQueryRequest) SetOpenConversationid(openConversationid2 string) {
	this.OpenConversationid = openConversationid2
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetOpenConversationid() string {
	return this.OpenConversationid
}
func (this *OapiSceneservicegroupGroupQueryRequest) SetOpenGroupsetid(openGroupsetid2 string) {
	this.OpenGroupsetid = openGroupsetid2
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetOpenGroupsetid() string {
	return this.OpenGroupsetid
}
func (this *OapiSceneservicegroupGroupQueryRequest) SetOpenTeamid(openTeamid2 string) {
	this.OpenTeamid = openTeamid2
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetOpenTeamid() string {
	return this.OpenTeamid
}
func (this *OapiSceneservicegroupGroupQueryRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sceneservicegroup.group.query"
}
func (this *OapiSceneservicegroupGroupQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSceneservicegroupGroupQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSceneservicegroupGroupQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["group_name"] = this.GroupName
	txtParams["open_conversationid"] = this.OpenConversationid
	txtParams["open_groupsetid"] = this.OpenGroupsetid
	txtParams["open_teamid"] = this.OpenTeamid
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSceneservicegroupGroupQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSceneservicegroupGroupQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSceneservicegroupGroupQueryResponse struct {
	taobao.TaobaoResponse
	Result Page `json:"result,omitempty"`
}
type GroupResponse struct {
	GroupName          string `json:"group_name,omitempty"`
	OpenConversationid string `json:"open_conversationid,omitempty"`
	OpenGroupsetid     string `json:"open_groupsetid,omitempty"`
	OpenTeamid         string `json:"open_teamid,omitempty"`
}
