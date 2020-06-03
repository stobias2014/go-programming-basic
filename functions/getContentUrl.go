// Simple program that gets content type of a url

package main

import(
    "fmt";
    "net/http"
)

func getContentType(url string) (string, error) {
    resp, err := http.Get(url)

    if err != nil {
        return "", err    
    }    
    
    defer resp.Body.Close()
     
    return resp.Header.Get("Content-Type"), nil  
}

func main() {

    url := "https://www.google.com/?client=safari"

    content,err := getContentType(url)

    if err != nil {
        fmt.Errorf("ERROR: %s", err)
    } else {
        fmt.Println("Content ", content)   
    }

}
