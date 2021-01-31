package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavDauSummaryGetRequest() *OapiKacDatavDauSummaryGetRequest {
	return &OapiKacDatavDauSummaryGetRequest{}
}

type OapiKacDatavDauSummaryGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavDauSummaryGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavDauSummaryGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavDauSummaryGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavDauSummaryGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavDauSummaryGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavDauSummaryGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.dau.summary.get"
}
func (this *OapiKacDatavDauSummaryGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavDauSummaryGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavDauSummaryGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavDauSummaryGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavDauSummaryGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavDauSummaryGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavDauSummaryGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavDauSummaryGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DauSummaryRequest struct {
	DataId string `json:"data_id,omitempty"`
}
type OapiKacDatavDauSummaryGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64              `json:"errcode,omitempty"`
	Errmsg  string             `json:"errmsg,omitempty"`
	Result  DauSummaryResponse `json:"result,omitempty"`
}
type DauSummaryResponse struct {
	ActivatedCount   int64 `json:"activated_count,omitempty"`
	AppActiveUsers   int64 `json:"app_active_users,omitempty"`
	ChatUserCount    int64 `json:"chat_user_count,omitempty"`
	ContactsCount    int64 `json:"contacts_count,omitempty"`
	DailyActiveUsers int64 `json:"daily_active_users,omitempty"`
	PcActiveUsers    int64 `json:"pc_active_users,omitempty"`
}
