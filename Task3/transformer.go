package main

import "strings"

func checkForQuestionMark(sentences []string, words *[][]string)  {
	for i,v:= range sentences{
		if strings.HasSuffix(v,"?"){
			for j,v:=range (*words)[i]{
				for letter:=1;letter<len(v);letter++{
					if v[0]==v[letter]{
						(*words)[i][j]=v[:letter]+v[letter+1:]
					}
				}
			}
		}
	}
}

func checkForExclamationPoint(sentences []string, words *[][]string)  {
	for i,v:= range sentences{
		if strings.HasSuffix(v,"!"){
			buf:=(*words)[i][0]
			(*words)[i][0]=(*words)[i][len((*words)[i])-1][:len((*words)[i][len((*words)[i])-1])-1]
			(*words)[i][len((*words)[i])-1]=buf+"!"
		}
	}
}
