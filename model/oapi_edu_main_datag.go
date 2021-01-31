package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduMainDataGetRequest() *OapiEduMainDataGetRequest {
	return &OapiEduMainDataGetRequest{}
}

type OapiEduMainDataGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduMainDataGetResponse
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduMainDataGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduMainDataGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduMainDataGetRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiEduMainDataGetRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiEduMainDataGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.main.data.get"
}
func (this *OapiEduMainDataGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduMainDataGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduMainDataGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduMainDataGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduMainDataGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["stat_date"] = this.StatDate
	return txtParams
}
func (this *OapiEduMainDataGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.StatDate, "statDate"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduMainDataGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduMainDataGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduMainDataGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                                `json:"errcode,omitempty"`
	Errmsg  string                               `json:"errmsg,omitempty"`
	Result  OpenEduBureauStatisticalDataResponse `json:"result,omitempty"`
	Success bool                                 `json:"success,omitempty"`
}
type OpenEduBureauStatisticalDataResponse struct {
	ActClassCircleCnt1d   string `json:"act_class_circle_cnt_1d,omitempty"`
	ActClassCircleCnt7d   string `json:"act_class_circle_cnt_7d,omitempty"`
	ActClassGroupCnt1d    string `json:"act_class_group_cnt_1d,omitempty"`
	ActClassGroupCnt7d    string `json:"act_class_group_cnt_7d,omitempty"`
	ActPatriarchCnt1d     string `json:"act_patriarch_cnt_1d,omitempty"`
	ActPatriarchCnt7d     string `json:"act_patriarch_cnt_7d,omitempty"`
	ActTeacherCnt1d       string `json:"act_teacher_cnt_1d,omitempty"`
	ActTeacherCnt7d       string `json:"act_teacher_cnt_7d,omitempty"`
	AuthTeacherCntStd     string `json:"auth_teacher_cnt_std,omitempty"`
	ClassCardUserCnt1d    string `json:"class_card_user_cnt_1d,omitempty"`
	ClassCardUserCnt7d    string `json:"class_card_user_cnt_7d,omitempty"`
	ClassCircleUserCnt1d  string `json:"class_circle_user_cnt_1d,omitempty"`
	ClassCircleUserCnt7d  string `json:"class_circle_user_cnt_7d,omitempty"`
	ClassCntStd           string `json:"class_cnt_std,omitempty"`
	ClassGroupUserCnt1d   string `json:"class_group_user_cnt_1d,omitempty"`
	ClassGroupUserCnt7d   string `json:"class_group_user_cnt_7d,omitempty"`
	CorpId                string `json:"corp_id,omitempty"`
	PatriarchCntStd       string `json:"patriarch_cnt_std,omitempty"`
	RcvDingPatriarchCnt1d string `json:"rcv_ding_patriarch_cnt_1d,omitempty"`
	RcvDingPatriarchCnt7d string `json:"rcv_ding_patriarch_cnt_7d,omitempty"`
	SchoolCntStd          string `json:"school_cnt_std,omitempty"`
	SendCirclePostCnt1d   string `json:"send_circle_post_cnt_1d,omitempty"`
	SendCirclePostCnt7d   string `json:"send_circle_post_cnt_7d,omitempty"`
	StatDate              string `json:"stat_date,omitempty"`
	StudentCntStd         string `json:"student_cnt_std,omitempty"`
	TeacherCntStd         string `json:"teacher_cnt_std,omitempty"`
	TeacherSendDingCnt1d  string `json:"teacher_send_ding_cnt_1d,omitempty"`
	TeacherSendDingCnt7d  string `json:"teacher_send_ding_cnt_7d,omitempty"`
}
