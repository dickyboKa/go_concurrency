package main

import (
	"github.com/dickyboKa/go_concurrency/goroutineleak"
)

func main() {
	/*introduction.DataRace()
	introduction.PlayAroundWithChannel()
	introduction.UnderstandSelectStatement()
	confinemen.AdHocConfinemen()
	confinemen.LexicalConfinemen()
	confinemen.LexicalConfinemenBuffer()*/
	goroutineleak.ThisIsLeaking()
	goroutineleak.AvoidGoRoutineLeakWithForSelect()
}
