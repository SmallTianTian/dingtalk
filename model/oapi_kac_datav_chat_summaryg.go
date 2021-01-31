package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavChatSummaryGetRequest() *OapiKacDatavChatSummaryGetRequest {
	return &OapiKacDatavChatSummaryGetRequest{}
}

type OapiKacDatavChatSummaryGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavChatSummaryGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavChatSummaryGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavChatSummaryGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavChatSummaryGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavChatSummaryGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavChatSummaryGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.chat.summary.get"
}
func (this *OapiKacDatavChatSummaryGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavChatSummaryGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavChatSummaryGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavChatSummaryGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavChatSummaryGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavChatSummaryGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavChatSummaryGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavChatSummaryGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ImSummaryRequest struct {
	DataId string `json:"data_id,omitempty"`
}
type OapiKacDatavChatSummaryGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
	Result  ImSummaryResponse `json:"result,omitempty"`
}
type ImSummaryResponse struct {
	ActiveGroupCount   int64  `json:"active_group_count,omitempty"`
	ChatUserCount      int64  `json:"chat_user_count,omitempty"`
	GroupChatUserCount int64  `json:"group_chat_user_count,omitempty"`
	GroupCount         int64  `json:"group_count,omitempty"`
	GroupMessageCount  int64  `json:"group_message_count,omitempty"`
	MessageAvgCount    string `json:"message_avg_count,omitempty"`
	MessageTotalCount  int64  `json:"message_total_count,omitempty"`
	MessageUserCount   int64  `json:"message_user_count,omitempty"`
	NewGroupCount      int64  `json:"new_group_count,omitempty"`
	SingleMessageCount int64  `json:"single_message_count,omitempty"`
}
