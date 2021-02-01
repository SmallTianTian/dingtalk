package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSceneservicegroupGroupCreateRequest() *OapiSceneservicegroupGroupCreateRequest {
	return &OapiSceneservicegroupGroupCreateRequest{}
}

type OapiSceneservicegroupGroupCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSceneservicegroupGroupCreateResponse
	Bizid           string
	GroupName       string
	GroupTagids     string
	MemberStaffids  string
	OpenGroupsetid  string
	OpenTeamid      string
	OwnerStaffId    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSceneservicegroupGroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetBizid(bizid2 string) {
	this.Bizid = bizid2
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetBizid() string {
	return this.Bizid
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetGroupName(groupName2 string) {
	this.GroupName = groupName2
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetGroupName() string {
	return this.GroupName
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetGroupTagids(groupTagids2 string) {
	this.GroupTagids = groupTagids2
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetGroupTagids() string {
	return this.GroupTagids
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetMemberStaffids(memberStaffids2 string) {
	this.MemberStaffids = memberStaffids2
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetMemberStaffids() string {
	return this.MemberStaffids
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetOpenGroupsetid(openGroupsetid2 string) {
	this.OpenGroupsetid = openGroupsetid2
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetOpenGroupsetid() string {
	return this.OpenGroupsetid
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetOpenTeamid(openTeamid2 string) {
	this.OpenTeamid = openTeamid2
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetOpenTeamid() string {
	return this.OpenTeamid
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetOwnerStaffId(ownerStaffId2 string) {
	this.OwnerStaffId = ownerStaffId2
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetOwnerStaffId() string {
	return this.OwnerStaffId
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sceneservicegroup.group.create"
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSceneservicegroupGroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["bizid"] = this.Bizid
	txtParams["group_name"] = this.GroupName
	txtParams["group_tagids"] = this.GroupTagids
	txtParams["member_staffids"] = this.MemberStaffids
	txtParams["open_groupsetid"] = this.OpenGroupsetid
	txtParams["open_teamid"] = this.OpenTeamid
	txtParams["owner_staffId"] = this.OwnerStaffId
	return txtParams
}
func (this *OapiSceneservicegroupGroupCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupName, "groupName"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSceneservicegroupGroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSceneservicegroupGroupCreateResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
