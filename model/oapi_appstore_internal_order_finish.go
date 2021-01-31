package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAppstoreInternalOrderFinishRequest() *OapiAppstoreInternalOrderFinishRequest {
	return &OapiAppstoreInternalOrderFinishRequest{}
}

type OapiAppstoreInternalOrderFinishRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAppstoreInternalOrderFinishResponse
	BizOrderId      int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAppstoreInternalOrderFinishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreInternalOrderFinishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreInternalOrderFinishRequest) SetBizOrderId(bizOrderId2 int64) {
	this.BizOrderId = bizOrderId2
}
func (this *OapiAppstoreInternalOrderFinishRequest) GetBizOrderId() int64 {
	return this.BizOrderId
}
func (this *OapiAppstoreInternalOrderFinishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.internal.order.finish"
}
func (this *OapiAppstoreInternalOrderFinishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreInternalOrderFinishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreInternalOrderFinishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreInternalOrderFinishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreInternalOrderFinishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_order_id"] = this.BizOrderId
	return txtParams
}
func (this *OapiAppstoreInternalOrderFinishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizOrderId, "bizOrderId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAppstoreInternalOrderFinishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreInternalOrderFinishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreInternalOrderFinishResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
