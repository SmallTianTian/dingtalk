package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiBlackboardGetRequest() *OapiBlackboardGetRequest {
	return &OapiBlackboardGetRequest{}
}

type OapiBlackboardGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBlackboardGetResponse
	BlackboardId    string
	OperationUserid string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiBlackboardGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBlackboardGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBlackboardGetRequest) SetBlackboardId(blackboardId2 string) {
	this.BlackboardId = blackboardId2
}
func (this *OapiBlackboardGetRequest) GetBlackboardId() string {
	return this.BlackboardId
}
func (this *OapiBlackboardGetRequest) SetOperationUserid(operationUserid2 string) {
	this.OperationUserid = operationUserid2
}
func (this *OapiBlackboardGetRequest) GetOperationUserid() string {
	return this.OperationUserid
}
func (this *OapiBlackboardGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.blackboard.get"
}
func (this *OapiBlackboardGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBlackboardGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBlackboardGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBlackboardGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBlackboardGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["blackboard_id"] = this.BlackboardId
	txtParams["operation_userid"] = this.OperationUserid
	return txtParams
}
func (this *OapiBlackboardGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BlackboardId, "blackboardId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiBlackboardGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBlackboardGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiBlackboardGetResponse struct {
	taobao.TaobaoResponse
	Result  OapiBlackboardVo `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type OapiBlackboardVo struct {
	Author       string    `json:"author,omitempty"`
	CategoryId   string    `json:"category_id,omitempty"`
	Content      string    `json:"content,omitempty"`
	CoverpicUrl  string    `json:"coverpic_url,omitempty"`
	DepnameList  []string  `json:"depname_list,omitempty"`
	GmtCreate    time.Time `json:"gmt_create,omitempty"`
	GmtModified  time.Time `json:"gmt_modified,omitempty"`
	Id           string    `json:"id,omitempty"`
	PrivateLevel int64     `json:"private_level,omitempty"`
	ReadCount    int64     `json:"read_count,omitempty"`
	Title        string    `json:"title,omitempty"`
	UnreadCount  int64     `json:"unread_count,omitempty"`
	UsernameList []string  `json:"username_list,omitempty"`
}
