package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiDingpayOrderSyncstatusRequest() *OapiDingpayOrderSyncstatusRequest {
	return &OapiDingpayOrderSyncstatusRequest{}
}

type OapiDingpayOrderSyncstatusRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingpayOrderSyncstatusResponse
	OrderNos        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingpayOrderSyncstatusRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayOrderSyncstatusRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayOrderSyncstatusRequest) SetOrderNos(orderNos2 string) {
	this.OrderNos = orderNos2
}
func (this *OapiDingpayOrderSyncstatusRequest) GetOrderNos() string {
	return this.OrderNos
}
func (this *OapiDingpayOrderSyncstatusRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.order.syncstatus"
}
func (this *OapiDingpayOrderSyncstatusRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayOrderSyncstatusRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayOrderSyncstatusRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayOrderSyncstatusRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayOrderSyncstatusRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["order_nos"] = this.OrderNos
	return txtParams
}
func (this *OapiDingpayOrderSyncstatusRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OrderNos, "orderNos"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingpayOrderSyncstatusRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayOrderSyncstatusRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayOrderSyncstatusResponse struct {
	taobao.TaobaoResponse
	Errcode int64                       `json:"errcode,omitempty"`
	Errmsg  string                      `json:"errmsg,omitempty"`
	Result  OrderSyncStatusOpenResponse `json:"result,omitempty"`
	Success bool                        `json:"success,omitempty"`
}
type Orders struct {
	Amount                 int64     `json:"amount,omitempty"`
	ApplyPayOperatorUserid string    `json:"apply_pay_operator_userid,omitempty"`
	BizCode                string    `json:"biz_code,omitempty"`
	CreateOperatorUserid   string    `json:"create_operator_userid,omitempty"`
	Extension              string    `json:"extension,omitempty"`
	GmtApplyPay            time.Time `json:"gmt_apply_pay,omitempty"`
	GmtCreate              time.Time `json:"gmt_create,omitempty"`
	GmtModified            time.Time `json:"gmt_modified,omitempty"`
	GmtPay                 time.Time `json:"gmt_pay,omitempty"`
	OrderNo                string    `json:"order_no,omitempty"`
	OutBizNo               string    `json:"out_biz_no,omitempty"`
	PayChannel             string    `json:"pay_channel,omitempty"`
	PayChannelBizNo        string    `json:"pay_channel_biz_no,omitempty"`
	PayChannelPayerRealUid string    `json:"pay_channel_payer_real_uid,omitempty"`
	PayeeId                string    `json:"payee_id,omitempty"`
	PayeeUserType          string    `json:"payee_user_type,omitempty"`
	PayerId                string    `json:"payer_id,omitempty"`
	PayerUserType          string    `json:"payer_user_type,omitempty"`
	SourceAppId            string    `json:"source_app_id,omitempty"`
	Status                 string    `json:"status,omitempty"`
	Title                  string    `json:"title,omitempty"`
}
type OrderSyncStatusOpenResponse struct {
	Orders []Orders `json:"orders,omitempty"`
}
