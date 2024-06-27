package main

import (
	"fmt"
	"github.com/misfj/jrpc_study.git/pb"
)

// int64 id = 1 [(validate.rules).int64 = {gt: 0}];
// int32 age = 2 [(validate.rules).int32 = {gt:0, lt: 120}];
// uint32 code = 3 [(validate.rules).uint32 = {in: [1,2,3]}];
// float score = 4 [(validate.rules).float = {not_in: [0, 99.99]}];
// bool state = 5 [(validate.rules).bool.const = true];
// string path = 6 [(validate.rules).string.const = "/hello"];
// string phone = 7 [(validate.rules).string.len = 11];
// string explain = 8 [(validate.rules).string.min_len = 3];
// string name = 9 [(validate.rules).string = {min_len: 1, max_len: 10}];
// string card = 10 [(validate.rules).string.pattern = "(?i)^[0-9a-f]+$"];
func main() {
	fmt.Println("hello world")
	r := pb.Request{
		Id:      1,
		Age:     121,
		Code:    0,
		Score:   0,
		State:   false,
		Path:    "",
		Phone:   "",
		Explain: "",
		Name:    "",
		Card:    "",
	}
	err := r.Validate()
	fmt.Println(err)
}
