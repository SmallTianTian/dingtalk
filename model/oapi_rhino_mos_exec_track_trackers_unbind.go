package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecTrackTrackersUnbindRequest() *OapiRhinoMosExecTrackTrackersUnbindRequest {
	return &OapiRhinoMosExecTrackTrackersUnbindRequest{}
}

type OapiRhinoMosExecTrackTrackersUnbindRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecTrackTrackersUnbindResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.track.trackers.unbind"
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecTrackTrackersUnbindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SpecificEntityTypeWithTrackInfoReq struct {
	EntityType string   `json:"entity_type,omitempty"`
	OrderId    int64    `json:"order_id,omitempty"`
	TenantId   string   `json:"tenant_id,omitempty"`
	TrackIds   []string `json:"track_ids,omitempty"`
	TrackType  string   `json:"track_type,omitempty"`
	Userid     string   `json:"userid,omitempty"`
}
type OapiRhinoMosExecTrackTrackersUnbindResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
