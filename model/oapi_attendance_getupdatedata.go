package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGetupdatedataRequest() *OapiAttendanceGetupdatedataRequest {
	return &OapiAttendanceGetupdatedataRequest{}
}

type OapiAttendanceGetupdatedataRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGetupdatedataResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
	WorkDate        time.Time
}

func (this *OapiAttendanceGetupdatedataRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGetupdatedataRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGetupdatedataRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceGetupdatedataRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceGetupdatedataRequest) SetWorkDate(workDate2 time.Time) {
	this.WorkDate = workDate2
}
func (this *OapiAttendanceGetupdatedataRequest) GetWorkDate() time.Time {
	return this.WorkDate
}
func (this *OapiAttendanceGetupdatedataRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.getupdatedata"
}
func (this *OapiAttendanceGetupdatedataRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGetupdatedataRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGetupdatedataRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGetupdatedataRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGetupdatedataRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["work_date"] = this.WorkDate
	return txtParams
}
func (this *OapiAttendanceGetupdatedataRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGetupdatedataRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGetupdatedataRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGetupdatedataResponse struct {
	taobao.TaobaoResponse
	Result  AtCheckInfoForOpenVo `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type AtRestTimeVo struct {
	RestBeginTime int64 `json:"rest_begin_time,omitempty"`
	RestEndTime   int64 `json:"rest_end_time,omitempty"`
}
type AtClassSettingInfoForOpenVo struct {
	RestTimeVoList []AtRestTimeVo `json:"rest_time_vo_list,omitempty"`
}
