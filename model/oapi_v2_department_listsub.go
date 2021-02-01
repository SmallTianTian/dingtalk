package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiV2DepartmentListsubRequest() *OapiV2DepartmentListsubRequest {
	return &OapiV2DepartmentListsubRequest{}
}

type OapiV2DepartmentListsubRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2DepartmentListsubResponse
	DeptId          int64
	Language        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiV2DepartmentListsubRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2DepartmentListsubRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2DepartmentListsubRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiV2DepartmentListsubRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiV2DepartmentListsubRequest) SetLanguage(language2 string) {
	this.Language = language2
}
func (this *OapiV2DepartmentListsubRequest) GetLanguage() string {
	return this.Language
}
func (this *OapiV2DepartmentListsubRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.department.listsub"
}
func (this *OapiV2DepartmentListsubRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2DepartmentListsubRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2DepartmentListsubRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2DepartmentListsubRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2DepartmentListsubRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	txtParams["language"] = this.Language
	return txtParams
}
func (this *OapiV2DepartmentListsubRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiV2DepartmentListsubRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2DepartmentListsubRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2DepartmentListsubResponse struct {
	taobao.TaobaoResponse
	Result []DeptBaseResponse `json:"result,omitempty"`
}
type DeptBaseResponse struct {
	AutoAddUser      bool   `json:"auto_add_user,omitempty"`
	CreateDeptGroup  bool   `json:"create_dept_group,omitempty"`
	DeptId           int64  `json:"dept_id,omitempty"`
	Ext              string `json:"ext,omitempty"`
	FromUnionOrg     bool   `json:"from_union_org,omitempty"`
	Name             string `json:"name,omitempty"`
	ParentId         int64  `json:"parent_id,omitempty"`
	SourceIdentifier string `json:"source_identifier,omitempty"`
	Tags             string `json:"tags,omitempty"`
}
