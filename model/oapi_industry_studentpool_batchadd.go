package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiIndustryStudentpoolBatchaddRequest() *OapiIndustryStudentpoolBatchaddRequest {
	return &OapiIndustryStudentpoolBatchaddRequest{}
}

type OapiIndustryStudentpoolBatchaddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiIndustryStudentpoolBatchaddResponse
	BizCode         string
	StudentList     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiIndustryStudentpoolBatchaddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIndustryStudentpoolBatchaddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIndustryStudentpoolBatchaddRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiIndustryStudentpoolBatchaddRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiIndustryStudentpoolBatchaddRequest) SetStudentList(studentList2 string) {
	this.StudentList = studentList2
}
func (this *OapiIndustryStudentpoolBatchaddRequest) GetStudentList() string {
	return this.StudentList
}
func (this *OapiIndustryStudentpoolBatchaddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.industry.studentpool.batchadd"
}
func (this *OapiIndustryStudentpoolBatchaddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIndustryStudentpoolBatchaddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIndustryStudentpoolBatchaddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIndustryStudentpoolBatchaddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIndustryStudentpoolBatchaddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["student_list"] = this.StudentList
	return txtParams
}
func (this *OapiIndustryStudentpoolBatchaddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiIndustryStudentpoolBatchaddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIndustryStudentpoolBatchaddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PatriarchDto struct {
	Mobile   string `json:"mobile,omitempty"`
	Name     string `json:"name,omitempty"`
	Relation string `json:"relation,omitempty"`
}
type TrainingStudentDto struct {
	Mobile     string         `json:"mobile,omitempty"`
	Name       string         `json:"name,omitempty"`
	Number     string         `json:"number,omitempty"`
	PackageIds []int64        `json:"package_ids,omitempty"`
	Patriarchs []PatriarchDto `json:"patriarchs,omitempty"`
}
type OapiIndustryStudentpoolBatchaddResponse struct {
	taobao.TaobaoResponse
	Result  []string `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
