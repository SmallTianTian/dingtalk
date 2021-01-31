package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiBlackboardDeleteRequest() *OapiBlackboardDeleteRequest {
	return &OapiBlackboardDeleteRequest{}
}

type OapiBlackboardDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBlackboardDeleteResponse
	BlackboardId    string
	OperationUserid string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiBlackboardDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBlackboardDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBlackboardDeleteRequest) SetBlackboardId(blackboardId2 string) {
	this.BlackboardId = blackboardId2
}
func (this *OapiBlackboardDeleteRequest) GetBlackboardId() string {
	return this.BlackboardId
}
func (this *OapiBlackboardDeleteRequest) SetOperationUserid(operationUserid2 string) {
	this.OperationUserid = operationUserid2
}
func (this *OapiBlackboardDeleteRequest) GetOperationUserid() string {
	return this.OperationUserid
}
func (this *OapiBlackboardDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.blackboard.delete"
}
func (this *OapiBlackboardDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBlackboardDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBlackboardDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBlackboardDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBlackboardDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["blackboard_id"] = this.BlackboardId
	txtParams["operation_userid"] = this.OperationUserid
	return txtParams
}
func (this *OapiBlackboardDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BlackboardId, "blackboardId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiBlackboardDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBlackboardDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiBlackboardDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
