package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmAuthGroupMemberListRequest() *OapiCrmAuthGroupMemberListRequest {
	return &OapiCrmAuthGroupMemberListRequest{}
}

type OapiCrmAuthGroupMemberListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmAuthGroupMemberListResponse
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmAuthGroupMemberListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmAuthGroupMemberListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmAuthGroupMemberListRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiCrmAuthGroupMemberListRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiCrmAuthGroupMemberListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.auth.group.member.list"
}
func (this *OapiCrmAuthGroupMemberListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmAuthGroupMemberListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmAuthGroupMemberListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmAuthGroupMemberListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmAuthGroupMemberListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["role_id"] = this.RoleId
	return txtParams
}
func (this *OapiCrmAuthGroupMemberListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmAuthGroupMemberListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmAuthGroupMemberListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmAuthGroupMemberListResponse struct {
	taobao.TaobaoResponse
	Result Result `json:"result,omitempty"`
}
type Staff struct {
	Name   string `json:"name,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type Dept struct {
	DeptId int64  `json:"dept_id,omitempty"`
	Name   string `json:"name,omitempty"`
}
type Tag struct {
	Name  string `json:"name,omitempty"`
	TagId int64  `json:"tag_id,omitempty"`
}
type MemberInfo struct {
	AllMember bool    `json:"all_member,omitempty"`
	Dept      []Dept  `json:"dept,omitempty"`
	Staff     []Staff `json:"staff,omitempty"`
	Tag       []Tag   `json:"tag,omitempty"`
}
