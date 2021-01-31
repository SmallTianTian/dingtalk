package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiIndustryDepartmentListRequest() *OapiIndustryDepartmentListRequest {
	return &OapiIndustryDepartmentListRequest{}
}

type OapiIndustryDepartmentListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiIndustryDepartmentListResponse
	Cursor          int64
	DeptId          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiIndustryDepartmentListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIndustryDepartmentListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIndustryDepartmentListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiIndustryDepartmentListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiIndustryDepartmentListRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiIndustryDepartmentListRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiIndustryDepartmentListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiIndustryDepartmentListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiIndustryDepartmentListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.industry.department.list"
}
func (this *OapiIndustryDepartmentListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIndustryDepartmentListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIndustryDepartmentListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIndustryDepartmentListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIndustryDepartmentListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["dept_id"] = this.DeptId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiIndustryDepartmentListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiIndustryDepartmentListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIndustryDepartmentListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiIndustryDepartmentListResponse struct {
	taobao.TaobaoResponse
	Errcode int64         `json:"errcode,omitempty"`
	Errmsg  string        `json:"errmsg,omitempty"`
	Result  ResultWrapper `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
type ResultWrapper struct {
	Details    []OpenIndustryDeptInfo `json:"details,omitempty"`
	HasMore    bool                   `json:"has_more,omitempty"`
	NextCursor int64                  `json:"next_cursor,omitempty"`
}
