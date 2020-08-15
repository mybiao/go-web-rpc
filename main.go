package main

func main() {

}

/*
import (
	"fmt"
	"log"
	"time"

	"github.com/mybiao/test/mess"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	body := mess.Body{
		Able: 100,
		Ch:   []int32{20, 10, 40},
		Mark: "mark",
		Time: &timestamppb.Timestamp{Seconds: time.Now().Unix()},
	}
	msg := mess.Msg{
		Id:   1000,
		Op:   "control",
		Body: &body,
	}
	buf, err := proto.Marshal(&msg)
	if err != nil {
		log.Println(err)
	}
	var mm mess.Msg
	err = proto.Unmarshal(buf, &mm)
	if err != nil {
		log.Println(err)
	}
	t := mm.Body.Time.Seconds
	ti := time.Unix(t, 0)
	fmt.Println(ti)
	fmt.Println(mm.Body.Ch)
	fmt.Println(mm.Body)
}

*/
