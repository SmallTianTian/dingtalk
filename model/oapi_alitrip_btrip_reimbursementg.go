package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAlitripBtripReimbursementGetRequest() *OapiAlitripBtripReimbursementGetRequest {
	return &OapiAlitripBtripReimbursementGetRequest{}
}

type OapiAlitripBtripReimbursementGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiAlitripBtripReimbursementGetResponse
	Corpid           string
	ThirdpartyFlowId string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiAlitripBtripReimbursementGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripReimbursementGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripReimbursementGetRequest) SetCorpid(corpid2 string) {
	this.Corpid = corpid2
}
func (this *OapiAlitripBtripReimbursementGetRequest) GetCorpid() string {
	return this.Corpid
}
func (this *OapiAlitripBtripReimbursementGetRequest) SetThirdpartyFlowId(thirdpartyFlowId2 string) {
	this.ThirdpartyFlowId = thirdpartyFlowId2
}
func (this *OapiAlitripBtripReimbursementGetRequest) GetThirdpartyFlowId() string {
	return this.ThirdpartyFlowId
}
func (this *OapiAlitripBtripReimbursementGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.reimbursement.get"
}
func (this *OapiAlitripBtripReimbursementGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripReimbursementGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripReimbursementGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripReimbursementGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripReimbursementGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corpid"] = this.Corpid
	txtParams["thirdparty_flow_id"] = this.ThirdpartyFlowId
	return txtParams
}
func (this *OapiAlitripBtripReimbursementGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Corpid, "corpid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAlitripBtripReimbursementGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripReimbursementGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripReimbursementGetResponse struct {
	taobao.TaobaoResponse
	Module  int64 `json:"module,omitempty"`
	Success bool  `json:"success,omitempty"`
}
