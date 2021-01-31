package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanBankcardDeleteRequest() *OapiFinanceLoanBankcardDeleteRequest {
	return &OapiFinanceLoanBankcardDeleteRequest{}
}

type OapiFinanceLoanBankcardDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanBankcardDeleteResponse
	BankCode        string
	BankName        string
	BankcardId      int64
	BankcardMobile  string
	BankcardNo      string
	IdCardNo        string
	TopHttpMethod   string
	TopResponseType string
	UserMobile      string
}

func (this *OapiFinanceLoanBankcardDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetBankCode(bankCode2 string) {
	this.BankCode = bankCode2
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetBankCode() string {
	return this.BankCode
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetBankName(bankName2 string) {
	this.BankName = bankName2
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetBankName() string {
	return this.BankName
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetBankcardId(bankcardId2 int64) {
	this.BankcardId = bankcardId2
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetBankcardId() int64 {
	return this.BankcardId
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetBankcardMobile(bankcardMobile2 string) {
	this.BankcardMobile = bankcardMobile2
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetBankcardMobile() string {
	return this.BankcardMobile
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetBankcardNo(bankcardNo2 string) {
	this.BankcardNo = bankcardNo2
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetBankcardNo() string {
	return this.BankcardNo
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.bankcard.delete"
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanBankcardDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["bank_code"] = this.BankCode
	txtParams["bank_name"] = this.BankName
	txtParams["bankcard_id"] = this.BankcardId
	txtParams["bankcard_mobile"] = this.BankcardMobile
	txtParams["bankcard_no"] = this.BankcardNo
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["user_mobile"] = this.UserMobile
	return txtParams
}
func (this *OapiFinanceLoanBankcardDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BankCode, "bankCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanBankcardDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanBankcardDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
