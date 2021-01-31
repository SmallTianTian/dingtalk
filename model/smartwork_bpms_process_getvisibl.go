package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkBpmsProcessGetvisibleRequest() *SmartworkBpmsProcessGetvisibleRequest {
	return &SmartworkBpmsProcessGetvisibleRequest{}
}

type SmartworkBpmsProcessGetvisibleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkBpmsProcessGetvisibleResponse
	ProcessCodeList string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *SmartworkBpmsProcessGetvisibleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkBpmsProcessGetvisibleRequest) SetProcessCodeList(processCodeList2 string) {
	this.ProcessCodeList = processCodeList2
}
func (this *SmartworkBpmsProcessGetvisibleRequest) GetProcessCodeList() string {
	return this.ProcessCodeList
}
func (this *SmartworkBpmsProcessGetvisibleRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *SmartworkBpmsProcessGetvisibleRequest) GetUserid() string {
	return this.Userid
}
func (this *SmartworkBpmsProcessGetvisibleRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.bpms.process.getvisible"
}
func (this *SmartworkBpmsProcessGetvisibleRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkBpmsProcessGetvisibleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkBpmsProcessGetvisibleRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkBpmsProcessGetvisibleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkBpmsProcessGetvisibleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkBpmsProcessGetvisibleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["process_code_list"] = this.ProcessCodeList
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *SmartworkBpmsProcessGetvisibleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessCodeList, "processCodeList"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkBpmsProcessGetvisibleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkBpmsProcessGetvisibleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkBpmsProcessGetvisibleResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
