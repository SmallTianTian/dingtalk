package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiAtsJobGetRequest() *OapiAtsJobGetRequest {
	return &OapiAtsJobGetRequest{}
}

type OapiAtsJobGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsJobGetResponse
	BizCode         string
	JobId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsJobGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsJobGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsJobGetRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsJobGetRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsJobGetRequest) SetJobId(jobId2 string) {
	this.JobId = jobId2
}
func (this *OapiAtsJobGetRequest) GetJobId() string {
	return this.JobId
}
func (this *OapiAtsJobGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.job.get"
}
func (this *OapiAtsJobGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsJobGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsJobGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsJobGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsJobGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["job_id"] = this.JobId
	return txtParams
}
func (this *OapiAtsJobGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsJobGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsJobGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsJobGetResponse struct {
	taobao.TaobaoResponse
	Result JobSimpleVO `json:"result,omitempty"`
}
type JobAddressVO struct {
	Detail    string `json:"detail,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Name      string `json:"name,omitempty"`
}
type TopJobPartTimeExtDataVo struct {
	ContactNumber    string    `json:"contact_number,omitempty"`
	SalaryPeriod     string    `json:"salary_period,omitempty"`
	SettleType       string    `json:"settle_type,omitempty"`
	SpecifyWorkDate  bool      `json:"specify_work_date,omitempty"`
	SpecifyWorkTime  bool      `json:"specify_work_time,omitempty"`
	WorkBeginTimeMin int64     `json:"work_begin_time_min,omitempty"`
	WorkDateType     string    `json:"work_date_type,omitempty"`
	WorkEndDate      time.Time `json:"work_end_date,omitempty"`
	WorkEndTimeMin   int64     `json:"work_end_time_min,omitempty"`
	WorkStartDate    time.Time `json:"work_start_date,omitempty"`
}
type JobSimpleVO struct {
	Address          JobAddressVO            `json:"address,omitempty"`
	Campus           bool                    `json:"campus,omitempty"`
	Category         string                  `json:"category,omitempty"`
	City             string                  `json:"city,omitempty"`
	Corpid           string                  `json:"corpid,omitempty"`
	Description      string                  `json:"description,omitempty"`
	District         string                  `json:"district,omitempty"`
	HeadCount        int64                   `json:"head_count,omitempty"`
	JobCode          string                  `json:"job_code,omitempty"`
	JobId            string                  `json:"job_id,omitempty"`
	JobNature        string                  `json:"job_nature,omitempty"`
	MainDeptId       int64                   `json:"main_dept_id,omitempty"`
	MaxJobExperience int64                   `json:"max_job_experience,omitempty"`
	MaxSalary        int64                   `json:"max_salary,omitempty"`
	MinJobExperience int64                   `json:"min_job_experience,omitempty"`
	MinSalary        int64                   `json:"min_salary,omitempty"`
	Name             string                  `json:"name,omitempty"`
	PartTimeData     TopJobPartTimeExtDataVo `json:"part_time_data,omitempty"`
	Province         string                  `json:"province,omitempty"`
	RequiredEdu      int64                   `json:"required_edu,omitempty"`
	SalaryMonth      int64                   `json:"salary_month,omitempty"`
	SalaryNegotiable bool                    `json:"salary_negotiable,omitempty"`
	SalaryPeriod     string                  `json:"salary_period,omitempty"`
	Tags             []string                `json:"tags,omitempty"`
}
