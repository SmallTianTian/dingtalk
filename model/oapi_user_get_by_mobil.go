package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserGetByMobileRequest() *OapiUserGetByMobileRequest {
	return &OapiUserGetByMobileRequest{}
}

type OapiUserGetByMobileRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserGetByMobileResponse
	Mobile          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserGetByMobileRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserGetByMobileRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserGetByMobileRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiUserGetByMobileRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiUserGetByMobileRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.get_by_mobile"
}
func (this *OapiUserGetByMobileRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserGetByMobileRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserGetByMobileRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserGetByMobileRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserGetByMobileRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["mobile"] = this.Mobile
	return txtParams
}
func (this *OapiUserGetByMobileRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Mobile, "mobile"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserGetByMobileRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserGetByMobileRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserGetByMobileResponse struct {
	taobao.TaobaoResponse
	Userid string `json:"userid,omitempty"`
}
