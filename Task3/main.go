package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main(){
	var text string
	filePath,exists:= os.LookupEnv("FILE_PATH")
	if !exists{
		log.Fatal("No FILE_PATH in .env")
	}
	file, err := os.Open(filePath)
	if err != nil{
		log.Fatal("No file with such name")
	}
	defer file.Close()
	data := make([]byte, 100000)
	for{
		n, err := file.Read(data)
		if err == io.EOF{
			break
		}
		text=string(data[:n])
	}

	sentences:=separateSentences(text)
	words:=separateWords(sentences)

	checkForQuestionMark(sentences,&words)

	checkForExclamationPoint(sentences,&words)

	var newText string
	for _,v:=range words{
		for _,v:=range v{
			newText+=v+" "
		}
	}
	newPath,exists:=os.LookupEnv("NEW_FILE_PATH")
	if !exists{
		log.Fatal("No NEW_FILE_PATH in .env")
	}
	newFile,err:=os.Create(newPath)
	if err != nil {
		log.Fatal("Unable to create file")
	}
	newFile.WriteString(newText)
}
