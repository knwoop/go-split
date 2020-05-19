package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func Split(r io.Reader, hasHeader bool, partSize1, partSize2 int) (io.Reader, io.Reader, error) {
	scanner := bufio.NewScanner(r)
	var b1, b2 []byte
	if hasHeader {
		if !scanner.Scan() {
			return nil, nil, fmt.Errorf("error no data")
		}
		l := scanner.Bytes()
		b1 = append(append(b1, l...), '\n')
		b2 = append(append(b2, l...), '\n')
	}

	dst1 := 0
	dst2 := 0
	for {
		if !scanner.Scan() {
			break
		}
		l := scanner.Bytes()
		if dst1 < partSize1 {
			b1 = append(append(b1, l...), '\n')
			dst1++
		} else {
			b2 = append(append(b2, l...), '\n')
			dst2++
		}
	}

	total := dst1 + dst2
	if (partSize1 + partSize2) != total {
		const msg = "error row number(%d) and args total number of partSize1(%d) and partSize2(%d) are not equal"
		return nil, nil, fmt.Errorf(msg, total, partSize1, partSize2)
	}

	return bytes.NewReader(b1), bytes.NewReader(b2), nil
}
