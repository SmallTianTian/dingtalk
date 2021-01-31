package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavDeptDauListRequest() *OapiKacDatavDeptDauListRequest {
	return &OapiKacDatavDeptDauListRequest{}
}

type OapiKacDatavDeptDauListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavDeptDauListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavDeptDauListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavDeptDauListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavDeptDauListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavDeptDauListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavDeptDauListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.dept.dau.list"
}
func (this *OapiKacDatavDeptDauListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavDeptDauListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavDeptDauListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavDeptDauListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavDeptDauListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavDeptDauListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavDeptDauListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavDeptDauListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DeptDauSummaryRequest struct {
	Cursor int64  `json:"cursor,omitempty"`
	DataId string `json:"data_id,omitempty"`
	Size   int64  `json:"size,omitempty"`
}
type OapiKacDatavDeptDauListResponse struct {
	taobao.TaobaoResponse
	Result DeptDauSummaryResponse `json:"result,omitempty"`
}
type DeptDauSummaryVo struct {
	AppActiveUsers   int64  `json:"app_active_users,omitempty"`
	ContactsNumber   int64  `json:"contacts_number,omitempty"`
	DailyActiveUsers int64  `json:"daily_active_users,omitempty"`
	DeptId           int64  `json:"dept_id,omitempty"`
	DeptName         string `json:"dept_name,omitempty"`
	PcActiveUsers    int64  `json:"pc_active_users,omitempty"`
}
type DeptDauSummaryResponse struct {
	Data       []DeptDauSummaryVo `json:"data,omitempty"`
	HasMore    bool               `json:"has_more,omitempty"`
	NextCursor int64              `json:"next_cursor,omitempty"`
}
