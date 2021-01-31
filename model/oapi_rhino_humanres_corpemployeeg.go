package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoHumanresCorpemployeeGetRequest() *OapiRhinoHumanresCorpemployeeGetRequest {
	return &OapiRhinoHumanresCorpemployeeGetRequest{}
}

type OapiRhinoHumanresCorpemployeeGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiRhinoHumanresCorpemployeeGetResponse
	QueryCorpEmployeeParam string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiRhinoHumanresCorpemployeeGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) SetQueryCorpEmployeeParam(queryCorpEmployeeParam2 string) {
	this.QueryCorpEmployeeParam = queryCorpEmployeeParam2
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) GetQueryCorpEmployeeParam() string {
	return this.QueryCorpEmployeeParam
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.humanres.corpemployee.get"
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["query_corp_employee_param"] = this.QueryCorpEmployeeParam
	return txtParams
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoHumanresCorpemployeeGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type QueryCorpEmployeeDto struct {
	TenantId string `json:"tenant_id,omitempty"`
	Userid   string `json:"userid,omitempty"`
	WorkNo   string `json:"work_no,omitempty"`
}
type OapiRhinoHumanresCorpemployeeGetResponse struct {
	taobao.TaobaoResponse
	CorpEmployeeInfo Model  `json:"corp_employee_info,omitempty"`
	Errcode          int64  `json:"errcode,omitempty"`
	Errmsg           string `json:"errmsg,omitempty"`
	ExternalMsgInfo  string `json:"external_msg_info,omitempty"`
	StatusCode       int64  `json:"status_code,omitempty"`
	Success          bool   `json:"success,omitempty"`
}
type Model struct {
	BuMail         string    `json:"bu_mail,omitempty"`
	BucId          int64     `json:"buc_id,omitempty"`
	CorpId         string    `json:"corp_id,omitempty"`
	DepartureDate  time.Time `json:"departure_date,omitempty"`
	DingtalkNo     string    `json:"dingtalk_no,omitempty"`
	DingtalkUserId string    `json:"dingtalk_user_id,omitempty"`
	HireDate       time.Time `json:"hire_date,omitempty"`
	Id             int64     `json:"id,omitempty"`
	ImgUrl         string    `json:"img_url,omitempty"`
	Job            string    `json:"job,omitempty"`
	Mobile         string    `json:"mobile,omitempty"`
	Name           string    `json:"name,omitempty"`
	Status         int64     `json:"status,omitempty"`
	UicId          int64     `json:"uic_id,omitempty"`
	WorkNo         string    `json:"work_no,omitempty"`
	WorkStatus     int64     `json:"work_status,omitempty"`
}
