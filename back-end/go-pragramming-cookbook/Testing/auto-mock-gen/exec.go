package auto_mock_gen

type Controller struct {
	GetSetter
}

func (c *Controller) GetThenSet(key, value string) error {
	val, err := c.Get(key)
	if err != nil {
		return err
	}

	// Update the value
	if val != value {
		return c.Set(key, value)
	}

	return nil
}
