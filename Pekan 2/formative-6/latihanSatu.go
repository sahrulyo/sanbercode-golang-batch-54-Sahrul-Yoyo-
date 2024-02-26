//poniter function

//pas by value

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeUddressToIndonesia(uddress Address) {
	uddress.Country = "Indonesia"
}
func main() {
	/*address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)*/

	//operator & and * can be used to get the memory
	//untuk merubah  nilai objek dan mengakses  properti dari sebuah objek
	/*address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2) //ada tanda *(pointer) */

	//operator *
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1 //asa 3 address yang mengacu kepada address1

	address2.City = "Bandung"
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // tambah & untuk menyimpan alamat
	//*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // untuk merubah address 1
	//untuk merubah semua memori di address 1 & 2

	fmt.Println(address1) //address 1 berubah
	fmt.Println(address2)
	fmt.Println(address3)

	//function new
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)

	//Pointer di Function -- belum selesai
	/*var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	uddress := Address{"Subang", "Jawa Barat", ""}
	ChangeUddressToIndonesia(alamat)
	fmt.Println(alamat) //tidak berubah*/

}
