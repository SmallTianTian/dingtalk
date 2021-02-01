package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEnterpriseMainorgTotaldataStatRequest() *OapiEnterpriseMainorgTotaldataStatRequest {
	return &OapiEnterpriseMainorgTotaldataStatRequest{}
}

type OapiEnterpriseMainorgTotaldataStatRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEnterpriseMainorgTotaldataStatResponse
	CorpId          string
	OrderBy         string
	PageSize        int64
	PageStart       int64
	ReturnFields    string
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) SetCorpId(corpId2 string) {
	this.CorpId = corpId2
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetCorpId() string {
	return this.CorpId
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) SetOrderBy(orderBy2 string) {
	this.OrderBy = orderBy2
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetOrderBy() string {
	return this.OrderBy
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) SetPageStart(pageStart2 int64) {
	this.PageStart = pageStart2
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetPageStart() int64 {
	return this.PageStart
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) SetReturnFields(returnFields2 string) {
	this.ReturnFields = returnFields2
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetReturnFields() string {
	return this.ReturnFields
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetApiMethodName() string {
	return "dingtalk.oapi.enterprise.mainorg.totaldata.stat"
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corp_id"] = this.CorpId
	txtParams["order_by"] = this.OrderBy
	txtParams["page_size"] = this.PageSize
	txtParams["page_start"] = this.PageStart
	txtParams["return_fields"] = this.ReturnFields
	txtParams["stat_date"] = this.StatDate
	return txtParams
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpId, "corpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEnterpriseMainorgTotaldataStatRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEnterpriseMainorgTotaldataStatResponse struct {
	taobao.TaobaoResponse
	Result  []MainStatDataVo `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type MainStatDataVo struct {
	ActiveMbrCntStd       string `json:"active_mbr_cnt_std,omitempty"`
	AllGroupCnt           string `json:"all_group_cnt,omitempty"`
	CarbonAmount1d        string `json:"carbon_amount_1d,omitempty"`
	CarbonAmount1w        string `json:"carbon_amount_1w,omitempty"`
	CorpId                string `json:"corp_id,omitempty"`
	DeptGroupCnt          string `json:"dept_group_cnt,omitempty"`
	DingSaveHour1d        string `json:"ding_save_hour_1d,omitempty"`
	DingSaveHour1w        string `json:"ding_save_hour_1w,omitempty"`
	EndMessageUserCnt1w   string `json:"end_message_user_cnt_1w,omitempty"`
	InnerGroupCnt         string `json:"innerGroupCnt,omitempty"`
	LiveLaunchSuccCnt1d   string `json:"liveLaunchSuccCnt1d,omitempty"`
	LiveLaunchSuccCnt1w   string `json:"liveLaunchSuccCnt1w,omitempty"`
	MbrCntStd             string `json:"mbr_cnt_std,omitempty"`
	OnlineConferenceCnt1d string `json:"online_conference_cnt_1d,omitempty"`
	OnlineConferenceCnt7d string `json:"online_conference_cnt_7d,omitempty"`
	OnlineOrgNt           string `json:"online_org_nt,omitempty"`
	OrgOnlineRatio        string `json:"org_online_ratio,omitempty"`
	RealOrgCnt            string `json:"real_org_cnt,omitempty"`
	ReceiveDingUserCnt1d  string `json:"receive_ding_user_cnt_1d,omitempty"`
	ReceiveDingUserCnt1w  string `json:"receive_ding_user_cnt_1w,omitempty"`
	RelOrgNt              string `json:"rel_org_nt,omitempty"`
	SendMessageUserCnt1d  string `json:"send_message_user_cnt_1d,omitempty"`
	StatDate              string `json:"stat_date,omitempty"`
}
