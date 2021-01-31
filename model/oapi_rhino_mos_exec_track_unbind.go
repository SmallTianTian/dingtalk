package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecTrackUnbindRequest() *OapiRhinoMosExecTrackUnbindRequest {
	return &OapiRhinoMosExecTrackUnbindRequest{}
}

type OapiRhinoMosExecTrackUnbindRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecTrackUnbindResponse
	Entities        string
	EntityType      string
	OrderId         int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	TrackId         string
	TrackType       string
	Userid          string
	WorkstationCode string
}

func (this *OapiRhinoMosExecTrackUnbindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetEntities(entities2 string) {
	this.Entities = entities2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetEntities() string {
	return this.Entities
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetEntityType(entityType2 string) {
	this.EntityType = entityType2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetEntityType() string {
	return this.EntityType
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetTrackId(trackId2 string) {
	this.TrackId = trackId2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetTrackId() string {
	return this.TrackId
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetTrackType(trackType2 string) {
	this.TrackType = trackType2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetTrackType() string {
	return this.TrackType
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetWorkstationCode(workstationCode2 string) {
	this.WorkstationCode = workstationCode2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetWorkstationCode() string {
	return this.WorkstationCode
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.track.unbind"
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecTrackUnbindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["entities"] = this.Entities
	txtParams["entity_type"] = this.EntityType
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams["track_id"] = this.TrackId
	txtParams["track_type"] = this.TrackType
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["workstation_code"] = this.WorkstationCode
	return txtParams
}
func (this *OapiRhinoMosExecTrackUnbindRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Entities, "entities"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecTrackUnbindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecTrackUnbindResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
