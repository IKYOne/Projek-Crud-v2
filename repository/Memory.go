package repository

import (
	"github.com/IKYOne/CRUD-Full.git/entity"
	"github.com/IKYOne/CRUD-Full.git/mahasiswaError"
)

type MemoryRepository struct {
	db map[string]*entity.Mahasiswa
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		db: make(map[string]*entity.Mahasiswa),
	}
}

func (m *MemoryRepository) Create(id string, data *entity.Mahasiswa) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa) {
	if len(id) == 0 {
		return nil, mahasiswaError.ErrorMahasiswa{
			StatusCode: 400,
			Message:    "ID Wajib Di isi",
		}
	}

	m.db[id] = data

	return m.db[id], mahasiswaError.ErrorMahasiswa{}
}

func (m *MemoryRepository) Read(id string) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa) {
	hasil, ada := m.db[id]

	if !ada {
		return nil, mahasiswaError.ErrorMahasiswa{
			StatusCode: 404,
			Message:    "Gagal: NOT FOUND",
		}

	}

	return hasil, mahasiswaError.ErrorMahasiswa{}

}

func (m *MemoryRepository) Update(id string, name string, nim string) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa) {
	if len(name) == 0 || len(nim) == 0 {
		return nil, mahasiswaError.ErrorMahasiswa{
			StatusCode: 400,
			Message:    "Data wajib Di Isi",
		}
	}

	hasil, data := m.db[id]

	if !data {
		return nil, mahasiswaError.ErrorMahasiswa{
			StatusCode: 404,
			Message:    "Gagal: NOT FOUND",
		}
	}

	hasil.Nama = name
	hasil.Nim = nim

	return hasil, mahasiswaError.ErrorMahasiswa{}

}

func (m *MemoryRepository) Delete(id string) (*entity.Mahasiswa, mahasiswaError.ErrorMahasiswa) {
	_, ada := m.db[id]
	if !ada {
		return nil, mahasiswaError.ErrorMahasiswa{
			StatusCode: 404,
			Message:    "Gagal: NOT FOUND",
		}

	}

	delete(m.db, id)
	return nil, mahasiswaError.ErrorMahasiswa{}
}
