package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2UserGetuserinfoRequest() *OapiV2UserGetuserinfoRequest {
	return &OapiV2UserGetuserinfoRequest{}
}

type OapiV2UserGetuserinfoRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2UserGetuserinfoResponse
	Code            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiV2UserGetuserinfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2UserGetuserinfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2UserGetuserinfoRequest) SetCode(code2 string) {
	this.Code = code2
}
func (this *OapiV2UserGetuserinfoRequest) GetCode() string {
	return this.Code
}
func (this *OapiV2UserGetuserinfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.user.getuserinfo"
}
func (this *OapiV2UserGetuserinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2UserGetuserinfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2UserGetuserinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2UserGetuserinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2UserGetuserinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.ERROR_CODE] = this.Code
	return txtParams
}
func (this *OapiV2UserGetuserinfoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Code, taobao.ERROR_CODE); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2UserGetuserinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2UserGetuserinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2UserGetuserinfoResponse struct {
	taobao.TaobaoResponse
	Errcode int64                 `json:"errcode,omitempty"`
	Errmsg  string                `json:"errmsg,omitempty"`
	Result  UserGetByCodeResponse `json:"result,omitempty"`
}
type UserGetByCodeResponse struct {
	AssociatedUnionid string `json:"associated_unionid,omitempty"`
	DeviceId          string `json:"device_id,omitempty"`
	Name              string `json:"name,omitempty"`
	Sys               bool   `json:"sys,omitempty"`
	SysLevel          int64  `json:"sys_level,omitempty"`
	Unionid           string `json:"unionid,omitempty"`
	Userid            string `json:"userid,omitempty"`
}
