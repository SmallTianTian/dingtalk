package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDepartmentGetRequest() *OapiDepartmentGetRequest {
	return &OapiDepartmentGetRequest{}
}

type OapiDepartmentGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDepartmentGetResponse
	Id              string
	Lang            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDepartmentGetRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiDepartmentGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDepartmentGetRequest) SetId(id2 string) {
	this.Id = id2
}
func (this *OapiDepartmentGetRequest) GetId() string {
	return this.Id
}
func (this *OapiDepartmentGetRequest) SetLang(lang2 string) {
	this.Lang = lang2
}
func (this *OapiDepartmentGetRequest) GetLang() string {
	return this.Lang
}
func (this *OapiDepartmentGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.department.get"
}
func (this *OapiDepartmentGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDepartmentGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDepartmentGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDepartmentGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDepartmentGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["id"] = this.Id
	txtParams["lang"] = this.Lang
	return txtParams
}
func (this *OapiDepartmentGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDepartmentGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDepartmentGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDepartmentGetResponse struct {
	taobao.TaobaoResponse
	AutoAddUser           bool   `json:"autoAddUser,omitempty"`
	CreateDeptGroup       bool   `json:"createDeptGroup,omitempty"`
	DeptHiding            bool   `json:"deptHiding,omitempty"`
	DeptManagerUseridList string `json:"deptManagerUseridList,omitempty"`
	DeptPerimits          string `json:"deptPerimits,omitempty"`
	DeptPermits           string `json:"deptPermits,omitempty"`

	GroupContainSubDept bool   `json:"groupContainSubDept,omitempty"`
	Id                  int64  `json:"id,omitempty"`
	IsFromUnionOrg      bool   `json:"isFromUnionOrg,omitempty"`
	Name                string `json:"name,omitempty"`
	Order               int64  `json:"order,omitempty"`
	OrgDeptOwner        string `json:"orgDeptOwner,omitempty"`
	OuterDept           bool   `json:"outerDept,omitempty"`
	OuterPermitDepts    string `json:"outerPermitDepts,omitempty"`
	OuterPermitUsers    string `json:"outerPermitUsers,omitempty"`
	Parentid            int64  `json:"parentid,omitempty"`
	SourceIdentifier    string `json:"sourceIdentifier,omitempty"`
	UserPerimits        string `json:"userPerimits,omitempty"`
	UserPermits         string `json:"userPermits,omitempty"`
}
