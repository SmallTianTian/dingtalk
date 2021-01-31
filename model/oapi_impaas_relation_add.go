package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasRelationAddRequest() *OapiImpaasRelationAddRequest {
	return &OapiImpaasRelationAddRequest{}
}

type OapiImpaasRelationAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasRelationAddResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasRelationAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasRelationAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasRelationAddRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasRelationAddRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasRelationAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.relation.add"
}
func (this *OapiImpaasRelationAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasRelationAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasRelationAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasRelationAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasRelationAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasRelationAddRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasRelationAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasRelationAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AddRelationReq struct {
	BeginTime   int64  `json:"begin_time,omitempty"`
	DstImOpenid string `json:"dst_im_openid,omitempty"`
	EndTime     int64  `json:"end_time,omitempty"`
	IsDoubleWay bool   `json:"is_double_way,omitempty"`
	SrcImOpenid string `json:"src_im_openid,omitempty"`
}
type OapiImpaasRelationAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
