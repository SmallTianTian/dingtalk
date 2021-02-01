package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkStudentReportSubmitRequest() *OapiEduHomeworkStudentReportSubmitRequest {
	return &OapiEduHomeworkStudentReportSubmitRequest{}
}

type OapiEduHomeworkStudentReportSubmitRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkStudentReportSubmitResponse
	Attributes      string
	BizCode         string
	ClassId         int64
	HwId            int64
	HwReport        string
	HwResult        string
	StudentId       string
	StudentName     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduHomeworkStudentReportSubmitRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetAttributes(attributes2 string) {
	this.Attributes = attributes2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetAttributes() string {
	return this.Attributes
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetHwId(hwId2 int64) {
	this.HwId = hwId2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetHwId() int64 {
	return this.HwId
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetHwReport(hwReport2 string) {
	this.HwReport = hwReport2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetHwReport() string {
	return this.HwReport
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetHwResult(hwResult2 string) {
	this.HwResult = hwResult2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetHwResult() string {
	return this.HwResult
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetStudentId(studentId2 string) {
	this.StudentId = studentId2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetStudentId() string {
	return this.StudentId
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetStudentName(studentName2 string) {
	this.StudentName = studentName2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetStudentName() string {
	return this.StudentName
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.student.report.submit"
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["attributes"] = this.Attributes
	txtParams["biz_code"] = this.BizCode
	txtParams["class_id"] = this.ClassId
	txtParams["hw_id"] = this.HwId
	txtParams["hw_report"] = this.HwReport
	txtParams["hw_result"] = this.HwResult
	txtParams["student_id"] = this.StudentId
	txtParams["student_name"] = this.StudentName
	return txtParams
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkStudentReportSubmitRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkStudentReportSubmitResponse struct {
	taobao.TaobaoResponse
	Result  int64 `json:"result,omitempty"`
	Success bool  `json:"success,omitempty"`
}
