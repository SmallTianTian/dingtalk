package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserSeniorSettingRequest() *OapiUserSeniorSettingRequest {
	return &OapiUserSeniorSettingRequest{}
}

type OapiUserSeniorSettingRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserSeniorSettingResponse
	Open            bool
	PermitStaffIds  string
	ProtectScenes   string
	SeniorStaffId   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserSeniorSettingRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserSeniorSettingRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserSeniorSettingRequest) SetOpen(open2 bool) {
	this.Open = open2
}
func (this *OapiUserSeniorSettingRequest) GetOpen() bool {
	return this.Open
}
func (this *OapiUserSeniorSettingRequest) SetPermitStaffIds(permitStaffIds2 string) {
	this.PermitStaffIds = permitStaffIds2
}
func (this *OapiUserSeniorSettingRequest) GetPermitStaffIds() string {
	return this.PermitStaffIds
}
func (this *OapiUserSeniorSettingRequest) SetProtectScenes(protectScenes2 string) {
	this.ProtectScenes = protectScenes2
}
func (this *OapiUserSeniorSettingRequest) GetProtectScenes() string {
	return this.ProtectScenes
}
func (this *OapiUserSeniorSettingRequest) SetSeniorStaffId(seniorStaffId2 string) {
	this.SeniorStaffId = seniorStaffId2
}
func (this *OapiUserSeniorSettingRequest) GetSeniorStaffId() string {
	return this.SeniorStaffId
}
func (this *OapiUserSeniorSettingRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.senior.setting"
}
func (this *OapiUserSeniorSettingRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserSeniorSettingRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserSeniorSettingRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserSeniorSettingRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserSeniorSettingRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open"] = this.Open
	txtParams["permit_staffIds"] = this.PermitStaffIds
	txtParams["protect_scenes"] = this.ProtectScenes
	txtParams["senior_staffId"] = this.SeniorStaffId
	return txtParams
}
func (this *OapiUserSeniorSettingRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.PermitStaffIds, 999, "permitStaffIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserSeniorSettingRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserSeniorSettingRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserSeniorSettingResponse struct {
	taobao.TaobaoResponse
}
