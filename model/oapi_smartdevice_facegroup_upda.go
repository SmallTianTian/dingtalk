package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacegroupUpdateRequest() *OapiSmartdeviceFacegroupUpdateRequest {
	return &OapiSmartdeviceFacegroupUpdateRequest{}
}

type OapiSmartdeviceFacegroupUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacegroupUpdateResponse
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

func (this *OapiSmartdeviceFacegroupUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetBgImgUrl(bgImgUrl2 string) {
	this.BgImgUrl = bgImgUrl2
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetBgImgUrl() string {
	return this.BgImgUrl
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetGreetingMsg(greetingMsg2 string) {
	this.GreetingMsg = greetingMsg2
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetGreetingMsg() string {
	return this.GreetingMsg
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetStatus(status2 int64) {
	this.Status = status2
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetStatus() int64 {
	return this.Status
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facegroup.update"
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetTextParams() map[string]interface{} {
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
func (this *OapiSmartdeviceFacegroupUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxLength(this.BgImgUrl, 512, "bgImgUrl"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacegroupUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacegroupUpdateResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
