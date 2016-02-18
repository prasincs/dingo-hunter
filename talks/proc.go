// From Go source code - https://github.com/golang/go
// File runtime/proc.go

// +build OMIT

func checkdead() {
...

	// -1 for sysmon
	run := sched.mcount - sched.nmidle - sched.nmidlelocked - 1
	if run > 0 {
		return
	}
...

	getg().m.throwing = -1 // do not dump full stacks
	throw("all goroutines are asleep - deadlock!")
}
