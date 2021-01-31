package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingpayBillBatchqueryRequest() *OapiDingpayBillBatchqueryRequest {
	return &OapiDingpayBillBatchqueryRequest{}
}

type OapiDingpayBillBatchqueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiDingpayBillBatchqueryResponse
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
	NextKey                string
	PageNum                int64
	PageSize               int64
	PayChannelList         string
	PayChannelPayerRealUid string
	PayeeId                string
	PayeeUserType          string
	PayerId                string
	PayerUserType          string
	ReceiptorTypeList      string
	Size                   int64
	StatusList             string
	TerminationReason      string
	Title                  string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiDingpayBillBatchqueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayBillBatchqueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayBillBatchqueryRequest) SetApplyPayOperatorUserid(applyPayOperatorUserid2 string) {
	this.ApplyPayOperatorUserid = applyPayOperatorUserid2
}
func (this *OapiDingpayBillBatchqueryRequest) GetApplyPayOperatorUserid() string {
	return this.ApplyPayOperatorUserid
}
func (this *OapiDingpayBillBatchqueryRequest) SetBillCategory(billCategory2 string) {
	this.BillCategory = billCategory2
}
func (this *OapiDingpayBillBatchqueryRequest) GetBillCategory() string {
	return this.BillCategory
}
func (this *OapiDingpayBillBatchqueryRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiDingpayBillBatchqueryRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiDingpayBillBatchqueryRequest) SetCreateOperatorUserid(createOperatorUserid2 string) {
	this.CreateOperatorUserid = createOperatorUserid2
}
func (this *OapiDingpayBillBatchqueryRequest) GetCreateOperatorUserid() string {
	return this.CreateOperatorUserid
}
func (this *OapiDingpayBillBatchqueryRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayBillBatchqueryRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayBillBatchqueryRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiDingpayBillBatchqueryRequest) SetGmtApplyPayBeginTime(gmtApplyPayBeginTime2 time.Time) {
	this.GmtApplyPayBeginTime = gmtApplyPayBeginTime2
}
func (this *OapiDingpayBillBatchqueryRequest) GetGmtApplyPayBeginTime() time.Time {
	return this.GmtApplyPayBeginTime
}
func (this *OapiDingpayBillBatchqueryRequest) SetGmtApplyPayEndTime(gmtApplyPayEndTime2 time.Time) {
	this.GmtApplyPayEndTime = gmtApplyPayEndTime2
}
func (this *OapiDingpayBillBatchqueryRequest) GetGmtApplyPayEndTime() time.Time {
	return this.GmtApplyPayEndTime
}
func (this *OapiDingpayBillBatchqueryRequest) SetGmtCreateBeginTime(gmtCreateBeginTime2 time.Time) {
	this.GmtCreateBeginTime = gmtCreateBeginTime2
}
func (this *OapiDingpayBillBatchqueryRequest) GetGmtCreateBeginTime() time.Time {
	return this.GmtCreateBeginTime
}
func (this *OapiDingpayBillBatchqueryRequest) SetGmtCreateEndTime(gmtCreateEndTime2 time.Time) {
	this.GmtCreateEndTime = gmtCreateEndTime2
}
func (this *OapiDingpayBillBatchqueryRequest) GetGmtCreateEndTime() time.Time {
	return this.GmtCreateEndTime
}
func (this *OapiDingpayBillBatchqueryRequest) SetGmtPayBeginTime(gmtPayBeginTime2 time.Time) {
	this.GmtPayBeginTime = gmtPayBeginTime2
}
func (this *OapiDingpayBillBatchqueryRequest) GetGmtPayBeginTime() time.Time {
	return this.GmtPayBeginTime
}
func (this *OapiDingpayBillBatchqueryRequest) SetGmtPayEndTime(gmtPayEndTime2 time.Time) {
	this.GmtPayEndTime = gmtPayEndTime2
}
func (this *OapiDingpayBillBatchqueryRequest) GetGmtPayEndTime() time.Time {
	return this.GmtPayEndTime
}
func (this *OapiDingpayBillBatchqueryRequest) SetMaxAmount(maxAmount2 int64) {
	this.MaxAmount = maxAmount2
}
func (this *OapiDingpayBillBatchqueryRequest) GetMaxAmount() int64 {
	return this.MaxAmount
}
func (this *OapiDingpayBillBatchqueryRequest) SetMinAmount(minAmount2 int64) {
	this.MinAmount = minAmount2
}
func (this *OapiDingpayBillBatchqueryRequest) GetMinAmount() int64 {
	return this.MinAmount
}
func (this *OapiDingpayBillBatchqueryRequest) SetNextKey(nextKey2 string) {
	this.NextKey = nextKey2
}
func (this *OapiDingpayBillBatchqueryRequest) GetNextKey() string {
	return this.NextKey
}
func (this *OapiDingpayBillBatchqueryRequest) SetPageNum(pageNum2 int64) {
	this.PageNum = pageNum2
}
func (this *OapiDingpayBillBatchqueryRequest) GetPageNum() int64 {
	return this.PageNum
}
func (this *OapiDingpayBillBatchqueryRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiDingpayBillBatchqueryRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiDingpayBillBatchqueryRequest) SetPayChannelList(payChannelList2 string) {
	this.PayChannelList = payChannelList2
}
func (this *OapiDingpayBillBatchqueryRequest) GetPayChannelList() string {
	return this.PayChannelList
}
func (this *OapiDingpayBillBatchqueryRequest) SetPayChannelPayerRealUid(payChannelPayerRealUid2 string) {
	this.PayChannelPayerRealUid = payChannelPayerRealUid2
}
func (this *OapiDingpayBillBatchqueryRequest) GetPayChannelPayerRealUid() string {
	return this.PayChannelPayerRealUid
}
func (this *OapiDingpayBillBatchqueryRequest) SetPayeeId(payeeId2 string) {
	this.PayeeId = payeeId2
}
func (this *OapiDingpayBillBatchqueryRequest) GetPayeeId() string {
	return this.PayeeId
}
func (this *OapiDingpayBillBatchqueryRequest) SetPayeeUserType(payeeUserType2 string) {
	this.PayeeUserType = payeeUserType2
}
func (this *OapiDingpayBillBatchqueryRequest) GetPayeeUserType() string {
	return this.PayeeUserType
}
func (this *OapiDingpayBillBatchqueryRequest) SetPayerId(payerId2 string) {
	this.PayerId = payerId2
}
func (this *OapiDingpayBillBatchqueryRequest) GetPayerId() string {
	return this.PayerId
}
func (this *OapiDingpayBillBatchqueryRequest) SetPayerUserType(payerUserType2 string) {
	this.PayerUserType = payerUserType2
}
func (this *OapiDingpayBillBatchqueryRequest) GetPayerUserType() string {
	return this.PayerUserType
}
func (this *OapiDingpayBillBatchqueryRequest) SetReceiptorTypeList(receiptorTypeList2 string) {
	this.ReceiptorTypeList = receiptorTypeList2
}
func (this *OapiDingpayBillBatchqueryRequest) GetReceiptorTypeList() string {
	return this.ReceiptorTypeList
}
func (this *OapiDingpayBillBatchqueryRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiDingpayBillBatchqueryRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiDingpayBillBatchqueryRequest) SetStatusList(statusList2 string) {
	this.StatusList = statusList2
}
func (this *OapiDingpayBillBatchqueryRequest) GetStatusList() string {
	return this.StatusList
}
func (this *OapiDingpayBillBatchqueryRequest) SetTerminationReason(terminationReason2 string) {
	this.TerminationReason = terminationReason2
}
func (this *OapiDingpayBillBatchqueryRequest) GetTerminationReason() string {
	return this.TerminationReason
}
func (this *OapiDingpayBillBatchqueryRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiDingpayBillBatchqueryRequest) GetTitle() string {
	return this.Title
}
func (this *OapiDingpayBillBatchqueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.bill.batchquery"
}
func (this *OapiDingpayBillBatchqueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayBillBatchqueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayBillBatchqueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayBillBatchqueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayBillBatchqueryRequest) GetTextParams() map[string]interface{} {
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
	txtParams["next_key"] = this.NextKey
	txtParams["page_num"] = this.PageNum
	txtParams["page_size"] = this.PageSize
	txtParams["pay_channel_list"] = this.PayChannelList
	txtParams["pay_channel_payer_real_uid"] = this.PayChannelPayerRealUid
	txtParams["payee_id"] = this.PayeeId
	txtParams["payee_user_type"] = this.PayeeUserType
	txtParams["payer_id"] = this.PayerId
	txtParams["payer_user_type"] = this.PayerUserType
	txtParams["receiptor_type_list"] = this.ReceiptorTypeList
	txtParams["size"] = this.Size
	txtParams["status_list"] = this.StatusList
	txtParams["termination_reason"] = this.TerminationReason
	txtParams["title"] = this.Title
	return txtParams
}
func (this *OapiDingpayBillBatchqueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.PayChannelList, 20, "payChannelList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingpayBillBatchqueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayBillBatchqueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayBillBatchqueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64                      `json:"errcode,omitempty"`
	Errmsg  string                     `json:"errmsg,omitempty"`
	Result  BillBatchQueryOpenResponse `json:"result,omitempty"`
	Success bool                       `json:"success,omitempty"`
}
type DingPayBillOpenBo struct {
	Amount                 int64     `json:"amount,omitempty"`
	ApplyPayOperatorUserid string    `json:"apply_pay_operator_userid,omitempty"`
	BillCategory           string    `json:"bill_category,omitempty"`
	BillNo                 string    `json:"bill_no,omitempty"`
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
	PayChannelPayeeRealUid string    `json:"pay_channel_payee_real_uid,omitempty"`
	PayChannelPayerRealUid string    `json:"pay_channel_payer_real_uid,omitempty"`
	PayeeId                string    `json:"payee_id,omitempty"`
	PayeeUserType          string    `json:"payee_user_type,omitempty"`
	PayerId                string    `json:"payer_id,omitempty"`
	PayerUserType          string    `json:"payer_user_type,omitempty"`
	PrincipalId            string    `json:"principal_id,omitempty"`
	ReceiptorType          string    `json:"receiptor_type,omitempty"`
	SourceAppId            string    `json:"source_app_id,omitempty"`
	Status                 string    `json:"status,omitempty"`
	TerminationOperatorId  string    `json:"termination_operator_id,omitempty"`
	TerminationReason      string    `json:"termination_reason,omitempty"`
	Title                  string    `json:"title,omitempty"`
}
type BillBatchQueryOpenResponse struct {
	BillList       []DingPayBillOpenBo `json:"bill_list,omitempty"`
	CurrentPageNum int64               `json:"current_page_num,omitempty"`
	NextKey        string              `json:"next_key,omitempty"`
	PageSize       int64               `json:"page_size,omitempty"`
	TotalCount     int64               `json:"total_count,omitempty"`
	TotalPage      int64               `json:"total_page,omitempty"`
}
