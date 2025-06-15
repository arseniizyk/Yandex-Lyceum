package main

import "fmt"

func main() {
	var userFavorite string
	fmt.Scanln(&userFavorite)

	result := fmt.Sprintf("А ещё я люблю %s!", userFavorite)

	fmt.Println(result)
}
