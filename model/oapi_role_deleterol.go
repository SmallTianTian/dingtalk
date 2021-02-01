package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleDeleteroleRequest() *OapiRoleDeleteroleRequest {
	return &OapiRoleDeleteroleRequest{}
}

type OapiRoleDeleteroleRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleDeleteroleResponse
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleDeleteroleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleDeleteroleRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleDeleteroleRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiRoleDeleteroleRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiRoleDeleteroleRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.deleterole"
}
func (this *OapiRoleDeleteroleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleDeleteroleRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleDeleteroleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleDeleteroleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleDeleteroleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["role_id"] = this.RoleId
	return txtParams
}
func (this *OapiRoleDeleteroleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleDeleteroleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleDeleteroleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleDeleteroleResponse struct {
	taobao.TaobaoResponse
}
