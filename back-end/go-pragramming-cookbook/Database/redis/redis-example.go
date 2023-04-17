package redis

func RedisExample() error {
	client, err := Setup()
	if err != nil {
		return nil
	}

	key := "single"
	val := "single-value"
	SaveSingleData(client, key, val)
	if err := QuerySingleValue(client, key); err != nil {
		return err
	}

	listKey := "list"
	listVal := []int64{5, 2, 3, -1}
	if err := SaveListData(client, listKey, listVal); err != nil {
		return err
	}
	if err := QueryListValueBySort(client, listKey); err != nil {
		return err
	}

	return nil
}
