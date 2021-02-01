package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripApprovalUpdateRequest() *OapiAlitripBtripApprovalUpdateRequest {
	return &OapiAlitripBtripApprovalUpdateRequest{}
}

type OapiAlitripBtripApprovalUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripApprovalUpdateResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripApprovalUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripApprovalUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripApprovalUpdateRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripApprovalUpdateRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripApprovalUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.approval.update"
}
func (this *OapiAlitripBtripApprovalUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripApprovalUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripApprovalUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripApprovalUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripApprovalUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripApprovalUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripApprovalUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripApprovalUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiUpdateApplyRq struct {
	Corpid           string    `json:"corpid,omitempty"`
	Note             string    `json:"note,omitempty"`
	OperateTime      time.Time `json:"operate_time,omitempty"`
	Status           int64     `json:"status,omitempty"`
	ThirdpartApplyId string    `json:"thirdpart_apply_id,omitempty"`
	UserName         string    `json:"user_name,omitempty"`
	Userid           string    `json:"userid,omitempty"`
}
type OapiAlitripBtripApprovalUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
