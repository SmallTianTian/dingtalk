package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduClassStudentListRequest() *OapiEduClassStudentListRequest {
	return &OapiEduClassStudentListRequest{}
}

type OapiEduClassStudentListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduClassStudentListResponse
	ClassId         int64
	StudentParam    string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduClassStudentListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduClassStudentListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduClassStudentListRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduClassStudentListRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduClassStudentListRequest) SetStudentParam(studentParam2 string) {
	this.StudentParam = studentParam2
}
func (this *OapiEduClassStudentListRequest) GetStudentParam() string {
	return this.StudentParam
}
func (this *OapiEduClassStudentListRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduClassStudentListRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduClassStudentListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.class.student.list"
}
func (this *OapiEduClassStudentListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduClassStudentListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduClassStudentListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduClassStudentListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduClassStudentListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["student_param"] = this.StudentParam
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduClassStudentListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduClassStudentListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduClassStudentListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduClassStudentListResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
