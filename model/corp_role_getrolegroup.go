package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpRoleGetrolegroupRequest() *CorpRoleGetrolegroupRequest {
	return &CorpRoleGetrolegroupRequest{}
}

type CorpRoleGetrolegroupRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpRoleGetrolegroupResponse
	GroupId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpRoleGetrolegroupRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpRoleGetrolegroupRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *CorpRoleGetrolegroupRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *CorpRoleGetrolegroupRequest) GetApiMethodName() string {
	return "dingtalk.corp.role.getrolegroup"
}
func (this *CorpRoleGetrolegroupRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpRoleGetrolegroupRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpRoleGetrolegroupRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpRoleGetrolegroupRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpRoleGetrolegroupRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpRoleGetrolegroupRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_id"] = this.GroupId
	return txtParams
}
func (this *CorpRoleGetrolegroupRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupId, "groupId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpRoleGetrolegroupRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpRoleGetrolegroupRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpRoleGetrolegroupResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OpenRole struct {
	RoleId   int64  `json:"role_id,omitempty"`
	RoleName string `json:"role_name,omitempty"`
}
type OpenRoleGroup struct {
	GroupName string     `json:"group_name,omitempty"`
	Roles     []OpenRole `json:"roles,omitempty"`
}
