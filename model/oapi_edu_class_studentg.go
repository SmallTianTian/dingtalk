package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduClassStudentGetRequest() *OapiEduClassStudentGetRequest {
	return &OapiEduClassStudentGetRequest{}
}

type OapiEduClassStudentGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduClassStudentGetResponse
	ClassId         int64
	StudentParam    string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduClassStudentGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduClassStudentGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduClassStudentGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduClassStudentGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduClassStudentGetRequest) SetStudentParam(studentParam2 string) {
	this.StudentParam = studentParam2
}
func (this *OapiEduClassStudentGetRequest) GetStudentParam() string {
	return this.StudentParam
}
func (this *OapiEduClassStudentGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduClassStudentGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduClassStudentGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.class.student.get"
}
func (this *OapiEduClassStudentGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduClassStudentGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduClassStudentGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduClassStudentGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduClassStudentGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["student_param"] = this.StudentParam
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduClassStudentGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.StudentParam, 100, "studentParam"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduClassStudentGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduClassStudentGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenStudentParam struct {
	ClassId    string   `json:"class_id,omitempty"`
	StudentIds []string `json:"student_ids,omitempty"`
}
type OapiEduClassStudentGetResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
