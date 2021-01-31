package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiAttendanceAdvancedServiceIsboundRequest() *OapiAttendanceAdvancedServiceIsboundRequest {
	return &OapiAttendanceAdvancedServiceIsboundRequest{}
}

type OapiAttendanceAdvancedServiceIsboundRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceAdvancedServiceIsboundResponse
	OpUserid        string
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceAdvancedServiceIsboundRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) GetParam() string {
	return this.Param
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.advanced.service.isbound"
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_userid"] = this.OpUserid
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserid, "opUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceAdvancedServiceIsboundRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type IsBoundParam struct {
	EntityIds  []string `json:"entity_ids,omitempty"`
	EntityType string   `json:"entity_type,omitempty"`
	ServiceId  int64    `json:"service_id,omitempty"`
}
type OapiAttendanceAdvancedServiceIsboundResponse struct {
	taobao.TaobaoResponse
	Errcode int64                       `json:"errcode,omitempty"`
	Errmsg  string                      `json:"errmsg,omitempty"`
	Result  []AdvancedServiceInstanceVo `json:"result,omitempty"`
	Success bool                        `json:"success,omitempty"`
}
type AdvancedServiceInstanceVo struct {
	CorpId     string    `json:"corp_id,omitempty"`
	EndTime    time.Time `json:"end_time,omitempty"`
	EntityId   string    `json:"entity_id,omitempty"`
	EntityType string    `json:"entity_type,omitempty"`
	Id         int64     `json:"id,omitempty"`
	IsDeleted  string    `json:"is_deleted,omitempty"`
	Name       string    `json:"name,omitempty"`
	ServiceId  int64     `json:"service_id,omitempty"`
	StartTime  time.Time `json:"start_time,omitempty"`
}
