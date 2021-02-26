package run_env

import (
	"lu-short/common/utils/id_util"
	"os"
)

// build : go build -ldflags "-X 'lu-short/common/run_env.VERSION=`cat ./version`' -X 'lu-short/common/run_env.BUILD_TIME=`date "+%Y-%m-%d %H:%M:%S"`' -X 'lu-short/common/run_env.GO_VERSION=`go version`' -X 'lu-short/common/run_env.GIT_COMMIT=${GITHUB_SHA}'" -v .

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
	GIT_COMMIT string
)

// RunEnv is the app's run env
type RunEnv struct {
	AppName   string
	AppHost   string
	Version   string
	BuildTime string
	GoVersion string
	GitCommit string
}

var RunEnvLocal = &RunEnv{
	AppName:   "lu-short",
	AppHost:   GetHostName(),
	Version:   VERSION,
	BuildTime: BUILD_TIME,
	GoVersion: GO_VERSION,
	GitCommit: GIT_COMMIT,
}

var tempHostName = id_util.NextSnowflake()

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return tempHostName
	}
	return hostname
}

func (env RunEnv) GetShortGitCommitSha() string {
	if len(env.GitCommit) >= 6 {
		return env.GitCommit[0:6]
	}
	return env.GitCommit
}
