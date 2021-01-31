package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavVideoconfDetailListRequest() *OapiKacDatavVideoconfDetailListRequest {
	return &OapiKacDatavVideoconfDetailListRequest{}
}

type OapiKacDatavVideoconfDetailListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavVideoconfDetailListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavVideoconfDetailListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavVideoconfDetailListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavVideoconfDetailListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavVideoconfDetailListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavVideoconfDetailListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.videoconf.detail.list"
}
func (this *OapiKacDatavVideoconfDetailListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavVideoconfDetailListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavVideoconfDetailListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavVideoconfDetailListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavVideoconfDetailListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavVideoconfDetailListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavVideoconfDetailListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavVideoconfDetailListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiKacDatavVideoconfDetailListResponse struct {
	taobao.TaobaoResponse
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
	Result  McsDetailResponse `json:"result,omitempty"`
}
type McsDetailVo struct {
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
type McsDetailResponse struct {
	Data       []McsDetailVo `json:"data,omitempty"`
	HasMore    bool          `json:"has_more,omitempty"`
	NextCursor int64         `json:"next_cursor,omitempty"`
}
