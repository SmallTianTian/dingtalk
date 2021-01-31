package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserGetbyunionidRequest() *OapiUserGetbyunionidRequest {
	return &OapiUserGetbyunionidRequest{}
}

type OapiUserGetbyunionidRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserGetbyunionidResponse
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiUserGetbyunionidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserGetbyunionidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserGetbyunionidRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiUserGetbyunionidRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiUserGetbyunionidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.getbyunionid"
}
func (this *OapiUserGetbyunionidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserGetbyunionidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserGetbyunionidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserGetbyunionidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserGetbyunionidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiUserGetbyunionidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Unionid, "unionid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserGetbyunionidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserGetbyunionidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserGetbyunionidResponse struct {
	taobao.TaobaoResponse
	Result UserGetByUnionIdResponse `json:"result,omitempty"`
}
type UserGetByUnionIdResponse struct {
	ContactType int64  `json:"contact_type,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
