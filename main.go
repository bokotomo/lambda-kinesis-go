package main
 
import (
  "fmt"
  "github.com/aws/aws-lambda-go/lambda"
)
 
type MyEvent struct {
  Name string `json:"name"`
}
 
type MyResponse struct {
  Message string `json:"sucess"`
}
 
func hello(event MyEvent) (MyResponse, error) {
  return MyResponse{Message: fmt.Sprintf("name is %s!!", event.Name)}, nil
}
 
func main() {
  lambda.Start(hello)
}