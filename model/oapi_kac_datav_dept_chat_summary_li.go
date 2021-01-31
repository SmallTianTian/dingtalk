package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavDeptChatSummaryListRequest() *OapiKacDatavDeptChatSummaryListRequest {
	return &OapiKacDatavDeptChatSummaryListRequest{}
}

type OapiKacDatavDeptChatSummaryListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavDeptChatSummaryListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavDeptChatSummaryListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavDeptChatSummaryListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavDeptChatSummaryListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavDeptChatSummaryListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavDeptChatSummaryListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.dept.chat.summary.list"
}
func (this *OapiKacDatavDeptChatSummaryListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavDeptChatSummaryListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavDeptChatSummaryListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavDeptChatSummaryListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavDeptChatSummaryListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavDeptChatSummaryListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavDeptChatSummaryListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavDeptChatSummaryListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiKacDatavDeptChatSummaryListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                 `json:"errcode,omitempty"`
	Errmsg  string                `json:"errmsg,omitempty"`
	Result  ImDeptSummaryResponse `json:"result,omitempty"`
}
type ImDeptSummaryVo struct {
	DeptId                   int64  `json:"dept_id,omitempty"`
	DeptName                 string `json:"dept_name,omitempty"`
	GroupMessageCnt          int64  `json:"group_message_cnt,omitempty"`
	MessageCnt               int64  `json:"message_cnt,omitempty"`
	SendGroupMessageUserCnt  int64  `json:"send_group_message_user_cnt,omitempty"`
	SendMessageAvgCnt        string `json:"send_message_avg_cnt,omitempty"`
	SendMessageUserCnt       int64  `json:"send_message_user_cnt,omitempty"`
	SendSingleMessageUserCnt int64  `json:"send_single_message_user_cnt,omitempty"`
	SingleMessageCnt         int64  `json:"single_message_cnt,omitempty"`
}
type ImDeptSummaryResponse struct {
	Data       []ImDeptSummaryVo `json:"data,omitempty"`
	HasMore    bool              `json:"has_more,omitempty"`
	NextCursor int64             `json:"next_cursor,omitempty"`
}
