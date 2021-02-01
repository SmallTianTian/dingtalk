package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiIndustryUserListRequest() *OapiIndustryUserListRequest {
	return &OapiIndustryUserListRequest{}
}

type OapiIndustryUserListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiIndustryUserListResponse
	Cursor          int64
	DeptId          int64
	Role            string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiIndustryUserListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIndustryUserListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIndustryUserListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiIndustryUserListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiIndustryUserListRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiIndustryUserListRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiIndustryUserListRequest) SetRole(role2 string) {
	this.Role = role2
}
func (this *OapiIndustryUserListRequest) GetRole() string {
	return this.Role
}
func (this *OapiIndustryUserListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiIndustryUserListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiIndustryUserListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.industry.user.list"
}
func (this *OapiIndustryUserListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIndustryUserListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIndustryUserListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIndustryUserListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIndustryUserListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["dept_id"] = this.DeptId
	txtParams["role"] = this.Role
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiIndustryUserListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiIndustryUserListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIndustryUserListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiIndustryUserListResponse struct {
	taobao.TaobaoResponse
	Result  ResultWrapper `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
