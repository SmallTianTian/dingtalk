package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiOrgUserteaminviteAcceptRequest() *OapiOrgUserteaminviteAcceptRequest {
	return &OapiOrgUserteaminviteAcceptRequest{}
}

type OapiOrgUserteaminviteAcceptRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOrgUserteaminviteAcceptResponse
	Mobile          string
	StateCode       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiOrgUserteaminviteAcceptRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgUserteaminviteAcceptRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgUserteaminviteAcceptRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiOrgUserteaminviteAcceptRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiOrgUserteaminviteAcceptRequest) SetStateCode(stateCode2 string) {
	this.StateCode = stateCode2
}
func (this *OapiOrgUserteaminviteAcceptRequest) GetStateCode() string {
	return this.StateCode
}
func (this *OapiOrgUserteaminviteAcceptRequest) GetApiMethodName() string {
	return "dingtalk.oapi.org.userteaminvite.accept"
}
func (this *OapiOrgUserteaminviteAcceptRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgUserteaminviteAcceptRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgUserteaminviteAcceptRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgUserteaminviteAcceptRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgUserteaminviteAcceptRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["mobile"] = this.Mobile
	txtParams["state_code"] = this.StateCode
	return txtParams
}
func (this *OapiOrgUserteaminviteAcceptRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Mobile, "mobile"); err != nil {
		return err
	}
	return nil
}
func (this *OapiOrgUserteaminviteAcceptRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgUserteaminviteAcceptRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgUserteaminviteAcceptResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
}
