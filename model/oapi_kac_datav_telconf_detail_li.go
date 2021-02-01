package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavTelconfDetailListRequest() *OapiKacDatavTelconfDetailListRequest {
	return &OapiKacDatavTelconfDetailListRequest{}
}

type OapiKacDatavTelconfDetailListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavTelconfDetailListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavTelconfDetailListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavTelconfDetailListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavTelconfDetailListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavTelconfDetailListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavTelconfDetailListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.telconf.detail.list"
}
func (this *OapiKacDatavTelconfDetailListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavTelconfDetailListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavTelconfDetailListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavTelconfDetailListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavTelconfDetailListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavTelconfDetailListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavTelconfDetailListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavTelconfDetailListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiKacDatavTelconfDetailListResponse struct {
	taobao.TaobaoResponse
	Result TelConferenceDetailResponse `json:"result,omitempty"`
}
type TelConferenceDetailVo struct {
	ConfId        string `json:"conf_id,omitempty"`
	ConfLenMin    string `json:"conf_len_min,omitempty"`
	DeptId        int64  `json:"dept_id,omitempty"`
	DeptName      string `json:"dept_name,omitempty"`
	EndTime       string `json:"end_time,omitempty"`
	JoinUserCount int64  `json:"join_user_count,omitempty"`
	StaffJobNum   string `json:"staff_job_num,omitempty"`
	StaffName     string `json:"staff_name,omitempty"`
	StartTime     string `json:"start_time,omitempty"`
	Userid        string `json:"userid,omitempty"`
}
type TelConferenceDetailResponse struct {
	Data       []TelConferenceDetailVo `json:"data,omitempty"`
	HasMore    bool                    `json:"has_more,omitempty"`
	NextCursor int64                   `json:"next_cursor,omitempty"`
}
