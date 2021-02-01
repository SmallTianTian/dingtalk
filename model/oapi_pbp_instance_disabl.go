package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPbpInstanceDisableRequest() *OapiPbpInstanceDisableRequest {
	return &OapiPbpInstanceDisableRequest{}
}

type OapiPbpInstanceDisableRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpInstanceDisableResponse
	BizId           string
	BizInstId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpInstanceDisableRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpInstanceDisableRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpInstanceDisableRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiPbpInstanceDisableRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiPbpInstanceDisableRequest) SetBizInstId(bizInstId2 string) {
	this.BizInstId = bizInstId2
}
func (this *OapiPbpInstanceDisableRequest) GetBizInstId() string {
	return this.BizInstId
}
func (this *OapiPbpInstanceDisableRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.instance.disable"
}
func (this *OapiPbpInstanceDisableRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpInstanceDisableRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpInstanceDisableRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpInstanceDisableRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpInstanceDisableRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["biz_inst_id"] = this.BizInstId
	return txtParams
}
func (this *OapiPbpInstanceDisableRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizInstId, "bizInstId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPbpInstanceDisableRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpInstanceDisableRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPbpInstanceDisableResponse struct {
	taobao.TaobaoResponse
}
