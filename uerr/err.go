package uerr

type ErrFunc func() error

// First returns the first error in the chain and stops the chain execution if the one occurs.
func First(fn ...ErrFunc) error {
	var err error

	for _, f := range fn {
		if err = f(); err != nil {
			break
		}
	}

	return err
}
