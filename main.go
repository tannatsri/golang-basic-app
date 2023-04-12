package main

import(
	"fmt"
	"log"
	"net/http"
)



func helloHanlder(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path != "/hello") {
		http.Error(w, "404 not fount", http.StatusNotFound)
		return
	}

	if(r.Method != "GET") {
		http.Error(w, "404 not fount", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, " -- hello")
}






func formHandler(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path != "/form") {
		http.Error(w, "404 not fount", http.StatusNotFound)
		return
	}

	if(r.Method != "GET") {
		http.Error(w, "404 not fount", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, " -- form")
}



func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHanlder)


	fmt.Printf("starting server at port 9090")
	if err:= http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}