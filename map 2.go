package main

import "fmt"

func mergeMaps(map1, map2 map[string]string) map[string]string {
	result := make(map[string]string)

	// Copiar os elementos do primeiro mapa para o resultado
	for key, value := range map1 {
		result[key] = value
	}

	for key, value := range map2 {
		result[key] = value
	}

	return result
}

func main() {
	map1 := map[string]string{
		"chave1": "valor1",
		"chave2": "valor2",
	}

	map2 := map[string]string{
		"chave2": "novo valor2",
		"chave3": "valor3",
	}

	mergedMap := mergeMaps(map1, map2)

	for key, value := range mergedMap {
		fmt.Printf("%s: %s\n", key, value)
	}
}
