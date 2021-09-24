package client

import (
	"time"

	executorv1 "github.com/kubeshop/kubtest-operator/apis/executor/v1"
	scriptv1 "github.com/kubeshop/kubtest-operator/apis/script/v1"
	"github.com/kubeshop/kubtest/pkg/api/v1/kubtest"
)

const (
	WatchInterval = time.Second

	ExecutorTypeRest = "rest"
	ExecutorTypeJob  = "job"
)

type ExecuteOptions struct {
	ID           string
	ScriptName   string
	ScriptSpec   scriptv1.ScriptSpec
	ExecutorName string
	ExecutorSpec executorv1.ExecutorSpec
	Request      kubtest.ExecutionRequest
}

func NewExecuteOptions() ExecuteOptions {
	options := ExecuteOptions{}
	return options
}
