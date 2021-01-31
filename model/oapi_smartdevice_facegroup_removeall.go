package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacegroupRemoveallRequest() *OapiSmartdeviceFacegroupRemoveallRequest {
	return &OapiSmartdeviceFacegroupRemoveallRequest{}
}

type OapiSmartdeviceFacegroupRemoveallRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacegroupRemoveallResponse
	BizId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceFacegroupRemoveallRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facegroup.removeall"
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	return txtParams
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizId, "bizId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacegroupRemoveallRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacegroupRemoveallResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
