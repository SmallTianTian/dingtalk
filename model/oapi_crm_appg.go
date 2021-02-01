package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmAppGetRequest() *OapiCrmAppGetRequest {
	return &OapiCrmAppGetRequest{}
}

type OapiCrmAppGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmAppGetResponse
	BizKey          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmAppGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmAppGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmAppGetRequest) SetBizKey(bizKey2 string) {
	this.BizKey = bizKey2
}
func (this *OapiCrmAppGetRequest) GetBizKey() string {
	return this.BizKey
}
func (this *OapiCrmAppGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.app.get"
}
func (this *OapiCrmAppGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmAppGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmAppGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmAppGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmAppGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_key"] = this.BizKey
	return txtParams
}
func (this *OapiCrmAppGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizKey, "bizKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmAppGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmAppGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmAppGetResponse struct {
	taobao.TaobaoResponse
	Result GetCrmMicroAppResponse `json:"result,omitempty"`
}
