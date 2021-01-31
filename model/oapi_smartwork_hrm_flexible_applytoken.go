package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmFlexibleApplytokenRequest() *OapiSmartworkHrmFlexibleApplytokenRequest {
	return &OapiSmartworkHrmFlexibleApplytokenRequest{}
}

type OapiSmartworkHrmFlexibleApplytokenRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmFlexibleApplytokenResponse
	Agentid         int64
	OptUserId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmFlexibleApplytokenRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) SetOptUserId(optUserId2 string) {
	this.OptUserId = optUserId2
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) GetOptUserId() string {
	return this.OptUserId
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.flexible.applytoken"
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["opt_user_id"] = this.OptUserId
	return txtParams
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmFlexibleApplytokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmFlexibleApplytokenResponse struct {
	taobao.TaobaoResponse
	Errcode int64           `json:"errcode,omitempty"`
	Errmsg  string          `json:"errmsg,omitempty"`
	Result  FlexibleTokenVo `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
type FlexibleTokenVo struct {
	Token string `json:"token,omitempty"`
}
