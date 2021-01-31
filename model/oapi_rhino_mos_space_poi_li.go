package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpacePoiListRequest() *OapiRhinoMosSpacePoiListRequest {
	return &OapiRhinoMosSpacePoiListRequest{}
}

type OapiRhinoMosSpacePoiListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpacePoiListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpacePoiListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpacePoiListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpacePoiListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpacePoiListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpacePoiListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.poi.list"
}
func (this *OapiRhinoMosSpacePoiListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpacePoiListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpacePoiListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpacePoiListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpacePoiListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpacePoiListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpacePoiListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpacePoiListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SpacePoiConditionListReq struct {
	PoiCodeList []string `json:"poi_code_list,omitempty"`
	TenantId    string   `json:"tenant_id,omitempty"`
	Userid      string   `json:"userid,omitempty"`
}
type OapiRhinoMosSpacePoiListResponse struct {
	taobao.TaobaoResponse
	Model []SpacePoiDto `json:"model,omitempty"`
}
