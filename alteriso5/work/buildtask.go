package work

import (
	"fmt"
	"log/slog"
)

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
	slog.Info("Running task", "name", t.name)

	err := t.task(*w)
	if err != nil {
		return fmt.Errorf("error on %s: %v", t.name, err)
	}
	return nil
}
