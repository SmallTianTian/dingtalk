package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleGetroleRequest() *OapiRoleGetroleRequest {
	return &OapiRoleGetroleRequest{}
}

type OapiRoleGetroleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleGetroleResponse
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleGetroleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleGetroleRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleGetroleRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiRoleGetroleRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiRoleGetroleRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.getrole"
}
func (this *OapiRoleGetroleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleGetroleRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleGetroleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleGetroleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleGetroleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["roleId"] = this.RoleId
	return txtParams
}
func (this *OapiRoleGetroleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleGetroleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleGetroleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleGetroleResponse struct {
	taobao.TaobaoResponse
	Errcode int64    `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	Role    OpenRole `json:"role,omitempty"`
}
