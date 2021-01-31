package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoDeviceUniquecodeGetRequest() *OapiRhinoDeviceUniquecodeGetRequest {
	return &OapiRhinoDeviceUniquecodeGetRequest{}
}

type OapiRhinoDeviceUniquecodeGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoDeviceUniquecodeGetResponse
	TopHttpMethod   string
	TopResponseType string
	UniqueCode      string
	Userid          string
}

func (this *OapiRhinoDeviceUniquecodeGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) SetUniqueCode(uniqueCode2 string) {
	this.UniqueCode = uniqueCode2
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) GetUniqueCode() string {
	return this.UniqueCode
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.device.uniquecode.get"
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["unique_code"] = this.UniqueCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UniqueCode, "uniqueCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoDeviceUniquecodeGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoDeviceUniquecodeGetResponse struct {
	taobao.TaobaoResponse
	Errcode         int64              `json:"errcode,omitempty"`
	Errmsg          string             `json:"errmsg,omitempty"`
	ExternalMsgInfo string             `json:"external_msg_info,omitempty"`
	Model           MosDeviceTenantDto `json:"model,omitempty"`
	Success         bool               `json:"success,omitempty"`
}
type MosDeviceTenantDto struct {
	Attributes      string `json:"attributes,omitempty"`
	Deleted         bool   `json:"deleted,omitempty"`
	GlobalModelId   string `json:"global_model_id,omitempty"`
	Id              int64  `json:"id,omitempty"`
	InstanceId      string `json:"instance_id,omitempty"`
	ModelId         int64  `json:"model_id,omitempty"`
	SupplierModelId int64  `json:"supplier_model_id,omitempty"`
	SupplierSn      string `json:"supplier_sn,omitempty"`
	TenantId        string `json:"tenant_id,omitempty"`
	Version         string `json:"version,omitempty"`
}
