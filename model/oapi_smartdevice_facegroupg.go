package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacegroupGetRequest() *OapiSmartdeviceFacegroupGetRequest {
	return &OapiSmartdeviceFacegroupGetRequest{}
}

type OapiSmartdeviceFacegroupGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacegroupGetResponse
	BizId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceFacegroupGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacegroupGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacegroupGetRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiSmartdeviceFacegroupGetRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiSmartdeviceFacegroupGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facegroup.get"
}
func (this *OapiSmartdeviceFacegroupGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacegroupGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacegroupGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacegroupGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacegroupGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	return txtParams
}
func (this *OapiSmartdeviceFacegroupGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizId, "bizId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacegroupGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacegroupGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacegroupGetResponse struct {
	taobao.TaobaoResponse
	Result  FaceGroupDto `json:"result,omitempty"`
	Success bool         `json:"success,omitempty"`
}
type FaceGroupDto struct {
	BgImgUrl    string `json:"bg_img_url,omitempty"`
	BizId       string `json:"biz_id,omitempty"`
	EndTime     int64  `json:"end_time,omitempty"`
	GreetingMsg string `json:"greeting_msg,omitempty"`
	StartTime   int64  `json:"start_time,omitempty"`
	Status      int64  `json:"status,omitempty"`
	Title       string `json:"title,omitempty"`
}
