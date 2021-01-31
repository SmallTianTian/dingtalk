package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacegroupCreateRequest() *OapiSmartdeviceFacegroupCreateRequest {
	return &OapiSmartdeviceFacegroupCreateRequest{}
}

type OapiSmartdeviceFacegroupCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacegroupCreateResponse
	BgImgUrl        string
	BizId           string
	EndTime         int64
	GreetingMsg     string
	StartTime       int64
	Status          int64
	Title           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceFacegroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetBgImgUrl(bgImgUrl2 string) {
	this.BgImgUrl = bgImgUrl2
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetBgImgUrl() string {
	return this.BgImgUrl
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetGreetingMsg(greetingMsg2 string) {
	this.GreetingMsg = greetingMsg2
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetGreetingMsg() string {
	return this.GreetingMsg
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetStatus(status2 int64) {
	this.Status = status2
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetStatus() int64 {
	return this.Status
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facegroup.create"
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacegroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["bg_img_url"] = this.BgImgUrl
	txtParams["biz_id"] = this.BizId
	txtParams["end_time"] = this.EndTime
	txtParams["greeting_msg"] = this.GreetingMsg
	txtParams["start_time"] = this.StartTime
	txtParams["status"] = this.Status
	txtParams["title"] = this.Title
	return txtParams
}
func (this *OapiSmartdeviceFacegroupCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxLength(this.BgImgUrl, 512, "bgImgUrl"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacegroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacegroupCreateResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
