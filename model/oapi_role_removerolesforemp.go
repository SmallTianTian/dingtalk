package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleRemoverolesforempsRequest() *OapiRoleRemoverolesforempsRequest {
	return &OapiRoleRemoverolesforempsRequest{}
}

type OapiRoleRemoverolesforempsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleRemoverolesforempsResponse
	RoleIds         string
	TopHttpMethod   string
	TopResponseType string
	UserIds         string
}

func (this *OapiRoleRemoverolesforempsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleRemoverolesforempsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleRemoverolesforempsRequest) SetRoleIds(roleIds2 string) {
	this.RoleIds = roleIds2
}
func (this *OapiRoleRemoverolesforempsRequest) GetRoleIds() string {
	return this.RoleIds
}
func (this *OapiRoleRemoverolesforempsRequest) SetUserIds(userIds2 string) {
	this.UserIds = userIds2
}
func (this *OapiRoleRemoverolesforempsRequest) GetUserIds() string {
	return this.UserIds
}
func (this *OapiRoleRemoverolesforempsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.removerolesforemps"
}
func (this *OapiRoleRemoverolesforempsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleRemoverolesforempsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleRemoverolesforempsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleRemoverolesforempsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleRemoverolesforempsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["roleIds"] = this.RoleIds
	txtParams["userIds"] = this.UserIds
	return txtParams
}
func (this *OapiRoleRemoverolesforempsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleIds, "roleIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleRemoverolesforempsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleRemoverolesforempsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleRemoverolesforempsResponse struct {
	taobao.TaobaoResponse
}
