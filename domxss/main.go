package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// URL search parameters
// location.search returns the whole string after ?
// hash fragment
var paramC = "<h1>CANARY</h1>"
var fragmentC = "MYHASHFRAGMENTCANARY\">"


func generateUrls(url string)(string, []string){
	var paramsWCanary string
	var paramsIden []string

	cleanUrl, params := splitUrl(url)

	for i, p := range(params){
		p = fmt.Sprintf("%s=%s", p, paramC + fmt.Sprint(i))
		paramsWCanary = fmt.Sprintf("%s&%s", paramsWCanary, p)
		paramsIden = append(paramsIden, p)
	}
	cleanUrl = fmt.Sprintf("%s?%s#%s", cleanUrl, paramsWCanary, fragmentC)


	return cleanUrl, paramsIden
}


// func getPaths(url string)(string){
// 	baseUrl := strings.Join(strings.Split(url, "/")[0:3], "//")
// }

func splitUrl(url string)(string, []string){
	var parameters []string 
	 
	urlParts := strings.SplitN(url, "?", 2)
	searchstring := urlParts[1]
	params := strings.Split(searchstring, "&")
	for _, param := range params{
		parameters = append(parameters, strings.SplitN(param, "=", 2)[0])
	}

	return urlParts[0], parameters
}



func main() {

	// Open the file
	file, err := os.Open(os.Args[1]) // replace 'yourfile.txt' with your file's name
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Read line by line
		
		endUrl, _ := generateUrls(scanner.Text())
	
		dom := browse(endUrl)
		fmt.Println(dom)
		if strings.Contains(dom, fragmentC){
			fmt.Printf("XSS FOUND in hash fragment, URL:%s \n", endUrl)
		}
		// for _, p := range paramsIden {
		// 	if strings.Contains(dom, strings.SplitN(p, "=", 2)[1]){
		// 		fmt.Printf("XSS FOUND in %s, URL: %s", p, endUrl)
		// 	}
		// }
		if strings.Contains(dom, paramC){
			fmt.Printf("XSS found at %s \n", endUrl)
		}
	}

	// Check for errors during Scan. End of file is expected and not reported by Scan as an error.
	if err := scanner.Err(); err != nil {
		log.Fatalf("error while reading file: %s", err)
	}
}

	// url := "https://0aa60099043dd67d80214e4500db00e1.web-security-academy.net/product?productId=6&storeId=123"


