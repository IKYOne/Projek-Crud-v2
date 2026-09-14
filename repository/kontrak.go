package repository

import (
	"github.com/IKYOne/CRUD-Full.git/entity"
	"github.com/IKYOne/CRUD-Full.git/mahasiswaError"
)

type KotrakMahasiswa interface {
	Create(id string, data *entity.Mahasiswa) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa)
	Read(id string) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa)
	Update(id string, name string, nim string) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa)
	Delete(id string) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa)
}
