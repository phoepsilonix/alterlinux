package work

// Run each bootmodes
var makeBootModes *BuildTask = NewBuildTask("makeBootModes", func(w Work) error {
	return w.RunOnce(makeSysLinux)
})
