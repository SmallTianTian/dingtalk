package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpHealthStepinfoListRequest() *CorpHealthStepinfoListRequest {
	return &CorpHealthStepinfoListRequest{}
}

type CorpHealthStepinfoListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpHealthStepinfoListResponse
	ObjectId        string
	StatDates       string
	TopHttpMethod   string
	TopResponseType string
	Type            int64
}

func (this *CorpHealthStepinfoListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpHealthStepinfoListRequest) SetObjectId(objectId2 string) {
	this.ObjectId = objectId2
}
func (this *CorpHealthStepinfoListRequest) GetObjectId() string {
	return this.ObjectId
}
func (this *CorpHealthStepinfoListRequest) SetStatDates(statDates2 string) {
	this.StatDates = statDates2
}
func (this *CorpHealthStepinfoListRequest) GetStatDates() string {
	return this.StatDates
}
func (this *CorpHealthStepinfoListRequest) SetType(type2 int64) {
	this.Type = type2
}
func (this *CorpHealthStepinfoListRequest) GetType() int64 {
	return this.Type
}
func (this *CorpHealthStepinfoListRequest) GetApiMethodName() string {
	return "dingtalk.corp.health.stepinfo.list"
}
func (this *CorpHealthStepinfoListRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpHealthStepinfoListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpHealthStepinfoListRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpHealthStepinfoListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpHealthStepinfoListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpHealthStepinfoListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["object_id"] = this.ObjectId
	txtParams["stat_dates"] = this.StatDates
	txtParams["type"] = this.Type
	return txtParams
}
func (this *CorpHealthStepinfoListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ObjectId, "objectId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpHealthStepinfoListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpHealthStepinfoListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpHealthStepinfoListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type BasicStepInfoVo struct {
	StatDate  int64 `json:"stat_date,omitempty"`
	StepCount int64 `json:"step_count,omitempty"`
}
