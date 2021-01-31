package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpRoleSimplelistRequest() *CorpRoleSimplelistRequest {
	return &CorpRoleSimplelistRequest{}
}

type CorpRoleSimplelistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpRoleSimplelistResponse
	Offset          int64
	RoleId          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpRoleSimplelistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpRoleSimplelistRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *CorpRoleSimplelistRequest) GetOffset() int64 {
	return this.Offset
}
func (this *CorpRoleSimplelistRequest) SetRoleId(roleId2 int64) {
	this.RoleId = roleId2
}
func (this *CorpRoleSimplelistRequest) GetRoleId() int64 {
	return this.RoleId
}
func (this *CorpRoleSimplelistRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpRoleSimplelistRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpRoleSimplelistRequest) GetApiMethodName() string {
	return "dingtalk.corp.role.simplelist"
}
func (this *CorpRoleSimplelistRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpRoleSimplelistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpRoleSimplelistRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpRoleSimplelistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpRoleSimplelistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpRoleSimplelistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["role_id"] = this.RoleId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *CorpRoleSimplelistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleId, "roleId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpRoleSimplelistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpRoleSimplelistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpRoleSimplelistResponse struct {
	taobao.TaobaoResponse
	Result PageVo `json:"result,omitempty"`
}
type EmpSimpleList struct {
	Userid string `json:"userid,omitempty"`
}
