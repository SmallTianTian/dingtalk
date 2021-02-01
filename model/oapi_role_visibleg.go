package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleVisibleGetRequest() *OapiRoleVisibleGetRequest {
	return &OapiRoleVisibleGetRequest{}
}

type OapiRoleVisibleGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleVisibleGetResponse
	Offset          int64
	RoleId          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleVisibleGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleVisibleGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleVisibleGetRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiRoleVisibleGetRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiRoleVisibleGetRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *OapiRoleVisibleGetRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *OapiRoleVisibleGetRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiRoleVisibleGetRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiRoleVisibleGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.visible.get"
}
func (this *OapiRoleVisibleGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleVisibleGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleVisibleGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleVisibleGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleVisibleGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["role_id"] = this.RoleId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiRoleVisibleGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleVisibleGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleVisibleGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleVisibleGetResponse struct {
	taobao.TaobaoResponse
	Result PageVo `json:"result,omitempty"`
}
