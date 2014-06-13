package main

import "fmt"
import "strings"
import "net/http"
import "io/ioutil"

func main() {
	fmt.Println("GO Post Agent")
	var buf string
	buf = "a=1&b=2"
	resp, err := PostTo("http://xxx.xxx/go/p.php", buf, "application/x-www-form-urlencoded ")
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("----------------------------")
		fmt.Println(resp)
	}
}
func PostTo(aUrl string, aContent string, aContentType string) (string, error) {
	resp, err := http.Post("http://xparthas.com/go/p.php", aContentType, strings.NewReader(aContent))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	lRet := string(body[:])
	return lRet, nil
}
