package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2DepartmentUpdateRequest() *OapiV2DepartmentUpdateRequest {
	return &OapiV2DepartmentUpdateRequest{}
}

type OapiV2DepartmentUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiV2DepartmentUpdateResponse
	AutoAddUser            bool
	CreateDeptGroup        bool
	DeptId                 int64
	DeptManagerUseridList  string
	DeptPermits            string
	Extension              string
	GroupContainHiddenDept bool
	GroupContainOuterDept  bool
	GroupContainSubDept    bool
	HideDept               bool
	Language               string
	Name                   string
	Order                  int64
	OrgDeptOwner           string
	OuterDept              bool
	OuterDeptOnlySelf      bool
	OuterPermitDepts       string
	OuterPermitUsers       string
	ParentId               int64
	SourceIdentifier       string
	TopHttpMethod          string
	TopResponseType        string
	UserPermits            string
}

func (this *OapiV2DepartmentUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2DepartmentUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2DepartmentUpdateRequest) SetAutoAddUser(autoAddUser2 bool) {
	this.AutoAddUser = autoAddUser2
}
func (this *OapiV2DepartmentUpdateRequest) GetAutoAddUser() bool {
	return this.AutoAddUser
}
func (this *OapiV2DepartmentUpdateRequest) SetCreateDeptGroup(createDeptGroup2 bool) {
	this.CreateDeptGroup = createDeptGroup2
}
func (this *OapiV2DepartmentUpdateRequest) GetCreateDeptGroup() bool {
	return this.CreateDeptGroup
}
func (this *OapiV2DepartmentUpdateRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiV2DepartmentUpdateRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiV2DepartmentUpdateRequest) SetDeptManagerUseridList(deptManagerUseridList2 string) {
	this.DeptManagerUseridList = deptManagerUseridList2
}
func (this *OapiV2DepartmentUpdateRequest) GetDeptManagerUseridList() string {
	return this.DeptManagerUseridList
}
func (this *OapiV2DepartmentUpdateRequest) SetDeptPermits(deptPermits2 string) {
	this.DeptPermits = deptPermits2
}
func (this *OapiV2DepartmentUpdateRequest) GetDeptPermits() string {
	return this.DeptPermits
}
func (this *OapiV2DepartmentUpdateRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiV2DepartmentUpdateRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiV2DepartmentUpdateRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiV2DepartmentUpdateRequest) SetGroupContainHiddenDept(groupContainHiddenDept2 bool) {
	this.GroupContainHiddenDept = groupContainHiddenDept2
}
func (this *OapiV2DepartmentUpdateRequest) GetGroupContainHiddenDept() bool {
	return this.GroupContainHiddenDept
}
func (this *OapiV2DepartmentUpdateRequest) SetGroupContainOuterDept(groupContainOuterDept2 bool) {
	this.GroupContainOuterDept = groupContainOuterDept2
}
func (this *OapiV2DepartmentUpdateRequest) GetGroupContainOuterDept() bool {
	return this.GroupContainOuterDept
}
func (this *OapiV2DepartmentUpdateRequest) SetGroupContainSubDept(groupContainSubDept2 bool) {
	this.GroupContainSubDept = groupContainSubDept2
}
func (this *OapiV2DepartmentUpdateRequest) GetGroupContainSubDept() bool {
	return this.GroupContainSubDept
}
func (this *OapiV2DepartmentUpdateRequest) SetHideDept(hideDept2 bool) {
	this.HideDept = hideDept2
}
func (this *OapiV2DepartmentUpdateRequest) GetHideDept() bool {
	return this.HideDept
}
func (this *OapiV2DepartmentUpdateRequest) SetLanguage(language2 string) {
	this.Language = language2
}
func (this *OapiV2DepartmentUpdateRequest) GetLanguage() string {
	return this.Language
}
func (this *OapiV2DepartmentUpdateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiV2DepartmentUpdateRequest) GetName() string {
	return this.Name
}
func (this *OapiV2DepartmentUpdateRequest) SetOrder(order2 int64) {
	this.Order = order2
}
func (this *OapiV2DepartmentUpdateRequest) GetOrder() int64 {
	return this.Order
}
func (this *OapiV2DepartmentUpdateRequest) SetOrgDeptOwner(orgDeptOwner2 string) {
	this.OrgDeptOwner = orgDeptOwner2
}
func (this *OapiV2DepartmentUpdateRequest) GetOrgDeptOwner() string {
	return this.OrgDeptOwner
}
func (this *OapiV2DepartmentUpdateRequest) SetOuterDept(outerDept2 bool) {
	this.OuterDept = outerDept2
}
func (this *OapiV2DepartmentUpdateRequest) GetOuterDept() bool {
	return this.OuterDept
}
func (this *OapiV2DepartmentUpdateRequest) SetOuterDeptOnlySelf(outerDeptOnlySelf2 bool) {
	this.OuterDeptOnlySelf = outerDeptOnlySelf2
}
func (this *OapiV2DepartmentUpdateRequest) GetOuterDeptOnlySelf() bool {
	return this.OuterDeptOnlySelf
}
func (this *OapiV2DepartmentUpdateRequest) SetOuterPermitDepts(outerPermitDepts2 string) {
	this.OuterPermitDepts = outerPermitDepts2
}
func (this *OapiV2DepartmentUpdateRequest) GetOuterPermitDepts() string {
	return this.OuterPermitDepts
}
func (this *OapiV2DepartmentUpdateRequest) SetOuterPermitUsers(outerPermitUsers2 string) {
	this.OuterPermitUsers = outerPermitUsers2
}
func (this *OapiV2DepartmentUpdateRequest) GetOuterPermitUsers() string {
	return this.OuterPermitUsers
}
func (this *OapiV2DepartmentUpdateRequest) SetParentId(parentId2 int64) {
	this.ParentId = parentId2
}
func (this *OapiV2DepartmentUpdateRequest) GetParentId() int64 {
	return this.ParentId
}
func (this *OapiV2DepartmentUpdateRequest) SetSourceIdentifier(sourceIdentifier2 string) {
	this.SourceIdentifier = sourceIdentifier2
}
func (this *OapiV2DepartmentUpdateRequest) GetSourceIdentifier() string {
	return this.SourceIdentifier
}
func (this *OapiV2DepartmentUpdateRequest) SetUserPermits(userPermits2 string) {
	this.UserPermits = userPermits2
}
func (this *OapiV2DepartmentUpdateRequest) GetUserPermits() string {
	return this.UserPermits
}
func (this *OapiV2DepartmentUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.department.update"
}
func (this *OapiV2DepartmentUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2DepartmentUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2DepartmentUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2DepartmentUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2DepartmentUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["auto_add_user"] = this.AutoAddUser
	txtParams["create_dept_group"] = this.CreateDeptGroup
	txtParams["dept_id"] = this.DeptId
	txtParams["dept_manager_userid_list"] = this.DeptManagerUseridList
	txtParams["dept_permits"] = this.DeptPermits
	txtParams["extension"] = this.Extension
	txtParams["group_contain_hidden_dept"] = this.GroupContainHiddenDept
	txtParams["group_contain_outer_dept"] = this.GroupContainOuterDept
	txtParams["group_contain_sub_dept"] = this.GroupContainSubDept
	txtParams["hide_dept"] = this.HideDept
	txtParams["language"] = this.Language
	txtParams["name"] = this.Name
	txtParams["order"] = this.Order
	txtParams["org_dept_owner"] = this.OrgDeptOwner
	txtParams["outer_dept"] = this.OuterDept
	txtParams["outer_dept_only_self"] = this.OuterDeptOnlySelf
	txtParams["outer_permit_depts"] = this.OuterPermitDepts
	txtParams["outer_permit_users"] = this.OuterPermitUsers
	txtParams["parent_id"] = this.ParentId
	txtParams["source_identifier"] = this.SourceIdentifier
	txtParams["user_permits"] = this.UserPermits
	return txtParams
}
func (this *OapiV2DepartmentUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2DepartmentUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2DepartmentUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2DepartmentUpdateResponse struct {
	taobao.TaobaoResponse
}
