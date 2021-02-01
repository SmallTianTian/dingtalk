package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduGradeGetRequest() *OapiEduGradeGetRequest {
	return &OapiEduGradeGetRequest{}
}

type OapiEduGradeGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduGradeGetResponse
	GradeId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduGradeGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduGradeGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduGradeGetRequest) SetGradeId(gradeId2 int64) {
	this.GradeId = gradeId2
}
func (this *OapiEduGradeGetRequest) GetGradeId() int64 {
	return this.GradeId
}
func (this *OapiEduGradeGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.grade.get"
}
func (this *OapiEduGradeGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduGradeGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduGradeGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduGradeGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduGradeGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["grade_id"] = this.GradeId
	return txtParams
}
func (this *OapiEduGradeGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GradeId, "gradeId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduGradeGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduGradeGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduGradeGetResponse struct {
	taobao.TaobaoResponse
	Result  GradeResponse `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
type GradeResponse struct {
	CampusId   int64  `json:"campus_id,omitempty"`
	GradeLevel int64  `json:"grade_level,omitempty"`
	Name       string `json:"name,omitempty"`
	Nick       string `json:"nick,omitempty"`
	PeriodId   int64  `json:"period_id,omitempty"`
	StartYear  string `json:"start_year,omitempty"`
}
