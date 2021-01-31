package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanContactsListRequest() *OapiFinanceLoanContactsListRequest {
	return &OapiFinanceLoanContactsListRequest{}
}

type OapiFinanceLoanContactsListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanContactsListResponse
	IdCardNo        string
	TopHttpMethod   string
	TopResponseType string
	UserMobile      string
}

func (this *OapiFinanceLoanContactsListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanContactsListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanContactsListRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanContactsListRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanContactsListRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanContactsListRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanContactsListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.contacts.list"
}
func (this *OapiFinanceLoanContactsListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanContactsListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanContactsListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanContactsListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanContactsListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["user_mobile"] = this.UserMobile
	return txtParams
}
func (this *OapiFinanceLoanContactsListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.IdCardNo, "idCardNo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanContactsListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanContactsListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanContactsListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                      `json:"errcode,omitempty"`
	Errmsg  string                     `json:"errmsg,omitempty"`
	Result  OpenContactInfoQueryResult `json:"result,omitempty"`
	Success bool                       `json:"success,omitempty"`
}
type OpenContactInfo struct {
	ContactId    int64  `json:"contact_id,omitempty"`
	UserCategory string `json:"user_category,omitempty"`
	UserMobile   string `json:"user_mobile,omitempty"`
	UserName     string `json:"user_name,omitempty"`
	UserRelation string `json:"user_relation,omitempty"`
}
type OpenContactInfoQueryResult struct {
	ContactInfoList []OpenContactInfo `json:"contact_info_list,omitempty"`
}
