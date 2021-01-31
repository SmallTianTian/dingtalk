package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiAttendanceClassGetRequest() *OapiAttendanceClassGetRequest {
	return &OapiAttendanceClassGetRequest{}
}

type OapiAttendanceClassGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceClassGetResponse
	ClassId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceClassGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceClassGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceClassGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiAttendanceClassGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiAttendanceClassGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.class.get"
}
func (this *OapiAttendanceClassGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceClassGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceClassGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceClassGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceClassGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	return txtParams
}
func (this *OapiAttendanceClassGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceClassGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceClassGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceClassGetResponse struct {
	taobao.TaobaoResponse
	Result AtClassForTopVo `json:"result,omitempty"`
}
type AtTimeVo struct {
	Across    int64     `json:"across,omitempty"`
	BeginMin  int64     `json:"begin_min,omitempty"`
	CheckTime time.Time `json:"check_time,omitempty"`
	CheckType string    `json:"check_type,omitempty"`
	EndMin    int64     `json:"end_min,omitempty"`
}
type AtClassSettingForTopVo struct {
	ClassId       int64    `json:"class_id,omitempty"`
	Id            int64    `json:"id,omitempty"`
	RestBeginTime AtTimeVo `json:"rest_begin_time,omitempty"`
	RestEndTime   AtTimeVo `json:"rest_end_time,omitempty"`
}
type AtSectionVo struct {
	Times []AtTimeVo `json:"times,omitempty"`
}
type AtClassForTopVo struct {
	ClassSetting AtClassSettingForTopVo `json:"class_setting,omitempty"`
	CorpId       string                 `json:"corp_id,omitempty"`
	Id           int64                  `json:"id,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Sections     []AtSectionVo          `json:"sections,omitempty"`
	WorkDays     []int64                `json:"work_days,omitempty"`
}
