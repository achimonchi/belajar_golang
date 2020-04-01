package main

import "fmt"

func main() {
	// map disini seperti array assosiatif pada PHP
	// yaitu berupa data[key] = value
	// untuk membuatnya, bisa menggunakan make(map[tipe_data]tipe_data)
	x := make(map[string]string)

	x["106"] = "Reyhan Jovie"
	x["105"] = "Dwiputra"

	fmt.Println(x)

	// untuk mencetak datanya, bisa melakukan foreach
	fmt.Println()
	for i, v := range x {
		fmt.Println("NIM  ", i, " bernama ", v)
	}

	// bisa juga membuat map di dalam map.
	// hal ini akan mirip dengan json pada JS
	mahasiswa := map[string]map[string]string{
		"101": {
			"name":   "Reyhan",
			"gender": "Laki laki",
		},
		"102": {
			"name":   "Jovie",
			"gender": "Perempuan",
		},
	}

	fmt.Println()

	for i, v := range mahasiswa {
		fmt.Println("ID ", i, " bernama ", v["name"], " dan gendernya ", v["gender"])
	}

	// adapun, untuk menghapusnya menggunakan delete(map, key)
	delete(mahasiswa, "101")

	fmt.Println()
	for i, v := range mahasiswa {
		fmt.Println("ID ", i, " bernama ", v["name"], " dan gendernya ", v["gender"])
	}
}
