package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduDeptListRequest() *OapiEduDeptListRequest {
	return &OapiEduDeptListRequest{}
}

type OapiEduDeptListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduDeptListResponse
	PageNo          int64
	PageSize        int64
	SuperId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduDeptListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduDeptListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduDeptListRequest) SetPageNo(pageNo2 int64) {
	this.PageNo = pageNo2
}
func (this *OapiEduDeptListRequest) GetPageNo() int64 {
	return this.PageNo
}
func (this *OapiEduDeptListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEduDeptListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEduDeptListRequest) SetSuperId(superId2 int64) {
	this.SuperId = superId2
}
func (this *OapiEduDeptListRequest) GetSuperId() int64 {
	return this.SuperId
}
func (this *OapiEduDeptListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.dept.list"
}
func (this *OapiEduDeptListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduDeptListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduDeptListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduDeptListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduDeptListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["page_no"] = this.PageNo
	txtParams["page_size"] = this.PageSize
	txtParams["super_id"] = this.SuperId
	return txtParams
}
func (this *OapiEduDeptListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.PageNo, "pageNo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduDeptListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduDeptListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduDeptListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                   `json:"errcode,omitempty"`
	Errmsg  string                  `json:"errmsg,omitempty"`
	Result  OpenEduDeptListResponse `json:"result,omitempty"`
	Success bool                    `json:"success,omitempty"`
}
type OpenEduDeptDetails struct {
	Chain       string `json:"chain,omitempty"`
	ContactType string `json:"contact_type,omitempty"`
	DeptId      int64  `json:"dept_id,omitempty"`
	DeptType    string `json:"dept_type,omitempty"`
	Feature     string `json:"feature,omitempty"`
	Name        string `json:"name,omitempty"`
	Nick        string `json:"nick,omitempty"`
}
type OpenEduDeptListResponse struct {
	Details []OpenEduDeptDetails `json:"details,omitempty"`
	HasMore bool                 `json:"has_more,omitempty"`
	SuperId int64                `json:"super_id,omitempty"`
}
