package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDepartmentCreateRequest() *OapiDepartmentCreateRequest {
	return &OapiDepartmentCreateRequest{}
}

type OapiDepartmentCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiDepartmentCreateResponse
	CreateDeptGroup    bool
	DeptHiding         bool
	DeptPerimits       string
	DeptPermits        string
	Id                 string
	Name               string
	Order              string
	OuterDept          bool
	OuterDeptOnlySelf  bool
	OuterPermitDepts   string
	OuterPermitUsers   string
	ParentBalanceFirst bool
	Parentid           string
	ShareBalance       bool
	SourceIdentifier   string
	TopHttpMethod      string
	TopResponseType    string
	UserPerimits       string
	UserPermits        string
}

func (this *OapiDepartmentCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDepartmentCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDepartmentCreateRequest) SetCreateDeptGroup(createDeptGroup2 bool) {
	this.CreateDeptGroup = createDeptGroup2
}
func (this *OapiDepartmentCreateRequest) GetCreateDeptGroup() bool {
	return this.CreateDeptGroup
}
func (this *OapiDepartmentCreateRequest) SetDeptHiding(deptHiding2 bool) {
	this.DeptHiding = deptHiding2
}
func (this *OapiDepartmentCreateRequest) GetDeptHiding() bool {
	return this.DeptHiding
}
func (this *OapiDepartmentCreateRequest) SetDeptPerimits(deptPerimits2 string) {
	this.DeptPerimits = deptPerimits2
}
func (this *OapiDepartmentCreateRequest) GetDeptPerimits() string {
	return this.DeptPerimits
}
func (this *OapiDepartmentCreateRequest) SetDeptPermits(deptPermits2 string) {
	this.DeptPermits = deptPermits2
}
func (this *OapiDepartmentCreateRequest) GetDeptPermits() string {
	return this.DeptPermits
}
func (this *OapiDepartmentCreateRequest) SetId(id2 string) {
	this.Id = id2
}
func (this *OapiDepartmentCreateRequest) GetId() string {
	return this.Id
}
func (this *OapiDepartmentCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiDepartmentCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiDepartmentCreateRequest) SetOrder(order2 string) {
	this.Order = order2
}
func (this *OapiDepartmentCreateRequest) GetOrder() string {
	return this.Order
}
func (this *OapiDepartmentCreateRequest) SetOuterDept(outerDept2 bool) {
	this.OuterDept = outerDept2
}
func (this *OapiDepartmentCreateRequest) GetOuterDept() bool {
	return this.OuterDept
}
func (this *OapiDepartmentCreateRequest) SetOuterDeptOnlySelf(outerDeptOnlySelf2 bool) {
	this.OuterDeptOnlySelf = outerDeptOnlySelf2
}
func (this *OapiDepartmentCreateRequest) GetOuterDeptOnlySelf() bool {
	return this.OuterDeptOnlySelf
}
func (this *OapiDepartmentCreateRequest) SetOuterPermitDepts(outerPermitDepts2 string) {
	this.OuterPermitDepts = outerPermitDepts2
}
func (this *OapiDepartmentCreateRequest) GetOuterPermitDepts() string {
	return this.OuterPermitDepts
}
func (this *OapiDepartmentCreateRequest) SetOuterPermitUsers(outerPermitUsers2 string) {
	this.OuterPermitUsers = outerPermitUsers2
}
func (this *OapiDepartmentCreateRequest) GetOuterPermitUsers() string {
	return this.OuterPermitUsers
}
func (this *OapiDepartmentCreateRequest) SetParentBalanceFirst(parentBalanceFirst2 bool) {
	this.ParentBalanceFirst = parentBalanceFirst2
}
func (this *OapiDepartmentCreateRequest) GetParentBalanceFirst() bool {
	return this.ParentBalanceFirst
}
func (this *OapiDepartmentCreateRequest) SetParentid(parentid2 string) {
	this.Parentid = parentid2
}
func (this *OapiDepartmentCreateRequest) GetParentid() string {
	return this.Parentid
}
func (this *OapiDepartmentCreateRequest) SetShareBalance(shareBalance2 bool) {
	this.ShareBalance = shareBalance2
}
func (this *OapiDepartmentCreateRequest) GetShareBalance() bool {
	return this.ShareBalance
}
func (this *OapiDepartmentCreateRequest) SetSourceIdentifier(sourceIdentifier2 string) {
	this.SourceIdentifier = sourceIdentifier2
}
func (this *OapiDepartmentCreateRequest) GetSourceIdentifier() string {
	return this.SourceIdentifier
}
func (this *OapiDepartmentCreateRequest) SetUserPerimits(userPerimits2 string) {
	this.UserPerimits = userPerimits2
}
func (this *OapiDepartmentCreateRequest) GetUserPerimits() string {
	return this.UserPerimits
}
func (this *OapiDepartmentCreateRequest) SetUserPermits(userPermits2 string) {
	this.UserPermits = userPermits2
}
func (this *OapiDepartmentCreateRequest) GetUserPermits() string {
	return this.UserPermits
}
func (this *OapiDepartmentCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.department.create"
}
func (this *OapiDepartmentCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDepartmentCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDepartmentCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDepartmentCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDepartmentCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["createDeptGroup"] = this.CreateDeptGroup
	txtParams["deptHiding"] = this.DeptHiding
	txtParams["deptPerimits"] = this.DeptPerimits
	txtParams["deptPermits"] = this.DeptPermits
	txtParams["id"] = this.Id
	txtParams["name"] = this.Name
	txtParams["order"] = this.Order
	txtParams["outerDept"] = this.OuterDept
	txtParams["outerDeptOnlySelf"] = this.OuterDeptOnlySelf
	txtParams["outerPermitDepts"] = this.OuterPermitDepts
	txtParams["outerPermitUsers"] = this.OuterPermitUsers
	txtParams["parentBalanceFirst"] = this.ParentBalanceFirst
	txtParams["parentid"] = this.Parentid
	txtParams["shareBalance"] = this.ShareBalance
	txtParams["sourceIdentifier"] = this.SourceIdentifier
	txtParams["userPerimits"] = this.UserPerimits
	txtParams["userPermits"] = this.UserPermits
	return txtParams
}
func (this *OapiDepartmentCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDepartmentCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDepartmentCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDepartmentCreateResponse struct {
	taobao.TaobaoResponse
	Id int64 `json:"id,omitempty"`
}
