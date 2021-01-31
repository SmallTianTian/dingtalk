package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkrecordGetbyuseridRequest() *OapiWorkrecordGetbyuseridRequest {
	return &OapiWorkrecordGetbyuseridRequest{}
}

type OapiWorkrecordGetbyuseridRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkrecordGetbyuseridResponse
	Limit           int64
	Offset          int64
	Status          int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiWorkrecordGetbyuseridRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkrecordGetbyuseridRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkrecordGetbyuseridRequest) SetLimit(limit2 int64) {
	this.Limit = limit2
}
func (this *OapiWorkrecordGetbyuseridRequest) GetLimit() int64 {
	return this.Limit
}
func (this *OapiWorkrecordGetbyuseridRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiWorkrecordGetbyuseridRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiWorkrecordGetbyuseridRequest) SetStatus(status2 int64) {
	this.Status = status2
}
func (this *OapiWorkrecordGetbyuseridRequest) GetStatus() int64 {
	return this.Status
}
func (this *OapiWorkrecordGetbyuseridRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiWorkrecordGetbyuseridRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiWorkrecordGetbyuseridRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workrecord.getbyuserid"
}
func (this *OapiWorkrecordGetbyuseridRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkrecordGetbyuseridRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkrecordGetbyuseridRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkrecordGetbyuseridRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkrecordGetbyuseridRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["limit"] = this.Limit
	txtParams["offset"] = this.Offset
	txtParams["status"] = this.Status
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiWorkrecordGetbyuseridRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Limit, "limit"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkrecordGetbyuseridRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkrecordGetbyuseridRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkrecordGetbyuseridResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Records PageResult `json:"records,omitempty"`
}
