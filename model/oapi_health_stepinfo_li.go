package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHealthStepinfoListRequest() *OapiHealthStepinfoListRequest {
	return &OapiHealthStepinfoListRequest{}
}

type OapiHealthStepinfoListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHealthStepinfoListResponse
	ObjectId        string
	StatDates       string
	TopHttpMethod   string
	TopResponseType string
	Type            int64
}

func (this *OapiHealthStepinfoListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHealthStepinfoListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHealthStepinfoListRequest) SetObjectId(objectId2 string) {
	this.ObjectId = objectId2
}
func (this *OapiHealthStepinfoListRequest) GetObjectId() string {
	return this.ObjectId
}
func (this *OapiHealthStepinfoListRequest) SetStatDates(statDates2 string) {
	this.StatDates = statDates2
}
func (this *OapiHealthStepinfoListRequest) GetStatDates() string {
	return this.StatDates
}
func (this *OapiHealthStepinfoListRequest) SetType(type2 int64) {
	this.Type = type2
}
func (this *OapiHealthStepinfoListRequest) GetType() int64 {
	return this.Type
}
func (this *OapiHealthStepinfoListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.health.stepinfo.list"
}
func (this *OapiHealthStepinfoListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHealthStepinfoListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHealthStepinfoListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHealthStepinfoListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHealthStepinfoListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["object_id"] = this.ObjectId
	txtParams["stat_dates"] = this.StatDates
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiHealthStepinfoListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ObjectId, "objectId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHealthStepinfoListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHealthStepinfoListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHealthStepinfoListResponse struct {
	taobao.TaobaoResponse

	StepinfoList []BasicStepInfoVo `json:"stepinfo_list,omitempty"`
}
