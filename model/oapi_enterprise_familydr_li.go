package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEnterpriseFamilydrListRequest() *OapiEnterpriseFamilydrListRequest {
	return &OapiEnterpriseFamilydrListRequest{}
}

type OapiEnterpriseFamilydrListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEnterpriseFamilydrListResponse
	CorpId          string
	OrderBy         string
	PageSize        int64
	PageStart       int64
	ReturnFields    string
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEnterpriseFamilydrListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEnterpriseFamilydrListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEnterpriseFamilydrListRequest) SetCorpId(corpId2 string) {
	this.CorpId = corpId2
}
func (this *OapiEnterpriseFamilydrListRequest) GetCorpId() string {
	return this.CorpId
}
func (this *OapiEnterpriseFamilydrListRequest) SetOrderBy(orderBy2 string) {
	this.OrderBy = orderBy2
}
func (this *OapiEnterpriseFamilydrListRequest) GetOrderBy() string {
	return this.OrderBy
}
func (this *OapiEnterpriseFamilydrListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEnterpriseFamilydrListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEnterpriseFamilydrListRequest) SetPageStart(pageStart2 int64) {
	this.PageStart = pageStart2
}
func (this *OapiEnterpriseFamilydrListRequest) GetPageStart() int64 {
	return this.PageStart
}
func (this *OapiEnterpriseFamilydrListRequest) SetReturnFields(returnFields2 string) {
	this.ReturnFields = returnFields2
}
func (this *OapiEnterpriseFamilydrListRequest) GetReturnFields() string {
	return this.ReturnFields
}
func (this *OapiEnterpriseFamilydrListRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiEnterpriseFamilydrListRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiEnterpriseFamilydrListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.enterprise.familydr.list"
}
func (this *OapiEnterpriseFamilydrListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEnterpriseFamilydrListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEnterpriseFamilydrListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEnterpriseFamilydrListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEnterpriseFamilydrListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corp_id"] = this.CorpId
	txtParams["order_by"] = this.OrderBy
	txtParams["page_size"] = this.PageSize
	txtParams["page_start"] = this.PageStart
	txtParams["return_fields"] = this.ReturnFields
	txtParams["stat_date"] = this.StatDate
	return txtParams
}
func (this *OapiEnterpriseFamilydrListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpId, "corpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEnterpriseFamilydrListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEnterpriseFamilydrListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OrderBy struct {
	Field string `json:"field,omitempty"`
	Order string `json:"order,omitempty"`
}
type OapiEnterpriseFamilydrListResponse struct {
	taobao.TaobaoResponse
	Result  []FamilyDoctorDataVo `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type FamilyDoctorDataVo struct {
	CorpId                  string `json:"corp_id,omitempty"`
	DeptNameLv2             string `json:"dept_name_lv2,omitempty"`
	DeptNameLv3             string `json:"dept_name_lv3,omitempty"`
	LiveLaunchSuccUserCnt1d string `json:"live_launch_succ_user_cnt_1d,omitempty"`
	LiveLaunchSuccUserCnt1w string `json:"live_launch_succ_user_cnt_1w,omitempty"`
	LivePlayUserCnt1d       string `json:"live_play_user_cnt_1d,omitempty"`
	LivePlayUserCnt1w       string `json:"live_play_user_cnt_1w,omitempty"`
	StatDate                string `json:"stat_date,omitempty"`
}
