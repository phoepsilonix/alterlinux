package work

type BuildTask struct {
	name string
	task func(work Work) error
}


func NewBuildTask(name string, task func(Work) error) *BuildTask {
	return &BuildTask{
		name: name,
		task: task,
	}
}

func (t *BuildTask) Name() string {
	return t.name
}

func (t *BuildTask) Run(w *Work) error {
	return t.task(*w)
}
