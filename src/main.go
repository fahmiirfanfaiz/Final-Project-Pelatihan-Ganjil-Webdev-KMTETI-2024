package main

import (
"fmt"
"net/http")

func main(){
	s := &http.Server{
		Addr: ":8080",
	}

	fmt.Println("HTTP Server running on port 8080")
	err := s.ListenAndServe()
	if err != nil{
		fmt.Println(err.Error()) 
	}

}