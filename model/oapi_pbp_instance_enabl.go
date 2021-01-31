package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPbpInstanceEnableRequest() *OapiPbpInstanceEnableRequest {
	return &OapiPbpInstanceEnableRequest{}
}

type OapiPbpInstanceEnableRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpInstanceEnableResponse
	BizId           string
	BizInstId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpInstanceEnableRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpInstanceEnableRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpInstanceEnableRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiPbpInstanceEnableRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiPbpInstanceEnableRequest) SetBizInstId(bizInstId2 string) {
	this.BizInstId = bizInstId2
}
func (this *OapiPbpInstanceEnableRequest) GetBizInstId() string {
	return this.BizInstId
}
func (this *OapiPbpInstanceEnableRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.instance.enable"
}
func (this *OapiPbpInstanceEnableRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpInstanceEnableRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpInstanceEnableRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpInstanceEnableRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpInstanceEnableRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["biz_inst_id"] = this.BizInstId
	return txtParams
}
func (this *OapiPbpInstanceEnableRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizId, "bizId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPbpInstanceEnableRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpInstanceEnableRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPbpInstanceEnableResponse struct {
	taobao.TaobaoResponse
}
