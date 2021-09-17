package main

import  ("path")

func main(){

	myPath := "main/main.css"

	// this method return two value, diectory name and file name
	dir, file := path.Split(myPath)

	// when we dont want get directory value, only file name
	// _, file := path.Split(myPath)

	println(dir, file);

}