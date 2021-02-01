package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduTeacherGetRequest() *OapiEduTeacherGetRequest {
	return &OapiEduTeacherGetRequest{}
}

type OapiEduTeacherGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduTeacherGetResponse
	ClassId         int64
	TeacherUserid   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduTeacherGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduTeacherGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduTeacherGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduTeacherGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduTeacherGetRequest) SetTeacherUserid(teacherUserid2 string) {
	this.TeacherUserid = teacherUserid2
}
func (this *OapiEduTeacherGetRequest) GetTeacherUserid() string {
	return this.TeacherUserid
}
func (this *OapiEduTeacherGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.teacher.get"
}
func (this *OapiEduTeacherGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduTeacherGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduTeacherGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduTeacherGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduTeacherGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["teacher_userid"] = this.TeacherUserid
	return txtParams
}
func (this *OapiEduTeacherGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduTeacherGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduTeacherGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduTeacherGetResponse struct {
	taobao.TaobaoResponse
	Result  TeacherRespone `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type TeacherRespone struct {
	CampusId    int64  `json:"campus_id,omitempty"`
	ClassId     int64  `json:"class_id,omitempty"`
	GradeId     int64  `json:"grade_id,omitempty"`
	IsAdviser   int64  `json:"is_adviser,omitempty"`
	TeacherName string `json:"teacher_name,omitempty"`
}
