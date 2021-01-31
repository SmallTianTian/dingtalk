package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecTrackTrackconditionListRequest() *OapiRhinoMosExecTrackTrackconditionListRequest {
	return &OapiRhinoMosExecTrackTrackconditionListRequest{}
}

type OapiRhinoMosExecTrackTrackconditionListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecTrackTrackconditionListResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecTrackTrackconditionListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.track.trackcondition.list"
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecTrackTrackconditionListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecTrackTrackconditionListResponse struct {
	taobao.TaobaoResponse
	Model   PageResult `json:"model,omitempty"`
	Success bool       `json:"success,omitempty"`
}
