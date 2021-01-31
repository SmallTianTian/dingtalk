package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpSmartdeviceGetfaceRequest() *CorpSmartdeviceGetfaceRequest {
	return &CorpSmartdeviceGetfaceRequest{}
}

type CorpSmartdeviceGetfaceRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpSmartdeviceGetfaceResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *CorpSmartdeviceGetfaceRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpSmartdeviceGetfaceRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *CorpSmartdeviceGetfaceRequest) GetUserid() string {
	return this.Userid
}
func (this *CorpSmartdeviceGetfaceRequest) GetApiMethodName() string {
	return "dingtalk.corp.smartdevice.getface"
}
func (this *CorpSmartdeviceGetfaceRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpSmartdeviceGetfaceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpSmartdeviceGetfaceRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpSmartdeviceGetfaceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpSmartdeviceGetfaceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpSmartdeviceGetfaceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *CorpSmartdeviceGetfaceRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpSmartdeviceGetfaceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpSmartdeviceGetfaceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpSmartdeviceGetfaceResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
