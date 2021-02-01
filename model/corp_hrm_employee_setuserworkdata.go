package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpHrmEmployeeSetuserworkdataRequest() *CorpHrmEmployeeSetuserworkdataRequest {
	return &CorpHrmEmployeeSetuserworkdataRequest{}
}

type CorpHrmEmployeeSetuserworkdataRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                CorpHrmEmployeeSetuserworkdataResponse
	HrmApiUserDataModel string
	OpUserid            string
	TopHttpMethod       string
	TopResponseType     string
}

func (this *CorpHrmEmployeeSetuserworkdataRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) SetHrmApiUserDataModel(hrmApiUserDataModel2 string) {
	this.HrmApiUserDataModel = hrmApiUserDataModel2
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) GetHrmApiUserDataModel() string {
	return this.HrmApiUserDataModel
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) GetApiMethodName() string {
	return "dingtalk.corp.hrm.employee.setuserworkdata"
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["hrm_api_user_data_model"] = this.HrmApiUserDataModel
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserid, "opUserid"); err != nil {
		return err
	}
	return nil
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpHrmEmployeeSetuserworkdataRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type HrmApiUserDataModel struct {
	DataDesc  string `json:"data_desc,omitempty"`
	DataValue string `json:"data_value,omitempty"`
	Userid    string `json:"userid,omitempty"`
}
type CorpHrmEmployeeSetuserworkdataResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
