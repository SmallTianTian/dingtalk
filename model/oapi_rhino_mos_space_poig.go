package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpacePoiGetRequest() *OapiRhinoMosSpacePoiGetRequest {
	return &OapiRhinoMosSpacePoiGetRequest{}
}

type OapiRhinoMosSpacePoiGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpacePoiGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpacePoiGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpacePoiGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpacePoiGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpacePoiGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpacePoiGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.poi.get"
}
func (this *OapiRhinoMosSpacePoiGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpacePoiGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpacePoiGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpacePoiGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpacePoiGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpacePoiGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpacePoiGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpacePoiGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SpacePoiConditionReq struct {
	CategoryCode    string `json:"category_code,omitempty"`
	CategorySubCode string `json:"category_sub_code,omitempty"`
	PoiCode         string `json:"poi_code,omitempty"`
	TenantId        string `json:"tenant_id,omitempty"`
	Userid          string `json:"userid,omitempty"`
}
type OapiRhinoMosSpacePoiGetResponse struct {
	taobao.TaobaoResponse
	Model SpacePoiDto `json:"model,omitempty"`
}
type SpacePoiDto struct {
	CategoryCode    string `json:"category_code,omitempty"`
	CategorySubCode string `json:"category_sub_code,omitempty"`
	PoiCode         string `json:"poi_code,omitempty"`
	PoiName         string `json:"poi_name,omitempty"`
	TenantId        string `json:"tenant_id,omitempty"`
}
