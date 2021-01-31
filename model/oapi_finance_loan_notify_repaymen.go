package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanNotifyRepaymentRequest() *OapiFinanceLoanNotifyRepaymentRequest {
	return &OapiFinanceLoanNotifyRepaymentRequest{}
}

type OapiFinanceLoanNotifyRepaymentRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiFinanceLoanNotifyRepaymentResponse
	Amount                 int64
	AvailableAmount        int64
	BankName               string
	BankcardNo             string
	CurrentIntOvdDays      int64
	CurrentOvdTerms        string
	CurrentPaidInterest    int64
	CurrentPaidPenalty     int64
	CurrentPaidPrincipal   int64
	CurrentPaidTotalAmount int64
	CurrentPrinOvdDays     int64
	FailReason             string
	IdCardNo               string
	LoanOrderNo            string
	OpenChannelName        string
	OpenProductCode        string
	OpenProductName        string
	OpenProductType        string
	PaidInterest           int64
	PaidPenalty            int64
	PaidPrincipal          int64
	PaidTotalAmount        int64
	PayableInterest        int64
	PayablePenalty         int64
	PayablePrincipal       int64
	PayableTotalAmount     int64
	RepayRealDate          string
	RepayType              string
	RepaymentNo            string
	RepaymentTerms         string
	Status                 string
	TopHttpMethod          string
	TopResponseType        string
	Type                   string
	UnpaidInterest         int64
	UnpaidPenalty          int64
	UnpaidPrincipal        int64
	UnpaidTotalAmount      int64
	UserMobile             string
}

func (this *OapiFinanceLoanNotifyRepaymentRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetAmount(amount2 int64) {
	this.Amount = amount2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetAmount() int64 {
	return this.Amount
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetAvailableAmount(availableAmount2 int64) {
	this.AvailableAmount = availableAmount2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetAvailableAmount() int64 {
	return this.AvailableAmount
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetBankName(bankName2 string) {
	this.BankName = bankName2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetBankName() string {
	return this.BankName
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetBankcardNo(bankcardNo2 string) {
	this.BankcardNo = bankcardNo2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetBankcardNo() string {
	return this.BankcardNo
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetCurrentIntOvdDays(currentIntOvdDays2 int64) {
	this.CurrentIntOvdDays = currentIntOvdDays2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetCurrentIntOvdDays() int64 {
	return this.CurrentIntOvdDays
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetCurrentOvdTerms(currentOvdTerms2 string) {
	this.CurrentOvdTerms = currentOvdTerms2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetCurrentOvdTerms() string {
	return this.CurrentOvdTerms
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetCurrentPaidInterest(currentPaidInterest2 int64) {
	this.CurrentPaidInterest = currentPaidInterest2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetCurrentPaidInterest() int64 {
	return this.CurrentPaidInterest
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetCurrentPaidPenalty(currentPaidPenalty2 int64) {
	this.CurrentPaidPenalty = currentPaidPenalty2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetCurrentPaidPenalty() int64 {
	return this.CurrentPaidPenalty
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetCurrentPaidPrincipal(currentPaidPrincipal2 int64) {
	this.CurrentPaidPrincipal = currentPaidPrincipal2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetCurrentPaidPrincipal() int64 {
	return this.CurrentPaidPrincipal
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetCurrentPaidTotalAmount(currentPaidTotalAmount2 int64) {
	this.CurrentPaidTotalAmount = currentPaidTotalAmount2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetCurrentPaidTotalAmount() int64 {
	return this.CurrentPaidTotalAmount
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetCurrentPrinOvdDays(currentPrinOvdDays2 int64) {
	this.CurrentPrinOvdDays = currentPrinOvdDays2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetCurrentPrinOvdDays() int64 {
	return this.CurrentPrinOvdDays
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetFailReason(failReason2 string) {
	this.FailReason = failReason2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetFailReason() string {
	return this.FailReason
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetLoanOrderNo(loanOrderNo2 string) {
	this.LoanOrderNo = loanOrderNo2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetLoanOrderNo() string {
	return this.LoanOrderNo
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetOpenChannelName(openChannelName2 string) {
	this.OpenChannelName = openChannelName2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetOpenChannelName() string {
	return this.OpenChannelName
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetOpenProductCode(openProductCode2 string) {
	this.OpenProductCode = openProductCode2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetOpenProductCode() string {
	return this.OpenProductCode
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetOpenProductName(openProductName2 string) {
	this.OpenProductName = openProductName2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetOpenProductName() string {
	return this.OpenProductName
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetOpenProductType(openProductType2 string) {
	this.OpenProductType = openProductType2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetOpenProductType() string {
	return this.OpenProductType
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetPaidInterest(paidInterest2 int64) {
	this.PaidInterest = paidInterest2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetPaidInterest() int64 {
	return this.PaidInterest
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetPaidPenalty(paidPenalty2 int64) {
	this.PaidPenalty = paidPenalty2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetPaidPenalty() int64 {
	return this.PaidPenalty
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetPaidPrincipal(paidPrincipal2 int64) {
	this.PaidPrincipal = paidPrincipal2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetPaidPrincipal() int64 {
	return this.PaidPrincipal
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetPaidTotalAmount(paidTotalAmount2 int64) {
	this.PaidTotalAmount = paidTotalAmount2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetPaidTotalAmount() int64 {
	return this.PaidTotalAmount
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetPayableInterest(payableInterest2 int64) {
	this.PayableInterest = payableInterest2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetPayableInterest() int64 {
	return this.PayableInterest
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetPayablePenalty(payablePenalty2 int64) {
	this.PayablePenalty = payablePenalty2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetPayablePenalty() int64 {
	return this.PayablePenalty
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetPayablePrincipal(payablePrincipal2 int64) {
	this.PayablePrincipal = payablePrincipal2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetPayablePrincipal() int64 {
	return this.PayablePrincipal
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetPayableTotalAmount(payableTotalAmount2 int64) {
	this.PayableTotalAmount = payableTotalAmount2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetPayableTotalAmount() int64 {
	return this.PayableTotalAmount
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetRepayRealDate(repayRealDate2 string) {
	this.RepayRealDate = repayRealDate2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetRepayRealDate() string {
	return this.RepayRealDate
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetRepayType(repayType2 string) {
	this.RepayType = repayType2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetRepayType() string {
	return this.RepayType
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetRepaymentNo(repaymentNo2 string) {
	this.RepaymentNo = repaymentNo2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetRepaymentNo() string {
	return this.RepaymentNo
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetRepaymentTerms(repaymentTerms2 string) {
	this.RepaymentTerms = repaymentTerms2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetRepaymentTerms() string {
	return this.RepaymentTerms
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetStatus(status2 string) {
	this.Status = status2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetStatus() string {
	return this.Status
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetType() string {
	return this.Type
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetUnpaidInterest(unpaidInterest2 int64) {
	this.UnpaidInterest = unpaidInterest2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetUnpaidInterest() int64 {
	return this.UnpaidInterest
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetUnpaidPenalty(unpaidPenalty2 int64) {
	this.UnpaidPenalty = unpaidPenalty2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetUnpaidPenalty() int64 {
	return this.UnpaidPenalty
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetUnpaidPrincipal(unpaidPrincipal2 int64) {
	this.UnpaidPrincipal = unpaidPrincipal2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetUnpaidPrincipal() int64 {
	return this.UnpaidPrincipal
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetUnpaidTotalAmount(unpaidTotalAmount2 int64) {
	this.UnpaidTotalAmount = unpaidTotalAmount2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetUnpaidTotalAmount() int64 {
	return this.UnpaidTotalAmount
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.notify.repayment"
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["amount"] = this.Amount
	txtParams["available_amount"] = this.AvailableAmount
	txtParams["bank_name"] = this.BankName
	txtParams["bankcard_no"] = this.BankcardNo
	txtParams["current_int_ovd_days"] = this.CurrentIntOvdDays
	txtParams["current_ovd_terms"] = this.CurrentOvdTerms
	txtParams["current_paid_interest"] = this.CurrentPaidInterest
	txtParams["current_paid_penalty"] = this.CurrentPaidPenalty
	txtParams["current_paid_principal"] = this.CurrentPaidPrincipal
	txtParams["current_paid_total_amount"] = this.CurrentPaidTotalAmount
	txtParams["current_prin_ovd_days"] = this.CurrentPrinOvdDays
	txtParams["fail_reason"] = this.FailReason
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["loan_order_no"] = this.LoanOrderNo
	txtParams["open_channel_name"] = this.OpenChannelName
	txtParams["open_product_code"] = this.OpenProductCode
	txtParams["open_product_name"] = this.OpenProductName
	txtParams["open_product_type"] = this.OpenProductType
	txtParams["paid_interest"] = this.PaidInterest
	txtParams["paid_penalty"] = this.PaidPenalty
	txtParams["paid_principal"] = this.PaidPrincipal
	txtParams["paid_total_amount"] = this.PaidTotalAmount
	txtParams["payable_interest"] = this.PayableInterest
	txtParams["payable_penalty"] = this.PayablePenalty
	txtParams["payable_principal"] = this.PayablePrincipal
	txtParams["payable_total_amount"] = this.PayableTotalAmount
	txtParams["repay_real_date"] = this.RepayRealDate
	txtParams["repay_type"] = this.RepayType
	txtParams["repayment_no"] = this.RepaymentNo
	txtParams["repayment_terms"] = this.RepaymentTerms
	txtParams["status"] = this.Status
	txtParams["type"] = this.Type
	txtParams["unpaid_interest"] = this.UnpaidInterest
	txtParams["unpaid_penalty"] = this.UnpaidPenalty
	txtParams["unpaid_principal"] = this.UnpaidPrincipal
	txtParams["unpaid_total_amount"] = this.UnpaidTotalAmount
	txtParams["user_mobile"] = this.UserMobile
	return txtParams
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Amount, "amount"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanNotifyRepaymentRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanNotifyRepaymentResponse struct {
	taobao.TaobaoResponse
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
