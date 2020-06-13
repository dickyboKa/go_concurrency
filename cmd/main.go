package main

import (
	"github.com/dickyboKa/go_concurrency/confinemen"
	"github.com/dickyboKa/go_concurrency/ctxpackage"
	"github.com/dickyboKa/go_concurrency/errorhandling"
	"github.com/dickyboKa/go_concurrency/goroutineleak"
	"github.com/dickyboKa/go_concurrency/introduction"
	"github.com/dickyboKa/go_concurrency/pipeline"
	"github.com/dickyboKa/go_concurrency/theorchannel"
)

func main() {
	introduction.DataRace()
	introduction.PlayAroundWithChannel()
	introduction.UnderstandSelectStatement()
	confinemen.AdHocConfinemen()
	confinemen.LexicalConfinemen()
	confinemen.LexicalConfinemenBuffer()
	goroutineleak.GouRoutineLeakReadChannel()
	goroutineleak.AvoidGouRoutineLeakReadChannel()
	goroutineleak.GouRoutineLeakWriteChannel()
	theorchannel.TheORChannelExperiment()
	errorhandling.ExperimentWithErrorHandling()
	pipeline.ExperimentWithPipeline()
	pipeline.InefficientPrimeNumber()
	ctxpackage.ExperimentWithCtx()
}
