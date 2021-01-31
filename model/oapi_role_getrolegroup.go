package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRoleGetrolegroupRequest() *OapiRoleGetrolegroupRequest {
	return &OapiRoleGetrolegroupRequest{}
}

type OapiRoleGetrolegroupRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRoleGetrolegroupResponse
	GroupId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRoleGetrolegroupRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRoleGetrolegroupRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRoleGetrolegroupRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *OapiRoleGetrolegroupRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *OapiRoleGetrolegroupRequest) GetApiMethodName() string {
	return "dingtalk.oapi.role.getrolegroup"
}
func (this *OapiRoleGetrolegroupRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRoleGetrolegroupRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRoleGetrolegroupRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRoleGetrolegroupRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRoleGetrolegroupRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_id"] = this.GroupId
	return txtParams
}
func (this *OapiRoleGetrolegroupRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupId, "groupId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRoleGetrolegroupRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRoleGetrolegroupRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRoleGetrolegroupResponse struct {
	taobao.TaobaoResponse
	Errcode   int64         `json:"errcode,omitempty"`
	Errmsg    string        `json:"errmsg,omitempty"`
	RoleGroup OpenRoleGroup `json:"role_group,omitempty"`
}
