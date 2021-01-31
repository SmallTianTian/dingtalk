package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasRelationDelRequest() *OapiImpaasRelationDelRequest {
	return &OapiImpaasRelationDelRequest{}
}

type OapiImpaasRelationDelRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasRelationDelResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasRelationDelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasRelationDelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasRelationDelRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasRelationDelRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasRelationDelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.relation.del"
}
func (this *OapiImpaasRelationDelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasRelationDelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasRelationDelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasRelationDelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasRelationDelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasRelationDelRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasRelationDelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasRelationDelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DelRelationReq struct {
	DstImOpenid string `json:"dst_im_openid,omitempty"`
	IsDoubleWay bool   `json:"is_double_way,omitempty"`
	SrcImOpenid string `json:"src_im_openid,omitempty"`
}
type OapiImpaasRelationDelResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
