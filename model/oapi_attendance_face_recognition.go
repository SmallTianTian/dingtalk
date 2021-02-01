package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceFaceRecognitionRequest() *OapiAttendanceFaceRecognitionRequest {
	return &OapiAttendanceFaceRecognitionRequest{}
}

type OapiAttendanceFaceRecognitionRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceFaceRecognitionResponse
	MediaId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceFaceRecognitionRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceFaceRecognitionRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceFaceRecognitionRequest) SetMediaId(mediaId2 string) {
	this.MediaId = mediaId2
}
func (this *OapiAttendanceFaceRecognitionRequest) GetMediaId() string {
	return this.MediaId
}
func (this *OapiAttendanceFaceRecognitionRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.face.recognition"
}
func (this *OapiAttendanceFaceRecognitionRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceFaceRecognitionRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceFaceRecognitionRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceFaceRecognitionRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceFaceRecognitionRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["media_id"] = this.MediaId
	return txtParams
}
func (this *OapiAttendanceFaceRecognitionRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MediaId, "mediaId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceFaceRecognitionRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceFaceRecognitionRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceFaceRecognitionResponse struct {
	taobao.TaobaoResponse
	Result TopUserInfoVO `json:"result,omitempty"`
}
type TopUserInfoVO struct {
	Userid string `json:"userid,omitempty"`
}
