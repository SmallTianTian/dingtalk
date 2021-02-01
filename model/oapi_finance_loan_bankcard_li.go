package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanBankcardListRequest() *OapiFinanceLoanBankcardListRequest {
	return &OapiFinanceLoanBankcardListRequest{}
}

type OapiFinanceLoanBankcardListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanBankcardListResponse
	IdCardNo        string
	TopHttpMethod   string
	TopResponseType string
	UserMobile      string
}

func (this *OapiFinanceLoanBankcardListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanBankcardListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanBankcardListRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanBankcardListRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanBankcardListRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanBankcardListRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanBankcardListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.bankcard.list"
}
func (this *OapiFinanceLoanBankcardListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanBankcardListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanBankcardListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanBankcardListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanBankcardListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["user_mobile"] = this.UserMobile
	return txtParams
}
func (this *OapiFinanceLoanBankcardListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.IdCardNo, "idCardNo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanBankcardListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanBankcardListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanBankcardListResponse struct {
	taobao.TaobaoResponse
	Result  OpenBankcardQueryResult `json:"result,omitempty"`
	Success bool                    `json:"success,omitempty"`
}
type OpenBankcardInfo struct {
	BankCode       string `json:"bank_code,omitempty"`
	BankName       string `json:"bank_name,omitempty"`
	BankcardId     int64  `json:"bankcard_id,omitempty"`
	BankcardMobile string `json:"bankcard_mobile,omitempty"`
	BankcardNo     string `json:"bankcard_no,omitempty"`
}
type OpenBankcardQueryResult struct {
	BankcardList []OpenBankcardInfo `json:"bankcard_list,omitempty"`
}
