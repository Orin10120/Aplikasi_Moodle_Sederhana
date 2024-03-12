package main

import "fmt"

const NMAX = 250

type user [NMAX]struct {
	guru  konten
	siswa konten
}

type konten [NMAX]struct {
	tugas    string
	jwbTugas string
	nilTugas int
	quiz     string
	nilQuiz  int
	jwbQuiz  string
	name     string
	pass     string
}

type isiForum [NMAX]struct {
	pertanyaan string
	jawaban    string
}

func registrasi(arrUser *user, a, b *int) {
	var tipe int
	var name, pass string

	fmt.Println("---------------------")
	fmt.Println("Menu Registrasi")
	fmt.Println("Pilih tipe user Registrasi : ")
	fmt.Println("1. Guru")
	fmt.Println("2. Siswa")
	fmt.Println("3. Kembali")
	fmt.Println("---------------------")
	fmt.Scan(&tipe)

	for tipe != 3 {
		fmt.Println("Masukkan Username dan Password anda : ")

		if tipe == 1 {
			fmt.Scan(&name, &pass)
			arrUser[0].guru[*a].name = name
			arrUser[0].guru[*a].pass = pass
			*a++
		} else if tipe == 2 {
			fmt.Scan(&name, &pass)
			arrUser[0].siswa[*b].name = name
			arrUser[0].siswa[*b].pass = pass
			*b++
		}

		fmt.Println("Username dan Password anda berhasil disimpan")

		tipe = 3
	}
}

func login(arrForum *isiForum, arrUser *user, g, s int, c, d, e *int) {
	var tipe, idSiswa int
	var ketemu = false
	var username, password string

	fmt.Println("---------------------")
	fmt.Println("Menu Login")
	fmt.Println("Pilih tipe user Login : ")
	fmt.Println("1. Guru")
	fmt.Println("2. Siswa")
	fmt.Println("3. Kembali")
	fmt.Println("---------------------")
	fmt.Scan(&tipe)

	for tipe != 3 {
		fmt.Println("Masukkan Username dan Password anda : ")
		//sekuensial search
		if tipe == 1 {
			fmt.Scan(&username, &password)
			for i := 0; i < g && !ketemu; i++ {
				if username == arrUser[0].guru[i].name && password == arrUser[0].guru[i].pass {
					ketemu = true
				}
			}
			if ketemu {
				fmt.Println("Login berhasil")
				menuGuru(arrForum, arrUser, g, s, idSiswa, c, d, e)
			} else {
				fmt.Println("Username atau password salah")
			}
			//sekuensial search
		} else if tipe == 2 {
			fmt.Scan(&username, &password)
			for i := 0; i < s && !ketemu; i++ {
				if username == arrUser[0].siswa[i].name && password == arrUser[0].siswa[i].pass {
					ketemu = true
					idSiswa = i
				}
			}
			if ketemu {
				fmt.Println("Login berhasil")
				menuSiswa(arrForum, arrUser, c, d, e, idSiswa, s)
			} else {
				fmt.Println("Username atau password salah")
			}

		} else {
			fmt.Println("User tidak ditemukan")
		}
		tipe = 3
	}
}

func menuGuru(arrForum *isiForum, arrUser *user, g, s, idSiswa int, c, d, e *int) {
	var input int

	fmt.Println("---------------------")
	fmt.Println("Menu Guru")
	fmt.Println("1. Menu Tugas")
	fmt.Println("2. Menu Quiz")
	fmt.Println("3. Buka Forum")
	fmt.Println("4. Penilaian")
	fmt.Println("5. Kembali")
	fmt.Println("---------------------")
	fmt.Scan(&input)

	for input != 5 {
		if input == 1 {
			menuTugas(arrUser, d)
		} else if input == 2 {
			menuQuiz(arrUser, c)
		} else if input == 3 {
			forum(arrForum, e)
		} else if input == 4 {
			penilaian(arrUser, g, s, idSiswa, d)
		} else {
			fmt.Println("Input anda tidak valid")
		}

		fmt.Println("---------------------")
		fmt.Println("Menu Guru")
		fmt.Println("1. Menu Tugas")
		fmt.Println("2. Menu Quiz")
		fmt.Println("3. Buka Forum")
		fmt.Println("4. Penilaian")
		fmt.Println("5. Kembali")
		fmt.Println("---------------------")
		fmt.Scan(&input)
	}
}

func menuTugas(arrUser *user, n *int) {
	var input int

	fmt.Println("---------------------")
	fmt.Println("Menu Tugas")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Tugas")
	fmt.Println("3. Delete Tugas")
	fmt.Println("4. Edit Tugas")
	fmt.Println("5. Kembali")
	fmt.Println("---------------------")

	fmt.Scan(&input)

	for input != 5 {
		if input == 1 {
			tambahTugas(arrUser, n)
		} else if input == 2 {
			lihatTugas(arrUser, n)
		} else if input == 3 {
			deleteTugas(arrUser, n)
		} else if input == 4 {
			editTugas(arrUser, n)
		} else {
			fmt.Println("Input anda tidak valid")
		}

		fmt.Println("---------------------")
		fmt.Println("Menu Tugas")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Lihat Tugas")
		fmt.Println("3. Delete Tugas")
		fmt.Println("4. Edit Tugas")
		fmt.Println("5. Kembali")
		fmt.Println("---------------------")
		fmt.Scan(&input)
	}
}

func menuQuiz(arrUser *user, n *int) {
	var input int

	fmt.Println("---------------------")
	fmt.Println("Menu Quiz")
	fmt.Println("1. Tambah Quiz")
	fmt.Println("2. Lihat Quiz")
	fmt.Println("3. Delete Quiz")
	fmt.Println("4. Edit Quiz")
	fmt.Println("5. Kembali")
	fmt.Println("---------------------")
	fmt.Scan(&input)

	for input != 5 {
		if input == 1 {
			tambahQuiz(arrUser, n)
		} else if input == 2 {
			lihatQuiz(arrUser, n)
		} else if input == 3 {
			deleteQuiz(arrUser, n)
		} else if input == 4 {
			editQuiz(arrUser, n)
		}

		fmt.Println("---------------------")
		fmt.Println("Menu Quiz")
		fmt.Println("1. Tambah Quiz")
		fmt.Println("2. Lihat Quiz")
		fmt.Println("3. Delete Quiz")
		fmt.Println("4. Edit Quiz")
		fmt.Println("5. Kembali")
		fmt.Println("---------------------")
		fmt.Scan(&input)
	}
}

func lihatTugas(A *user, n *int) {
	var count = 1

	fmt.Println("---------------------")
	fmt.Println("Soal - soal Tugas : ")
	for i := 0; i < *n; i++ {
		fmt.Println(count, " ", A[i].guru[i].tugas)
		count++
	}
	fmt.Println("---------------------")
}

func tambahTugas(A *user, n *int) {
	var soal string

	fmt.Println("---------------------")
	fmt.Println("Silahkan tambahkan soal tugas, untuk berhenti silahkan input 'STOP' ")
	fmt.Println("(Gunakan '_' sebagai pengganti spasi)")
	fmt.Println("Soal Tugas : ")
	fmt.Scan(&soal)
	for soal != "STOP" && *n < NMAX {
		A[*n].guru[*n].tugas = soal
		*n++
		fmt.Println("Soal Tugas : ")
		fmt.Scan(&soal)
	}

	fmt.Println("Soal berhasil disimpan")
	fmt.Println("---------------------")
}

func editTugas(A *user, n *int) {
	var index int
	fmt.Println("---------------------")
	fmt.Println("Edit Tugas")
	fmt.Println("Masukkan nomor index tugas yang ingin diedit:")
	fmt.Scan(&index)

	index -= 1

	if index >= 0 && index < *n {
		var soalBaru string
		fmt.Println("Masukkan soal tugas yang baru:")
		fmt.Scan(&soalBaru)

		A[index].guru[index].tugas = soalBaru

		fmt.Println("Tugas berhasil diubah")
	} else {
		fmt.Println("Nomor index tugas tidak valid")
	}

	fmt.Println("---------------------")
}

func deleteTugas(A *user, n *int) {
	var index int
	fmt.Println("---------------------")
	fmt.Println("Delete Tugas")
	fmt.Println("Masukkan nomor index tugas yang ingin dihapus:")
	fmt.Scan(&index)

	index -= 1

	if index >= 0 && index < *n {
		for i := index; i < *n-1; i++ {
			A[i].guru[i].tugas = A[i+1].guru[i+1].tugas
		}
		*n--

		fmt.Println("Tugas berhasil dihapus.")
	} else {
		fmt.Println("Nomor index tugas tidak valid.")
	}

	fmt.Println("---------------------")
}

func editQuiz(A *user, n *int) {
	var index int
	fmt.Println("---------------------")
	fmt.Println("Edit Quiz")
	fmt.Println("Masukkan nomor index Quiz yang ingin diedit:")
	fmt.Scan(&index)

	index -= 1

	if index >= 0 && index < *n {
		var quizBaru string
		var jawaban string

		fmt.Println("Masukkan soal Quiz yang baru:")
		fmt.Scan(&quizBaru)
		fmt.Println("Masukkan jawaban Quiz yang baru:")
		fmt.Scan(&jawaban)

		A[index].guru[index].quiz = quizBaru
		A[index].guru[index].jwbQuiz = jawaban

		fmt.Println("Quiz berhasil diubah.")
	} else {
		fmt.Println("Nomor index Quiz tidak valid.")
	}

	fmt.Println("---------------------")
}

func deleteQuiz(A *user, n *int) {
	var index int
	fmt.Println("---------------------")
	fmt.Println("Delete Quiz")
	fmt.Println("Masukkan nomor index Quiz yang ingin dihapus:")
	fmt.Scan(&index)

	index -= 1

	if index >= 0 && index < *n {
		for i := index; i < *n-1; i++ {
			A[i].guru[i].quiz = A[i+1].guru[i+1].quiz
			A[i].guru[i].jwbQuiz = A[i+1].guru[i+1].jwbQuiz
		}
		*n--

		fmt.Println("Tugas berhasil dihapus.")
	} else {
		fmt.Println("Nomor index tugas tidak valid.")
	}

	fmt.Println("---------------------")
}

func tambahQuiz(A *user, n *int) {
	var quiz string
	var jawaban string

	fmt.Println("---------------------")
	fmt.Println("Silahkan tambahkan soal Quiz, untuk berhenti silahkan input 'STOP' ")
	fmt.Println("Soal Quiz : ")
	fmt.Scan(&quiz)
	fmt.Println("Masukan jawaban pertanyaan tersebut : ")
	fmt.Scan(&jawaban)

	for quiz != "STOP" && *n < NMAX {
		A[*n].guru[*n].quiz = quiz
		A[*n].guru[*n].jwbQuiz = jawaban
		*n++
		fmt.Println("Soal Quiz : ")
		fmt.Scan(&quiz)
		fmt.Println("Masukan jawaban pertanyaan tersebut : ")
		fmt.Scan(&jawaban)
	}

	fmt.Println("Soal berhasil disimpan")
	fmt.Println("---------------------")
}

func lihatQuiz(A *user, n *int) {
	var count = 1

	fmt.Println("---------------------")
	fmt.Println("Soal - soal Quiz : ")
	for i := 0; i < *n; i++ {
		fmt.Println(count, " ", A[i].guru[i].quiz)
		fmt.Println("jawaban : ", A[i].guru[i].jwbQuiz)
		count++
	}
	fmt.Println("---------------------")
}

func menuSiswa(arrForum *isiForum, arrUser *user, t, q, f *int, idSiswa, s int) {
	var input int
	count := 1

	fmt.Println("---------------------")
	fmt.Println("Menu Siswa")
	fmt.Println("1. Mulai Tugas")
	fmt.Println("2. Mulai Quiz")
	fmt.Println("3. Buka Forum")
	fmt.Println("4. Lihat Nilai")
	fmt.Println("5. Kembali")
	fmt.Println("---------------------")
	fmt.Scan(&input)

	for input != 5 {
		if input == 1 {
			mTugas(arrUser, t, idSiswa)
		} else if input == 2 {
			mQuiz(arrUser, q, idSiswa)
		} else if input == 3 {
			forum(arrForum, f)
		} else if input == 4 {
			fmt.Println("Nilai - nilai Tugas anda : ")
			for i := 0; i < *t; i++ {
				fmt.Println(count, " ", arrUser[i].guru[i].tugas)
				fmt.Println("Nilai anda :", arrUser[i].siswa[i].nilTugas)
				count++
			}
			fmt.Println("Nilai - nilai Quiz anda : ", arrUser[0].siswa[idSiswa].nilQuiz)
			// Insertion Sort
			for j := 1; j < s; j++ {
				temp := arrUser[0].siswa[j]
				x := j
				for x > 0 && arrUser[0].siswa[x-1].nilQuiz < temp.nilQuiz {
					arrUser[0].siswa[x] = arrUser[0].siswa[x-1]
					x--
				}
				arrUser[0].siswa[x] = temp
			}
			fmt.Println("Ranking berdasarkan nilai quiz")
			for i := 0; i < s; i++ {
				fmt.Println(arrUser[0].siswa[i].name, arrUser[0].siswa[i].nilQuiz)
			}
		} else {
			fmt.Println("Input anda tidak valid")
		}

		fmt.Println("---------------------")
		fmt.Println("Menu Siswa")
		fmt.Println("1. Mulai Tugas")
		fmt.Println("2. Mulai Quiz")
		fmt.Println("3. Buka Forum")
		fmt.Println("4. Lihat Nilai")
		fmt.Println("5. Kembali")
		fmt.Println("---------------------")
		fmt.Scan(&input)
	}
}

func mTugas(A *user, n *int, idSiswa int) {
	var jawaban string
	var index int
	count := 1

	fmt.Println("---------------------")
	fmt.Println("Selamat mengerjakan :) ")

	for i := 0; i < *n; i++ {
		fmt.Println(count, " ", A[i].guru[i].tugas)
		count++
	}
	fmt.Println("Pilih Tugas mana yang ingin anda jawab : ")
	fmt.Println("(Jika sudah silahkan input '0')")
	fmt.Println("(Gunakan '_' sebagai pengganti spasi)")
	fmt.Scan(&index)
	index--

	for index != -1 {
		fmt.Println("Masukkan jawaban anda : ")
		fmt.Scan(&jawaban)
		A[index].siswa[idSiswa].jwbTugas = jawaban
		fmt.Println("Pilih Tugas mana yang ingin anda jawab : ")
		fmt.Println("(Jika sudah silahkan input '0')")
		fmt.Println("(Gunakan '_' sebagai pengganti spasi)")
		fmt.Scan(&index)
		index--
	}
	fmt.Println("---------------------")
}

func mQuiz(A *user, n *int, idSiswa int) {
	var jawaban string
	var nilai int

	fmt.Println("---------------------")
	fmt.Println("Selamat mengerjakan ")
	for i := 0; i < *n; i++ {
		fmt.Println(i+1, " ", A[i].guru[i].quiz)
		fmt.Println("Masukan jawaban anda : ")
		fmt.Scan(&jawaban)
		nilai++
		if jawaban == A[i].guru[i].jwbQuiz {
			nilai += 1
		}
		A[0].siswa[idSiswa].nilQuiz = nilai
	}
	fmt.Println("---------------------")
	fmt.Println("Nilai anda adalah : ", nilai, "benar dari", *n, "soal")
	fmt.Println("---------------------")
}

func forum(A *isiForum, n *int) {
	var input int

	fmt.Println("---------------------")
	for i := 0; i < *n; i++ {
		fmt.Println(i+1, A[i].pertanyaan)
		fmt.Println("jawaban forum : ")
		fmt.Println(A[i].jawaban)
		fmt.Println(" ")
		fmt.Println("---------------------")
	}
	fmt.Println("---------------------")
	fmt.Println("Menu Forum : ")
	fmt.Println("1. Menambah pertanyaan")
	fmt.Println("2. Menambah jawaban")
	fmt.Println("3. Kembali")
	fmt.Println("---------------------")
	fmt.Scan(&input)

	for input != 3 {
		if input == 1 {
			tanyaForum(A, n)
		} else if input == 2 {
			jawabForum(A, n)
		} else {
			fmt.Print("Masukan anda tidak valid")
		}

		fmt.Println("---------------------")
		for i := 0; i < *n; i++ {
			fmt.Println(i+1, A[i].pertanyaan)
			fmt.Println("jawaban forum : ")
			fmt.Println(A[i].jawaban)
			fmt.Println(" ")
			fmt.Println("---------------------")
		}
		fmt.Println("---------------------")
		fmt.Println("Menu Forum : ")
		fmt.Println("1. Menambah pertanyaan")
		fmt.Println("2. Menambah jawaban")
		fmt.Println("3. Kembali")
		fmt.Println("---------------------")
		fmt.Scan(&input)
	}
}

func tanyaForum(A *isiForum, n *int) {
	var tanya string

	fmt.Println("Masukkan pertanyaan anda : ")
	fmt.Scan(&tanya)
	A[*n].pertanyaan = tanya
	*n++
}

func jawabForum(A *isiForum, n *int) {
	var jawab string
	var index int

	fmt.Println("Pilih index pertanyaan yang ingin dijawab : ")
	fmt.Scan(&index)

	index -= 1
	if index >= 0 && index < *n {
		fmt.Println("Masukkan jawaban pertanyaan tersebut :")
		fmt.Scan(&jawab)

		A[index].jawaban = jawab

		fmt.Println("Jawaban berhasil ditambah")
	} else {
		fmt.Println("Nomor index Forum tidak valid")
	}

}

func penilaian(T *user, g, s, idSiswa int, t *int) {
	var input, index, nilai int
	count := 1

	fmt.Println("---------------------")
	fmt.Println("Menu Penilaian")
	fmt.Println("1. Nilai Tugas")
	fmt.Println("2. Nilai Quiz")
	fmt.Println("3. Kembali")
	fmt.Println("---------------------")
	fmt.Scan(&input)

	for input != 3 {
		if input == 1 {
			for i := 0; i < s; i++ {
				fmt.Println(count, " ", T[0].siswa[i].name)
				count++
			}
			fmt.Println("Pilih siswa mana yang ingin dinilai : ")
			fmt.Scan(&index)
			index--
			for j := 0; j < *t; j++ {
				fmt.Println("Soal tugas : ")
				fmt.Println(T[index].guru[index].tugas)
				fmt.Println(T[index].siswa[idSiswa].jwbTugas)
				fmt.Println("Masukkan nilai jawaban : ")
				fmt.Scan(&nilai)
				T[index].siswa[index].nilTugas = nilai
				index++
			}
		} else if input == 2 {
			//Selection Sort
			var idx, i int
			var pass = 1
			for pass <= s-1 {
				idx = pass - 1
				i = pass
				for i < s {
					if T[0].siswa[idx].name > T[0].siswa[i].name {
						idx = i
					}
					i++
				}
				temp := T[0].siswa[pass-1]
				T[0].siswa[pass-1] = T[0].siswa[idx]
				T[0].siswa[idx] = temp
				pass++
			}

			for i := 0; i < s; i++ {
				fmt.Println(T[0].siswa[i].name, T[0].siswa[i].nilQuiz)
			}
			//Binary Search
			var input string
			fmt.Println("Nama siapa yang ingin dicari : ")
			fmt.Println("(Masukan '0' jika ingin kembali)")
			fmt.Scan(&input)

			if input != "0" {

				var left, right, mid, idTemu int
				idTemu = -1
				left = 0
				right = s - 1
				for left <= right && idTemu == -1 {
					mid = (left + right) / 2
					if input < T[0].siswa[mid].name {
						right = mid - 1
					} else if input > T[0].siswa[mid].name {
						left = mid + 1
					} else {
						idTemu = mid
					}
				}
				if idTemu == -1 {
					fmt.Println("Nama tidak ditemukan")
				} else {
					fmt.Println(T[0].siswa[idTemu].name, T[0].siswa[idTemu].nilQuiz)
				}
			}

		} else {
			fmt.Println("Input anda tidak valid")
		}
		fmt.Println("---------------------")
		fmt.Println("Menu Penilaian")
		fmt.Println("1. Nilai Tugas")
		fmt.Println("2. Nilai Quiz")
		fmt.Println("3. Kembali")
		fmt.Println("---------------------")
		fmt.Scan(&input)
	}
}

func main() {
	var f, t, q, g, s int

	var arrForum isiForum
	var arrUser user

	menuUtama(&arrForum, &arrUser, g, s, &f, &t, &q)
}

func menuUtama(arrForum *isiForum, arrUser *user, g, s int, f, t, q *int) {
	var aksi int

	for aksi != 3 {
		fmt.Println("========= APLIKASI MOODLE SEDERHANA ==========")
		fmt.Println("---------------------")
		fmt.Println("Menu utama")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Println("---------------------")

		fmt.Scan(&aksi)

		if aksi == 1 {
			registrasi(arrUser, &g, &s)
		} else if aksi == 2 {
			login(arrForum, arrUser, g, s, f, t, q)
		}
	}
}
