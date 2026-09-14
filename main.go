package main

import (
	"fmt"

	"github.com/IKYOne/CRUD-Full.git/entity"
	"github.com/IKYOne/CRUD-Full.git/repository"
	"github.com/IKYOne/CRUD-Full.git/service"
)

func DataMahasiswa(daftar []entity.Mahasiswa, filData func(entity.Mahasiswa) bool) []entity.Mahasiswa {
	hasil := []entity.Mahasiswa{}
	for _, ada := range daftar {
		if filData(ada) == true {
			hasil = append(hasil, ada)
		}
	}

	return hasil
}

func main() {
	memRepo := repository.NewMemoryRepository()
	Service := service.NewService(memRepo)

	user := entity.Mahasiswa{
		Nama: "zikril",
		Nim:  "22415013",
	}

	fmt.Println(Service.CreateService("B01", &user))
	fmt.Println(Service.ReadService("B01"))

	user1 := []entity.Mahasiswa{
		{Nama: "ikyone", Nim: "20000000"},
	}

	hasil := DataMahasiswa(user1, func(m entity.Mahasiswa) bool {
		return m.Nama == "ikyone"
	})

	fmt.Println(hasil)
	fmt.Println(Service.UpdateService("B01", "Zikril", "23012340"))
	fmt.Println(Service.DeleteService("B01"))
}
