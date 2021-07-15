package other

type T struct {
	S []byte
}

func Copy(s []byte) T {
	return T{S: s}
}
