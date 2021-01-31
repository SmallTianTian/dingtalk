package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpEncryptionKeyListRequest() *CorpEncryptionKeyListRequest {
	return &CorpEncryptionKeyListRequest{}
}

type CorpEncryptionKeyListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpEncryptionKeyListResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpEncryptionKeyListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpEncryptionKeyListRequest) GetApiMethodName() string {
	return "dingtalk.corp.encryption.key.list"
}
func (this *CorpEncryptionKeyListRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpEncryptionKeyListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpEncryptionKeyListRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpEncryptionKeyListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpEncryptionKeyListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpEncryptionKeyListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *CorpEncryptionKeyListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpEncryptionKeyListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpEncryptionKeyListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpEncryptionKeyListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
