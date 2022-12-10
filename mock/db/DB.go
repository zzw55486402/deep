package db

type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}
	return -1
}

/*
	如果 GetFromDB() 方法长这个样子
	func GetFromDB(key string) int {
		db := NewDB()
		if value, err := db.Get(key); err == nil {
			return value
		}

		return -1
	}
	对 DB 接口的 mock 并不能作用于 GetFromDB() 内部，这样写是没办法进行测试的。
	那如果将接口 db DB 通过参数传递到 GetFromDB()，那么就可以轻而易举地传入 Mock 对象了。
*/
