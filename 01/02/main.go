package main

import (
  "encoding/hex"
  "fmt"
)

func xorBytes(a, b []byte) []byte {
  n := len(a)
  dst := make([]byte, n)

  for i := 0; i < n; i++ {
    dst[i] = a[i] ^ b[i]
  }

  return dst
}

func main() {
  str1 := "1c0111001f010100061a024b53535009181c"
  bytes1, _ := hex.DecodeString(str1)

  str2 := "686974207468652062756c6c277320657965"
  bytes2, _ := hex.DecodeString(str2)

  xoredBytes := xorBytes(bytes1, bytes2)
  hexString := hex.EncodeToString(xoredBytes)

  fmt.Print(hexString)
}
