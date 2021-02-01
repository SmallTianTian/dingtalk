package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduPeriodGetRequest() *OapiEduPeriodGetRequest {
	return &OapiEduPeriodGetRequest{}
}

type OapiEduPeriodGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduPeriodGetResponse
	PeriodId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduPeriodGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduPeriodGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduPeriodGetRequest) SetPeriodId(periodId2 int64) {
	this.PeriodId = periodId2
}
func (this *OapiEduPeriodGetRequest) GetPeriodId() int64 {
	return this.PeriodId
}
func (this *OapiEduPeriodGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.period.get"
}
func (this *OapiEduPeriodGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduPeriodGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduPeriodGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduPeriodGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduPeriodGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["period_id"] = this.PeriodId
	return txtParams
}
func (this *OapiEduPeriodGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.PeriodId, "periodId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduPeriodGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduPeriodGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduPeriodGetResponse struct {
	taobao.TaobaoResponse
	Result  PeriodResponse `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type PeriodResponse struct {
	CampusId   int64  `json:"campus_id,omitempty"`
	Name       string `json:"name,omitempty"`
	NameMode   string `json:"name_mode,omitempty"`
	Nick       string `json:"nick,omitempty"`
	PeriodId   int64  `json:"period_id,omitempty"`
	PeriodType string `json:"period_type,omitempty"`
}
