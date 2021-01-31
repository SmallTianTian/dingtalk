package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserSimplelistRequest() *OapiUserSimplelistRequest {
	return &OapiUserSimplelistRequest{}
}

type OapiUserSimplelistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserSimplelistResponse
	DepartmentId    int64
	Lang            string
	Offset          int64
	Order           string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserSimplelistRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserSimplelistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserSimplelistRequest) SetDepartmentId(departmentId2 int64) {
	this.DepartmentId = departmentId2
}
func (this *OapiUserSimplelistRequest) GetDepartmentId() int64 {
	return this.DepartmentId
}
func (this *OapiUserSimplelistRequest) SetLang(lang2 string) {
	this.Lang = lang2
}
func (this *OapiUserSimplelistRequest) GetLang() string {
	return this.Lang
}
func (this *OapiUserSimplelistRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiUserSimplelistRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiUserSimplelistRequest) SetOrder(order2 string) {
	this.Order = order2
}
func (this *OapiUserSimplelistRequest) GetOrder() string {
	return this.Order
}
func (this *OapiUserSimplelistRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiUserSimplelistRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiUserSimplelistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.simplelist"
}
func (this *OapiUserSimplelistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserSimplelistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserSimplelistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserSimplelistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserSimplelistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["department_id"] = this.DepartmentId
	txtParams["lang"] = this.Lang
	txtParams["offset"] = this.Offset
	txtParams["order"] = this.Order
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiUserSimplelistRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserSimplelistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserSimplelistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserSimplelistResponse struct {
	taobao.TaobaoResponse
	Errcode  int64      `json:"errcode,omitempty"`
	Errmsg   string     `json:"errmsg,omitempty"`
	HasMore  bool       `json:"hasMore,omitempty"`
	Userlist []Userlist `json:"userlist,omitempty"`
}
