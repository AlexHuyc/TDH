package entity

import (
	"fmt"
	"strconv"
	"tdhdbamonithr/src/util"
	"time"
)

type JsonQuery1 struct {
	Data []struct {
		AccumulableDurations          interface{} `json:"accumulableDurations"`
		ActiveTasks                   int64       `json:"activeTasks"`
		CompletionTime                int64       `json:"completionTime"`
		Description                   string      `json:"description"`
		DetailedDescription           string      `json:"detailedDescription"`
		DurationBlockMap              struct{}    `json:"durationBlockMap"`
		Durations                     interface{} `json:"durations"`
		ExecutionFinishTime           int64       `json:"executionFinishTime"`
		ExecutionStartTime            int64       `json:"executionStartTime"`
		ExtraJobID                    int64       `json:"extraJobId"`
		FetchComputeTotal             int64       `json:"fetchComputeTotal"`
		FetchCount                    int64       `json:"fetchCount"`
		FetchResultStartTime          int64       `json:"fetchResultStartTime"`
		FetchRowCount                 int64       `json:"fetchRowCount"`
		FileCommitFinishTime          int64       `json:"fileCommitFinishTime"`
		Group                         string      `json:"group"`
		Jobs                          string      `json:"jobs"`
		JobsActive                    int64       `json:"jobsActive"`
		JobsCompleted                 int64       `json:"jobsCompleted"`
		JobsFailed                    int64       `json:"jobsFailed"`
		LogicalPlanString             interface{} `json:"logicalPlanString"`
		LogicalPlanStrings            interface{} `json:"logicalPlanStrings"`
		MemoryPerTask                 int64       `json:"memoryPerTask"`
		Message                       string `json:"message"`
		MessageNameStr                interface{} `json:"messageNameStr"`
		Meta                          struct{}    `json:"meta"`
		Metrics                       interface{} `json:"metrics"`
		Mode                          string      `json:"mode"`
		NonLinearAccumulableDurations interface{} `json:"nonLinearAccumulableDurations"`
		NonLinearDurations            interface{} `json:"nonLinearDurations"`
		PhysicalPlanString            interface{} `json:"physicalPlanString"`
		PhysicalPlanStrings           interface{} `json:"physicalPlanStrings"`
		PlanInputs                    interface{} `json:"planInputs"`
		PlanOutputs                   interface{} `json:"planOutputs"`
		RealFetchResultStartTime      int64       `json:"realFetchResultStartTime"`
		SessionID                     int64       `json:"sessionId"`
		SessionName                   string      `json:"sessionName"`
		SqlBlockJSON                  interface{} `json:"sqlBlockJson"`
		SqlID                         int64       `json:"sqlId"`
		SqlName                       interface{} `json:"sqlName"`
		SqlType                       interface{} `json:"sqlType"`
		Stages                        interface{} `json:"stages"`
		StagesActive                  int64       `json:"stagesActive"`
		StagesCompleted               int64       `json:"stagesCompleted"`
		StagesFailed                  int64       `json:"stagesFailed"`
		StagesSkipped                 int64       `json:"stagesSkipped"`
		State                         string      `json:"state"`
		SubmissionTime                int64       `json:"submissionTime"`
		TableScans                    interface{} `json:"tableScans"`
		Tags                          string      `json:"tags"`
		TasksActive                   int64       `json:"tasksActive"`
		TasksCompleted                int64       `json:"tasksCompleted"`
		TasksExpected                 int64       `json:"tasksExpected"`
		TasksFailed                   int64       `json:"tasksFailed"`
		TasksSkipped                  int64       `json:"tasksSkipped"`
		TotalTaskTime                 int64       `json:"totalTaskTime"`
		User                          string      `json:"user"`
	} `json:"data"`
	Error []string `json:"error"`
	Info  []interface{} `json:"info"`
	Query struct {
		DataKey  string      `json:"dataKey"`
		DataSize int64       `json:"dataSize"`
		Host     interface{} `json:"host"`
		ID       int64       `json:"id"`
		Port     interface{} `json:"port"`
		StringID interface{} `json:"stringId"`
	} `json:"query"`
	Warning           []interface{} `json:"warning"`
	WatchmanTimestamp int64         `json:"watchmanTimestamp"`
}

type JsonQueryStageInfo struct {
	Data struct {
		SqlID                         int64         `json:"sqlId"`
		Stages                        []struct {
			//Stages []struct {
			//	StageID        int64       `json:"stageId"`
			//} `json:"attempts"`
			StageID        int64       `json:"stageId"`
		} `json:"stages"`
	} `json:"data"`
}

type Query struct {
	ServerKey 			string
	SqlID 				int64
	State				string
	Stages              []int64
	TaskInfo			[]StageTaskInfo
	User				string
	Description			string
	SubmissionTime		int64
	CompletionTime      int64
	Message				string
	CrawlMessage		string

}

func GetQueriesList (queies JsonQuery1,serverKey string) map[string]Query{

	querymap  := make(map[string]Query)
	for _,v :=range queies.Data{
		if util.FilterBySQL(v.Description) ||
			util.FilterByUnixtime(v.SubmissionTime,40,"secound") ||
			util.FilterByState(v.State){
		}else {
		var query Query
		query.ServerKey=serverKey
		query.SqlID=v.SqlID
		query.State=v.State
		query.CompletionTime=v.CompletionTime
		query.Description=v.Description
		query.SubmissionTime=v.SubmissionTime
		query.User=v.User
		query.Message=v.Message
		querymap [queies.Query.DataKey+"||"+strconv.FormatInt(v.SqlID,10)]=query
	}
	}
	return querymap
}

func QueryToStringBySeparator (query Query,separator string,wftasklist map[string]Wftaskinfo)string {
		fmt.Println()
	return "query"+	separator	+
		query.ServerKey	+"||"+util.Int64ToString(query.SqlID)+"||"+util.Int64ToString(query.SubmissionTime)	+	separator	+
		query.ServerKey	+	separator	+
		util.Int64ToString(query.SqlID)	+	separator	+
		query.State	+	separator	+
		StagesListToString(query.Stages) +	separator	+
		TaskInfoListSplitToString(query.ServerKey,query.TaskInfo) +	separator	+
		query.User +	separator	+
		query.Description +	separator	+
		time.Unix(query.SubmissionTime/1000,0).Format("2006-01-02") +	separator	+
		util.Int64ToString(query.SubmissionTime) +	separator	+
		util.Int64ToString(query.CompletionTime) +	separator	+
		GetSimilaryListInFo(util.CleanNewlineChart(query.Description),wftasklist,separator)  +	separator	+
		query.Message +	separator	+
		query.CrawlMessage


}
