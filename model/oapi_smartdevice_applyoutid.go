package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceApplyoutidRequest() *OapiSmartdeviceApplyoutidRequest {
	return &OapiSmartdeviceApplyoutidRequest{}
}

type OapiSmartdeviceApplyoutidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceApplyoutidResponse
	DevServId       int64
	Sn              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceApplyoutidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceApplyoutidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceApplyoutidRequest) SetDevServId(devServId2 int64) {
	this.DevServId = devServId2
}
func (this *OapiSmartdeviceApplyoutidRequest) GetDevServId() int64 {
	return this.DevServId
}
func (this *OapiSmartdeviceApplyoutidRequest) SetSn(sn2 string) {
	this.Sn = sn2
}
func (this *OapiSmartdeviceApplyoutidRequest) GetSn() string {
	return this.Sn
}
func (this *OapiSmartdeviceApplyoutidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.applyoutid"
}
func (this *OapiSmartdeviceApplyoutidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceApplyoutidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceApplyoutidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceApplyoutidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceApplyoutidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dev_serv_id"] = this.DevServId
	txtParams["sn"] = this.Sn
	return txtParams
}
func (this *OapiSmartdeviceApplyoutidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DevServId, "devServId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceApplyoutidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceApplyoutidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceApplyoutidResponse struct {
	taobao.TaobaoResponse
	Result  OutDeviceIdVo `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
type OutDeviceIdVo struct {
	Outdid string `json:"outdid,omitempty"`
}
