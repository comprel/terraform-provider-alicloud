package r_kvstore

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

// ItemsItem is a nested struct in r_kvstore response
type ItemsItem struct {
	Id              int    `json:"Id" xml:"Id"`
	InsName         string `json:"InsName" xml:"InsName"`
	DbType          string `json:"DbType" xml:"DbType"`
	StartTime       string `json:"StartTime" xml:"StartTime"`
	SwitchTime      string `json:"SwitchTime" xml:"SwitchTime"`
	Deadline        string `json:"Deadline" xml:"Deadline"`
	Status          int    `json:"Status" xml:"Status"`
	CreatedTime     string `json:"CreatedTime" xml:"CreatedTime"`
	ModifiedTime    string `json:"ModifiedTime" xml:"ModifiedTime"`
	ResultInfo      string `json:"ResultInfo" xml:"ResultInfo"`
	PrepareInterval string `json:"PrepareInterval" xml:"PrepareInterval"`
	TaskParams      string `json:"TaskParams" xml:"TaskParams"`
	TaskType        string `json:"TaskType" xml:"TaskType"`
	Region          string `json:"Region" xml:"Region"`
}
