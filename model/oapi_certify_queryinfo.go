package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCertifyQueryinfoRequest() *OapiCertifyQueryinfoRequest {
	return &OapiCertifyQueryinfoRequest{}
}

type OapiCertifyQueryinfoRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCertifyQueryinfoResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiCertifyQueryinfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCertifyQueryinfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCertifyQueryinfoRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiCertifyQueryinfoRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiCertifyQueryinfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.certify.queryinfo"
}
func (this *OapiCertifyQueryinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCertifyQueryinfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCertifyQueryinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCertifyQueryinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCertifyQueryinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiCertifyQueryinfoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiCertifyQueryinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCertifyQueryinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCertifyQueryinfoResponse struct {
	taobao.TaobaoResponse
	Result YunQi2018CertifyVO `json:"result,omitempty"`
}
type YunQi2018CertifyVO struct {
	CertifyFaceImage string `json:"certify_face_image,omitempty"`
	HasCertify       bool   `json:"has_certify,omitempty"`
	NeedEnterFace    bool   `json:"need_enter_face,omitempty"`
}
