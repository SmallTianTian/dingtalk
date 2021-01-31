package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoHumanresProductionteamQueryRequest() *OapiRhinoHumanresProductionteamQueryRequest {
	return &OapiRhinoHumanresProductionteamQueryRequest{}
}

type OapiRhinoHumanresProductionteamQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                     OapiRhinoHumanresProductionteamQueryResponse
	QueryProductionTeamParam string
	TopHttpMethod            string
	TopResponseType          string
}

func (this *OapiRhinoHumanresProductionteamQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) SetQueryProductionTeamParam(queryProductionTeamParam2 string) {
	this.QueryProductionTeamParam = queryProductionTeamParam2
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) GetQueryProductionTeamParam() string {
	return this.QueryProductionTeamParam
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.humanres.productionteam.query"
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["query_production_team_param"] = this.QueryProductionTeamParam
	return txtParams
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoHumanresProductionteamQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type QueryProductionTeamDto struct {
	EndTime             time.Time `json:"end_time,omitempty"`
	StartTime           time.Time `json:"start_time,omitempty"`
	TenantId            string    `json:"tenant_id,omitempty"`
	Userid              string    `json:"userid,omitempty"`
	WorkshopSectionCode string    `json:"workshop_section_code,omitempty"`
}
type OapiRhinoHumanresProductionteamQueryResponse struct {
	taobao.TaobaoResponse
	Errcode         int64   `json:"errcode,omitempty"`
	Errmsg          string  `json:"errmsg,omitempty"`
	ExternalMsgInfo string  `json:"external_msg_info,omitempty"`
	Model           []Model `json:"model,omitempty"`
	StatusCode      int64   `json:"status_code,omitempty"`
	Success         bool    `json:"success,omitempty"`
}
type EmpList struct {
	BuMail         string    `json:"bu_mail,omitempty"`
	BucId          int64     `json:"buc_id,omitempty"`
	DepartureDate  time.Time `json:"departure_date,omitempty"`
	DingtalkNo     string    `json:"dingtalk_no,omitempty"`
	DingtalkUserId string    `json:"dingtalk_user_id,omitempty"`
	HireDate       time.Time `json:"hire_date,omitempty"`
	ImgUrl         string    `json:"img_url,omitempty"`
	Job            string    `json:"job,omitempty"`
	Name           string    `json:"name,omitempty"`
	Status         int64     `json:"status,omitempty"`
	UicId          int64     `json:"uic_id,omitempty"`
	WorkNo         string    `json:"work_no,omitempty"`
	WorkStatus     int64     `json:"work_status,omitempty"`
}
type WeekdayConfigList struct {
	Day                int64     `json:"day,omitempty"`
	EndTime            time.Time `json:"end_time,omitempty"`
	ProductionTeamCode string    `json:"production_team_code,omitempty"`
	StartTime          time.Time `json:"start_time,omitempty"`
	Type               int64     `json:"type,omitempty"`
	WeekDay            string    `json:"week_day,omitempty"`
}
