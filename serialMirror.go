package main

import (
    "log"
    "fmt"
    "serial"
    "time"
)

func main() {
    c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600/*115200*/, ReadTimeout: time.Second * 5}
    s, err := serial.OpenPort(c)
    if err != nil {
        log.Fatal(err)
    }
    /*
    n, err := s.Write([]byte("HELLO"))
    if err != nil {
      log.Fatal(err)
    }
  */

    buf := make([]byte, 32)
    char := make([]byte, 1)
    i := 0

    char[0] = '#' // something different from 0x0

    for char[0] != 0x0 {
        _, err = s.Read(char)

        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("char: %c (hex: %X)\n", char[0], char[0])
        buf[i] = char[0]
        i++
    }

    char[0] = '#'

    for i = 0; char[0] != 0x0 ; i++{
        char[0] = buf[i]
        _, err = s.Write(char)

        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("written char: %c (hex: %X)\n", char[0], char[0])
    }
}
