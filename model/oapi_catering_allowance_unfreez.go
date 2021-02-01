package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringAllowanceUnfreezeRequest() *OapiCateringAllowanceUnfreezeRequest {
	return &OapiCateringAllowanceUnfreezeRequest{}
}

type OapiCateringAllowanceUnfreezeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringAllowanceUnfreezeResponse
	OrderId         string
	RuleCode        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCateringAllowanceUnfreezeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringAllowanceUnfreezeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringAllowanceUnfreezeRequest) SetOrderId(orderId2 string) {
	this.OrderId = orderId2
}
func (this *OapiCateringAllowanceUnfreezeRequest) GetOrderId() string {
	return this.OrderId
}
func (this *OapiCateringAllowanceUnfreezeRequest) SetRuleCode(ruleCode2 string) {
	this.RuleCode = ruleCode2
}
func (this *OapiCateringAllowanceUnfreezeRequest) GetRuleCode() string {
	return this.RuleCode
}
func (this *OapiCateringAllowanceUnfreezeRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringAllowanceUnfreezeRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringAllowanceUnfreezeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.allowance.unfreeze"
}
func (this *OapiCateringAllowanceUnfreezeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringAllowanceUnfreezeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringAllowanceUnfreezeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringAllowanceUnfreezeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringAllowanceUnfreezeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["order_id"] = this.OrderId
	txtParams["rule_code"] = this.RuleCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringAllowanceUnfreezeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OrderId, "orderId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringAllowanceUnfreezeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringAllowanceUnfreezeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringAllowanceUnfreezeResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
