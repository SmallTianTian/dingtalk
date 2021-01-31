package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2DepartmentCreateRequest() *OapiV2DepartmentCreateRequest {
	return &OapiV2DepartmentCreateRequest{}
}

type OapiV2DepartmentCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiV2DepartmentCreateResponse
	CreateDeptGroup   bool
	DeptPermits       string
	Extension         string
	HideDept          bool
	Name              string
	Order             int64
	OuterDept         bool
	OuterDeptOnlySelf bool
	OuterPermitDepts  string
	OuterPermitUsers  string
	ParentId          int64
	SourceIdentifier  string
	TopHttpMethod     string
	TopResponseType   string
	UserPermits       string
}

func (this *OapiV2DepartmentCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2DepartmentCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2DepartmentCreateRequest) SetCreateDeptGroup(createDeptGroup2 bool) {
	this.CreateDeptGroup = createDeptGroup2
}
func (this *OapiV2DepartmentCreateRequest) GetCreateDeptGroup() bool {
	return this.CreateDeptGroup
}
func (this *OapiV2DepartmentCreateRequest) SetDeptPermits(deptPermits2 string) {
	this.DeptPermits = deptPermits2
}
func (this *OapiV2DepartmentCreateRequest) GetDeptPermits() string {
	return this.DeptPermits
}
func (this *OapiV2DepartmentCreateRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiV2DepartmentCreateRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiV2DepartmentCreateRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiV2DepartmentCreateRequest) SetHideDept(hideDept2 bool) {
	this.HideDept = hideDept2
}
func (this *OapiV2DepartmentCreateRequest) GetHideDept() bool {
	return this.HideDept
}
func (this *OapiV2DepartmentCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiV2DepartmentCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiV2DepartmentCreateRequest) SetOrder(order2 int64) {
	this.Order = order2
}
func (this *OapiV2DepartmentCreateRequest) GetOrder() int64 {
	return this.Order
}
func (this *OapiV2DepartmentCreateRequest) SetOuterDept(outerDept2 bool) {
	this.OuterDept = outerDept2
}
func (this *OapiV2DepartmentCreateRequest) GetOuterDept() bool {
	return this.OuterDept
}
func (this *OapiV2DepartmentCreateRequest) SetOuterDeptOnlySelf(outerDeptOnlySelf2 bool) {
	this.OuterDeptOnlySelf = outerDeptOnlySelf2
}
func (this *OapiV2DepartmentCreateRequest) GetOuterDeptOnlySelf() bool {
	return this.OuterDeptOnlySelf
}
func (this *OapiV2DepartmentCreateRequest) SetOuterPermitDepts(outerPermitDepts2 string) {
	this.OuterPermitDepts = outerPermitDepts2
}
func (this *OapiV2DepartmentCreateRequest) GetOuterPermitDepts() string {
	return this.OuterPermitDepts
}
func (this *OapiV2DepartmentCreateRequest) SetOuterPermitUsers(outerPermitUsers2 string) {
	this.OuterPermitUsers = outerPermitUsers2
}
func (this *OapiV2DepartmentCreateRequest) GetOuterPermitUsers() string {
	return this.OuterPermitUsers
}
func (this *OapiV2DepartmentCreateRequest) SetParentId(parentId2 int64) {
	this.ParentId = parentId2
}
func (this *OapiV2DepartmentCreateRequest) GetParentId() int64 {
	return this.ParentId
}
func (this *OapiV2DepartmentCreateRequest) SetSourceIdentifier(sourceIdentifier2 string) {
	this.SourceIdentifier = sourceIdentifier2
}
func (this *OapiV2DepartmentCreateRequest) GetSourceIdentifier() string {
	return this.SourceIdentifier
}
func (this *OapiV2DepartmentCreateRequest) SetUserPermits(userPermits2 string) {
	this.UserPermits = userPermits2
}
func (this *OapiV2DepartmentCreateRequest) GetUserPermits() string {
	return this.UserPermits
}
func (this *OapiV2DepartmentCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.department.create"
}
func (this *OapiV2DepartmentCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2DepartmentCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2DepartmentCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2DepartmentCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2DepartmentCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["create_dept_group"] = this.CreateDeptGroup
	txtParams["dept_permits"] = this.DeptPermits
	txtParams["extension"] = this.Extension
	txtParams["hide_dept"] = this.HideDept
	txtParams["name"] = this.Name
	txtParams["order"] = this.Order
	txtParams["outer_dept"] = this.OuterDept
	txtParams["outer_dept_only_self"] = this.OuterDeptOnlySelf
	txtParams["outer_permit_depts"] = this.OuterPermitDepts
	txtParams["outer_permit_users"] = this.OuterPermitUsers
	txtParams["parent_id"] = this.ParentId
	txtParams["source_identifier"] = this.SourceIdentifier
	txtParams["user_permits"] = this.UserPermits
	return txtParams
}
func (this *OapiV2DepartmentCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DeptPermits, 200, "deptPermits"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2DepartmentCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2DepartmentCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2DepartmentCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64              `json:"errcode,omitempty"`
	Errmsg  string             `json:"errmsg,omitempty"`
	Result  DeptCreateResponse `json:"result,omitempty"`
}
type DeptCreateResponse struct {
	DeptId int64 `json:"dept_id,omitempty"`
}
