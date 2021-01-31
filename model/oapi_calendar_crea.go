package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCalendarCreateRequest() *OapiCalendarCreateRequest {
	return &OapiCalendarCreateRequest{}
}

type OapiCalendarCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCalendarCreateResponse
	CreateVo        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCalendarCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCalendarCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCalendarCreateRequest) SetCreateVo(createVo2 string) {
	this.CreateVo = createVo2
}
func (this *OapiCalendarCreateRequest) GetCreateVo() string {
	return this.CreateVo
}
func (this *OapiCalendarCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.calendar.create"
}
func (this *OapiCalendarCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCalendarCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCalendarCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCalendarCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCalendarCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["create_vo"] = this.CreateVo
	return txtParams
}
func (this *OapiCalendarCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCalendarCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCalendarCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCalendarCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64                    `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Result  CorpCalendarCreateResult `json:"result,omitempty"`
}
