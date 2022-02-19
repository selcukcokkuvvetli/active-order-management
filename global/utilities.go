package global

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func BoolToPSQLBit(value bool) string {
	if value {
		return "1"
	} else {
		return "0"
	}
}