package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringAllowanceUnfreezeBetaRequest() *OapiCateringAllowanceUnfreezeBetaRequest {
	return &OapiCateringAllowanceUnfreezeBetaRequest{}
}

type OapiCateringAllowanceUnfreezeBetaRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringAllowanceUnfreezeBetaResponse
	OrderId         string
	RuleCode        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) SetOrderId(orderId2 string) {
	this.OrderId = orderId2
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetOrderId() string {
	return this.OrderId
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) SetRuleCode(ruleCode2 string) {
	this.RuleCode = ruleCode2
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetRuleCode() string {
	return this.RuleCode
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.allowance.unfreeze.beta"
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["order_id"] = this.OrderId
	txtParams["rule_code"] = this.RuleCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OrderId, "orderId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringAllowanceUnfreezeBetaRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringAllowanceUnfreezeBetaResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
