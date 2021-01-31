package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleVisibleDeleteRequest() *OapiRoleVisibleDeleteRequest {
	return &OapiRoleVisibleDeleteRequest{}
}

type OapiRoleVisibleDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleVisibleDeleteResponse
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleVisibleDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleVisibleDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleVisibleDeleteRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiRoleVisibleDeleteRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiRoleVisibleDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.visible.delete"
}
func (this *OapiRoleVisibleDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleVisibleDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleVisibleDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleVisibleDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleVisibleDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["role_id"] = this.RoleId
	return txtParams
}
func (this *OapiRoleVisibleDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleVisibleDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleVisibleDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleVisibleDeleteResponse struct {
	taobao.TaobaoResponse
}
