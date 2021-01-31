package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduSchoolInitRequest() *OapiEduSchoolInitRequest {
	return &OapiEduSchoolInitRequest{}
}

type OapiEduSchoolInitRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduSchoolInitResponse
	Campus          string
	Operator        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduSchoolInitRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduSchoolInitRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduSchoolInitRequest) SetCampus(campus2 string) {
	this.Campus = campus2
}
func (this *OapiEduSchoolInitRequest) GetCampus() string {
	return this.Campus
}
func (this *OapiEduSchoolInitRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiEduSchoolInitRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiEduSchoolInitRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.school.init"
}
func (this *OapiEduSchoolInitRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduSchoolInitRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduSchoolInitRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduSchoolInitRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduSchoolInitRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["campus"] = this.Campus
	txtParams["operator"] = this.Operator
	return txtParams
}
func (this *OapiEduSchoolInitRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Operator, "operator"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduSchoolInitRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduSchoolInitRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCampus struct {
	Name    string       `json:"name,omitempty"`
	Periods []OpenPeriod `json:"periods,omitempty"`
}
type OapiEduSchoolInitResponse struct {
	taobao.TaobaoResponse
	Result  OpenEduSchoolInitResponse `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type OpenEduSchoolInitResponse struct {
	CampusList []int64 `json:"campus_list,omitempty"`
	Effected   string  `json:"effected,omitempty"`
}
