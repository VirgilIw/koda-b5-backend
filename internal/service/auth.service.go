package service

import (
	"errors"
	"strings"

	"github.com/virgilIw/koda-b5-backend/internal/dto"
)

// AuthService adalah service untuk menangani logika bisnis authentication
type AuthService struct {
	userData []dto.Register // menyimpan data user yang sudah register
}

// NewAuthService membuat instance baru dari AuthService
func NewAuthService() *AuthService {
	return &AuthService{
		userData: make([]dto.Register, 0), // inisialisasi slice kosong
	}
}

// Register melakukan proses registrasi user baru
func (s *AuthService) ServiceRegister(register dto.Register) ([]dto.Register, error) {
	// Validasi 1: Cek email kosong
	if register.Email == "" {
		return nil, errors.New("email kosong")
	}

	// Validasi 2: Cek format email harus ada @
	if !strings.Contains(register.Email, "@") {
		return nil, errors.New("email harus ada @")
	}

	// Validasi 3: Cek apakah email sudah terdaftar
	for _, user := range s.userData {
		if user.Email == register.Email {
			return nil, errors.New("email sudah terdaftar")
		}
	}

	// Validasi 4: Cek password kosong
	if register.Password == "" {
		return nil, errors.New("password kosong")
	}

	// Validasi 5: Cek panjang password minimal 6 karakter
	if len(register.Password) < 6 {
		return nil, errors.New("password minimal 6")
	}

	// Semua validasi lolos, simpan user baru
	s.userData = append(s.userData, register)
	
	// Return semua data user
	return s.userData, nil
}

// Login melakukan proses login user
func (s *AuthService) ServiceLogin(authen dto.Authentication) error {
	// Validasi 1: Cek email kosong
	if authen.Email == "" {
		return errors.New("email kosong")
	}

	// Validasi 2: Cek format email harus ada @
	if !strings.Contains(authen.Email, "@") {
		return errors.New("email harus ada @")
	}

	// Validasi 3: Cek password kosong
	if authen.Password == "" {
		return errors.New("password kosong")
	}

	// Validasi 4: Cek panjang password minimal 6 karakter
	if len(authen.Password) < 6 {
		return errors.New("password minimal 6")
	}

	// Cari user dengan email dan password yang cocok
	for _, user := range s.userData {
		if user.Email == authen.Email && user.Password == authen.Password {
			return nil // Login berhasil
		}
	}

	// Kalau tidak ketemu, berarti email/password salah
	return errors.New("email atau password salah / belum terdaftar")
}