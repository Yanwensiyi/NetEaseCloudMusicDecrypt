package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "regexp"
  "strings"
)

func main() {
  /*
    var ln string
    fmt.Scanln(&ln)
    fns, err := ioutil.ReadDir(ln)
    if err != nil {
      fmt.Println("!!!")
    }
    pat, err := regexp.Compile(".+\\.uc")
    if err != nil {
      fmt.Println("!!!!!!")
    }
    for _, f := range fns {
      if f.IsDir() {
        continue
      }
      if !pat.MatchString(f.Name()) {
        continue
      }
      fmt.Println(f.Name())
    }
    for {
    }
  */
  for {
    var c string
    fmt.Println("NetEaseCloudMusicDecryptor V1.0\r\n1:Decrypt a file\r\n2:Decrypt files\r\n0:Exit")
    fmt.Scanln(&c)
    switch c {
    case "0":
      os.Exit(0)
    case "1":
      decryptFile()
    case "2":
      decryptFiles()
    default:
      fmt.Println("Invalid choice.Choose an exist item.")
    }
    fmt.Println("\n\n\n\n")
  }
}

func decryptFiles() {
  var ln, dst, flt string
  fmt.Print("Source directory:")
  fmt.Scanln(&ln)
  fns, err := ioutil.ReadDir(ln)
  if err != nil {
    fmt.Println("Invalid source directory name.")
    return
  }
  if ln[len(ln)-1] != '/' && ln[len(ln)-1] != '\\' {
    ln += "/"
  }
  fmt.Print("Destination directory(Leave blank to use default output directory):")
  fmt.Scanln(&dst)
  if dst == "" {
    dst = "Decrypted/"
  }
  _, err = ioutil.ReadDir(dst)
  if err != nil {
    err = os.Mkdir(dst, 0777)
    if err != nil {
      fmt.Println("Create output directory failed.")
      return
    }
  }
  if dst[len(dst)-1] != '/' && dst[len(dst)-1] != '\\' {
    dst += "/"
  }
  fmt.Print("Filename filter(A regular expression,leave blank to use .+\\.uc)")
  fmt.Scanln(&flt)
  if flt == "" {
    flt = ".+\\.uc"
  }
  pat, err := regexp.Compile(flt)
  if err != nil {
    fmt.Println("Invalid filter.")
    return
  }
  for _, f := range fns {
    if f.IsDir() {
      continue
    }
    if !pat.MatchString(f.Name()) {
      continue
    }
    decryptFile(ln+f.Name(), dst+f.Name()[:strings.LastIndex(f.Name(), ".")]+".mp3")
  }
}

func decryptFile(fn ...string) {
  var ln, dst string
  if len(fn) == 2 {
    ln = fn[0]
    dst = fn[1]
  } else {
    fmt.Print("Source filename:")
    fmt.Scanln(&ln)
    fmt.Print("Destination filename(Leave blank to use origin file name):")
    fmt.Scanln(&dst)
    if dst == "" {
      i := strings.LastIndex(ln, "\\")
      if i == -1 {
        i = strings.LastIndex(ln, "/")
        if i == -1 {
          dst = ln[:strings.LastIndex(ln, ".")] + ".mp3"
        } else {
          dst = ln[i+1:strings.LastIndex(ln, ".")] + ".mp3"
        }
      } else {
        dst = ln[i+1:strings.LastIndex(ln, ".")] + ".mp3"
      }
    }
    dst = "Decrypted/" + dst
  }
  d, e := ioutil.ReadFile(ln)
  if e != nil {
    fmt.Println("Invalid file.")
    return
  }
  if ioutil.WriteFile(dst, decode(d), 0777) != nil {
    err := os.Mkdir("Decrypted", 0777)
    if err != nil {
      fmt.Println("Create output directory failed.")
      return
    }
    if ioutil.WriteFile(dst, decode(d), 0777) != nil {
      fmt.Println("Write out decrypted file failed.")
      return
    }
  }
  fmt.Println("Decryption success.")
}

func decode(d []byte) []byte {
  r := make([]byte, len(d))
  for i := 0; i < len(d); i++ {
    r[i] = d[i] ^ 0xa3
  }
  return r
}
