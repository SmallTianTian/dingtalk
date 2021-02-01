package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavVideoliveDetailListRequest() *OapiKacDatavVideoliveDetailListRequest {
	return &OapiKacDatavVideoliveDetailListRequest{}
}

type OapiKacDatavVideoliveDetailListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavVideoliveDetailListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavVideoliveDetailListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavVideoliveDetailListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavVideoliveDetailListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavVideoliveDetailListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavVideoliveDetailListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.videolive.detail.list"
}
func (this *OapiKacDatavVideoliveDetailListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavVideoliveDetailListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavVideoliveDetailListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavVideoliveDetailListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavVideoliveDetailListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavVideoliveDetailListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavVideoliveDetailListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavVideoliveDetailListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type VideoLiveDetailRequest struct {
	Cursor int64  `json:"cursor,omitempty"`
	DataId string `json:"data_id,omitempty"`
	Size   int64  `json:"size,omitempty"`
}
type OapiKacDatavVideoliveDetailListResponse struct {
	taobao.TaobaoResponse
	Result VideoLiveDetailResponse `json:"result,omitempty"`
}
type VideoLiveDetailVo struct {
	Cid                  string `json:"cid,omitempty"`
	DeptId               int64  `json:"dept_id,omitempty"`
	DeptName             string `json:"dept_name,omitempty"`
	GroupName            string `json:"group_name,omitempty"`
	GroupUserCount       int64  `json:"group_user_count,omitempty"`
	LiveWatchCount       int64  `json:"live_watch_count,omitempty"`
	LiveWatchDuration    string `json:"live_watch_duration,omitempty"`
	LiveWatchDurationMin string `json:"live_watch_duration_min,omitempty"`
	LiveWatchEndTime     string `json:"live_watch_end_time,omitempty"`
	LiveWatchStartTime   string `json:"live_watch_start_time,omitempty"`
	LiveWatchTitle       string `json:"live_watch_title,omitempty"`
	LiveWatchUserCount   int64  `json:"live_watch_user_count,omitempty"`
	StaffJobNum          string `json:"staff_job_num,omitempty"`
	StaffName            string `json:"staff_name,omitempty"`
	Userid               string `json:"userid,omitempty"`
	Uuid                 string `json:"uuid,omitempty"`
}
type VideoLiveDetailResponse struct {
	Data       []VideoLiveDetailVo `json:"data,omitempty"`
	HasMore    bool                `json:"has_more,omitempty"`
	NextCursor int64               `json:"next_cursor,omitempty"`
}
