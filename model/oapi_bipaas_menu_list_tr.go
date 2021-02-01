package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiBipaasMenuListTreeRequest() *OapiBipaasMenuListTreeRequest {
	return &OapiBipaasMenuListTreeRequest{}
}

type OapiBipaasMenuListTreeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp             OapiBipaasMenuListTreeResponse
	AntcloudTenantId string
	Published        bool
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiBipaasMenuListTreeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBipaasMenuListTreeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBipaasMenuListTreeRequest) SetAntcloudTenantId(antcloudTenantId2 string) {
	this.AntcloudTenantId = antcloudTenantId2
}
func (this *OapiBipaasMenuListTreeRequest) GetAntcloudTenantId() string {
	return this.AntcloudTenantId
}
func (this *OapiBipaasMenuListTreeRequest) SetPublished(published2 bool) {
	this.Published = published2
}
func (this *OapiBipaasMenuListTreeRequest) GetPublished() bool {
	return this.Published
}
func (this *OapiBipaasMenuListTreeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.bipaas.menu.list_tree"
}
func (this *OapiBipaasMenuListTreeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBipaasMenuListTreeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBipaasMenuListTreeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBipaasMenuListTreeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBipaasMenuListTreeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["antcloud_tenant_id"] = this.AntcloudTenantId
	txtParams["published"] = this.Published
	return txtParams
}
func (this *OapiBipaasMenuListTreeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AntcloudTenantId, "antcloudTenantId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiBipaasMenuListTreeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBipaasMenuListTreeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiBipaasMenuListTreeResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
