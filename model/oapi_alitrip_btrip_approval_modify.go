package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripApprovalModifyRequest() *OapiAlitripBtripApprovalModifyRequest {
	return &OapiAlitripBtripApprovalModifyRequest{}
}

type OapiAlitripBtripApprovalModifyRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripApprovalModifyResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripApprovalModifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripApprovalModifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripApprovalModifyRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripApprovalModifyRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripApprovalModifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.approval.modify"
}
func (this *OapiAlitripBtripApprovalModifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripApprovalModifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripApprovalModifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripApprovalModifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripApprovalModifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripApprovalModifyRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripApprovalModifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripApprovalModifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiNewApplyRq struct {
	CorpName            string              `json:"corp_name,omitempty"`
	Corpid              string              `json:"corpid,omitempty"`
	DeptName            string              `json:"dept_name,omitempty"`
	Deptid              string              `json:"deptid,omitempty"`
	ItineraryList       []OpenItineraryInfo `json:"itinerary_list,omitempty"`
	Status              int64               `json:"status,omitempty"`
	ThirdpartApplyId    string              `json:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId string              `json:"thirdpart_business_id,omitempty"`
	TravelerList        []OpenUserInfo      `json:"traveler_list,omitempty"`
	TripCause           string              `json:"trip_cause,omitempty"`
	TripDay             int64               `json:"trip_day,omitempty"`
	TripTitle           string              `json:"trip_title,omitempty"`
	UserName            string              `json:"user_name,omitempty"`
	Userid              string              `json:"userid,omitempty"`
}
type OapiAlitripBtripApprovalModifyResponse struct {
	taobao.TaobaoResponse
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
	Module  OpenApiNewApplyRs `json:"module,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type OpenApiNewApplyRs struct {
	ApplyId          int64  `json:"apply_id,omitempty"`
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty"`
}
