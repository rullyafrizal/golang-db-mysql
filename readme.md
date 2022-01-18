# Golang with Database MySQL

## Package Database
- Secara default, Go memiliki sebuah package database
- Package database adalah package yang berisikan kumpulan standard interface untuk mengakses/berkomunikasi dengan database
- Hal ini bisa menjadikan kode program yang kita buat untuk akses database apapun bisa memakai kode yang sama

### Cara kerja
- **Aplikasi** call **Database Interface** call **Database Driver** call **DBMS**

## Koneksi ke Database
- Untuk melakukan koneksi database bisa membuat object **sql.DB** menggunakan function `sql.Open(driver, dataSourceName)`
- Driver adalah nama driver yang digunakan untuk mengakses database. ex: `"mysql"`
- DataSourceName adalah url yang digunakan untuk mengakses database. ex: `"username:password@tcp(host:port)/db_name`
- Jika sudah tidak digunakan lagi, maka bisa menutup koneksi menggunakan function `Close()`

## Database Pooling (Manajemen Koneksi Database)
- **sql.DB** di Go sebenarnya bukan sebuah koneksi, melainkan sebuah pool ke database, atau dikenal sebagai konsep database pooling
- Di dalam **sql.DB**, Go melakukan manajemen koneksi ke db secara otomatis. Hal ini menjadikan kita tidak perlu melakukan manajemen koneksi db secara manual
- Dengan kemampuan database pooling ini, kita bisa tentukan minimal dan maksimal jumlah koneksi yang dibuat oleh program, sehingga tidak terjadi koneksi terlalu banyak, karena ada batasan maksimal koneksi yang bisa ditangani oleh database yang kita gunakan

## Pengaturan Database Pooling
| **Method** | **Description** |
| ----------- | ----------- |
| `(DB) SetMaxIdleConns(number)` | Pengaturan berapa jumlah koneksi minimal yang dibuat |
| `(DB) SetMaxOpenConns(number)` | Pengaturan berapa jumlah koneksi maksimal yang dibuat |
| `(DB) SetConnMaxIdleTime(duration)` | Pengaturan berapa lama koneksi yang sudah tidak dipakai akan dihapus |
| `(DB) SetConnMaxLifetime(duration)` | Pengaturan berapa lama koneksi boleh digunakan |

## Eksekusi Perintah SQL
- Di Go juga disediakan function yang bisa digunakan untuk eksekusi perintah SQL dengan menggunakan `(DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
- Ketika kita kirim perintah SQL dengan memakai function di atas, kita butuh passing context sebagai parameter agar kita bisa mengirim sinyal cancel jika kita ingin batalkan pengiriman perintah SQL-nya

## Query SQL
- Untuk operasi yang bukan query, kita pakai perintah Exec, namun jika kita ingin mengambil data dari database yang menghasilkan result, kita bisa memakai function `(DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)`

## Rows
- Hasil Query function adalah sebuah struct sql.Rows
- Rows digunakan untuk iterasi terhadap hasil Query
- Kita bisa memakai function `(Rows) Next() bool` untuk iterasi terhadap data hasil query, jika return false maka data sudah habis
- Untuk membaca tiap daata, kita bisa pakai `(Rows) Scan(columns...) error`
- Jangan lupa untuk selalu menutup dengan menggunakan `(Rows) Close() error`

## Mapping Tipe Data
https://docs.google.com/presentation/d/15pvN3L3HTgA9aIMNkm03PzzIwlff0WDE6hOWWut9pg8/edit#slide=id.gbcfff2dc3b_0_116

## Nullable Type
- Golang database tidak mengerti dengan tipe data NULL di database
- Untuk mengakses data yang NULL, kita bisa menggunakan tipe data Nullable
- Untuk mengakses data yang null, kita bisa menggunakan struct `sql.NullInt64`, `sql.NullString`, `sql.NullBool`, `sql.NullFloat64`, `sql.NullTime`

## Menghindari SQL Injection
- Untuk menghindari SQL Injection, kita bisa menggunakan parameter ketiga di function `(DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)`
- Parameter ketiga adalah `args...interface{}`
- Untuk menandai sebuah SQL membutuhkan parameter, kita bisa menggunakan tanda tanya `?`
- Contoh : `INSERT INTO users (name, email) VALUES (?, ?)`

## Auto Increment
- Kita bisa menggunakan function `(Result) LastInsertId() (int64, error)` untuk mendapatkan nilai auto increment
  

## Query atau Exec dengan Parameter
- Saat kita pakai function Query atau Exec yang menggunakan parameter, sebenarnya implementasi di bawahnya menggunakan Prepare Statement
- Jadi tahapan pertama statement-nya disiapkan terlebih dahulu, setelah itu baru di isi dengan parameter
- Terkadang ada kasus di mana kita ingin melakukan beberapa hal yang sama sekaligus, hanya berbeda di parameternya. Misal bulk insert
- Pembuatan Prepare Statement bisa dilakukan dengan manual, tanpa harus menggunakan Query atau Exec dengan parameter
- Pembuatan Prepare Statement dengan menggunakan function `(DB) PrepareContext(ctx context, query string) (*Stmt, error)`
- atau bisa juga tanpa context dengan menggunakan function `(DB) Prepare(query string) (*Stmt, error)`
- Prepare statement direpresentasikan dengan struct `sql.Stmt`
- Sama seperti resurce sql lainnya, Stmt harus di Close() jika sudah tidak digunakan

## Database Transaction
- Secara default, semua eksekusi SQL di Golang yang kita kirim akan otomatis di-commit, istilahnya auto-commit
- Namun untuk menggunakan fitur transaksi di SQL, kita tidak boleh commit secara otomatis
- Untuk memulai transaksi, kita bisa menggunakan function `(DB) Begin() (*Tx, error)`, di mana akan menghasilkan struct Tx yang merupakan representasi Transaction
- Struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, hampir semua function di DB juga tersedia di Tx seperti Exec, Query, atau Prepare
- Setelah selesai transaksi, kita bisa menggunakan function `(Tx) Commit() error` untuk commit data
- Jika kita ingin rollback, kita bisa menggunakan function `(Tx) Rollback() error`

## Repository Pattern
- Repository Pattern adalah sebuah pattern yang mengatur hubungan antara model dan database
- Repository Pattern memungkinkan kita untuk mengatur hubungan antara model dengan database secara efektif dan efisien
- Pattern ini biasa digunakan untuk menghubungkan business logic dengan semua perintah SQL ke database
- Semuq perintah SQL cukup ditulis di Repository, sedangkan jika dalam business logic butuh akses ke database maka tinggal panggil Repository saja


   