package workers

type Job interface {
	Jid() string
	Queue() string
	Args() *Args
	OriginalJson() string
	ToJson() string
	Equals(interface{}) bool
}

type JobFunc func(job Job)
