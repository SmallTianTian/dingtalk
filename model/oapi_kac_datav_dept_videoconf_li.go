package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavDeptVideoconfListRequest() *OapiKacDatavDeptVideoconfListRequest {
	return &OapiKacDatavDeptVideoconfListRequest{}
}

type OapiKacDatavDeptVideoconfListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavDeptVideoconfListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavDeptVideoconfListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavDeptVideoconfListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavDeptVideoconfListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavDeptVideoconfListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavDeptVideoconfListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.dept.videoconf.list"
}
func (this *OapiKacDatavDeptVideoconfListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavDeptVideoconfListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavDeptVideoconfListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavDeptVideoconfListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavDeptVideoconfListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavDeptVideoconfListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavDeptVideoconfListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavDeptVideoconfListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type McsSummaryRequest struct {
	Cursor int64  `json:"cursor,omitempty"`
	DataId string `json:"data_id,omitempty"`
	Size   int64  `json:"size,omitempty"`
}
type OapiKacDatavDeptVideoconfListResponse struct {
	taobao.TaobaoResponse
	Result McsDeptSummaryResponse `json:"result,omitempty"`
}
type McsDeptSummaryVo struct {
	DeptId         int64  `json:"dept_id,omitempty"`
	DeptName       string `json:"dept_name,omitempty"`
	JoinCount      int64  `json:"join_count,omitempty"`
	StartAvgLenMin string `json:"start_avg_len_min,omitempty"`
	StartCount     int64  `json:"start_count,omitempty"`
	StartLenMin    string `json:"start_len_min,omitempty"`
}
type McsDeptSummaryResponse struct {
	Data       []McsDeptSummaryVo `json:"data,omitempty"`
	HasMore    bool               `json:"has_more,omitempty"`
	NextCursor int64              `json:"next_cursor,omitempty"`
}
