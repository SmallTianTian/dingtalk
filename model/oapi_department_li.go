package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDepartmentListRequest() *OapiDepartmentListRequest {
	return &OapiDepartmentListRequest{}
}

type OapiDepartmentListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDepartmentListResponse
	FetchChild      bool
	Id              string
	Lang            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDepartmentListRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiDepartmentListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDepartmentListRequest) SetFetchChild(fetchChild2 bool) {
	this.FetchChild = fetchChild2
}
func (this *OapiDepartmentListRequest) GetFetchChild() bool {
	return this.FetchChild
}
func (this *OapiDepartmentListRequest) SetId(id2 string) {
	this.Id = id2
}
func (this *OapiDepartmentListRequest) GetId() string {
	return this.Id
}
func (this *OapiDepartmentListRequest) SetLang(lang2 string) {
	this.Lang = lang2
}
func (this *OapiDepartmentListRequest) GetLang() string {
	return this.Lang
}
func (this *OapiDepartmentListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.department.list"
}
func (this *OapiDepartmentListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDepartmentListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDepartmentListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDepartmentListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDepartmentListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["fetch_child"] = this.FetchChild
	txtParams["id"] = this.Id
	txtParams["lang"] = this.Lang
	return txtParams
}
func (this *OapiDepartmentListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDepartmentListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDepartmentListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDepartmentListResponse struct {
	taobao.TaobaoResponse
	Department []Department `json:"department,omitempty"`
	Errcode    int64        `json:"errcode,omitempty"`
	Errmsg     string       `json:"errmsg,omitempty"`
}
type Department struct {
	AutoAddUser      bool   `json:"autoAddUser,omitempty"`
	CreateDeptGroup  bool   `json:"createDeptGroup,omitempty"`
	Id               int64  `json:"id,omitempty"`
	IsFromUnionOrg   bool   `json:"isFromUnionOrg,omitempty"`
	Name             string `json:"name,omitempty"`
	Parentid         int64  `json:"parentid,omitempty"`
	SourceIdentifier string `json:"sourceIdentifier,omitempty"`
}
