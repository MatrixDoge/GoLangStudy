/*
	Author: Parthas.Menethil
	Email:  xywhpanda@gmail.com
*/
package main
import(
	"fmt"
	"encoding/json"
)
type Dummy1 struct {
	Name string
	Age	 int
	Desc string
}
type Dummy2 struct {
	name string
	age	 int
	desc string
}
func main(){
	lTest := Dummy1{"Foo", 22, "Bar"}
	lContent, lErr := json.Marshal(lTest)
	if lErr != nil {
		fmt.Println("Err:", lErr)
	}
	fmt.Println("Json", string(lContent))
	fmt.Println("===========================")
	lTest2 := Dummy2{"Foo2", 23, "Bar2"}
	lContent2, lErr2 := json.Marshal(lTest2)
	if lErr2 != nil {
		fmt.Println("Err:", lErr2)
	}
	fmt.Println("Json", string(lContent2))
}