package goutil

type ErrFunc func() error

// FirstErr returns the first error in the chain and stops the chain execution if the one occurs.
func FirstErr(fn ...ErrFunc) error {
	var err error

	for _, f := range fn {
		if err = f(); err != nil {
			break
		}
	}

	return err
}
