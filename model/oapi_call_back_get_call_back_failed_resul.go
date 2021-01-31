package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCallBackGetCallBackFailedResultRequest() *OapiCallBackGetCallBackFailedResultRequest {
	return &OapiCallBackGetCallBackFailedResultRequest{}
}

type OapiCallBackGetCallBackFailedResultRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallBackGetCallBackFailedResultResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCallBackGetCallBackFailedResultRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiCallBackGetCallBackFailedResultRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallBackGetCallBackFailedResultRequest) GetApiMethodName() string {
	return "dingtalk.oapi.call_back.get_call_back_failed_result"
}
func (this *OapiCallBackGetCallBackFailedResultRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallBackGetCallBackFailedResultRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallBackGetCallBackFailedResultRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallBackGetCallBackFailedResultRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallBackGetCallBackFailedResultRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiCallBackGetCallBackFailedResultRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCallBackGetCallBackFailedResultRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallBackGetCallBackFailedResultRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallBackGetCallBackFailedResultResponse struct {
	taobao.TaobaoResponse
	Errcode    int64    `json:"errcode,omitempty"`
	Errmsg     string   `json:"errmsg,omitempty"`
	FailedList []Failed `json:"failed_list,omitempty"`
	HasMore    bool     `json:"has_more,omitempty"`
}
type Failed struct {
	BpmsInstanceChange string `json:"bpms_instance_change,omitempty"`
	BpmsTaskChange     string `json:"bpms_task_change,omitempty"`
	CallBackTag        string `json:"call_back_tag,omitempty"`
	CheckIn            string `json:"check_in,omitempty"`
	Data               string `json:"data,omitempty"`
	EventTime          int64  `json:"event_time,omitempty"`
	OrgAdminAdd        string `json:"org_admin_add,omitempty"`
	OrgAdminRemove     string `json:"org_admin_remove,omitempty"`
	OrgChange          string `json:"org_change,omitempty"`
	OrgDeptCreate      string `json:"org_dept_create,omitempty"`
	OrgDeptModify      string `json:"org_dept_modify,omitempty"`
	OrgDeptRemove      string `json:"org_dept_remove,omitempty"`
	UserAddOrg         string `json:"user_add_org,omitempty"`
	UserLeaveOrg       string `json:"user_leave_org,omitempty"`
	UserModifyOrg      string `json:"user_modify_org,omitempty"`
}
