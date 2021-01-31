package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiIndustryDepartmentGetRequest() *OapiIndustryDepartmentGetRequest {
	return &OapiIndustryDepartmentGetRequest{}
}

type OapiIndustryDepartmentGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiIndustryDepartmentGetResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiIndustryDepartmentGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIndustryDepartmentGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIndustryDepartmentGetRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiIndustryDepartmentGetRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiIndustryDepartmentGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.industry.department.get"
}
func (this *OapiIndustryDepartmentGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIndustryDepartmentGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIndustryDepartmentGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIndustryDepartmentGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIndustryDepartmentGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	return txtParams
}
func (this *OapiIndustryDepartmentGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiIndustryDepartmentGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIndustryDepartmentGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiIndustryDepartmentGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  OpenIndustryDeptInfo `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type OpenIndustryDeptInfo struct {
	ContactType string `json:"contact_type,omitempty"`
	DeptType    string `json:"dept_type,omitempty"`
	Feature     string `json:"feature,omitempty"`
	Name        string `json:"name,omitempty"`
	SuperId     int64  `json:"super_id,omitempty"`
}
