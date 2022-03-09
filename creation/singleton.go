package creation

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}