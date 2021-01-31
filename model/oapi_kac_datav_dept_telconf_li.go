package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavDeptTelconfListRequest() *OapiKacDatavDeptTelconfListRequest {
	return &OapiKacDatavDeptTelconfListRequest{}
}

type OapiKacDatavDeptTelconfListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavDeptTelconfListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavDeptTelconfListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavDeptTelconfListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavDeptTelconfListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavDeptTelconfListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavDeptTelconfListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.dept.telconf.list"
}
func (this *OapiKacDatavDeptTelconfListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavDeptTelconfListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavDeptTelconfListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavDeptTelconfListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavDeptTelconfListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavDeptTelconfListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavDeptTelconfListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavDeptTelconfListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TelConferenceSummaryRequest struct {
	Cursor int64  `json:"cursor,omitempty"`
	DataId string `json:"data_id,omitempty"`
	Size   int64  `json:"size,omitempty"`
}
type OapiKacDatavDeptTelconfListResponse struct {
	taobao.TaobaoResponse
	Result TelConferenceDeptSummaryResponse `json:"result,omitempty"`
}
type TelConferenceDeptSummaryVo struct {
	DeptId         int64  `json:"dept_id,omitempty"`
	DeptName       string `json:"dept_name,omitempty"`
	JoinCount      int64  `json:"join_count,omitempty"`
	StartAvgLenMin string `json:"start_avg_len_min,omitempty"`
	StartCount     int64  `json:"start_count,omitempty"`
	StartLenMin    string `json:"start_len_min,omitempty"`
}
type TelConferenceDeptSummaryResponse struct {
	Data       []TelConferenceDeptSummaryVo `json:"data,omitempty"`
	HasMore    bool                         `json:"has_more,omitempty"`
	NextCursor int64                        `json:"next_cursor,omitempty"`
}
