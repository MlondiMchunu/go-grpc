// main.go

package main

import (
	"fmt"
	"log"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/koyeb/go-grpc/protogen/golang/orders"
	"github.com/koyeb/go-grpc/protogen/golang/product"
)

func main() {
	orderItem := orders.Order{
		OrderId:    10,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2021, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: product.ProductType_DRINK},
		},
	}

	bytes, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatal("deserialization error:", err)
	}

	fmt.Println(string(bytes))
}
