package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCustomizeConfigSetRequest() *OapiCustomizeConfigSetRequest {
	return &OapiCustomizeConfigSetRequest{}
}

type OapiCustomizeConfigSetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomizeConfigSetResponse
	ActiveId        string
	ActiveType      string
	Biz             string
	RuleName        string
	TopHttpMethod   string
	TopResponseType string
	Type            string
}

func (this *OapiCustomizeConfigSetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomizeConfigSetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomizeConfigSetRequest) SetActiveId(activeId2 string) {
	this.ActiveId = activeId2
}
func (this *OapiCustomizeConfigSetRequest) GetActiveId() string {
	return this.ActiveId
}
func (this *OapiCustomizeConfigSetRequest) SetActiveType(activeType2 string) {
	this.ActiveType = activeType2
}
func (this *OapiCustomizeConfigSetRequest) GetActiveType() string {
	return this.ActiveType
}
func (this *OapiCustomizeConfigSetRequest) SetBiz(biz2 string) {
	this.Biz = biz2
}
func (this *OapiCustomizeConfigSetRequest) GetBiz() string {
	return this.Biz
}
func (this *OapiCustomizeConfigSetRequest) SetRuleName(ruleName2 string) {
	this.RuleName = ruleName2
}
func (this *OapiCustomizeConfigSetRequest) GetRuleName() string {
	return this.RuleName
}
func (this *OapiCustomizeConfigSetRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiCustomizeConfigSetRequest) GetType() string {
	return this.Type
}
func (this *OapiCustomizeConfigSetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customize.config.set"
}
func (this *OapiCustomizeConfigSetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomizeConfigSetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomizeConfigSetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomizeConfigSetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomizeConfigSetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["active_id"] = this.ActiveId
	txtParams["active_type"] = this.ActiveType
	txtParams["biz"] = this.Biz
	txtParams["rule_name"] = this.RuleName
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiCustomizeConfigSetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ActiveId, "activeId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCustomizeConfigSetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomizeConfigSetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCustomizeConfigSetResponse struct {
	taobao.TaobaoResponse
}
