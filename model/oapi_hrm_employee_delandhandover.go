package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHrmEmployeeDelandhandoverRequest() *OapiHrmEmployeeDelandhandoverRequest {
	return &OapiHrmEmployeeDelandhandoverRequest{}
}

type OapiHrmEmployeeDelandhandoverRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                       OapiHrmEmployeeDelandhandoverResponse
	DismissionInfoWithHandOver string
	OpUserid                   string
	TopHttpMethod              string
	TopResponseType            string
}

func (this *OapiHrmEmployeeDelandhandoverRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHrmEmployeeDelandhandoverRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHrmEmployeeDelandhandoverRequest) SetDismissionInfoWithHandOver(dismissionInfoWithHandOver2 string) {
	this.DismissionInfoWithHandOver = dismissionInfoWithHandOver2
}
func (this *OapiHrmEmployeeDelandhandoverRequest) GetDismissionInfoWithHandOver() string {
	return this.DismissionInfoWithHandOver
}
func (this *OapiHrmEmployeeDelandhandoverRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiHrmEmployeeDelandhandoverRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiHrmEmployeeDelandhandoverRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hrm.employee.delandhandover"
}
func (this *OapiHrmEmployeeDelandhandoverRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHrmEmployeeDelandhandoverRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHrmEmployeeDelandhandoverRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHrmEmployeeDelandhandoverRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHrmEmployeeDelandhandoverRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dismission_info_with_hand_over"] = this.DismissionInfoWithHandOver
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiHrmEmployeeDelandhandoverRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserid, "opUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHrmEmployeeDelandhandoverRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHrmEmployeeDelandhandoverRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHrmEmployeeDelandhandoverResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
