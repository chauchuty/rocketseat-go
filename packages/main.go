package main

import (
	"encoding/json"
)

func main() {
	json.Marshal(map[string]string{"hello": "world"})
}
