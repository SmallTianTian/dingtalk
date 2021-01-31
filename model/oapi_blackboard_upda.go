package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiBlackboardUpdateRequest() *OapiBlackboardUpdateRequest {
	return &OapiBlackboardUpdateRequest{}
}

type OapiBlackboardUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBlackboardUpdateResponse
	TopHttpMethod   string
	TopResponseType string
	UpdateRequest   string
}

func (this *OapiBlackboardUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBlackboardUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBlackboardUpdateRequest) SetUpdateRequest(updateRequest2 string) {
	this.UpdateRequest = updateRequest2
}
func (this *OapiBlackboardUpdateRequest) GetUpdateRequest() string {
	return this.UpdateRequest
}
func (this *OapiBlackboardUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.blackboard.update"
}
func (this *OapiBlackboardUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBlackboardUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBlackboardUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBlackboardUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBlackboardUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["update_request"] = this.UpdateRequest
	return txtParams
}
func (this *OapiBlackboardUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiBlackboardUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBlackboardUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUpdateBlackboardVo struct {
	Author          string `json:"author,omitempty"`
	BlackboardId    string `json:"blackboard_id,omitempty"`
	CategoryId      string `json:"category_id,omitempty"`
	Content         string `json:"content,omitempty"`
	CoverpicMediaid string `json:"coverpic_mediaid,omitempty"`
	Ding            bool   `json:"ding,omitempty"`
	Notify          bool   `json:"notify,omitempty"`
	OperationUserid string `json:"operation_userid,omitempty"`
	Title           string `json:"title,omitempty"`
}
type OapiBlackboardUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
