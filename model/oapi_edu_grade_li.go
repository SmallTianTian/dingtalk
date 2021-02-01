package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduGradeListRequest() *OapiEduGradeListRequest {
	return &OapiEduGradeListRequest{}
}

type OapiEduGradeListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduGradeListResponse
	PeriodId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduGradeListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduGradeListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduGradeListRequest) SetPeriodId(periodId2 int64) {
	this.PeriodId = periodId2
}
func (this *OapiEduGradeListRequest) GetPeriodId() int64 {
	return this.PeriodId
}
func (this *OapiEduGradeListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.grade.list"
}
func (this *OapiEduGradeListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduGradeListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduGradeListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduGradeListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduGradeListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["period_id"] = this.PeriodId
	return txtParams
}
func (this *OapiEduGradeListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.PeriodId, "periodId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduGradeListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduGradeListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduGradeListResponse struct {
	taobao.TaobaoResponse
	Result  []GradeResponse `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
