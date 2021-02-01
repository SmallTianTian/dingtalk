package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsEvaluateJobmatchCancelRequest() *OapiAtsEvaluateJobmatchCancelRequest {
	return &OapiAtsEvaluateJobmatchCancelRequest{}
}

type OapiAtsEvaluateJobmatchCancelRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsEvaluateJobmatchCancelResponse
	BizCode         string
	ExtData         string
	OuterEvaluateId string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsEvaluateJobmatchCancelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) SetExtData(extData2 string) {
	this.ExtData = extData2
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) GetExtData() string {
	return this.ExtData
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) SetOuterEvaluateId(outerEvaluateId2 string) {
	this.OuterEvaluateId = outerEvaluateId2
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) GetOuterEvaluateId() string {
	return this.OuterEvaluateId
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.evaluate.jobmatch.cancel"
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["ext_data"] = this.ExtData
	txtParams["outer_evaluate_id"] = this.OuterEvaluateId
	return txtParams
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsEvaluateJobmatchCancelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsEvaluateJobmatchCancelResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
