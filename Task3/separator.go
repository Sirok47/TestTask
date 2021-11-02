package main

import (
	"strings"
)

func separateSentences(text string) []string {
	sentences:=make([]string,strings.Count(text,".")+strings.Count(text,"!")+strings.Count(text,"?"))
	index:=0
	sentences[0]=text
	for i:=0;;i++ {
		if strings.Count(sentences[i],".")+strings.Count(sentences[i],"!")+strings.Count(sentences[i],"?")==1{
			break
		}
		index=strings.IndexAny(sentences[i], ".!?")
		if index==-1 {
			break
		}
		sentences[i+1]=sentences[i][index+1:]
		for strings.HasPrefix(sentences[i+1]," "){
			sentences[i+1]=sentences[i+1][1:]
		}
		sentences[i]=sentences[i][:index+1]
	}
	return sentences
}

func separateWords(sentences []string) [][]string {
	words:=make([][]string,len(sentences))
	for i,v:=range sentences{
		words[i]=make([]string,strings.Count(v," ")+1)
		words[i][0]=sentences[i][:len(v)]
		for j,v:=range words[i]{
			index := strings.Index(v, " ")
			if index == -1 {
				break
			}
			words[i][j+1] = words[i][j][index+1:]
			words[i][j]=words[i][j][:index]
		}
	}
	return words
}

