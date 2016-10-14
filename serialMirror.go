package main

import (
    "log"
    io "fmt"
    "serial"
    t "time"
)

var (
	printf   = io.Printf // Closures for refactored code
	printfln = io.Println
	fatal    = log.Fatal
)

func readChars(s *serial.Port, n int) (buff []byte, amount int) {
  buff = make([]byte, n)
  char := make([]byte, n) // temp "buffer"
  amount = 0

  for i := 0; i < n; i++ {
    n, stat := s.Read(char)
    if stat != nil {
    return
    } else if n > 1 {
      buff = char
      amount = n
      return
    } else {
      buff[i] = char[0]
      amount++
    }
  }
  return
}

func writeChars(s *serial.Port, b []byte, n, tim int) (amount int) {
  char := make([]byte, 1)
  amount = 0

  for i := 0; i < n - 1 ; i++ {
    char[0] = b[i]
    _, err := s.Write(char)

    if err == nil {
      printf("written char: %c (hex: %X)\n", char[0], char[0])
      t.Sleep(t.Millisecond * t.Duration(tim)) // wait after send
      amount++
    } else {
      return
    }
  }
  return
}

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600, ReadTimeout: t.Second * 5}
  s, err := serial.OpenPort(c)
  if err != nil {
    fatal(err)
  }

	for {
		buf := make([]byte, 32)
		buf, an := readChars(s, 32)

		if an != 0 {
		  for i := 0; buf[i] != 0x0; i++ {
		  	printf("char: %c (hex: %X; dec: %d)\n", buf[i], buf[i], buf[i])
		  }

			n := writeChars(s, buf, an, 250)
			if n != (an - 1) {
				printf("Written chars < buff-length (length is %d, written %d)\n", an, n)
			}
		} else {
			println("OK, next;")
		}

		t.Sleep(t.Duration(250) * t.Millisecond)
	}
}
