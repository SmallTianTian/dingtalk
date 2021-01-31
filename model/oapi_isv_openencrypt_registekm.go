package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiIsvOpenencryptRegistekmsRequest() *OapiIsvOpenencryptRegistekmsRequest {
	return &OapiIsvOpenencryptRegistekmsRequest{}
}

type OapiIsvOpenencryptRegistekmsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiIsvOpenencryptRegistekmsResponse
	TopHttpMethod   string
	TopKmsMeta      string
	TopResponseType string
}

func (this *OapiIsvOpenencryptRegistekmsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIsvOpenencryptRegistekmsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIsvOpenencryptRegistekmsRequest) SetTopKmsMeta(topKmsMeta2 string) {
	this.TopKmsMeta = topKmsMeta2
}
func (this *OapiIsvOpenencryptRegistekmsRequest) GetTopKmsMeta() string {
	return this.TopKmsMeta
}
func (this *OapiIsvOpenencryptRegistekmsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.isv.openencrypt.registekms"
}
func (this *OapiIsvOpenencryptRegistekmsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIsvOpenencryptRegistekmsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIsvOpenencryptRegistekmsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIsvOpenencryptRegistekmsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIsvOpenencryptRegistekmsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["top_kms_meta"] = this.TopKmsMeta
	return txtParams
}
func (this *OapiIsvOpenencryptRegistekmsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiIsvOpenencryptRegistekmsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIsvOpenencryptRegistekmsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopKMSMeta struct {
	Appid     int64  `json:"appid,omitempty"`
	Endpoint  string `json:"endpoint,omitempty"`
	Extension string `json:"extension,omitempty"`
	Requestid string `json:"requestid,omitempty"`
	Status    int64  `json:"status,omitempty"`
}
type OapiIsvOpenencryptRegistekmsResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
