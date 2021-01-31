package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduCardCreateRequest() *OapiEduCardCreateRequest {
	return &OapiEduCardCreateRequest{}
}

type OapiEduCardCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiEduCardCreateResponse
	Opencardcreateparam string
	TopHttpMethod       string
	TopResponseType     string
}

func (this *OapiEduCardCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCardCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCardCreateRequest) SetOpencardcreateparam(opencardcreateparam2 string) {
	this.Opencardcreateparam = opencardcreateparam2
}
func (this *OapiEduCardCreateRequest) GetOpencardcreateparam() string {
	return this.Opencardcreateparam
}
func (this *OapiEduCardCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.card.create"
}
func (this *OapiEduCardCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCardCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCardCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCardCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCardCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["opencardcreateparam"] = this.Opencardcreateparam
	return txtParams
}
func (this *OapiEduCardCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduCardCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCardCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCreateDetailItem struct {
	CanReissueCard        bool      `json:"can_reissue_card,omitempty"`
	CardCycle             int64     `json:"card_cycle,omitempty"`
	CardFrequency         []int64   `json:"card_frequency,omitempty"`
	CardRuleItemParamlist []string  `json:"card_rule_Item_paramlist,omitempty"`
	ClassIds              []string  `json:"class_ids,omitempty"`
	ClassNames            []string  `json:"class_names,omitempty"`
	ClassSelectedStudents string    `json:"class_selected_students,omitempty"`
	Content               string    `json:"content,omitempty"`
	EffectDate            time.Time `json:"effect_date,omitempty"`
	Medias                string    `json:"medias,omitempty"`
	NeedMetering          string    `json:"need_metering,omitempty"`
	RemindHour            int64     `json:"remind_hour,omitempty"`
	RemindMinute          int64     `json:"remind_minute,omitempty"`
	TargetRole            string    `json:"target_role,omitempty"`
	TemplateId            int64     `json:"template_id,omitempty"`
	Title                 string    `json:"title,omitempty"`
	UnitOfMeasurement     string    `json:"unit_of_measurement,omitempty"`
}
type OpenCardCreateParam struct {
	CardBizcode string               `json:"card_bizcode,omitempty"`
	Data        OpenCreateDetailItem `json:"data,omitempty"`
	Identifier  string               `json:"identifier,omitempty"`
	Jsversion   int64                `json:"jsversion,omitempty"`
	Sourcetype  string               `json:"sourcetype,omitempty"`
	Userid      string               `json:"userid,omitempty"`
}
type OapiEduCardCreateResponse struct {
	taobao.TaobaoResponse
	Result  CreateCardResponse `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type CreateCardResponse struct {
	Cardid int64 `json:"cardid,omitempty"`
}
