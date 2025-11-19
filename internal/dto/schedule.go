package dto

type ListScheduleRequest struct {
	TeacherNik string `json:"teacher_nik" query:"teacher_nik"`
	Kelas      string `json:"kelas" query:"kelas"`
	Date       string `json:"date" query:"date"`
	StartDate  string `json:"start_date" query:"start_date"`
	EndDate    string `json:"end_date" query:"end_date"`
}

type Periode struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type JadwalTeacher struct {
	Date        string `json:"date"`
	ClassName   string `json:"class_name"`
	SubjectCode string `json:"subject_code"`
	JamKe       int    `json:"jam_ke"`
	TimeStart   string `json:"time_start"`
	TimeEnd     string `json:"time_end"`
}

type ListScheduleTeacherResponse struct {
	TeacherName string          `json:"teacher_name"`
	Periode     Periode         `json:"periode"`
	TotalJp     int             `json:"total_jp"`
	Jadwal      []JadwalTeacher `json:"jadwal"`
}

type JadwalStudent struct {
	JamKe       int    `json:"jam_ke"`
	SubjectCode string `json:"subject_code"`
	TeacherName string `json:"teacher_name"`
	TimeStart   string `json:"time_start"`
	TimeEnd     string `json:"time_end"`
}

type ListScheduleStudentResponse struct {
	ClassName string          `json:"class_name"`
	Date      string          `json:"date"`
	Jadwal    []JadwalStudent `json:"jadwal"`
}
