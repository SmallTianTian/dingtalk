package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceWorkstationGetRequest() *OapiRhinoMosSpaceWorkstationGetRequest {
	return &OapiRhinoMosSpaceWorkstationGetRequest{}
}

type OapiRhinoMosSpaceWorkstationGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceWorkstationGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceWorkstationGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.workstation.get"
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceWorkstationGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProdWorkstationConditionReq struct {
	CategoryCode        string `json:"category_code,omitempty"`
	CategorySubCode     string `json:"category_sub_code,omitempty"`
	PoiCode             string `json:"poi_code,omitempty"`
	ProdWorkstationCode string `json:"prod_workstation_code,omitempty"`
	TenantId            string `json:"tenant_id,omitempty"`
	Userid              string `json:"userid,omitempty"`
}
type OapiRhinoMosSpaceWorkstationGetResponse struct {
	taobao.TaobaoResponse
	Model ProdWorkstationDto `json:"model,omitempty"`
}
type ProdWorkstationDto struct {
	CategoryCode        string `json:"category_code,omitempty"`
	CategorySubCode     string `json:"category_sub_code,omitempty"`
	PoiCode             string `json:"poi_code,omitempty"`
	ProdWorkstationCode string `json:"prod_workstation_code,omitempty"`
	ProdWorkstationName string `json:"prod_workstation_name,omitempty"`
	TenantId            string `json:"tenant_id,omitempty"`
}
