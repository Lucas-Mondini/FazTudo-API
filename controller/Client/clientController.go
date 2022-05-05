package clientController

import (
	"fmt"
)

func GetClient(index string) map[string]string {
	var client [2]string
	client[0] = "Lucas"
	client[1] = "Ewerton"

	re := make(map[string]string)

	for i, s := range client {
		re[fmt.Sprint("client_", i)] = s
	}

	return re
}

func GetAllClients() map[string]string {
	var client [2]string
	client[0] = "Lucas"
	client[1] = "Ewerton"

	re := make(map[string]string)

	for i, s := range client {
		re[fmt.Sprint("client_", i)] = s
	}

	return re
}
