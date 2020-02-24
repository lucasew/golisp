package enforce

func Validate(e ...error) error {
	for _, err := range e {
		if err != nil {
			return err
		}
	}
	return nil
}
