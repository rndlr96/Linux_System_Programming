package main

import (
	"fmt"
	"io/ioutil"
	"os"
//  "sync"
  "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeAns(Answer_Dir string, ansMap map[string][]string){
  files, err := ioutil.ReadDir(Answer_Dir)
  check(err)
  for _, f := range files {
    fd, _ := ioutil.ReadDir("./"+Answer_Dir+"/"+f.Name())

    if strings.Split(fd[0].Name(), ".")[1] == "txt" { // txt file
      file, _ := ioutil.ReadFile("./"+Answer_Dir+"/"+f.Name()+"/"+f.Name()+".txt")
      for _, val := range strings.Split(string(file), " : ") {
        ansMap[fd[0].Name()] = append(ansMap[fd[0].Name()], val)
      }
    }
  }
}

func main() {
	Answer_Dir := "./" + os.Args[1]
	//	Question_Dir := "./" + os.Args[2]

  var ansMap = make(map[string][]string)
  makeAns(Answer_Dir, ansMap)



  fmt.Println(ansMap["9-5.txt"])
}
