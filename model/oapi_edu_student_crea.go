package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduStudentCreateRequest() *OapiEduStudentCreateRequest {
	return &OapiEduStudentCreateRequest{}
}

type OapiEduStudentCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduStudentCreateResponse
	BizId           string
	ClassId         int64
	Name            string
	Operator        string
	StudentNo       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduStudentCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduStudentCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduStudentCreateRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiEduStudentCreateRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiEduStudentCreateRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduStudentCreateRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduStudentCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiEduStudentCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiEduStudentCreateRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiEduStudentCreateRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiEduStudentCreateRequest) SetStudentNo(studentNo2 string) {
	this.StudentNo = studentNo2
}
func (this *OapiEduStudentCreateRequest) GetStudentNo() string {
	return this.StudentNo
}
func (this *OapiEduStudentCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.student.create"
}
func (this *OapiEduStudentCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduStudentCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduStudentCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduStudentCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduStudentCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["class_id"] = this.ClassId
	txtParams["name"] = this.Name
	txtParams["operator"] = this.Operator
	txtParams["student_no"] = this.StudentNo
	return txtParams
}
func (this *OapiEduStudentCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduStudentCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduStudentCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduStudentCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
