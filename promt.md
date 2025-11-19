Buatkan UI lengkap (tanpa backend) untuk sebuah sistem jadwal pelajaran sekolah berdasarkan spesifikasi dokumen berikut:

## Tujuan UI:
Hanya membuat tampilan antarmuka (frontend-only, tanpa integrasi database) untuk fitur-fitur:
- CRUD Jadwal Pelajaran
- Upload Excel Jadwal
- Export Rekap Jam Pelajaran (JP)
- Halaman Jadwal Siswa
- Halaman Jadwal Guru
- Halaman Rekap JP Yayasan

## Teknologi yang wajib digunakan:
- Next.js
- TypeScript
- TailwindCSS
- Layout modern
- Tema warna retro (retro palette)
- Fully responsive: desktop, tablet, mobile

## Style & Guideline UI:
- Nuansa retro modern (warna pastel retro seperti teal, cream, brown, mustard).
- Clean, flat, minimalis, tidak terlalu banyak gradient.
- Navigasi sidebar untuk desktop, bottom navbar untuk mobile.
- Komponen reusable (Button, Input, Card, Table, FileUploader).
- Semua halaman harus memiliki versi mobile-friendly.

## Halaman yang harus dibuat:
1. Dashboard
   - Ringkasan total jadwal, total guru, total kelas.
   - Grafik ringan (dummy) seperti bar chart JP guru.
2. Manajemen Jadwal (CRUD)
   - Tabel jadwal (class_code, class_name, subject_code, teacher_name, date, jam_ke, time_start, time_end).
   - Modal: Create / Edit Jadwal.
   - Tombol delete dengan confirm dialog.
3. Upload Jadwal (Excel .xlsx)
   - UI upload file dengan drag-and-drop.
   - Preview file name setelah dipilih.
4. Export Rekap JP
   - Form pilih tanggal mulai & akhir.
   - Button "Export".
5. Halaman Siswa
   - Form (class_code & date).
   - Tampilkan jadwal siswa per hari.

6. Halaman Guru
   - Form (teacher_nik, start_date, end_date).
   - Tabel jadwal mengajar + perhitungan total jam pelajaran (JP).
7. Halaman Yayasan (Rekap JP Guru)
   - Form filter tanggal.
   - Tabel rekap JP per guru.

## Output yang saya minta dari AI:
- Struktur folder Next.js
- Semua file page.tsx
- Semua komponen (Button, Card, Table, Input, Modal, FileUploader, dll.)
- Tailwind config jika perlu custom theme retro
- Contoh dummy data agar UI bisa dirender
- Semua halaman harus responsive
- Pastikan UI tidak melakukan koneksi backend (hanya mock/dummy)

## HARUS dipastikan:
- Tidak membuat backend
- Tidak membuat database
- Tidak membuat autentikasi real
- Semua hanya UI/UX mock