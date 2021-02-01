package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsResumeAddRequest() *OapiAtsResumeAddRequest {
	return &OapiAtsResumeAddRequest{}
}

type OapiAtsResumeAddRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsResumeAddResponse
	BizCode         string
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsResumeAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsResumeAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsResumeAddRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsResumeAddRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsResumeAddRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiAtsResumeAddRequest) GetParam() string {
	return this.Param
}
func (this *OapiAtsResumeAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.resume.add"
}
func (this *OapiAtsResumeAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsResumeAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsResumeAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsResumeAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsResumeAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiAtsResumeAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsResumeAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsResumeAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type EducationInfo struct {
	Education  string `json:"education,omitempty"`
	EndDate    string `json:"end_date,omitempty"`
	School     string `json:"school,omitempty"`
	Speciality string `json:"speciality,omitempty"`
	StartDate  string `json:"start_date,omitempty"`
	Summary    string `json:"summary,omitempty"`
}
type ExperienceInfo struct {
	Company   string `json:"company,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Title     string `json:"title,omitempty"`
}
type TrainingInfo struct {
	Certificate          string `json:"certificate,omitempty"`
	DescriptionInDetails string `json:"description_in_details,omitempty"`
	EndDate              string `json:"end_date,omitempty"`
	StartDate            string `json:"start_date,omitempty"`
	TrainingInstitution  string `json:"training_institution,omitempty"`
}
type ProjectInfo struct {
	EndDate            string `json:"end_date,omitempty"`
	ProjectDescription string `json:"project_description,omitempty"`
	ProjectName        string `json:"project_name,omitempty"`
	StartDate          string `json:"start_date,omitempty"`
}
type ResumeDetailInfo struct {
	AdvancedDegree     string           `json:"advanced_degree,omitempty"`
	AimSalary          string           `json:"aim_salary,omitempty"`
	BeginWorkTime      string           `json:"begin_work_time,omitempty"`
	Birth              string           `json:"birth,omitempty"`
	Education          string           `json:"education,omitempty"`
	EducationInfoList  []EducationInfo  `json:"education_info_list,omitempty"`
	Email              string           `json:"email,omitempty"`
	ExperienceInfoList []ExperienceInfo `json:"experience_info_list,omitempty"`
	ForwardLocation    string           `json:"forward_location,omitempty"`
	GradeOfEnglish     string           `json:"grade_of_english,omitempty"`
	GraduateTime       string           `json:"graduate_time,omitempty"`
	JobStatus          string           `json:"job_status,omitempty"`
	Married            string           `json:"married,omitempty"`
	Name               string           `json:"name,omitempty"`
	Nationality        string           `json:"nationality,omitempty"`
	NativePlace        string           `json:"native_place,omitempty"`
	NowLocation        string           `json:"now_location,omitempty"`
	PhoneNum           string           `json:"phone_num,omitempty"`
	Political          string           `json:"political,omitempty"`
	ProjectInfoList    []ProjectInfo    `json:"project_info_list,omitempty"`
	Salary             string           `json:"salary,omitempty"`
	School             string           `json:"school,omitempty"`
	SchoolType         string           `json:"school_type,omitempty"`
	Sex                string           `json:"sex,omitempty"`
	Speciality         string           `json:"speciality,omitempty"`
	StudentType        string           `json:"student_type,omitempty"`
	TitleStandard      string           `json:"title_standard,omitempty"`
	TrainingInfoList   []TrainingInfo   `json:"training_info_list,omitempty"`
	VocationStandard   string           `json:"vocation_standard,omitempty"`
}
type CollResumeMailParam struct {
	Channel          string           `json:"channel,omitempty"`
	OptUserId        string           `json:"opt_user_id,omitempty"`
	ResumeDetailInfo ResumeDetailInfo `json:"resume_detail_info,omitempty"`
}
type OapiAtsResumeAddResponse struct {
	taobao.TaobaoResponse
	Result TopCollectResumeResult `json:"result,omitempty"`
}
type TopCollectResumeResult struct {
	MobileJumpUrl string `json:"mobile_jump_url,omitempty"`
	PcJumpUrl     string `json:"pc_jump_url,omitempty"`
	ResumeId      string `json:"resume_id,omitempty"`
}
