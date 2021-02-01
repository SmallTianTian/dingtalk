package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingpayBillBatchquerycountRequest() *OapiDingpayBillBatchquerycountRequest {
	return &OapiDingpayBillBatchquerycountRequest{}
}

type OapiDingpayBillBatchquerycountRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiDingpayBillBatchquerycountResponse
	ApplyPayOperatorUserid string
	BillCategory           string
	BizCode                string
	CreateOperatorUserid   string
	Extension              string
	GmtApplyPayBeginTime   time.Time
	GmtApplyPayEndTime     time.Time
	GmtCreateBeginTime     time.Time
	GmtCreateEndTime       time.Time
	GmtPayBeginTime        time.Time
	GmtPayEndTime          time.Time
	MaxAmount              int64
	MinAmount              int64
	PayChannelList         string
	PayChannelPayerRealUid string
	PayeeId                string
	PayeeUserType          string
	PayerId                string
	PayerUserType          string
	ReceiptorTypeList      string
	StatusList             string
	TerminationReason      string
	Title                  string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiDingpayBillBatchquerycountRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayBillBatchquerycountRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayBillBatchquerycountRequest) SetApplyPayOperatorUserid(applyPayOperatorUserid2 string) {
	this.ApplyPayOperatorUserid = applyPayOperatorUserid2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetApplyPayOperatorUserid() string {
	return this.ApplyPayOperatorUserid
}
func (this *OapiDingpayBillBatchquerycountRequest) SetBillCategory(billCategory2 string) {
	this.BillCategory = billCategory2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetBillCategory() string {
	return this.BillCategory
}
func (this *OapiDingpayBillBatchquerycountRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiDingpayBillBatchquerycountRequest) SetCreateOperatorUserid(createOperatorUserid2 string) {
	this.CreateOperatorUserid = createOperatorUserid2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetCreateOperatorUserid() string {
	return this.CreateOperatorUserid
}
func (this *OapiDingpayBillBatchquerycountRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayBillBatchquerycountRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiDingpayBillBatchquerycountRequest) SetGmtApplyPayBeginTime(gmtApplyPayBeginTime2 time.Time) {
	this.GmtApplyPayBeginTime = gmtApplyPayBeginTime2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetGmtApplyPayBeginTime() time.Time {
	return this.GmtApplyPayBeginTime
}
func (this *OapiDingpayBillBatchquerycountRequest) SetGmtApplyPayEndTime(gmtApplyPayEndTime2 time.Time) {
	this.GmtApplyPayEndTime = gmtApplyPayEndTime2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetGmtApplyPayEndTime() time.Time {
	return this.GmtApplyPayEndTime
}
func (this *OapiDingpayBillBatchquerycountRequest) SetGmtCreateBeginTime(gmtCreateBeginTime2 time.Time) {
	this.GmtCreateBeginTime = gmtCreateBeginTime2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetGmtCreateBeginTime() time.Time {
	return this.GmtCreateBeginTime
}
func (this *OapiDingpayBillBatchquerycountRequest) SetGmtCreateEndTime(gmtCreateEndTime2 time.Time) {
	this.GmtCreateEndTime = gmtCreateEndTime2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetGmtCreateEndTime() time.Time {
	return this.GmtCreateEndTime
}
func (this *OapiDingpayBillBatchquerycountRequest) SetGmtPayBeginTime(gmtPayBeginTime2 time.Time) {
	this.GmtPayBeginTime = gmtPayBeginTime2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetGmtPayBeginTime() time.Time {
	return this.GmtPayBeginTime
}
func (this *OapiDingpayBillBatchquerycountRequest) SetGmtPayEndTime(gmtPayEndTime2 time.Time) {
	this.GmtPayEndTime = gmtPayEndTime2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetGmtPayEndTime() time.Time {
	return this.GmtPayEndTime
}
func (this *OapiDingpayBillBatchquerycountRequest) SetMaxAmount(maxAmount2 int64) {
	this.MaxAmount = maxAmount2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetMaxAmount() int64 {
	return this.MaxAmount
}
func (this *OapiDingpayBillBatchquerycountRequest) SetMinAmount(minAmount2 int64) {
	this.MinAmount = minAmount2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetMinAmount() int64 {
	return this.MinAmount
}
func (this *OapiDingpayBillBatchquerycountRequest) SetPayChannelList(payChannelList2 string) {
	this.PayChannelList = payChannelList2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetPayChannelList() string {
	return this.PayChannelList
}
func (this *OapiDingpayBillBatchquerycountRequest) SetPayChannelPayerRealUid(payChannelPayerRealUid2 string) {
	this.PayChannelPayerRealUid = payChannelPayerRealUid2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetPayChannelPayerRealUid() string {
	return this.PayChannelPayerRealUid
}
func (this *OapiDingpayBillBatchquerycountRequest) SetPayeeId(payeeId2 string) {
	this.PayeeId = payeeId2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetPayeeId() string {
	return this.PayeeId
}
func (this *OapiDingpayBillBatchquerycountRequest) SetPayeeUserType(payeeUserType2 string) {
	this.PayeeUserType = payeeUserType2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetPayeeUserType() string {
	return this.PayeeUserType
}
func (this *OapiDingpayBillBatchquerycountRequest) SetPayerId(payerId2 string) {
	this.PayerId = payerId2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetPayerId() string {
	return this.PayerId
}
func (this *OapiDingpayBillBatchquerycountRequest) SetPayerUserType(payerUserType2 string) {
	this.PayerUserType = payerUserType2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetPayerUserType() string {
	return this.PayerUserType
}
func (this *OapiDingpayBillBatchquerycountRequest) SetReceiptorTypeList(receiptorTypeList2 string) {
	this.ReceiptorTypeList = receiptorTypeList2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetReceiptorTypeList() string {
	return this.ReceiptorTypeList
}
func (this *OapiDingpayBillBatchquerycountRequest) SetStatusList(statusList2 string) {
	this.StatusList = statusList2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetStatusList() string {
	return this.StatusList
}
func (this *OapiDingpayBillBatchquerycountRequest) SetTerminationReason(terminationReason2 string) {
	this.TerminationReason = terminationReason2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetTerminationReason() string {
	return this.TerminationReason
}
func (this *OapiDingpayBillBatchquerycountRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetTitle() string {
	return this.Title
}
func (this *OapiDingpayBillBatchquerycountRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.bill.batchquerycount"
}
func (this *OapiDingpayBillBatchquerycountRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayBillBatchquerycountRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayBillBatchquerycountRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayBillBatchquerycountRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayBillBatchquerycountRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["apply_pay_operator_userid"] = this.ApplyPayOperatorUserid
	txtParams["bill_category"] = this.BillCategory
	txtParams["biz_code"] = this.BizCode
	txtParams["create_operator_userid"] = this.CreateOperatorUserid
	txtParams["extension"] = this.Extension
	txtParams["gmt_apply_pay_begin_time"] = this.GmtApplyPayBeginTime
	txtParams["gmt_apply_pay_end_time"] = this.GmtApplyPayEndTime
	txtParams["gmt_create_begin_time"] = this.GmtCreateBeginTime
	txtParams["gmt_create_end_time"] = this.GmtCreateEndTime
	txtParams["gmt_pay_begin_time"] = this.GmtPayBeginTime
	txtParams["gmt_pay_end_time"] = this.GmtPayEndTime
	txtParams["max_amount"] = this.MaxAmount
	txtParams["min_amount"] = this.MinAmount
	txtParams["pay_channel_list"] = this.PayChannelList
	txtParams["pay_channel_payer_real_uid"] = this.PayChannelPayerRealUid
	txtParams["payee_id"] = this.PayeeId
	txtParams["payee_user_type"] = this.PayeeUserType
	txtParams["payer_id"] = this.PayerId
	txtParams["payer_user_type"] = this.PayerUserType
	txtParams["receiptor_type_list"] = this.ReceiptorTypeList
	txtParams["status_list"] = this.StatusList
	txtParams["termination_reason"] = this.TerminationReason
	txtParams["title"] = this.Title
	return txtParams
}
func (this *OapiDingpayBillBatchquerycountRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.PayChannelList, 20, "payChannelList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingpayBillBatchquerycountRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayBillBatchquerycountRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayBillBatchquerycountResponse struct {
	taobao.TaobaoResponse
	Result  int64 `json:"result,omitempty"`
	Success bool  `json:"success,omitempty"`
}
