package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoTransportMaplocationQueryRequest() *OapiRhinoTransportMaplocationQueryRequest {
	return &OapiRhinoTransportMaplocationQueryRequest{}
}

type OapiRhinoTransportMaplocationQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoTransportMaplocationQueryResponse
	IncludeConfig   bool
	PoiCodeList     string
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoTransportMaplocationQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoTransportMaplocationQueryRequest) SetIncludeConfig(includeConfig2 bool) {
	this.IncludeConfig = includeConfig2
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetIncludeConfig() bool {
	return this.IncludeConfig
}
func (this *OapiRhinoTransportMaplocationQueryRequest) SetPoiCodeList(poiCodeList2 string) {
	this.PoiCodeList = poiCodeList2
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetPoiCodeList() string {
	return this.PoiCodeList
}
func (this *OapiRhinoTransportMaplocationQueryRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoTransportMaplocationQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.transport.maplocation.query"
}
func (this *OapiRhinoTransportMaplocationQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoTransportMaplocationQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoTransportMaplocationQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["include_config"] = this.IncludeConfig
	txtParams["poi_code_list"] = this.PoiCodeList
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoTransportMaplocationQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.PoiCodeList, "poiCodeList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoTransportMaplocationQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoTransportMaplocationQueryResponse struct {
	taobao.TaobaoResponse
	Errcode         int64                     `json:"errcode,omitempty"`
	Errmsg          string                    `json:"errmsg,omitempty"`
	ExternalMsgInfo string                    `json:"external_msg_info,omitempty"`
	Model           []MapLocationDetailResult `json:"model,omitempty"`
	Success         bool                      `json:"success,omitempty"`
}
type Configs struct {
	LocationId    int64  `json:"location_id,omitempty"`
	TenantId      string `json:"tenant_id,omitempty"`
	TransCmdCode  string `json:"trans_cmd_code,omitempty"`
	TransTypeCode string `json:"trans_type_code,omitempty"`
}
type MapLocationDetailResult struct {
	Configs    []Configs `json:"configs,omitempty"`
	LocationId int64     `json:"location_id,omitempty"`
	OwnPoiCode string    `json:"own_poi_code,omitempty"`
	TenantId   string    `json:"tenant_id,omitempty"`
}
