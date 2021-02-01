package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpUserPersonainfoGetRequest() *CorpUserPersonainfoGetRequest {
	return &CorpUserPersonainfoGetRequest{}
}

type CorpUserPersonainfoGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpUserPersonainfoGetResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *CorpUserPersonainfoGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpUserPersonainfoGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *CorpUserPersonainfoGetRequest) GetUserid() string {
	return this.Userid
}
func (this *CorpUserPersonainfoGetRequest) GetApiMethodName() string {
	return "dingtalk.corp.user.personainfo.get"
}
func (this *CorpUserPersonainfoGetRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpUserPersonainfoGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpUserPersonainfoGetRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpUserPersonainfoGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpUserPersonainfoGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpUserPersonainfoGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *CorpUserPersonainfoGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *CorpUserPersonainfoGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpUserPersonainfoGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpUserPersonainfoGetResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OpenPersonaInfo struct {
	Dob         time.Time `json:"dob,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Industry    string    `json:"industry,omitempty"`
	Title       string    `json:"title,omitempty"`
	Umids       []string  `json:"umids,omitempty"`
	WorkStation string    `json:"work_station,omitempty"`
}
