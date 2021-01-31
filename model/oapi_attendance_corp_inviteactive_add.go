package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceCorpInviteactiveAddRequest() *OapiAttendanceCorpInviteactiveAddRequest {
	return &OapiAttendanceCorpInviteactiveAddRequest{}
}

type OapiAttendanceCorpInviteactiveAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceCorpInviteactiveAddResponse
	AdminMobile     string
	InvitedMobile   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceCorpInviteactiveAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) SetAdminMobile(adminMobile2 string) {
	this.AdminMobile = adminMobile2
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) GetAdminMobile() string {
	return this.AdminMobile
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) SetInvitedMobile(invitedMobile2 string) {
	this.InvitedMobile = invitedMobile2
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) GetInvitedMobile() string {
	return this.InvitedMobile
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.corp.inviteactive.add"
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["admin_mobile"] = this.AdminMobile
	txtParams["invited_mobile"] = this.InvitedMobile
	return txtParams
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AdminMobile, "adminMobile"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceCorpInviteactiveAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceCorpInviteactiveAddResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
