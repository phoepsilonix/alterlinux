package work

var validate *BuildTask = NewBuildTask("validate", func(w Work) error {
	return nil
})
