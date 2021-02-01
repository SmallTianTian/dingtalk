package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduSubDataGetRequest() *OapiEduSubDataGetRequest {
	return &OapiEduSubDataGetRequest{}
}

type OapiEduSubDataGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduSubDataGetResponse
	Orders          string
	PageNum         int64
	PageSize        int64
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduSubDataGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduSubDataGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduSubDataGetRequest) SetOrders(orders2 string) {
	this.Orders = orders2
}
func (this *OapiEduSubDataGetRequest) GetOrders() string {
	return this.Orders
}
func (this *OapiEduSubDataGetRequest) SetPageNum(pageNum2 int64) {
	this.PageNum = pageNum2
}
func (this *OapiEduSubDataGetRequest) GetPageNum() int64 {
	return this.PageNum
}
func (this *OapiEduSubDataGetRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEduSubDataGetRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEduSubDataGetRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiEduSubDataGetRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiEduSubDataGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.sub.data.get"
}
func (this *OapiEduSubDataGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduSubDataGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduSubDataGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduSubDataGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduSubDataGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["orders"] = this.Orders
	txtParams["page_num"] = this.PageNum
	txtParams["page_size"] = this.PageSize
	txtParams["stat_date"] = this.StatDate
	return txtParams
}
func (this *OapiEduSubDataGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Orders, 20, "orders"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduSubDataGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduSubDataGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenOrderDataRequest struct {
	FieldName string `json:"field_name,omitempty"`
	Order     string `json:"order,omitempty"`
}
type OapiEduSubDataGetResponse struct {
	taobao.TaobaoResponse
	Result  []OpenEduSchoolStatisticalDataResponse `json:"result,omitempty"`
	Success bool                                   `json:"success,omitempty"`
}
type OpenEduSchoolStatisticalDataResponse struct {
	ActClassCircleCnt1d         string `json:"act_class_circle_cnt_1d,omitempty"`
	ActClassCircleCnt7d         string `json:"act_class_circle_cnt_7d,omitempty"`
	ActClassGroupCnt1d          string `json:"act_class_group_cnt_1d,omitempty"`
	ActClassGroupCnt7d          string `json:"act_class_group_cnt_7d,omitempty"`
	ActPatriarchCnt1d           string `json:"act_patriarch_cnt_1d,omitempty"`
	ActPatriarchCnt7d           string `json:"act_patriarch_cnt_7d,omitempty"`
	ActTeacherCnt1d             string `json:"act_teacher_cnt_1d,omitempty"`
	ActTeacherCnt7d             string `json:"act_teacher_cnt_7d,omitempty"`
	ActUsrCnt1d                 string `json:"act_usr_cnt_1d,omitempty"`
	ActUsrCnt7d                 string `json:"act_usr_cnt_7d,omitempty"`
	ActUsrRatio1d               string `json:"act_usr_ratio_1d,omitempty"`
	ActUsrRatio7d               string `json:"act_usr_ratio_7d,omitempty"`
	ActiveMbrCntStd             string `json:"active_mbr_cnt_std,omitempty"`
	AuthTeacherCntStd           string `json:"auth_teacher_cnt_std,omitempty"`
	ClassCardUserCnt1d          string `json:"class_card_user_cnt_1d,omitempty"`
	ClassCardUserCnt7d          string `json:"class_card_user_cnt_7d,omitempty"`
	ClassCircleUserCnt1d        string `json:"class_circle_user_cnt_1d,omitempty"`
	ClassCircleUserCnt7d        string `json:"class_circle_user_cnt_7d,omitempty"`
	ClassCntStd                 string `json:"class_cnt_std,omitempty"`
	ClassGroupUserCnt1d         string `json:"class_group_user_cnt_1d,omitempty"`
	ClassGroupUserCnt7d         string `json:"class_group_user_cnt_7d,omitempty"`
	CorpId                      string `json:"corp_id,omitempty"`
	DingIndex1d                 string `json:"ding_index_1d,omitempty"`
	DingIndex7d                 string `json:"ding_index_7d,omitempty"`
	LiveLaunchSuccCnt1d         string `json:"live_launch_succ_cnt_1d,omitempty"`
	LiveLaunchSuccCnt7d         string `json:"live_launch_succ_cnt_7d,omitempty"`
	LivePlayUserCnt1d           string `json:"live_play_user_cnt_1d,omitempty"`
	LivePlayUserCnt7d           string `json:"live_play_user_cnt_7d,omitempty"`
	LiveSuccTimeLen1d           string `json:"live_succ_time_len_1d,omitempty"`
	LiveSuccTimeLen7d           string `json:"live_succ_time_len_7d,omitempty"`
	MbrActiveRatio              string `json:"mbr_active_ratio,omitempty"`
	MbrCntStd                   string `json:"mbr_cnt_std,omitempty"`
	MultiPatriarchStudentCnt    string `json:"multi_patriarch_student_cnt,omitempty"`
	MultiPatriarchStudentRatio  string `json:"multi_patriarch_student_ratio,omitempty"`
	NonePatriarchStudentCnt     string `json:"none_patriarch_student_cnt,omitempty"`
	NonePatriarchStudentRatio   string `json:"none_patriarch_student_ratio,omitempty"`
	PatriarchCntStd             string `json:"patriarch_cnt_std,omitempty"`
	RcvDingPatriarchCnt1d       string `json:"rcv_ding_patriarch_cnt_1d,omitempty"`
	RcvDingPatriarchCnt7d       string `json:"rcv_ding_patriarch_cnt_7d,omitempty"`
	SendCirclePostCnt1d         string `json:"send_circle_post_cnt_1d,omitempty"`
	SendCirclePostCnt7d         string `json:"send_circle_post_cnt_7d,omitempty"`
	SendMsgUserCnt1d            string `json:"send_msg_user_cnt_1d,omitempty"`
	SendMsgUserCnt7d            string `json:"send_msg_user_cnt_7d,omitempty"`
	SendMsgUserRatio1d          string `json:"send_msg_user_ratio_1d,omitempty"`
	SendMsgUserRatio7d          string `json:"send_msg_user_ratio_7d,omitempty"`
	SinglePatriarchStudentCnt   string `json:"single_patriarch_student_cnt,omitempty"`
	SinglePatriarchStudentRatio string `json:"single_patriarch_student_ratio,omitempty"`
	StatDate                    string `json:"stat_date,omitempty"`
	StudentCntStd               string `json:"student_cnt_std,omitempty"`
	SubCorpAreaLat              string `json:"sub_corp_area_lat,omitempty"`
	SubCorpAreaLng              string `json:"sub_corp_area_lng,omitempty"`
	SubCorpId                   string `json:"sub_corp_id,omitempty"`
	SubCorpName                 string `json:"sub_corp_name,omitempty"`
	TeacherCntStd               string `json:"teacher_cnt_std,omitempty"`
	TeacherSendDingCnt1d        string `json:"teacher_send_ding_cnt_1d,omitempty"`
	TeacherSendDingCnt7d        string `json:"teacher_send_ding_cnt_7d,omitempty"`
	TwoPatriarchStudentCnt      string `json:"two_patriarch_student_cnt,omitempty"`
	TwoPatriarchStudentRatio    string `json:"two_patriarch_student_ratio,omitempty"`
}
