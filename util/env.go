package util

import "os"

const WorkerFlagEnv = "TK_IS_WORKER"

func IsWorker() bool {
	return os.Getenv(WorkerFlagEnv) == "1"
}
