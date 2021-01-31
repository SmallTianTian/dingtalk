package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceCorpInviteactiveOpenRequest() *OapiAttendanceCorpInviteactiveOpenRequest {
	return &OapiAttendanceCorpInviteactiveOpenRequest{}
}

type OapiAttendanceCorpInviteactiveOpenRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceCorpInviteactiveOpenResponse
	AdminName       string
	AdminPhone      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceCorpInviteactiveOpenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) SetAdminName(adminName2 string) {
	this.AdminName = adminName2
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) GetAdminName() string {
	return this.AdminName
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) SetAdminPhone(adminPhone2 string) {
	this.AdminPhone = adminPhone2
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) GetAdminPhone() string {
	return this.AdminPhone
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.corp.inviteactive.open"
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["admin_name"] = this.AdminName
	txtParams["admin_phone"] = this.AdminPhone
	return txtParams
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AdminName, "adminName"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceCorpInviteactiveOpenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceCorpInviteactiveOpenResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
