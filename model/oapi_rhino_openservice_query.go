package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoOpenserviceQueryRequest() *OapiRhinoOpenserviceQueryRequest {
	return &OapiRhinoOpenserviceQueryRequest{}
}

type OapiRhinoOpenserviceQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoOpenserviceQueryResponse
	Code            string
	GmtCreate       time.Time
	Id              int64
	TenentId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoOpenserviceQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoOpenserviceQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoOpenserviceQueryRequest) SetCode(code2 string) {
	this.Code = code2
}
func (this *OapiRhinoOpenserviceQueryRequest) GetCode() string {
	return this.Code
}
func (this *OapiRhinoOpenserviceQueryRequest) SetGmtCreate(gmtCreate2 time.Time) {
	this.GmtCreate = gmtCreate2
}
func (this *OapiRhinoOpenserviceQueryRequest) GetGmtCreate() time.Time {
	return this.GmtCreate
}
func (this *OapiRhinoOpenserviceQueryRequest) SetId(id2 int64) {
	this.Id = id2
}
func (this *OapiRhinoOpenserviceQueryRequest) GetId() int64 {
	return this.Id
}
func (this *OapiRhinoOpenserviceQueryRequest) SetTenentId(tenentId2 string) {
	this.TenentId = tenentId2
}
func (this *OapiRhinoOpenserviceQueryRequest) GetTenentId() string {
	return this.TenentId
}
func (this *OapiRhinoOpenserviceQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoOpenserviceQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoOpenserviceQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.openservice.query"
}
func (this *OapiRhinoOpenserviceQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoOpenserviceQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoOpenserviceQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoOpenserviceQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoOpenserviceQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.ERROR_CODE] = this.Code
	txtParams["gmt_create"] = this.GmtCreate
	txtParams["id"] = this.Id
	txtParams["tenent_id"] = this.TenentId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoOpenserviceQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoOpenserviceQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoOpenserviceQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoOpenserviceQueryResponse struct {
	taobao.TaobaoResponse
	Result Result `json:"result,omitempty"`
}
