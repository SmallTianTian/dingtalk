package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanNotifyRepaymentNoticeRequest() *OapiFinanceLoanNotifyRepaymentNoticeRequest {
	return &OapiFinanceLoanNotifyRepaymentNoticeRequest{}
}

type OapiFinanceLoanNotifyRepaymentNoticeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanNotifyRepaymentNoticeResponse
	IdCardNo        string
	LoanOrderNo     string
	OpenChannelName string
	OpenProductName string
	OvdTerms        string
	PaidInterest    int64
	PaidPenalty     int64
	PaidPrincipal   int64
	PaidTotalAmount int64
	RepayRealDate   string
	RepaymentTerms  string
	TopHttpMethod   string
	TopResponseType string
	UserMobile      string
}

func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetLoanOrderNo(loanOrderNo2 string) {
	this.LoanOrderNo = loanOrderNo2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetLoanOrderNo() string {
	return this.LoanOrderNo
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetOpenChannelName(openChannelName2 string) {
	this.OpenChannelName = openChannelName2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetOpenChannelName() string {
	return this.OpenChannelName
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetOpenProductName(openProductName2 string) {
	this.OpenProductName = openProductName2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetOpenProductName() string {
	return this.OpenProductName
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetOvdTerms(ovdTerms2 string) {
	this.OvdTerms = ovdTerms2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetOvdTerms() string {
	return this.OvdTerms
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetPaidInterest(paidInterest2 int64) {
	this.PaidInterest = paidInterest2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetPaidInterest() int64 {
	return this.PaidInterest
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetPaidPenalty(paidPenalty2 int64) {
	this.PaidPenalty = paidPenalty2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetPaidPenalty() int64 {
	return this.PaidPenalty
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetPaidPrincipal(paidPrincipal2 int64) {
	this.PaidPrincipal = paidPrincipal2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetPaidPrincipal() int64 {
	return this.PaidPrincipal
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetPaidTotalAmount(paidTotalAmount2 int64) {
	this.PaidTotalAmount = paidTotalAmount2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetPaidTotalAmount() int64 {
	return this.PaidTotalAmount
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetRepayRealDate(repayRealDate2 string) {
	this.RepayRealDate = repayRealDate2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetRepayRealDate() string {
	return this.RepayRealDate
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetRepaymentTerms(repaymentTerms2 string) {
	this.RepaymentTerms = repaymentTerms2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetRepaymentTerms() string {
	return this.RepaymentTerms
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.notify.repayment.notice"
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["loan_order_no"] = this.LoanOrderNo
	txtParams["open_channel_name"] = this.OpenChannelName
	txtParams["open_product_name"] = this.OpenProductName
	txtParams["ovd_terms"] = this.OvdTerms
	txtParams["paid_interest"] = this.PaidInterest
	txtParams["paid_penalty"] = this.PaidPenalty
	txtParams["paid_principal"] = this.PaidPrincipal
	txtParams["paid_total_amount"] = this.PaidTotalAmount
	txtParams["repay_real_date"] = this.RepayRealDate
	txtParams["repayment_terms"] = this.RepaymentTerms
	txtParams["user_mobile"] = this.UserMobile
	return txtParams
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.IdCardNo, "idCardNo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanNotifyRepaymentNoticeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanNotifyRepaymentNoticeResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
