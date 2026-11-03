package framer

import (
	"bytes"
	"strings"
)


type Framer struct{
	Buffer []byte
}

func Newframer() *Framer {
	return &Framer{
		Buffer: []byte{},
	}
}


func (f *Framer) Feed(data []byte) []string {
	f.Buffer = append(f.Buffer, data...)
	validMessages := []string{}
	for {
		indexOfDelimiter := bytes.IndexByte(f.Buffer,'$')
		if indexOfDelimiter == -1 {
			break
		}
		validMessage := strings.TrimSpace(string(f.Buffer[:indexOfDelimiter]))
		
		validMessages = append(validMessages, validMessage)


		f.Buffer = f.Buffer[indexOfDelimiter+1:]


	}

	return validMessages

}