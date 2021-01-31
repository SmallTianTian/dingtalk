package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavMicroappDetailListRequest() *OapiKacDatavMicroappDetailListRequest {
	return &OapiKacDatavMicroappDetailListRequest{}
}

type OapiKacDatavMicroappDetailListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavMicroappDetailListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavMicroappDetailListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavMicroappDetailListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavMicroappDetailListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavMicroappDetailListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavMicroappDetailListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.microapp.detail.list"
}
func (this *OapiKacDatavMicroappDetailListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavMicroappDetailListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavMicroappDetailListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavMicroappDetailListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavMicroappDetailListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavMicroappDetailListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavMicroappDetailListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavMicroappDetailListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type MicroAppSummaryRequest struct {
	Cursor int64  `json:"cursor,omitempty"`
	DataId string `json:"data_id,omitempty"`
	Size   int64  `json:"size,omitempty"`
}
type OapiKacDatavMicroappDetailListResponse struct {
	taobao.TaobaoResponse
	Result MicroAppSummaryResponse `json:"result,omitempty"`
}
type MicroAppSummaryVo struct {
	MicroAppName      string `json:"micro_app_name,omitempty"`
	MicroAppType      int64  `json:"micro_app_type,omitempty"`
	MicroAppUserCount int64  `json:"micro_app_user_count,omitempty"`
}
type MicroAppSummaryResponse struct {
	Data       []MicroAppSummaryVo `json:"data,omitempty"`
	HasMore    bool                `json:"has_more,omitempty"`
	NextCursor int64               `json:"next_cursor,omitempty"`
}
