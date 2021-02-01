package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiInspectTaskListRequest() *OapiInspectTaskListRequest {
	return &OapiInspectTaskListRequest{}
}

type OapiInspectTaskListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiInspectTaskListResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiInspectTaskListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiInspectTaskListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiInspectTaskListRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiInspectTaskListRequest) GetParam() string {
	return this.Param
}
func (this *OapiInspectTaskListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.inspect.task.list"
}
func (this *OapiInspectTaskListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiInspectTaskListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiInspectTaskListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiInspectTaskListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiInspectTaskListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiInspectTaskListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiInspectTaskListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiInspectTaskListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopInspectTaskDeptQueryPram struct {
	Cursor    int64   `json:"cursor,omitempty"`
	DeptId    string  `json:"dept_id,omitempty"`
	EndDate   int64   `json:"end_date,omitempty"`
	Size      int64   `json:"size,omitempty"`
	StartDate int64   `json:"start_date,omitempty"`
	Status    []int64 `json:"status,omitempty"`
}
type OapiInspectTaskListResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
type TopInspectTaskVo struct {
	CheckInTime  int64  `json:"check_in_time,omitempty"`
	CheckOutTime int64  `json:"check_out_time,omitempty"`
	Duration     string `json:"duration,omitempty"`
	PositionId   string `json:"position_id,omitempty"`
	PositionName string `json:"position_name,omitempty"`
	Status       string `json:"status,omitempty"`
	TaskId       string `json:"task_id,omitempty"`
	Userid       string `json:"userid,omitempty"`
	WorkDate     int64  `json:"work_date,omitempty"`
}
