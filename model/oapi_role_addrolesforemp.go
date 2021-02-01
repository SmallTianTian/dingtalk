package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleAddrolesforempsRequest() *OapiRoleAddrolesforempsRequest {
	return &OapiRoleAddrolesforempsRequest{}
}

type OapiRoleAddrolesforempsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleAddrolesforempsResponse
	RoleIds         string
	TopHttpMethod   string
	TopResponseType string
	UserIds         string
}

func (this *OapiRoleAddrolesforempsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleAddrolesforempsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleAddrolesforempsRequest) SetRoleIds(roleIds2 string) {
	this.RoleIds = roleIds2
}
func (this *OapiRoleAddrolesforempsRequest) GetRoleIds() string {
	return this.RoleIds
}
func (this *OapiRoleAddrolesforempsRequest) SetUserIds(userIds2 string) {
	this.UserIds = userIds2
}
func (this *OapiRoleAddrolesforempsRequest) GetUserIds() string {
	return this.UserIds
}
func (this *OapiRoleAddrolesforempsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.addrolesforemps"
}
func (this *OapiRoleAddrolesforempsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleAddrolesforempsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleAddrolesforempsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleAddrolesforempsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleAddrolesforempsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["roleIds"] = this.RoleIds
	txtParams["userIds"] = this.UserIds
	return txtParams
}
func (this *OapiRoleAddrolesforempsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleIds, "roleIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleAddrolesforempsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleAddrolesforempsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleAddrolesforempsResponse struct {
	taobao.TaobaoResponse
}
