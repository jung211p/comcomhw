/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START container_hello_app]
package main

import (
	"strconv"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}



// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	defer func() {
		s:= recover()
		fmt.Println(s)
	}()

	key, _ := r.URL.Query()["input"]
	
	value := key[0]

	
	var max int
	max = 1

	input, err := strconv.Atoi(string(value))

	if err != nil {
		fmt.Fprintf(w, "Input number error.\n")
		return
	}

	if input == 0 || input == 1 || input < 0 {
		fmt.Fprintf(w, "Error. input is %d.\n", input)
		return
	}


	for i:= 2; i<= input; i++ {
		for input % i == 0 {
			input /= i
			if max < i {
				max = i
			}
		}
	}

	fmt.Fprintf(w, "%d\n", max)
	
	

}

// [END container_hello_app]
