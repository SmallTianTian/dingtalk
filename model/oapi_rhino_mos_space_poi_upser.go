package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpacePoiUpsertRequest() *OapiRhinoMosSpacePoiUpsertRequest {
	return &OapiRhinoMosSpacePoiUpsertRequest{}
}

type OapiRhinoMosSpacePoiUpsertRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpacePoiUpsertResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpacePoiUpsertRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.poi.upsert"
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpacePoiUpsertRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SpacePoiUpsertReq struct {
	CategoryCode    string `json:"category_code,omitempty"`
	CategorySubCode string `json:"category_sub_code,omitempty"`
	PoiCode         string `json:"poi_code,omitempty"`
	PoiName         string `json:"poi_name,omitempty"`
	TenantId        string `json:"tenant_id,omitempty"`
	Userid          string `json:"userid,omitempty"`
}
type OapiRhinoMosSpacePoiUpsertResponse struct {
	taobao.TaobaoResponse
	Errcode int64       `json:"errcode,omitempty"`
	Errmsg  string      `json:"errmsg,omitempty"`
	Model   SpacePoiDto `json:"model,omitempty"`
}
