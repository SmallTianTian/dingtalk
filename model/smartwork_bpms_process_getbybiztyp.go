package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkBpmsProcessGetbybiztypeRequest() *SmartworkBpmsProcessGetbybiztypeRequest {
	return &SmartworkBpmsProcessGetbybiztypeRequest{}
}

type SmartworkBpmsProcessGetbybiztypeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkBpmsProcessGetbybiztypeResponse
	BizType         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *SmartworkBpmsProcessGetbybiztypeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) SetBizType(bizType2 string) {
	this.BizType = bizType2
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) GetBizType() string {
	return this.BizType
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.bpms.process.getbybiztype"
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_type"] = this.BizType
	return txtParams
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizType, "bizType"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkBpmsProcessGetbybiztypeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkBpmsProcessGetbybiztypeResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
