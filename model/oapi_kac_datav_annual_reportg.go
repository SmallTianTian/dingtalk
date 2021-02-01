package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiKacDatavAnnualReportGetRequest() *OapiKacDatavAnnualReportGetRequest {
	return &OapiKacDatavAnnualReportGetRequest{}
}

type OapiKacDatavAnnualReportGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavAnnualReportGetResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
	Type            int64
	UserId          string
	Year            string
}

func (this *OapiKacDatavAnnualReportGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavAnnualReportGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavAnnualReportGetRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiKacDatavAnnualReportGetRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiKacDatavAnnualReportGetRequest) SetType(type2 int64) {
	this.Type = type2
}
func (this *OapiKacDatavAnnualReportGetRequest) GetType() int64 {
	return this.Type
}
func (this *OapiKacDatavAnnualReportGetRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *OapiKacDatavAnnualReportGetRequest) GetUserId() string {
	return this.UserId
}
func (this *OapiKacDatavAnnualReportGetRequest) SetYear(year2 string) {
	this.Year = year2
}
func (this *OapiKacDatavAnnualReportGetRequest) GetYear() string {
	return this.Year
}
func (this *OapiKacDatavAnnualReportGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.annual_report.get"
}
func (this *OapiKacDatavAnnualReportGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavAnnualReportGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavAnnualReportGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavAnnualReportGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavAnnualReportGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	txtParams["type"] = this.Type
	txtParams["user_id"] = this.UserId
	txtParams["year"] = this.Year
	return txtParams
}
func (this *OapiKacDatavAnnualReportGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Type, "type"); err != nil {
		return err
	}
	return nil
}
func (this *OapiKacDatavAnnualReportGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavAnnualReportGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiKacDatavAnnualReportGetResponse struct {
	taobao.TaobaoResponse
	Result AnnualReportResponse `json:"result,omitempty"`
}
type AnnualReportResponse struct {
	AtCheckDays1y              int64  `json:"at_check_days_1y,omitempty"`
	GeneralFormUserCnt1y       int64  `json:"general_form_user_cnt_1y,omitempty"`
	InnerGroupCnt1y            int64  `json:"inner_group_cnt_1y,omitempty"`
	JoinSuccVideoConfLen1y     string `json:"join_succ_video_conf_len_1y,omitempty"`
	JoinSuccVideoConfNum1y     int64  `json:"join_succ_video_conf_num_1y,omitempty"`
	JoinSuccVideoConfUserCnt1y int64  `json:"join_succ_video_conf_user_cnt_1y,omitempty"`
	JoinSuccVoipConfLen1y      string `json:"join_succ_voip_conf_len_1y,omitempty"`
	JoinSuccVoipConfNum1y      int64  `json:"join_succ_voip_conf_num_1y,omitempty"`
	JoinSuccVoipConfUserCnt1y  int64  `json:"join_succ_voip_conf_user_cnt_1y,omitempty"`
	LiveJoinSuccCnt1y          int64  `json:"live_join_succ_cnt_1y,omitempty"`
	LiveJoinSuccTimeLen1y      string `json:"live_join_succ_time_len_1y,omitempty"`
	OutsideDays1y              int64  `json:"outside_days_1y,omitempty"`
	ProcessInstOperateCnt1y    int64  `json:"process_inst_operate_cnt_1y,omitempty"`
	ProcessInstSubmitCnt1y     int64  `json:"process_inst_submit_cnt_1y,omitempty"`
	UseMicroAppCnt1y           int64  `json:"use_micro_app_cnt_1y,omitempty"`
	UseMicroUserCnt1y          int64  `json:"use_micro_user_cnt_1y,omitempty"`
}
