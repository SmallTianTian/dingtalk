package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCampusGetRequest() *OapiEduCampusGetRequest {
	return &OapiEduCampusGetRequest{}
}

type OapiEduCampusGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCampusGetResponse
	CampusId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCampusGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCampusGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCampusGetRequest) SetCampusId(campusId2 int64) {
	this.CampusId = campusId2
}
func (this *OapiEduCampusGetRequest) GetCampusId() int64 {
	return this.CampusId
}
func (this *OapiEduCampusGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.campus.get"
}
func (this *OapiEduCampusGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCampusGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCampusGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCampusGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCampusGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["campus_id"] = this.CampusId
	return txtParams
}
func (this *OapiEduCampusGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CampusId, "campusId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCampusGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCampusGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCampusGetResponse struct {
	taobao.TaobaoResponse
	Result  CampusResponse `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type CampusResponse struct {
	CampusId int64  `json:"campus_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Nick     string `json:"nick,omitempty"`
}
