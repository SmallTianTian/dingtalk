package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecTrackBindRequest() *OapiRhinoMosExecTrackBindRequest {
	return &OapiRhinoMosExecTrackBindRequest{}
}

type OapiRhinoMosExecTrackBindRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecTrackBindResponse
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

func (this *OapiRhinoMosExecTrackBindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecTrackBindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecTrackBindRequest) SetEntities(entities2 string) {
	this.Entities = entities2
}
func (this *OapiRhinoMosExecTrackBindRequest) GetEntities() string {
	return this.Entities
}
func (this *OapiRhinoMosExecTrackBindRequest) SetEntityType(entityType2 string) {
	this.EntityType = entityType2
}
func (this *OapiRhinoMosExecTrackBindRequest) GetEntityType() string {
	return this.EntityType
}
func (this *OapiRhinoMosExecTrackBindRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecTrackBindRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecTrackBindRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecTrackBindRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecTrackBindRequest) SetTrackId(trackId2 string) {
	this.TrackId = trackId2
}
func (this *OapiRhinoMosExecTrackBindRequest) GetTrackId() string {
	return this.TrackId
}
func (this *OapiRhinoMosExecTrackBindRequest) SetTrackType(trackType2 string) {
	this.TrackType = trackType2
}
func (this *OapiRhinoMosExecTrackBindRequest) GetTrackType() string {
	return this.TrackType
}
func (this *OapiRhinoMosExecTrackBindRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecTrackBindRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecTrackBindRequest) SetWorkstationCode(workstationCode2 string) {
	this.WorkstationCode = workstationCode2
}
func (this *OapiRhinoMosExecTrackBindRequest) GetWorkstationCode() string {
	return this.WorkstationCode
}
func (this *OapiRhinoMosExecTrackBindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.track.bind"
}
func (this *OapiRhinoMosExecTrackBindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecTrackBindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecTrackBindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecTrackBindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecTrackBindRequest) GetTextParams() map[string]interface{} {
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
func (this *OapiRhinoMosExecTrackBindRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Entities, "entities"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecTrackBindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecTrackBindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecTrackBindResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Model   bool   `json:"model,omitempty"`
	Success bool   `json:"success,omitempty"`
}
