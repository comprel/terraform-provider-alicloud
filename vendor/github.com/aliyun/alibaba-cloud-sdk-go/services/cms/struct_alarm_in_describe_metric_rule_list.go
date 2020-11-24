package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AlarmInDescribeMetricRuleList is a nested struct in cms response
type AlarmInDescribeMetricRuleList struct {
	RuleId              string      `json:"RuleId" xml:"RuleId"`
	Namespace           string      `json:"Namespace" xml:"Namespace"`
	MetricName          string      `json:"MetricName" xml:"MetricName"`
	Period              string      `json:"Period" xml:"Period"`
	EffectiveInterval   string      `json:"EffectiveInterval" xml:"EffectiveInterval"`
	NoEffectiveInterval string      `json:"NoEffectiveInterval" xml:"NoEffectiveInterval"`
	SilenceTime         int         `json:"SilenceTime" xml:"SilenceTime"`
	EnableState         bool        `json:"EnableState" xml:"EnableState"`
	AlertState          string      `json:"AlertState" xml:"AlertState"`
	ContactGroups       string      `json:"ContactGroups" xml:"ContactGroups"`
	Webhook             string      `json:"Webhook" xml:"Webhook"`
	MailSubject         string      `json:"MailSubject" xml:"MailSubject"`
	RuleName            string      `json:"RuleName" xml:"RuleName"`
	Resources           string      `json:"Resources" xml:"Resources"`
	GroupId             string      `json:"GroupId" xml:"GroupId"`
	GroupName           string      `json:"GroupName" xml:"GroupName"`
	Dimensions          string      `json:"Dimensions" xml:"Dimensions"`
	SourceType          string      `json:"SourceType" xml:"SourceType"`
	Escalations         Escalations `json:"Escalations" xml:"Escalations"`
}