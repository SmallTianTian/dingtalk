package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavDeptVideoliveListRequest() *OapiKacDatavDeptVideoliveListRequest {
	return &OapiKacDatavDeptVideoliveListRequest{}
}

type OapiKacDatavDeptVideoliveListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavDeptVideoliveListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavDeptVideoliveListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavDeptVideoliveListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavDeptVideoliveListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavDeptVideoliveListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavDeptVideoliveListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.dept.videolive.list"
}
func (this *OapiKacDatavDeptVideoliveListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavDeptVideoliveListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavDeptVideoliveListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavDeptVideoliveListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavDeptVideoliveListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavDeptVideoliveListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavDeptVideoliveListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavDeptVideoliveListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type VideoLiveDeptSummaryRequest struct {
	Cursor int64  `json:"cursor,omitempty"`
	DataId string `json:"data_id,omitempty"`
	Size   int64  `json:"size,omitempty"`
}
type OapiKacDatavDeptVideoliveListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                        `json:"errcode,omitempty"`
	Errmsg  string                       `json:"errmsg,omitempty"`
	Result  VideoLiveDeptSummaryResponse `json:"result,omitempty"`
}
type VideoLiveDeptSummaryVo struct {
	DeptId            int64  `json:"dept_id,omitempty"`
	DeptName          string `json:"dept_name,omitempty"`
	LiveLaunchCount   int64  `json:"live_launch_count,omitempty"`
	LivePlayUserCount int64  `json:"live_play_user_count,omitempty"`
	LiveTimeLength    string `json:"live_time_length,omitempty"`
}
type VideoLiveDeptSummaryResponse struct {
	Data       []VideoLiveDeptSummaryVo `json:"data,omitempty"`
	HasMore    bool                     `json:"has_more,omitempty"`
	NextCursor int64                    `json:"next_cursor,omitempty"`
}
