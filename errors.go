package goutil

type ErrFunc func() error

func FirstErr(fn ...ErrFunc) error {
	var err error

	for _, f := range fn {
		if err = f(); err != nil {
			break
		}
	}

	return err
}
