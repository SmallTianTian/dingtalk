package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleVisibleSetRequest() *OapiRoleVisibleSetRequest {
	return &OapiRoleVisibleSetRequest{}
}

type OapiRoleVisibleSetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleVisibleSetResponse
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
	VisibleRoleIds  string
}

func (this *OapiRoleVisibleSetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleVisibleSetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleVisibleSetRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiRoleVisibleSetRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiRoleVisibleSetRequest) SetVisibleRoleIds(visibleRoleIds2 string) {
	this.VisibleRoleIds = visibleRoleIds2
}
func (this *OapiRoleVisibleSetRequest) GetVisibleRoleIds() string {
	return this.VisibleRoleIds
}
func (this *OapiRoleVisibleSetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.visible.set"
}
func (this *OapiRoleVisibleSetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleVisibleSetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleVisibleSetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleVisibleSetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleVisibleSetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["role_id"] = this.RoleId
	txtParams["visible_role_ids"] = this.VisibleRoleIds
	return txtParams
}
func (this *OapiRoleVisibleSetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleVisibleSetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleVisibleSetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleVisibleSetResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
