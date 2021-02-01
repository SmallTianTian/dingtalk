package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserGetUseridByUnionidRequest() *OapiUserGetUseridByUnionidRequest {
	return &OapiUserGetUseridByUnionidRequest{}
}

type OapiUserGetUseridByUnionidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserGetUseridByUnionidResponse
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiUserGetUseridByUnionidRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserGetUseridByUnionidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserGetUseridByUnionidRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiUserGetUseridByUnionidRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiUserGetUseridByUnionidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.getUseridByUnionid"
}
func (this *OapiUserGetUseridByUnionidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserGetUseridByUnionidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserGetUseridByUnionidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserGetUseridByUnionidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserGetUseridByUnionidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiUserGetUseridByUnionidRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserGetUseridByUnionidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserGetUseridByUnionidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserGetUseridByUnionidResponse struct {
	taobao.TaobaoResponse
	ContactType int64 `json:"contactType,omitempty"`

	Userid string `json:"userid,omitempty"`
}
