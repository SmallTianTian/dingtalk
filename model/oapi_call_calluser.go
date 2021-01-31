package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCallCalluserRequest() *OapiCallCalluserRequest {
	return &OapiCallCalluserRequest{}
}

type OapiCallCalluserRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallCalluserResponse
	AuthedCorpId    string
	AuthedStaffId   string
	StaffId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCallCalluserRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCallCalluserRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallCalluserRequest) SetAuthedCorpId(authedCorpId2 string) {
	this.AuthedCorpId = authedCorpId2
}
func (this *OapiCallCalluserRequest) GetAuthedCorpId() string {
	return this.AuthedCorpId
}
func (this *OapiCallCalluserRequest) SetAuthedStaffId(authedStaffId2 string) {
	this.AuthedStaffId = authedStaffId2
}
func (this *OapiCallCalluserRequest) GetAuthedStaffId() string {
	return this.AuthedStaffId
}
func (this *OapiCallCalluserRequest) SetStaffId(staffId2 string) {
	this.StaffId = staffId2
}
func (this *OapiCallCalluserRequest) GetStaffId() string {
	return this.StaffId
}
func (this *OapiCallCalluserRequest) GetApiMethodName() string {
	return "dingtalk.oapi.call.calluser"
}
func (this *OapiCallCalluserRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallCalluserRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallCalluserRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallCalluserRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallCalluserRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["authed_corp_id"] = this.AuthedCorpId
	txtParams["authed_staff_id"] = this.AuthedStaffId
	txtParams["staff_id"] = this.StaffId
	return txtParams
}
func (this *OapiCallCalluserRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AuthedCorpId, "authedCorpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCallCalluserRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallCalluserRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallCalluserResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
