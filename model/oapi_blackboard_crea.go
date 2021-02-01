package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiBlackboardCreateRequest() *OapiBlackboardCreateRequest {
	return &OapiBlackboardCreateRequest{}
}

type OapiBlackboardCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBlackboardCreateResponse
	CreateRequest   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiBlackboardCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBlackboardCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBlackboardCreateRequest) SetCreateRequest(createRequest2 string) {
	this.CreateRequest = createRequest2
}
func (this *OapiBlackboardCreateRequest) GetCreateRequest() string {
	return this.CreateRequest
}
func (this *OapiBlackboardCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.blackboard.create"
}
func (this *OapiBlackboardCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBlackboardCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBlackboardCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBlackboardCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBlackboardCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["create_request"] = this.CreateRequest
	return txtParams
}
func (this *OapiBlackboardCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiBlackboardCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBlackboardCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BlackboardReceiverOpenVo struct {
	DeptidList []int64  `json:"deptid_list,omitempty"`
	UseridList []string `json:"userid_list,omitempty"`
}
type OapiCreateBlackboardVo struct {
	Author             string                   `json:"author,omitempty"`
	BlackboardReceiver BlackboardReceiverOpenVo `json:"blackboard_receiver,omitempty"`
	CategoryId         string                   `json:"category_id,omitempty"`
	Content            string                   `json:"content,omitempty"`
	CoverpicMediaid    string                   `json:"coverpic_mediaid,omitempty"`
	Ding               bool                     `json:"ding,omitempty"`
	OperationUserid    string                   `json:"operation_userid,omitempty"`
	PrivateLevel       int64                    `json:"private_level,omitempty"`
	PushTop            bool                     `json:"push_top,omitempty"`
	Title              string                   `json:"title,omitempty"`
}
type OapiBlackboardCreateResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
