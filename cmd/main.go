package main

import (
	"github.com/dickyboKa/go_concurrency/confinemen"
	"github.com/dickyboKa/go_concurrency/goroutineleak"
	"github.com/dickyboKa/go_concurrency/introduction"
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
}
