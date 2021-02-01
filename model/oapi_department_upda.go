package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDepartmentUpdateRequest() *OapiDepartmentUpdateRequest {
	return &OapiDepartmentUpdateRequest{}
}

type OapiDepartmentUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiDepartmentUpdateResponse
	AutoAddUser            bool
	CreateDeptGroup        bool
	DeptHiding             bool
	DeptManagerUseridList  string
	DeptPerimits           string
	DeptPermits            string
	GroupContainHiddenDept bool
	GroupContainOuterDept  bool
	GroupContainSubDept    bool
	Id                     int64
	Lang                   string
	Name                   string
	Order                  string
	OrgDeptOwner           string
	OuterDept              bool
	OuterDeptOnlySelf      bool
	OuterPermitDepts       string
	OuterPermitUsers       string
	Parentid               string
	SourceIdentifier       string
	TopHttpMethod          string
	TopResponseType        string
	UserPerimits           string
	UserPermits            string
}

func (this *OapiDepartmentUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDepartmentUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDepartmentUpdateRequest) SetAutoAddUser(autoAddUser2 bool) {
	this.AutoAddUser = autoAddUser2
}
func (this *OapiDepartmentUpdateRequest) GetAutoAddUser() bool {
	return this.AutoAddUser
}
func (this *OapiDepartmentUpdateRequest) SetCreateDeptGroup(createDeptGroup2 bool) {
	this.CreateDeptGroup = createDeptGroup2
}
func (this *OapiDepartmentUpdateRequest) GetCreateDeptGroup() bool {
	return this.CreateDeptGroup
}
func (this *OapiDepartmentUpdateRequest) SetDeptHiding(deptHiding2 bool) {
	this.DeptHiding = deptHiding2
}
func (this *OapiDepartmentUpdateRequest) GetDeptHiding() bool {
	return this.DeptHiding
}
func (this *OapiDepartmentUpdateRequest) SetDeptManagerUseridList(deptManagerUseridList2 string) {
	this.DeptManagerUseridList = deptManagerUseridList2
}
func (this *OapiDepartmentUpdateRequest) GetDeptManagerUseridList() string {
	return this.DeptManagerUseridList
}
func (this *OapiDepartmentUpdateRequest) SetDeptPerimits(deptPerimits2 string) {
	this.DeptPerimits = deptPerimits2
}
func (this *OapiDepartmentUpdateRequest) GetDeptPerimits() string {
	return this.DeptPerimits
}
func (this *OapiDepartmentUpdateRequest) SetDeptPermits(deptPermits2 string) {
	this.DeptPermits = deptPermits2
}
func (this *OapiDepartmentUpdateRequest) GetDeptPermits() string {
	return this.DeptPermits
}
func (this *OapiDepartmentUpdateRequest) SetGroupContainHiddenDept(groupContainHiddenDept2 bool) {
	this.GroupContainHiddenDept = groupContainHiddenDept2
}
func (this *OapiDepartmentUpdateRequest) GetGroupContainHiddenDept() bool {
	return this.GroupContainHiddenDept
}
func (this *OapiDepartmentUpdateRequest) SetGroupContainOuterDept(groupContainOuterDept2 bool) {
	this.GroupContainOuterDept = groupContainOuterDept2
}
func (this *OapiDepartmentUpdateRequest) GetGroupContainOuterDept() bool {
	return this.GroupContainOuterDept
}
func (this *OapiDepartmentUpdateRequest) SetGroupContainSubDept(groupContainSubDept2 bool) {
	this.GroupContainSubDept = groupContainSubDept2
}
func (this *OapiDepartmentUpdateRequest) GetGroupContainSubDept() bool {
	return this.GroupContainSubDept
}
func (this *OapiDepartmentUpdateRequest) SetId(id2 int64) {
	this.Id = id2
}
func (this *OapiDepartmentUpdateRequest) GetId() int64 {
	return this.Id
}
func (this *OapiDepartmentUpdateRequest) SetLang(lang2 string) {
	this.Lang = lang2
}
func (this *OapiDepartmentUpdateRequest) GetLang() string {
	return this.Lang
}
func (this *OapiDepartmentUpdateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiDepartmentUpdateRequest) GetName() string {
	return this.Name
}
func (this *OapiDepartmentUpdateRequest) SetOrder(order2 string) {
	this.Order = order2
}
func (this *OapiDepartmentUpdateRequest) GetOrder() string {
	return this.Order
}
func (this *OapiDepartmentUpdateRequest) SetOrgDeptOwner(orgDeptOwner2 string) {
	this.OrgDeptOwner = orgDeptOwner2
}
func (this *OapiDepartmentUpdateRequest) GetOrgDeptOwner() string {
	return this.OrgDeptOwner
}
func (this *OapiDepartmentUpdateRequest) SetOuterDept(outerDept2 bool) {
	this.OuterDept = outerDept2
}
func (this *OapiDepartmentUpdateRequest) GetOuterDept() bool {
	return this.OuterDept
}
func (this *OapiDepartmentUpdateRequest) SetOuterDeptOnlySelf(outerDeptOnlySelf2 bool) {
	this.OuterDeptOnlySelf = outerDeptOnlySelf2
}
func (this *OapiDepartmentUpdateRequest) GetOuterDeptOnlySelf() bool {
	return this.OuterDeptOnlySelf
}
func (this *OapiDepartmentUpdateRequest) SetOuterPermitDepts(outerPermitDepts2 string) {
	this.OuterPermitDepts = outerPermitDepts2
}
func (this *OapiDepartmentUpdateRequest) GetOuterPermitDepts() string {
	return this.OuterPermitDepts
}
func (this *OapiDepartmentUpdateRequest) SetOuterPermitUsers(outerPermitUsers2 string) {
	this.OuterPermitUsers = outerPermitUsers2
}
func (this *OapiDepartmentUpdateRequest) GetOuterPermitUsers() string {
	return this.OuterPermitUsers
}
func (this *OapiDepartmentUpdateRequest) SetParentid(parentid2 string) {
	this.Parentid = parentid2
}
func (this *OapiDepartmentUpdateRequest) GetParentid() string {
	return this.Parentid
}
func (this *OapiDepartmentUpdateRequest) SetSourceIdentifier(sourceIdentifier2 string) {
	this.SourceIdentifier = sourceIdentifier2
}
func (this *OapiDepartmentUpdateRequest) GetSourceIdentifier() string {
	return this.SourceIdentifier
}
func (this *OapiDepartmentUpdateRequest) SetUserPerimits(userPerimits2 string) {
	this.UserPerimits = userPerimits2
}
func (this *OapiDepartmentUpdateRequest) GetUserPerimits() string {
	return this.UserPerimits
}
func (this *OapiDepartmentUpdateRequest) SetUserPermits(userPermits2 string) {
	this.UserPermits = userPermits2
}
func (this *OapiDepartmentUpdateRequest) GetUserPermits() string {
	return this.UserPermits
}
func (this *OapiDepartmentUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.department.update"
}
func (this *OapiDepartmentUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDepartmentUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDepartmentUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDepartmentUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDepartmentUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["autoAddUser"] = this.AutoAddUser
	txtParams["createDeptGroup"] = this.CreateDeptGroup
	txtParams["deptHiding"] = this.DeptHiding
	txtParams["deptManagerUseridList"] = this.DeptManagerUseridList
	txtParams["deptPerimits"] = this.DeptPerimits
	txtParams["deptPermits"] = this.DeptPermits
	txtParams["groupContainHiddenDept"] = this.GroupContainHiddenDept
	txtParams["groupContainOuterDept"] = this.GroupContainOuterDept
	txtParams["groupContainSubDept"] = this.GroupContainSubDept
	txtParams["id"] = this.Id
	txtParams["lang"] = this.Lang
	txtParams["name"] = this.Name
	txtParams["order"] = this.Order
	txtParams["orgDeptOwner"] = this.OrgDeptOwner
	txtParams["outerDept"] = this.OuterDept
	txtParams["outerDeptOnlySelf"] = this.OuterDeptOnlySelf
	txtParams["outerPermitDepts"] = this.OuterPermitDepts
	txtParams["outerPermitUsers"] = this.OuterPermitUsers
	txtParams["parentid"] = this.Parentid
	txtParams["sourceIdentifier"] = this.SourceIdentifier
	txtParams["userPerimits"] = this.UserPerimits
	txtParams["userPermits"] = this.UserPermits
	return txtParams
}
func (this *OapiDepartmentUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDepartmentUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDepartmentUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDepartmentUpdateResponse struct {
	taobao.TaobaoResponse
	Id int64 `json:"id,omitempty"`
}
