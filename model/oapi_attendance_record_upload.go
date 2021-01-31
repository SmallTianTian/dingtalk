package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceRecordUploadRequest() *OapiAttendanceRecordUploadRequest {
	return &OapiAttendanceRecordUploadRequest{}
}

type OapiAttendanceRecordUploadRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceRecordUploadResponse
	DeviceId        string
	DeviceName      string
	PhotoUrl        string
	TopHttpMethod   string
	TopResponseType string
	UserCheckTime   int64
	Userid          string
}

func (this *OapiAttendanceRecordUploadRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceRecordUploadRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceRecordUploadRequest) SetDeviceId(deviceId2 string) {
	this.DeviceId = deviceId2
}
func (this *OapiAttendanceRecordUploadRequest) GetDeviceId() string {
	return this.DeviceId
}
func (this *OapiAttendanceRecordUploadRequest) SetDeviceName(deviceName2 string) {
	this.DeviceName = deviceName2
}
func (this *OapiAttendanceRecordUploadRequest) GetDeviceName() string {
	return this.DeviceName
}
func (this *OapiAttendanceRecordUploadRequest) SetPhotoUrl(photoUrl2 string) {
	this.PhotoUrl = photoUrl2
}
func (this *OapiAttendanceRecordUploadRequest) GetPhotoUrl() string {
	return this.PhotoUrl
}
func (this *OapiAttendanceRecordUploadRequest) SetUserCheckTime(userCheckTime2 int64) {
	this.UserCheckTime = userCheckTime2
}
func (this *OapiAttendanceRecordUploadRequest) GetUserCheckTime() int64 {
	return this.UserCheckTime
}
func (this *OapiAttendanceRecordUploadRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceRecordUploadRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceRecordUploadRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.record.upload"
}
func (this *OapiAttendanceRecordUploadRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceRecordUploadRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceRecordUploadRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceRecordUploadRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceRecordUploadRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_id"] = this.DeviceId
	txtParams["device_name"] = this.DeviceName
	txtParams["photo_url"] = this.PhotoUrl
	txtParams["user_check_time"] = this.UserCheckTime
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAttendanceRecordUploadRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeviceId, "deviceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceRecordUploadRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceRecordUploadRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceRecordUploadResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
