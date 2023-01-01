package main

import (
	"fmt"
	// "lvd/test/packs_singleton"
	"lvd/test/builder"
	// "time"
)

func main() {

	// // Test for Singleton
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		fmt.Printf("%p\n", packs_singleton.GetInstance())
	// 	}()
	// }
	// time.Sleep(time.Second * 10)

	// Test for Builder
	norBuilder := builder.GetBuilder("normal")
	iglBuilder := builder.GetBuilder("igloo")

	dr := builder.NewDirector(norBuilder)
	norHouse := dr.BuildHouse()

	fmt.Printf("normal House, window: %s\n", norHouse.GetWindowType())
	fmt.Printf("normal House, floors: %d\n", norHouse.GetNumFoor())
	
	dr.SetBuilder(iglBuilder)
	igHouse := dr.BuildHouse()
	fmt.Printf("Igloo House, window: %s\n", igHouse.GetWindowType())
	fmt.Printf("Igloo House, floors: %d\n", igHouse.GetNumFoor())
}
