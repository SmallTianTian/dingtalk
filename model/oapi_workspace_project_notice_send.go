package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiWorkspaceProjectNoticeSendRequest() *OapiWorkspaceProjectNoticeSendRequest {
	return &OapiWorkspaceProjectNoticeSendRequest{}
}

type OapiWorkspaceProjectNoticeSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceProjectNoticeSendResponse
	SendNoticeReq   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceProjectNoticeSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceProjectNoticeSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceProjectNoticeSendRequest) SetSendNoticeReq(sendNoticeReq2 string) {
	this.SendNoticeReq = sendNoticeReq2
}
func (this *OapiWorkspaceProjectNoticeSendRequest) GetSendNoticeReq() string {
	return this.SendNoticeReq
}
func (this *OapiWorkspaceProjectNoticeSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.project.notice.send"
}
func (this *OapiWorkspaceProjectNoticeSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceProjectNoticeSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceProjectNoticeSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceProjectNoticeSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceProjectNoticeSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["send_notice_req"] = this.SendNoticeReq
	return txtParams
}
func (this *OapiWorkspaceProjectNoticeSendRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiWorkspaceProjectNoticeSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceProjectNoticeSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenSendNoticeRequestDto struct {
	Agentid         int64    `json:"agentid,omitempty"`
	MobileUrl       string   `json:"mobile_url,omitempty"`
	MsgButton       string   `json:"msg_button,omitempty"`
	MsgContent      string   `json:"msg_content,omitempty"`
	PcUrl           string   `json:"pc_url,omitempty"`
	ReceiverUserids []string `json:"receiver_userids,omitempty"`
	ShowRedPoint    bool     `json:"show_red_point,omitempty"`
	Uuid            string   `json:"uuid,omitempty"`
}
type OapiWorkspaceProjectNoticeSendResponse struct {
	taobao.TaobaoResponse
	Result  OpenSendNoticeResponseDto `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type OpenSendNoticeResponseDto struct {
	LimitUserids []string `json:"limit_userids,omitempty"`
}
