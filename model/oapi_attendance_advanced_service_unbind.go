package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceAdvancedServiceUnbindRequest() *OapiAttendanceAdvancedServiceUnbindRequest {
	return &OapiAttendanceAdvancedServiceUnbindRequest{}
}

type OapiAttendanceAdvancedServiceUnbindRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceAdvancedServiceUnbindResponse
	EntityId        string
	EntityType      string
	OpUserid        string
	ServiceId       int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) SetEntityId(entityId2 string) {
	this.EntityId = entityId2
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetEntityId() string {
	return this.EntityId
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) SetEntityType(entityType2 string) {
	this.EntityType = entityType2
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetEntityType() string {
	return this.EntityType
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) SetServiceId(serviceId2 int64) {
	this.ServiceId = serviceId2
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetServiceId() int64 {
	return this.ServiceId
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.advanced.service.unbind"
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["entity_id"] = this.EntityId
	txtParams["entity_type"] = this.EntityType
	txtParams["op_userid"] = this.OpUserid
	txtParams["service_id"] = this.ServiceId
	return txtParams
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EntityId, "entityId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceAdvancedServiceUnbindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceAdvancedServiceUnbindResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
