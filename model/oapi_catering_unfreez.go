package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringUnfreezeRequest() *OapiCateringUnfreezeRequest {
	return &OapiCateringUnfreezeRequest{}
}

type OapiCateringUnfreezeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringUnfreezeResponse
	OrderId         string
	RuleCode        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCateringUnfreezeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringUnfreezeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringUnfreezeRequest) SetOrderId(orderId2 string) {
	this.OrderId = orderId2
}
func (this *OapiCateringUnfreezeRequest) GetOrderId() string {
	return this.OrderId
}
func (this *OapiCateringUnfreezeRequest) SetRuleCode(ruleCode2 string) {
	this.RuleCode = ruleCode2
}
func (this *OapiCateringUnfreezeRequest) GetRuleCode() string {
	return this.RuleCode
}
func (this *OapiCateringUnfreezeRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCateringUnfreezeRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCateringUnfreezeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.unfreeze"
}
func (this *OapiCateringUnfreezeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringUnfreezeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringUnfreezeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringUnfreezeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringUnfreezeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["order_id"] = this.OrderId
	txtParams["rule_code"] = this.RuleCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCateringUnfreezeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OrderId, "orderId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringUnfreezeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringUnfreezeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringUnfreezeResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
}
