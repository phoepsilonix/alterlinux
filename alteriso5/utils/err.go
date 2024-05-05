package utils

func TryAll(funcs ...func() error) error {
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
