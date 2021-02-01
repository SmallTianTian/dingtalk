package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkBpmsProcessinstanceGetRequest() *SmartworkBpmsProcessinstanceGetRequest {
	return &SmartworkBpmsProcessinstanceGetRequest{}
}

type SmartworkBpmsProcessinstanceGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp              SmartworkBpmsProcessinstanceGetResponse
	ProcessInstanceId string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *SmartworkBpmsProcessinstanceGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkBpmsProcessinstanceGetRequest) SetProcessInstanceId(processInstanceId2 string) {
	this.ProcessInstanceId = processInstanceId2
}
func (this *SmartworkBpmsProcessinstanceGetRequest) GetProcessInstanceId() string {
	return this.ProcessInstanceId
}
func (this *SmartworkBpmsProcessinstanceGetRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.bpms.processinstance.get"
}
func (this *SmartworkBpmsProcessinstanceGetRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkBpmsProcessinstanceGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkBpmsProcessinstanceGetRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkBpmsProcessinstanceGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkBpmsProcessinstanceGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkBpmsProcessinstanceGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["process_instance_id"] = this.ProcessInstanceId
	return txtParams
}
func (this *SmartworkBpmsProcessinstanceGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessInstanceId, "processInstanceId"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkBpmsProcessinstanceGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkBpmsProcessinstanceGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkBpmsProcessinstanceGetResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
