package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringMealconfigGetRequest() *OapiCateringMealconfigGetRequest {
	return &OapiCateringMealconfigGetRequest{}
}

type OapiCateringMealconfigGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringMealconfigGetResponse
	MealDayOffset   int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringMealconfigGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringMealconfigGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringMealconfigGetRequest) SetMealDayOffset(mealDayOffset2 int64) {
	this.MealDayOffset = mealDayOffset2
}
func (this *OapiCateringMealconfigGetRequest) GetMealDayOffset() int64 {
	return this.MealDayOffset
}
func (this *OapiCateringMealconfigGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.mealconfig.get"
}
func (this *OapiCateringMealconfigGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringMealconfigGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringMealconfigGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringMealconfigGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringMealconfigGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["meal_day_offset"] = this.MealDayOffset
	return txtParams
}
func (this *OapiCateringMealconfigGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MealDayOffset, "mealDayOffset"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringMealconfigGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringMealconfigGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringMealconfigGetResponse struct {
	taobao.TaobaoResponse
	Result  GroupMealSettingVo `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
