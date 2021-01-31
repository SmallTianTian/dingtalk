package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceWorkstationListRequest() *OapiRhinoMosSpaceWorkstationListRequest {
	return &OapiRhinoMosSpaceWorkstationListRequest{}
}

type OapiRhinoMosSpaceWorkstationListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceWorkstationListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceWorkstationListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.workstation.list"
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceWorkstationListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProdWorkstationConditionListReq struct {
	ProdWorkstationCodeList []string `json:"prod_workstation_code_list,omitempty"`
	TenantId                string   `json:"tenant_id,omitempty"`
	Userid                  string   `json:"userid,omitempty"`
}
type OapiRhinoMosSpaceWorkstationListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Model   []ProdWorkstationDto `json:"model,omitempty"`
}
