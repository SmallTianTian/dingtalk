package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEnterpriseMicroappUsedataStatRequest() *OapiEnterpriseMicroappUsedataStatRequest {
	return &OapiEnterpriseMicroappUsedataStatRequest{}
}

type OapiEnterpriseMicroappUsedataStatRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEnterpriseMicroappUsedataStatResponse
	CorpId          string
	OrderBy         string
	PageSize        int64
	PageStart       int64
	ReturnFields    string
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEnterpriseMicroappUsedataStatRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) SetCorpId(corpId2 string) {
	this.CorpId = corpId2
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetCorpId() string {
	return this.CorpId
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) SetOrderBy(orderBy2 string) {
	this.OrderBy = orderBy2
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetOrderBy() string {
	return this.OrderBy
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) SetPageStart(pageStart2 int64) {
	this.PageStart = pageStart2
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetPageStart() int64 {
	return this.PageStart
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) SetReturnFields(returnFields2 string) {
	this.ReturnFields = returnFields2
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetReturnFields() string {
	return this.ReturnFields
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetApiMethodName() string {
	return "dingtalk.oapi.enterprise.microapp.usedata.stat"
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corp_id"] = this.CorpId
	txtParams["order_by"] = this.OrderBy
	txtParams["page_size"] = this.PageSize
	txtParams["page_start"] = this.PageStart
	txtParams["return_fields"] = this.ReturnFields
	txtParams["stat_date"] = this.StatDate
	return txtParams
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpId, "corpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEnterpriseMicroappUsedataStatRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEnterpriseMicroappUsedataStatResponse struct {
	taobao.TaobaoResponse
	Result  []MicroAppDataVO `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type MicroAppDataVO struct {
	AgentName          string `json:"agent_name,omitempty"`
	CorpId             string `json:"corp_id,omitempty"`
	OpenMicroUserCnt1d string `json:"open_micro_user_cnt_1d,omitempty"`
	OpenMicroUserCnt1w string `json:"open_micro_user_cnt_1w,omitempty"`
	StatDate           string `json:"stat_date,omitempty"`
}
