package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2UserGetbymobileRequest() *OapiV2UserGetbymobileRequest {
	return &OapiV2UserGetbymobileRequest{}
}

type OapiV2UserGetbymobileRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2UserGetbymobileResponse
	Mobile          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiV2UserGetbymobileRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2UserGetbymobileRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2UserGetbymobileRequest) SetMobile(mobile2 string) {
	this.Mobile = mobile2
}
func (this *OapiV2UserGetbymobileRequest) GetMobile() string {
	return this.Mobile
}
func (this *OapiV2UserGetbymobileRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.user.getbymobile"
}
func (this *OapiV2UserGetbymobileRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2UserGetbymobileRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2UserGetbymobileRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2UserGetbymobileRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2UserGetbymobileRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["mobile"] = this.Mobile
	return txtParams
}
func (this *OapiV2UserGetbymobileRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Mobile, "mobile"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2UserGetbymobileRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2UserGetbymobileRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2UserGetbymobileResponse struct {
	taobao.TaobaoResponse
	Errcode int64                   `json:"errcode,omitempty"`
	Errmsg  string                  `json:"errmsg,omitempty"`
	Result  UserGetByMobileResponse `json:"result,omitempty"`
}
type UserGetByMobileResponse struct {
	Userid string `json:"userid,omitempty"`
}
