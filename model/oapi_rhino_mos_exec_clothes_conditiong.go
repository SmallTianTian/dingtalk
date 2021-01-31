package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecClothesConditionGetRequest() *OapiRhinoMosExecClothesConditionGetRequest {
	return &OapiRhinoMosExecClothesConditionGetRequest{}
}

type OapiRhinoMosExecClothesConditionGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiRhinoMosExecClothesConditionGetResponse
	GetClothesConditionReq string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiRhinoMosExecClothesConditionGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) SetGetClothesConditionReq(getClothesConditionReq2 string) {
	this.GetClothesConditionReq = getClothesConditionReq2
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) GetGetClothesConditionReq() string {
	return this.GetClothesConditionReq
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.condition.get"
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["get_clothes_condition_req"] = this.GetClothesConditionReq
	return txtParams
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesConditionGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Source struct {
	SourceId   string `json:"source_id,omitempty"`
	SourceType string `json:"source_type,omitempty"`
}
type EntityCondition struct {
	EntityIds    []int64 `json:"entity_ids,omitempty"`
	EntitySource Source  `json:"entity_source,omitempty"`
	EntityType   string  `json:"entity_type,omitempty"`
}
type Page struct {
	PageSize int64 `json:"page_size,omitempty"`
	Start    int64 `json:"start,omitempty"`
}
type GetClothesConditionReq struct {
	Condition EntityCondition `json:"condition,omitempty"`
	OrderId   int64           `json:"order_id,omitempty"`
	Page      Page            `json:"page,omitempty"`
	SizeCodes []string        `json:"size_codes,omitempty"`
	Status    []string        `json:"status,omitempty"`
	TenantId  string          `json:"tenant_id,omitempty"`
	Userid    string          `json:"userid,omitempty"`
}
type OapiRhinoMosExecClothesConditionGetResponse struct {
	taobao.TaobaoResponse
	Errcode         int64      `json:"errcode,omitempty"`
	Errmsg          string     `json:"errmsg,omitempty"`
	ExternalMsgInfo string     `json:"external_msg_info,omitempty"`
	Model           PageResult `json:"model,omitempty"`
}
type ClothesDto struct {
	ColorId    int64  `json:"color_id,omitempty"`
	ColorName  string `json:"color_name,omitempty"`
	CreateType string `json:"create_type,omitempty"`
	Id         int64  `json:"id,omitempty"`
	OrderId    int64  `json:"order_id,omitempty"`
	SizeCode   string `json:"size_code,omitempty"`
	SizeName   string `json:"size_name,omitempty"`
	SourceId   string `json:"source_id,omitempty"`
	SourceType string `json:"source_type,omitempty"`
	Status     string `json:"status,omitempty"`
	TenantId   string `json:"tenant_id,omitempty"`
}
