package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRetailSellerOrgdetailQueryRequest() *OapiRetailSellerOrgdetailQueryRequest {
	return &OapiRetailSellerOrgdetailQueryRequest{}
}

type OapiRetailSellerOrgdetailQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRetailSellerOrgdetailQueryResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRetailSellerOrgdetailQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRetailSellerOrgdetailQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRetailSellerOrgdetailQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRetailSellerOrgdetailQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRetailSellerOrgdetailQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.retail.seller.orgdetail.query"
}
func (this *OapiRetailSellerOrgdetailQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRetailSellerOrgdetailQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRetailSellerOrgdetailQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRetailSellerOrgdetailQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRetailSellerOrgdetailQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRetailSellerOrgdetailQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiRetailSellerOrgdetailQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRetailSellerOrgdetailQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRetailSellerOrgdetailQueryResponse struct {
	taobao.TaobaoResponse
	Result  []OrgDto `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
type ShopEmpDto struct {
	EmpType    string `json:"emp_type,omitempty"`
	Name       string `json:"name,omitempty"`
	OuterId    string `json:"outer_id,omitempty"`
	OuterSubId string `json:"outer_sub_id,omitempty"`
	SellerNick string `json:"seller_nick,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
type SellerDto struct {
	SellerId    int64        `json:"seller_id,omitempty"`
	SellerNick  string       `json:"seller_nick,omitempty"`
	ShopEmpList []ShopEmpDto `json:"shop_emp_list,omitempty"`
	Type        string       `json:"type,omitempty"`
	Userid      string       `json:"userid,omitempty"`
}
type OrgDto struct {
	BindSellers []SellerDto `json:"bind_sellers,omitempty"`
	CorpId      string      `json:"corp_id,omitempty"`
	GmtCreate   time.Time   `json:"gmt_create,omitempty"`
	OrgName     string      `json:"org_name,omitempty"`
}
