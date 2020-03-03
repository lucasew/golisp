package enforce

func Validate(e ...func() error) error {
	for _, f := range e {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}
