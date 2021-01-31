package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiBipaasGenericRequest() *OapiBipaasGenericRequest {
	return &OapiBipaasGenericRequest{}
}

type OapiBipaasGenericRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBipaasGenericResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiBipaasGenericRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBipaasGenericRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBipaasGenericRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiBipaasGenericRequest) SetRequestString(request2 string) {
	this.Request = request2
}
func (this *OapiBipaasGenericRequest) GetRequest() string {
	return this.Request
}
func (this *OapiBipaasGenericRequest) GetApiMethodName() string {
	return "dingtalk.oapi.bipaas.generic"
}
func (this *OapiBipaasGenericRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBipaasGenericRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBipaasGenericRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBipaasGenericRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBipaasGenericRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiBipaasGenericRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Request, taobao.QM_ROOT_TAG); err != nil {
		return err
	}
	return nil
}
func (this *OapiBipaasGenericRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBipaasGenericRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiBipaasGenericResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
