package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiServiceaccountMenuUpdateRequest() *OapiServiceaccountMenuUpdateRequest {
	return &OapiServiceaccountMenuUpdateRequest{}
}

type OapiServiceaccountMenuUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceaccountMenuUpdateResponse
	Menu            string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiServiceaccountMenuUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceaccountMenuUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceaccountMenuUpdateRequest) SetMenu(menu2 string) {
	this.Menu = menu2
}
func (this *OapiServiceaccountMenuUpdateRequest) GetMenu() string {
	return this.Menu
}
func (this *OapiServiceaccountMenuUpdateRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiServiceaccountMenuUpdateRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiServiceaccountMenuUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.serviceaccount.menu.update"
}
func (this *OapiServiceaccountMenuUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceaccountMenuUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceaccountMenuUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceaccountMenuUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceaccountMenuUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["menu"] = this.Menu
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiServiceaccountMenuUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Unionid, "unionid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiServiceaccountMenuUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceaccountMenuUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceaccountMenuUpdateResponse struct {
	taobao.TaobaoResponse
}
