package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpSmartdeviceReceptionistPushinfoRequest() *CorpSmartdeviceReceptionistPushinfoRequest {
	return &CorpSmartdeviceReceptionistPushinfoRequest{}
}

type CorpSmartdeviceReceptionistPushinfoRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpSmartdeviceReceptionistPushinfoResponse
	DescContent     string
	DescTemplate    string
	MicroappAgentId int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) SetDescContent(descContent2 string) {
	this.DescContent = descContent2
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetDescContent() string {
	return this.DescContent
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) SetDescTemplate(descTemplate2 string) {
	this.DescTemplate = descTemplate2
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetDescTemplate() string {
	return this.DescTemplate
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetApiMethodName() string {
	return "dingtalk.corp.smartdevice.receptionist.pushinfo"
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["desc_content"] = this.DescContent
	txtParams["desc_template"] = this.DescTemplate
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	return txtParams
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DescContent, "descContent"); err != nil {
		return err
	}
	return nil
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpSmartdeviceReceptionistPushinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpSmartdeviceReceptionistPushinfoResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
