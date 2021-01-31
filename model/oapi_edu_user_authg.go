package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduUserAuthGetRequest() *OapiEduUserAuthGetRequest {
	return &OapiEduUserAuthGetRequest{}
}

type OapiEduUserAuthGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduUserAuthGetResponse
	Language        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduUserAuthGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduUserAuthGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduUserAuthGetRequest) SetLanguage(language2 string) {
	this.Language = language2
}
func (this *OapiEduUserAuthGetRequest) GetLanguage() string {
	return this.Language
}
func (this *OapiEduUserAuthGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduUserAuthGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduUserAuthGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.user.auth.get"
}
func (this *OapiEduUserAuthGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduUserAuthGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduUserAuthGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduUserAuthGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduUserAuthGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["language"] = this.Language
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduUserAuthGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxLength(this.Language, 6, "language"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduUserAuthGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduUserAuthGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduUserAuthGetResponse struct {
	taobao.TaobaoResponse
	Result UserGetRequest `json:"result,omitempty"`
}
type UserGetRequest struct {
	RealAuthed bool   `json:"real_authed,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
