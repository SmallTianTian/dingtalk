package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleSimplelistRequest() *OapiRoleSimplelistRequest {
	return &OapiRoleSimplelistRequest{}
}

type OapiRoleSimplelistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleSimplelistResponse
	Offset          int64
	RoleId          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleSimplelistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleSimplelistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleSimplelistRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiRoleSimplelistRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiRoleSimplelistRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiRoleSimplelistRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiRoleSimplelistRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiRoleSimplelistRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiRoleSimplelistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.simplelist"
}
func (this *OapiRoleSimplelistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleSimplelistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleSimplelistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleSimplelistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleSimplelistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["role_id"] = this.RoleId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiRoleSimplelistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleSimplelistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleSimplelistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleSimplelistResponse struct {
	taobao.TaobaoResponse
	Result PageVo `json:"result,omitempty"`
}
type OrgDeptVo struct {
	DeptId int64  `json:"dept_id,omitempty"`
	Name   string `json:"name,omitempty"`
}
type OpenEmpSimple struct {
	ManageScopes []OrgDeptVo `json:"manageScopes,omitempty"`
	Name         string      `json:"name,omitempty"`
	Userid       string      `json:"userid,omitempty"`
}
