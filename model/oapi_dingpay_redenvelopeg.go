package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingpayRedenvelopeGetRequest() *OapiDingpayRedenvelopeGetRequest {
	return &OapiDingpayRedenvelopeGetRequest{}
}

type OapiDingpayRedenvelopeGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingpayRedenvelopeGetResponse
	CorpBizNo       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingpayRedenvelopeGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayRedenvelopeGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayRedenvelopeGetRequest) SetCorpBizNo(corpBizNo2 string) {
	this.CorpBizNo = corpBizNo2
}
func (this *OapiDingpayRedenvelopeGetRequest) GetCorpBizNo() string {
	return this.CorpBizNo
}
func (this *OapiDingpayRedenvelopeGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.redenvelope.get"
}
func (this *OapiDingpayRedenvelopeGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayRedenvelopeGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayRedenvelopeGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayRedenvelopeGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayRedenvelopeGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corp_biz_no"] = this.CorpBizNo
	return txtParams
}
func (this *OapiDingpayRedenvelopeGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpBizNo, "corpBizNo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingpayRedenvelopeGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayRedenvelopeGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayRedenvelopeGetResponse struct {
	taobao.TaobaoResponse
	Result RedEnvelopeGetResult `json:"result,omitempty"`
}
type RedEnvelopeGetResult struct {
	CorpBizNo    string `json:"corp_biz_no,omitempty"`
	Greetings    string `json:"greetings,omitempty"`
	OrderNo      string `json:"order_no,omitempty"`
	PickAmount   string `json:"pick_amount,omitempty"`
	PickTime     string `json:"pick_time,omitempty"`
	ReceiverId   string `json:"receiver_id,omitempty"`
	RefundAmount string `json:"refund_amount,omitempty"`
	RefundTime   string `json:"refund_time,omitempty"`
	SendTime     string `json:"send_time,omitempty"`
	SenderId     string `json:"sender_id,omitempty"`
	Status       string `json:"status,omitempty"`
	ThemeId      string `json:"theme_id,omitempty"`
	TotalAmount  string `json:"total_amount,omitempty"`
	Type         string `json:"type,omitempty"`
}
