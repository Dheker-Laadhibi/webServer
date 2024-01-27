package main
import f "fmt"
import "net/http"

//http.ResponseWrite sends data  to the http client
// an http request is a  data-structure that represents the response a client http rerquest 

func handler (w http.ResponseWriter, r *http.Request){
	f.Fprintf(w,"hello %s !!" , r.URL.Path[1:])
	//r.url.path is the path compoonent of the requested url  in this case is "/world"  of http://localhost:8080/world
}

func main() {

http.HandleFunc("/", handler)
http.ListenAndServe(":8080" , nil)

}




/*
In simpler terms,
 think of ResponseWriter as the tool you use to build and send the response back to the client, 
while http.Request is the container that holds all the information about the incoming request that your server needs to process and respond to. 
They work together to handle the communication between your Go server and clients making HTTP requests.*/