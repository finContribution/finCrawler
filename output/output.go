package output

type Output interface {
	Convert(chan []byte)
	Send()
}
