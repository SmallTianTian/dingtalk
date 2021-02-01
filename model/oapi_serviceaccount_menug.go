package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiServiceaccountMenuGetRequest() *OapiServiceaccountMenuGetRequest {
	return &OapiServiceaccountMenuGetRequest{}
}

type OapiServiceaccountMenuGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceaccountMenuGetResponse
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiServiceaccountMenuGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceaccountMenuGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceaccountMenuGetRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiServiceaccountMenuGetRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiServiceaccountMenuGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.serviceaccount.menu.get"
}
func (this *OapiServiceaccountMenuGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceaccountMenuGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceaccountMenuGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceaccountMenuGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceaccountMenuGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiServiceaccountMenuGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Unionid, "unionid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiServiceaccountMenuGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceaccountMenuGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceaccountMenuGetResponse struct {
	taobao.TaobaoResponse
	Menu MenuConfigDTO `json:"menu,omitempty"`
}
type MenuSubButtonDTO struct {
	Key     string `json:"key,omitempty"`
	MediaId string `json:"media_id,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Url     string `json:"url,omitempty"`
}
type MenuButtonDTO struct {
	Key       string             `json:"key,omitempty"`
	MediaId   string             `json:"media_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	SubButton []MenuSubButtonDTO `json:"sub_button,omitempty"`
	Type      string             `json:"type,omitempty"`
	Url       string             `json:"url,omitempty"`
}
type MenuConfigDTO struct {
	Button      []MenuButtonDTO `json:"button,omitempty"`
	EnableInput bool            `json:"enable_input,omitempty"`
	Status      int64           `json:"status,omitempty"`
}
