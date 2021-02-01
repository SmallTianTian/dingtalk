package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduTypeDataGetRequest() *OapiEduTypeDataGetRequest {
	return &OapiEduTypeDataGetRequest{}
}

type OapiEduTypeDataGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduTypeDataGetResponse
	Orders          string
	PageNum         string
	PageSize        string
	StatDate        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduTypeDataGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduTypeDataGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduTypeDataGetRequest) SetOrders(orders2 string) {
	this.Orders = orders2
}
func (this *OapiEduTypeDataGetRequest) GetOrders() string {
	return this.Orders
}
func (this *OapiEduTypeDataGetRequest) SetPageNum(pageNum2 string) {
	this.PageNum = pageNum2
}
func (this *OapiEduTypeDataGetRequest) GetPageNum() string {
	return this.PageNum
}
func (this *OapiEduTypeDataGetRequest) SetPageSize(pageSize2 string) {
	this.PageSize = pageSize2
}
func (this *OapiEduTypeDataGetRequest) GetPageSize() string {
	return this.PageSize
}
func (this *OapiEduTypeDataGetRequest) SetStatDate(statDate2 string) {
	this.StatDate = statDate2
}
func (this *OapiEduTypeDataGetRequest) GetStatDate() string {
	return this.StatDate
}
func (this *OapiEduTypeDataGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.type.data.get"
}
func (this *OapiEduTypeDataGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduTypeDataGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduTypeDataGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduTypeDataGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduTypeDataGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["orders"] = this.Orders
	txtParams["page_num"] = this.PageNum
	txtParams["page_size"] = this.PageSize
	txtParams["stat_date"] = this.StatDate
	return txtParams
}
func (this *OapiEduTypeDataGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Orders, 20, "orders"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduTypeDataGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduTypeDataGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduTypeDataGetResponse struct {
	taobao.TaobaoResponse
	Result  []OpenEduSchoolTypeStatisticalDataResponse `json:"result,omitempty"`
	Success bool                                       `json:"success,omitempty"`
}
type OpenEduSchoolTypeStatisticalDataResponse struct {
	AuthTeacherCntStd           string `json:"auth_teacher_cnt_std,omitempty"`
	ClassCntStd                 string `json:"class_cnt_std,omitempty"`
	CorpId                      string `json:"corp_id,omitempty"`
	MultiPatriarchStudentCnt    string `json:"multi_patriarch_student_cnt,omitempty"`
	MultiPatriarchStudentRatio  string `json:"multi_patriarch_student_ratio,omitempty"`
	NonePatriarchStudentCnt     string `json:"none_patriarch_student_cnt,omitempty"`
	NonePatriarchStudentRatio   string `json:"none_patriarch_student_ratio,omitempty"`
	PatriarchCntStd             string `json:"patriarch_cnt_std,omitempty"`
	SchoolCntStd                string `json:"school_cnt_std,omitempty"`
	SchoolType                  string `json:"school_type,omitempty"`
	SinglePatriarchStudentCnt   string `json:"single_patriarch_student_cnt,omitempty"`
	SinglePatriarchStudentRatio string `json:"single_patriarch_student_ratio,omitempty"`
	StatDate                    string `json:"stat_date,omitempty"`
	StudentCntStd               string `json:"student_cnt_std,omitempty"`
	TeacherCntStd               string `json:"teacher_cnt_std,omitempty"`
	TwoPatriarchStudentCnt      string `json:"two_patriarch_student_cnt,omitempty"`
	TwoPatriarchStudentRatio    string `json:"two_patriarch_student_ratio,omitempty"`
}
