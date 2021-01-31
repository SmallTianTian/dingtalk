package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataCustomerQueryRequest() *OapiCrmObjectdataCustomerQueryRequest {
	return &OapiCrmObjectdataCustomerQueryRequest{}
}

type OapiCrmObjectdataCustomerQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiCrmObjectdataCustomerQueryResponse
	CurrentOperatorUserid string
	Cursor                string
	PageSize              int64
	QueryDsl              string
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiCrmObjectdataCustomerQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataCustomerQueryRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiCrmObjectdataCustomerQueryRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiCrmObjectdataCustomerQueryRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiCrmObjectdataCustomerQueryRequest) SetQueryDsl(queryDsl2 string) {
	this.QueryDsl = queryDsl2
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetQueryDsl() string {
	return this.QueryDsl
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.customer.query"
}
func (this *OapiCrmObjectdataCustomerQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataCustomerQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataCustomerQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["cursor"] = this.Cursor
	txtParams["page_size"] = this.PageSize
	txtParams["query_dsl"] = this.QueryDsl
	return txtParams
}
func (this *OapiCrmObjectdataCustomerQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.PageSize, "pageSize"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataCustomerQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataCustomerQueryResponse struct {
	taobao.TaobaoResponse
	Result IterablePage `json:"result,omitempty"`
}
