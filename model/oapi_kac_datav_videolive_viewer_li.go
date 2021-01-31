package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavVideoliveViewerListRequest() *OapiKacDatavVideoliveViewerListRequest {
	return &OapiKacDatavVideoliveViewerListRequest{}
}

type OapiKacDatavVideoliveViewerListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavVideoliveViewerListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavVideoliveViewerListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavVideoliveViewerListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavVideoliveViewerListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavVideoliveViewerListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavVideoliveViewerListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.videolive.viewer.list"
}
func (this *OapiKacDatavVideoliveViewerListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavVideoliveViewerListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavVideoliveViewerListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavVideoliveViewerListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavVideoliveViewerListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavVideoliveViewerListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavVideoliveViewerListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavVideoliveViewerListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GroupLiveViewerReq struct {
	Cid      string `json:"cid,omitempty"`
	Cursor   int64  `json:"cursor,omitempty"`
	LiveUuid string `json:"live_uuid,omitempty"`
	Size     int64  `json:"size,omitempty"`
}
type OapiKacDatavVideoliveViewerListResponse struct {
	taobao.TaobaoResponse
	Result GroupLiveViewerPageResult `json:"result,omitempty"`
}
type GroupLiveViewer struct {
	PlayDuration          int64  `json:"play_duration,omitempty"`
	PlayDurationMin       string `json:"play_duration_min,omitempty"`
	PlayRecordDuration    int64  `json:"play_record_duration,omitempty"`
	PlayRecordDurationMin string `json:"play_record_duration_min,omitempty"`
	Userid                string `json:"userid,omitempty"`
}
type GroupLiveViewerPageResult struct {
	Data       []GroupLiveViewer `json:"data,omitempty"`
	HasMore    bool              `json:"has_more,omitempty"`
	NextCursor int64             `json:"next_cursor,omitempty"`
}
