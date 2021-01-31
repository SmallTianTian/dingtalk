package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkStudentTopicRecordRequest() *OapiEduHomeworkStudentTopicRecordRequest {
	return &OapiEduHomeworkStudentTopicRecordRequest{}
}

type OapiEduHomeworkStudentTopicRecordRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                 OapiEduHomeworkStudentTopicRecordResponse
	StudentAnswerDetails string
	TopHttpMethod        string
	TopResponseType      string
}

func (this *OapiEduHomeworkStudentTopicRecordRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) SetStudentAnswerDetails(studentAnswerDetails2 string) {
	this.StudentAnswerDetails = studentAnswerDetails2
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) GetStudentAnswerDetails() string {
	return this.StudentAnswerDetails
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.student.topic.record"
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["student_answer_details"] = this.StudentAnswerDetails
	return txtParams
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.StudentAnswerDetails, 20, "studentAnswerDetails"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkStudentTopicRecordRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type StudentAnswerDetail struct {
	Attributes  string `json:"attributes,omitempty"`
	BizCode     string `json:"biz_code,omitempty"`
	ClassId     string `json:"class_id,omitempty"`
	HwId        int64  `json:"hw_id,omitempty"`
	IsRight     string `json:"is_right,omitempty"`
	QuestionId  string `json:"question_id,omitempty"`
	RedoTimes   int64  `json:"redo_times,omitempty"`
	SpendTime   int64  `json:"spend_time,omitempty"`
	StudentId   string `json:"student_id,omitempty"`
	StudentName string `json:"student_name,omitempty"`
}
type OapiEduHomeworkStudentTopicRecordResponse struct {
	taobao.TaobaoResponse
	Result  []int64 `json:"result,omitempty"`
	Success bool    `json:"success,omitempty"`
}
