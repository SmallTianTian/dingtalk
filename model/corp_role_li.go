package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpRoleListRequest() *CorpRoleListRequest {
	return &CorpRoleListRequest{}
}

type CorpRoleListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpRoleListResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpRoleListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpRoleListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *CorpRoleListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *CorpRoleListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpRoleListRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpRoleListRequest) GetApiMethodName() string {
	return "dingtalk.corp.role.list"
}
func (this *CorpRoleListRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpRoleListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpRoleListRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpRoleListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpRoleListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpRoleListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *CorpRoleListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpRoleListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpRoleListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpRoleListResponse struct {
	taobao.TaobaoResponse
	Result Result `json:"result,omitempty"`
}
type Roles struct {
	Id       int64  `json:"id,omitempty"`
	RoleName string `json:"role_name,omitempty"`
}
type RoleGroups struct {
	GroupName string  `json:"group_name,omitempty"`
	Roles     []Roles `json:"roles,omitempty"`
}
type Result struct {
	HasMore string       `json:"has_more,omitempty"`
	List    []RoleGroups `json:"list,omitempty"`
}
