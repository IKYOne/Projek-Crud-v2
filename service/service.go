package service

import (
	"fmt"

	"github.com/IKYOne/CRUD-Full.git/entity"
	"github.com/IKYOne/CRUD-Full.git/mahasiswaError"
	"github.com/IKYOne/CRUD-Full.git/repository"
)

// Filter Name
func filterData(name string) string {

	if name == "Anjing" || name == "Babi" {
		return "***"
	} else {
		return name
	}
}

// Filter Id
func filterId(id string) (string, mahasiswaError.ErrorMahasiswa) {
	if len(id) == 0 {
		return id, mahasiswaError.ErrorMahasiswa{
			StatusCode: 400,
			Message:    "Id tidak boleh kosong",
		}
	}

	if id == "bbb" || id == "BBB" {
		return id, mahasiswaError.ErrorMahasiswa{
			StatusCode: 422,
			Message:    "Input ID salah",
		}
	}

	return id, mahasiswaError.ErrorMahasiswa{}
}

type RepoService struct {
	repo repository.KotrakMahasiswa
}

func NewService(repo repository.KotrakMahasiswa) *RepoService {
	return &RepoService{
		repo: repo,
	}
}

func (r *RepoService) CreateService(id string, data *entity.Mahasiswa) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa) {
	fmt.Println("---CREATE---")
	Id, status := filterId(id)
	hasil, status := r.repo.Create(Id, data)
	return hasil, status
}

func (r *RepoService) ReadService(id string) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa) {
	fmt.Println("---READ---")
	hasil, status := r.repo.Read(id)
	return hasil, status
}

func (r *RepoService) UpdateService(id string, name string, nim string) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa) {
	Name := filterData(name)
	fmt.Println("---UPDATE---")
	hasil, status := r.repo.Update(id, Name, nim)
	return hasil, status
}

func (r *RepoService) DeleteService(id string) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa) {
	Id, _ := filterId(id)
	fmt.Println("---DELETE---")
	hasil, status := r.repo.Delete(Id)
	return hasil, status
}
