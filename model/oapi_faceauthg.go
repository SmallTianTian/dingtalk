package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFaceauthGetRequest() *OapiFaceauthGetRequest {
	return &OapiFaceauthGetRequest{}
}

type OapiFaceauthGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFaceauthGetResponse
	AppBizId        string
	TmpAuthCode     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiFaceauthGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFaceauthGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFaceauthGetRequest) SetAppBizId(appBizId2 string) {
	this.AppBizId = appBizId2
}
func (this *OapiFaceauthGetRequest) GetAppBizId() string {
	return this.AppBizId
}
func (this *OapiFaceauthGetRequest) SetTmpAuthCode(tmpAuthCode2 string) {
	this.TmpAuthCode = tmpAuthCode2
}
func (this *OapiFaceauthGetRequest) GetTmpAuthCode() string {
	return this.TmpAuthCode
}
func (this *OapiFaceauthGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.faceauth.get"
}
func (this *OapiFaceauthGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFaceauthGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFaceauthGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFaceauthGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFaceauthGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_biz_id"] = this.AppBizId
	txtParams["tmp_auth_code"] = this.TmpAuthCode
	return txtParams
}
func (this *OapiFaceauthGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppBizId, "appBizId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFaceauthGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFaceauthGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFaceauthGetResponse struct {
	taobao.TaobaoResponse
	Result Result `json:"result,omitempty"`
}
