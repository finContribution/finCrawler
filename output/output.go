package output

type Output interface {
	Convert(input chan []byte)
	Send()
}
