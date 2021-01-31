package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserCountRequest() *OapiUserCountRequest {
	return &OapiUserCountRequest{}
}

type OapiUserCountRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserCountResponse
	OnlyActive      bool
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserCountRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserCountRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserCountRequest) SetOnlyActive(onlyActive2 bool) {
	this.OnlyActive = onlyActive2
}
func (this *OapiUserCountRequest) GetOnlyActive() bool {
	return this.OnlyActive
}
func (this *OapiUserCountRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.count"
}
func (this *OapiUserCountRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserCountRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserCountRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserCountRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserCountRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["only_active"] = this.OnlyActive
	return txtParams
}
func (this *OapiUserCountRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OnlyActive, "onlyActive"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserCountRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserCountRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserCountResponse struct {
	taobao.TaobaoResponse
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
	Result  CountUserResponse `json:"result,omitempty"`
}
type CountUserResponse struct {
	Count int64 `json:"count,omitempty"`
}
