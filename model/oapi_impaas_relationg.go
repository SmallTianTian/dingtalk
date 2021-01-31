package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasRelationGetRequest() *OapiImpaasRelationGetRequest {
	return &OapiImpaasRelationGetRequest{}
}

type OapiImpaasRelationGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasRelationGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasRelationGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasRelationGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasRelationGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasRelationGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasRelationGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.relation.get"
}
func (this *OapiImpaasRelationGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasRelationGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasRelationGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasRelationGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasRelationGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasRelationGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasRelationGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasRelationGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GetRelationReq struct {
	DstImOpenids []string `json:"dst_im_openids,omitempty"`
	SrcImOpenid  string   `json:"src_im_openid,omitempty"`
}
type OapiImpaasRelationGetResponse struct {
	taobao.TaobaoResponse
	Result  []RelationModel `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
type RelationModel struct {
	BeginTime   int64  `json:"begin_time,omitempty"`
	DstImOpenid string `json:"dst_im_openid,omitempty"`
	EndTime     int64  `json:"end_time,omitempty"`
	SrcImOpenid string `json:"src_im_openid,omitempty"`
}
