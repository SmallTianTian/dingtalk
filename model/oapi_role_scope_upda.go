package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleScopeUpdateRequest() *OapiRoleScopeUpdateRequest {
	return &OapiRoleScopeUpdateRequest{}
}

type OapiRoleScopeUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleScopeUpdateResponse
	DeptIds         string
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRoleScopeUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleScopeUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleScopeUpdateRequest) SetDeptIds(deptIds2 string) {
	this.DeptIds = deptIds2
}
func (this *OapiRoleScopeUpdateRequest) GetDeptIds() string {
	return this.DeptIds
}
func (this *OapiRoleScopeUpdateRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiRoleScopeUpdateRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiRoleScopeUpdateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRoleScopeUpdateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRoleScopeUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.scope.update"
}
func (this *OapiRoleScopeUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleScopeUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleScopeUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleScopeUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleScopeUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_ids"] = this.DeptIds
	txtParams["role_id"] = this.RoleId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRoleScopeUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DeptIds, 200, "deptIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleScopeUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleScopeUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleScopeUpdateResponse struct {
	taobao.TaobaoResponse
}
