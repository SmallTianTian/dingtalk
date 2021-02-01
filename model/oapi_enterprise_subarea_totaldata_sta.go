package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEnterpriseSubareaTotaldataStatRequest() *OapiEnterpriseSubareaTotaldataStatRequest {
	return &OapiEnterpriseSubareaTotaldataStatRequest{}
}

type OapiEnterpriseSubareaTotaldataStatRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEnterpriseSubareaTotaldataStatResponse
	CorpId          string
	OrderBy         string
	PageSize        int64
	PageStart       int64
	ReturnFields    string
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) SetCorpId(corpId2 string) {
	this.CorpId = corpId2
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetCorpId() string {
	return this.CorpId
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) SetOrderBy(orderBy2 string) {
	this.OrderBy = orderBy2
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetOrderBy() string {
	return this.OrderBy
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) SetPageStart(pageStart2 int64) {
	this.PageStart = pageStart2
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetPageStart() int64 {
	return this.PageStart
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) SetReturnFields(returnFields2 string) {
	this.ReturnFields = returnFields2
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetReturnFields() string {
	return this.ReturnFields
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetApiMethodName() string {
	return "dingtalk.oapi.enterprise.subarea.totaldata.stat"
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corp_id"] = this.CorpId
	txtParams["order_by"] = this.OrderBy
	txtParams["page_size"] = this.PageSize
	txtParams["page_start"] = this.PageStart
	txtParams["return_fields"] = this.ReturnFields
	txtParams["stat_date"] = this.StatDate
	return txtParams
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpId, "corpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEnterpriseSubareaTotaldataStatRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEnterpriseSubareaTotaldataStatResponse struct {
	taobao.TaobaoResponse
	Result  []AreaStatDataVO `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type AreaStatDataVO struct {
	ActRatio1d           string `json:"act_ratio_1d,omitempty"`
	ActUsrCn1d           string `json:"act_usr_cn_1d,omitempty"`
	ActiveMbrCntStd      string `json:"active_mbr_cnt_std,omitempty"`
	ActiveMbrRatio       string `json:"active_mbr_ratio,omitempty"`
	CityLat              string `json:"city_lat,omitempty"`
	CityLng              string `json:"city_lng,omitempty"`
	CityName             string `json:"city_name,omitempty"`
	CorpId               string `json:"corp_id,omitempty"`
	CountyLat            string `json:"county_lat,omitempty"`
	CountyLng            string `json:"county_lng,omitempty"`
	CountyName           string `json:"county_name,omitempty"`
	MbrCntStd            string `json:"mbr_cnt_std,omitempty"`
	OnlineOrgCnt         string `json:"online_org_cnt,omitempty"`
	OrgOnlineRatio       string `json:"org_online_ratio,omitempty"`
	RealOrgCnt           string `json:"real_org_cnt,omitempty"`
	SendMessageCnt1d     string `json:"send_message_cnt_1d,omitempty"`
	SendMessageUserCnt1d string `json:"send_message_user_cnt_1d,omitempty"`
	StatDate             string `json:"stat_date,omitempty"`
}
