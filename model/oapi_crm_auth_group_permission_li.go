package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmAuthGroupPermissionListRequest() *OapiCrmAuthGroupPermissionListRequest {
	return &OapiCrmAuthGroupPermissionListRequest{}
}

type OapiCrmAuthGroupPermissionListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmAuthGroupPermissionListResponse
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmAuthGroupPermissionListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmAuthGroupPermissionListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmAuthGroupPermissionListRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiCrmAuthGroupPermissionListRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiCrmAuthGroupPermissionListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.auth.group.permission.list"
}
func (this *OapiCrmAuthGroupPermissionListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmAuthGroupPermissionListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmAuthGroupPermissionListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmAuthGroupPermissionListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmAuthGroupPermissionListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["role_id"] = this.RoleId
	return txtParams
}
func (this *OapiCrmAuthGroupPermissionListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmAuthGroupPermissionListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmAuthGroupPermissionListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmAuthGroupPermissionListResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Result `json:"result,omitempty"`
}
type OperateScope struct {
	HasAuth bool   `json:"has_auth,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
}
type DeptInfo struct {
	DeptId   int64  `json:"dept_id,omitempty"`
	DeptName string `json:"dept_name,omitempty"`
}
type StaffInfo struct {
	Name    string `json:"name,omitempty"`
	StaffId string `json:"staff_id,omitempty"`
}
type Ext struct {
	DeptInfo  []DeptInfo  `json:"dept_info,omitempty"`
	StaffInfo []StaffInfo `json:"staff_info,omitempty"`
}
type ManageScope struct {
	Ext     Ext    `json:"ext,omitempty"`
	HasAuth bool   `json:"has_auth,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
}
type AuthObjects struct {
	Label        string         `json:"label,omitempty"`
	ManageScope  []ManageScope  `json:"manage_scope,omitempty"`
	Name         string         `json:"name,omitempty"`
	OperateScope []OperateScope `json:"operate_scope,omitempty"`
}
