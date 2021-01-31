package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingmiCommonLoginAccesstokenRequest() *OapiDingmiCommonLoginAccesstokenRequest {
	return &OapiDingmiCommonLoginAccesstokenRequest{}
}

type OapiDingmiCommonLoginAccesstokenRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingmiCommonLoginAccesstokenResponse
	ForeignId       string
	Nick            string
	Source          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingmiCommonLoginAccesstokenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) SetForeignId(foreignId2 string) {
	this.ForeignId = foreignId2
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) GetForeignId() string {
	return this.ForeignId
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) SetNick(nick2 string) {
	this.Nick = nick2
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) GetNick() string {
	return this.Nick
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) SetSource(source2 int64) {
	this.Source = source2
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) GetSource() int64 {
	return this.Source
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingmi.common.login.accesstoken"
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["foreign_id"] = this.ForeignId
	txtParams["nick"] = this.Nick
	txtParams["source"] = this.Source
	return txtParams
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingmiCommonLoginAccesstokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingmiCommonLoginAccesstokenResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
