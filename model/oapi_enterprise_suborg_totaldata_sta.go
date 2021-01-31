package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEnterpriseSuborgTotaldataStatRequest() *OapiEnterpriseSuborgTotaldataStatRequest {
	return &OapiEnterpriseSuborgTotaldataStatRequest{}
}

type OapiEnterpriseSuborgTotaldataStatRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEnterpriseSuborgTotaldataStatResponse
	CorpId          string
	OrderBy         string
	PageSize        int64
	PageStart       int64
	ReturnFields    string
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) SetCorpId(corpId2 string) {
	this.CorpId = corpId2
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetCorpId() string {
	return this.CorpId
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) SetOrderBy(orderBy2 string) {
	this.OrderBy = orderBy2
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetOrderBy() string {
	return this.OrderBy
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) SetPageStart(pageStart2 int64) {
	this.PageStart = pageStart2
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetPageStart() int64 {
	return this.PageStart
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) SetReturnFields(returnFields2 string) {
	this.ReturnFields = returnFields2
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetReturnFields() string {
	return this.ReturnFields
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetApiMethodName() string {
	return "dingtalk.oapi.enterprise.suborg.totaldata.stat"
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corp_id"] = this.CorpId
	txtParams["order_by"] = this.OrderBy
	txtParams["page_size"] = this.PageSize
	txtParams["page_start"] = this.PageStart
	txtParams["return_fields"] = this.ReturnFields
	txtParams["stat_date"] = this.StatDate
	return txtParams
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpId, "corpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEnterpriseSuborgTotaldataStatRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEnterpriseSuborgTotaldataStatResponse struct {
	taobao.TaobaoResponse
	Result  []FamilyDoctorDataVo `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
