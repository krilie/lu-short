package sqlutil

func Like(str string) string {
	return "%" + str + "%"
}
