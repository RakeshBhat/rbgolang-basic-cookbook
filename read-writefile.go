package main

import (
	"log"
	"io/ioutil"
	"os"
)

func readFromFile(fileName string){
	data, err := ioutil.ReadFile(fileName)
	if( err != nil){
		log.Fatal(err)
	}

	log.Printf("Data read: %s \n", data)
}

func writeToFile(fileName string){
	//  open new file for writting only
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666)

	if(err != nil){
		log.Fatal(err)
	}

	defer file.Close()

	// Write bytes to file
	byteSlice := []byte("This is go reading/write file.")
	bytesWritten, err := file.Write(byteSlice)
	if(err != nil){
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}

func main() {
	fileName := "goReadWriteFile.txt"
	writeToFile(fileName)
	readFromFile(fileName)
}