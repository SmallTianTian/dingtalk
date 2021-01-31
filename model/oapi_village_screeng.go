package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiVillageScreenGetRequest() *OapiVillageScreenGetRequest {
	return &OapiVillageScreenGetRequest{}
}

type OapiVillageScreenGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiVillageScreenGetResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiVillageScreenGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiVillageScreenGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiVillageScreenGetRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiVillageScreenGetRequest) GetReq() string {
	return this.Req
}
func (this *OapiVillageScreenGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.village.screen.get"
}
func (this *OapiVillageScreenGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiVillageScreenGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiVillageScreenGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiVillageScreenGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiVillageScreenGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiVillageScreenGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiVillageScreenGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiVillageScreenGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DataVScreenRequest struct {
	Params string `json:"params,omitempty"`
}
type OapiVillageScreenGetResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
