package objutil

import "github.com/jinzhu/copier"

func Copy(to interface{}, from interface{}) error {
	return copier.Copy(to, from)
}

func MustCopy(to interface{}, from interface{}) (rto interface{}) {
	err := copier.Copy(to, from)
	if err != nil {
		panic(err)
	}
	return to
}
