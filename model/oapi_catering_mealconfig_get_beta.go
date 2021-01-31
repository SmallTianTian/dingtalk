package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringMealconfigGetBetaRequest() *OapiCateringMealconfigGetBetaRequest {
	return &OapiCateringMealconfigGetBetaRequest{}
}

type OapiCateringMealconfigGetBetaRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringMealconfigGetBetaResponse
	MealDayOffset   int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringMealconfigGetBetaRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringMealconfigGetBetaRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringMealconfigGetBetaRequest) SetMealDayOffset(mealDayOffset2 int64) {
	this.MealDayOffset = mealDayOffset2
}
func (this *OapiCateringMealconfigGetBetaRequest) GetMealDayOffset() int64 {
	return this.MealDayOffset
}
func (this *OapiCateringMealconfigGetBetaRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.mealconfig.get.beta"
}
func (this *OapiCateringMealconfigGetBetaRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringMealconfigGetBetaRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringMealconfigGetBetaRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringMealconfigGetBetaRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringMealconfigGetBetaRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["meal_day_offset"] = this.MealDayOffset
	return txtParams
}
func (this *OapiCateringMealconfigGetBetaRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MealDayOffset, "mealDayOffset"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringMealconfigGetBetaRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringMealconfigGetBetaRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringMealconfigGetBetaResponse struct {
	taobao.TaobaoResponse
	Result  GroupMealSettingVo `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type Addressvos struct {
	Address       string `json:"address,omitempty"`
	AddressDetail string `json:"address_detail,omitempty"`
	AddressId     int64  `json:"address_id,omitempty"`
}
type GroupMealSettingVo struct {
	Address           string         `json:"address,omitempty"`
	AddressDetail     string         `json:"address_detail,omitempty"`
	AddressId         int64          `json:"address_id,omitempty"`
	AddressList       []Addressvos   `json:"address_list,omitempty"`
	ComingMealDayList []string       `json:"coming_meal_day_list,omitempty"`
	CorpId            string         `json:"corp_id,omitempty"`
	MealItemList      []Mealitemlist `json:"meal_item_list,omitempty"`
	MealTime          int64          `json:"meal_time,omitempty"`
}
