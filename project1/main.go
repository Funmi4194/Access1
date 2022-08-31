package main

import (
	"fmt"
	"log"
	s "project/project11" //import the funcword.go into the code usin go mod init project 
	"strings"
)
func main(){
	word, err := s.GetWord("word.txt") //call the function getword fron funcword package
	if err != nil { // check if there is an error calling the file

		log.Fatal(err)
	}

	var wordCount int
	//creating a for range loop 
	for _, words := range word{  //

		words = strings.TrimSpace(words) // trims the text 
		// fmt.Println(words)// prints th words
		wordCount = wordCount + len(strings.Split(words, " "))
		
	}	
// numbers of  determined
	var character int
	for  i:= 0; i<len(word); i++{
		// word[i] = strings.TrimSpace(word[i])
	character = character + len(strings.Split(word[i], ""))

	}
	//numbers of characters determined 

	var characterX int
	for  i:= 0; i<len(word); i++{
		// word[i] = strings.TrimSpace(word[i])
		newWord := strings.Replace(word[i], " ", "", -1)
		// fmt.Println(newWord)
	characterX = characterX + len(strings.Split(newWord, ""))
	}	
	fmt.Printf("%2v    %39v\n", "Words", wordCount)
	fmt.Printf("%2v    %35v\n", "Characters", character)
	fmt.Printf("%2v    %18v\n", "Characters excluding spaces", characterX)


}