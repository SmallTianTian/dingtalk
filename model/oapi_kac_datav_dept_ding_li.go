package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavDeptDingListRequest() *OapiKacDatavDeptDingListRequest {
	return &OapiKacDatavDeptDingListRequest{}
}

type OapiKacDatavDeptDingListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavDeptDingListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavDeptDingListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavDeptDingListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavDeptDingListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavDeptDingListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavDeptDingListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.dept.ding.list"
}
func (this *OapiKacDatavDeptDingListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavDeptDingListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavDeptDingListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavDeptDingListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavDeptDingListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavDeptDingListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavDeptDingListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavDeptDingListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DingUsageSummaryRequest struct {
	Cursor int64  `json:"cursor,omitempty"`
	DataId string `json:"data_id,omitempty"`
	Size   int64  `json:"size,omitempty"`
}
type OapiKacDatavDeptDingListResponse struct {
	taobao.TaobaoResponse
	Result DingUsageDeptSummaryResponse `json:"result,omitempty"`
}
type DingUsageDeptSummaryVo struct {
	DeptId          int64  `json:"dept_id,omitempty"`
	DeptName        string `json:"dept_name,omitempty"`
	DingAppCnt      int64  `json:"ding_app_cnt,omitempty"`
	DingAppUserCnt  int64  `json:"ding_app_user_cnt,omitempty"`
	DingCallCnt     int64  `json:"ding_call_cnt,omitempty"`
	DingCallUserCnt int64  `json:"ding_call_user_cnt,omitempty"`
	DingCnt         int64  `json:"ding_cnt,omitempty"`
	DingSmsCnt      int64  `json:"ding_sms_cnt,omitempty"`
	DingSmsUserCnt  int64  `json:"ding_sms_user_cnt,omitempty"`
	DingUserCnt     int64  `json:"ding_user_cnt,omitempty"`
	DingVoipCnt     int64  `json:"ding_voip_cnt,omitempty"`
	DingVoipUserCnt int64  `json:"ding_voip_user_cnt,omitempty"`
}
type DingUsageDeptSummaryResponse struct {
	Data       []DingUsageDeptSummaryVo `json:"data,omitempty"`
	HasMore    bool                     `json:"has_more,omitempty"`
	NextCursor int64                    `json:"next_cursor,omitempty"`
}
