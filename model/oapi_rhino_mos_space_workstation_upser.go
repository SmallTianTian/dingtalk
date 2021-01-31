package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceWorkstationUpsertRequest() *OapiRhinoMosSpaceWorkstationUpsertRequest {
	return &OapiRhinoMosSpaceWorkstationUpsertRequest{}
}

type OapiRhinoMosSpaceWorkstationUpsertRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceWorkstationUpsertResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.workstation.upsert"
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceWorkstationUpsertRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProdWorkstationUpsertDto struct {
	CategoryCode        string `json:"category_code,omitempty"`
	CategorySubCode     string `json:"category_sub_code,omitempty"`
	PoiCode             string `json:"poi_code,omitempty"`
	ProdWorkstationCode string `json:"prod_workstation_code,omitempty"`
	ProdWorkstationName string `json:"prod_workstation_name,omitempty"`
	TenantId            string `json:"tenant_id,omitempty"`
	Userid              string `json:"userid,omitempty"`
}
type OapiRhinoMosSpaceWorkstationUpsertResponse struct {
	taobao.TaobaoResponse
	Errcode int64              `json:"errcode,omitempty"`
	Errmsg  string             `json:"errmsg,omitempty"`
	Model   ProdWorkstationDto `json:"model,omitempty"`
}
