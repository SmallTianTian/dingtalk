package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserGetOrgUserCountRequest() *OapiUserGetOrgUserCountRequest {
	return &OapiUserGetOrgUserCountRequest{}
}

type OapiUserGetOrgUserCountRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserGetOrgUserCountResponse
	OnlyActive      int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserGetOrgUserCountRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserGetOrgUserCountRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserGetOrgUserCountRequest) SetOnlyActive(onlyActive2 int64) {
	this.OnlyActive = onlyActive2
}
func (this *OapiUserGetOrgUserCountRequest) GetOnlyActive() int64 {
	return this.OnlyActive
}
func (this *OapiUserGetOrgUserCountRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.get_org_user_count"
}
func (this *OapiUserGetOrgUserCountRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserGetOrgUserCountRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserGetOrgUserCountRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserGetOrgUserCountRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserGetOrgUserCountRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["onlyActive"] = this.OnlyActive
	return txtParams
}
func (this *OapiUserGetOrgUserCountRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserGetOrgUserCountRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserGetOrgUserCountRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserGetOrgUserCountResponse struct {
	taobao.TaobaoResponse
	Count   int64  `json:"count,omitempty"`
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
