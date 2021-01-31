package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiEduHomeworkQueryRequest() *OapiEduHomeworkQueryRequest {
	return &OapiEduHomeworkQueryRequest{}
}

type OapiEduHomeworkQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkQueryResponse
	BizCode         string
	HwId            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduHomeworkQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkQueryRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkQueryRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkQueryRequest) SetHwId(hwId2 int64) {
	this.HwId = hwId2
}
func (this *OapiEduHomeworkQueryRequest) GetHwId() int64 {
	return this.HwId
}
func (this *OapiEduHomeworkQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.query"
}
func (this *OapiEduHomeworkQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["hw_id"] = this.HwId
	return txtParams
}
func (this *OapiEduHomeworkQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkQueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  OpenHwDetailResponse `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type OpenHwDetailResponse struct {
	Attributes       string    `json:"attributes,omitempty"`
	HwContent        string    `json:"hw_content,omitempty"`
	HwMedia          string    `json:"hw_media,omitempty"`
	HwPhoto          string    `json:"hw_photo,omitempty"`
	HwStatus         string    `json:"hw_status,omitempty"`
	HwTitle          string    `json:"hw_title,omitempty"`
	HwVideo          string    `json:"hw_video,omitempty"`
	ScheduledRelease string    `json:"scheduled_release,omitempty"`
	ScheduledTime    string    `json:"scheduled_time,omitempty"`
	SendTime         time.Time `json:"send_time,omitempty"`
	TeacherId        string    `json:"teacher_id,omitempty"`
	TeacherName      string    `json:"teacher_name,omitempty"`
}
