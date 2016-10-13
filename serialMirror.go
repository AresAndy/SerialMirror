package main

import (
    "log"
    "fmt"
    "serial"
    "time"
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

func writeChars(s *serial.Port, b []byte, n int, tim float32) (amount int) {
  char := make([]byte, 1)
  amount = 0

  for i := 0; i < n && b[n] != 0x0 ; i++ {
    char[0] = b[i]
    _, err := s.Write(char)

    if err == nil {
      fmt.Printf("written char: %c (hex: %X)\n", char[0], char[0])
      time.Sleep(time.Second * time.Duration(tim)) // just wait a bit before sending a new char
      amount++
    } else {
      return
    }
  }
  return
}

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600/*115200*/, ReadTimeout: time.Second * 5}
  s, err := serial.OpenPort(c)
  if err != nil {
    log.Fatal(err)
  }

	for {
		buf := make([]byte, 32)
		buf, an := readChars(s, 32)

		if an != 0 {
		  for i := 0; buf[i] != 0x0; i++ {
		  	fmt.Printf("char: %c (hex: %X; dec: %d)\n", buf[i], buf[i], buf[i])
		  }

			n := writeChars(s, buf, an, 0.25)
			if n != an {
				fmt.Println("Written chars < buff-length (length is %d, written %d)", an, n)
			}
		} else {
			fmt.Println("OK, next;")
		}

		time.Sleep(time.Second)
	}
}
