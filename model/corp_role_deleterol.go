package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpRoleDeleteroleRequest() *CorpRoleDeleteroleRequest {
	return &CorpRoleDeleteroleRequest{}
}

type CorpRoleDeleteroleRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpRoleDeleteroleResponse
	RoleId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpRoleDeleteroleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpRoleDeleteroleRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *CorpRoleDeleteroleRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *CorpRoleDeleteroleRequest) GetApiMethodName() string {
	return "dingtalk.corp.role.deleterole"
}
func (this *CorpRoleDeleteroleRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpRoleDeleteroleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpRoleDeleteroleRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpRoleDeleteroleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpRoleDeleteroleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpRoleDeleteroleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["role_id"] = this.RoleId
	return txtParams
}
func (this *CorpRoleDeleteroleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpRoleDeleteroleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpRoleDeleteroleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpRoleDeleteroleResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
