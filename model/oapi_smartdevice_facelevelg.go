package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacelevelGetRequest() *OapiSmartdeviceFacelevelGetRequest {
	return &OapiSmartdeviceFacelevelGetRequest{}
}

type OapiSmartdeviceFacelevelGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacelevelGetResponse
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiSmartdeviceFacelevelGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacelevelGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacelevelGetRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiSmartdeviceFacelevelGetRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiSmartdeviceFacelevelGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facelevel.get"
}
func (this *OapiSmartdeviceFacelevelGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacelevelGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacelevelGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacelevelGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacelevelGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiSmartdeviceFacelevelGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UseridList, "useridList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacelevelGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacelevelGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacelevelGetResponse struct {
	taobao.TaobaoResponse
	Result []FaceResultVo `json:"result,omitempty"`
}
type FaceResultVo struct {
	CertifyType      int64     `json:"certify_type,omitempty"`
	HasFace          bool      `json:"has_face,omitempty"`
	LastModified     time.Time `json:"last_modified,omitempty"`
	OperatorUserid   string    `json:"operator_userid,omitempty"`
	OperatorUsername string    `json:"operator_username,omitempty"`
	Userid           string    `json:"userid,omitempty"`
}
