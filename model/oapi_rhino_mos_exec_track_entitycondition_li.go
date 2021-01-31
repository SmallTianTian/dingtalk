package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecTrackEntityconditionListRequest() *OapiRhinoMosExecTrackEntityconditionListRequest {
	return &OapiRhinoMosExecTrackEntityconditionListRequest{}
}

type OapiRhinoMosExecTrackEntityconditionListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecTrackEntityconditionListResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecTrackEntityconditionListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.track.entitycondition.list"
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecTrackEntityconditionListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ListTrackRecordWithTrackIdsReq struct {
	EntityIds  []int64  `json:"entity_ids,omitempty"`
	EntityType string   `json:"entity_type,omitempty"`
	Page       Page     `json:"page,omitempty"`
	TenantId   string   `json:"tenant_id,omitempty"`
	TrackTypes []string `json:"track_types,omitempty"`
	Userid     string   `json:"userid,omitempty"`
}
type OapiRhinoMosExecTrackEntityconditionListResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Model   PageResult `json:"model,omitempty"`
	Success bool       `json:"success,omitempty"`
}
type TrackRecordDto struct {
	EffectEndTime              time.Time `json:"effect_end_time,omitempty"`
	EffectStartTime            time.Time `json:"effect_start_time,omitempty"`
	EffectStartWorkstationCode string    `json:"effect_start_workstation_code,omitempty"`
	EffectStatus               string    `json:"effect_status,omitempty"`
	EntityId                   int64     `json:"entity_id,omitempty"`
	EntityType                 string    `json:"entity_type,omitempty"`
	TenantId                   string    `json:"tenant_id,omitempty"`
	TrackId                    string    `json:"track_id,omitempty"`
	TrackType                  string    `json:"track_type,omitempty"`
}
