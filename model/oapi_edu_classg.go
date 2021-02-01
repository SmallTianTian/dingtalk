package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduClassGetRequest() *OapiEduClassGetRequest {
	return &OapiEduClassGetRequest{}
}

type OapiEduClassGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduClassGetResponse
	ClassId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduClassGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduClassGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduClassGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduClassGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduClassGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.class.get"
}
func (this *OapiEduClassGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduClassGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduClassGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduClassGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduClassGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	return txtParams
}
func (this *OapiEduClassGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduClassGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduClassGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduClassGetResponse struct {
	taobao.TaobaoResponse
	Result  ClassResponse `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
type ClassResponse struct {
	CampusId   int64  `json:"campus_id,omitempty"`
	ClassLevel int64  `json:"class_level,omitempty"`
	GradeId    int64  `json:"grade_id,omitempty"`
	GradeLevel int64  `json:"grade_level,omitempty"`
	Name       string `json:"name,omitempty"`
	Nick       string `json:"nick,omitempty"`
	PeriodId   int64  `json:"period_id,omitempty"`
}
