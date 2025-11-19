package services

import (
	"errors"

	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/dto"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/models"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ScheduleService struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewScheduleService(db *gorm.DB, log *logrus.Logger) *ScheduleService {
	return &ScheduleService{
		db:  db,
		log: log,
	}
}

func (s *ScheduleService) ListTeacher(ctx *fiber.Ctx, params dto.ListScheduleRequest) (dto.ListScheduleTeacherResponse, error) {
	var result dto.ListScheduleTeacherResponse
	var data models.Teacher

	if err := s.db.Preload("Schedules", func(db *gorm.DB) *gorm.DB {
		return db.Where("date >= ? AND date <= ?", params.StartDate, params.EndDate)
	}).
		Where("nik = ?", params.TeacherNik).
		Find(&data).Error; err != nil {
		s.log.Error(err.Error())
		ctx.Status(fiber.StatusInternalServerError)
		return result, errors.New("terjadi kesalahan sistem")
	}
	if data.ID == 0 {
		s.log.Error("guru tidak ditemukan")
		ctx.Status(fiber.StatusNotFound)
		return result, errors.New("guru tidak ditemukan")
	}

	jadwal := []dto.JadwalTeacher{}
	for _, v := range data.Schedules {
		jadwal = append(jadwal, dto.JadwalTeacher{
			Date:        v.Date.Format("2006-01-02"),
			ClassName:   v.ClassName,
			SubjectCode: v.SubjectCode,
			JamKe:       v.JamKe,
			TimeStart:   v.TimeStart,
			TimeEnd:     v.TimeEnd,
		})
	}

	result = dto.ListScheduleTeacherResponse{
		TeacherName: data.Name,
		Periode: dto.Periode{
			StartDate: params.StartDate,
			EndDate:   params.EndDate,
		},
		TotalJp: utils.HitungTotalJP(data.Schedules),
		Jadwal:  jadwal,
	}

	return result, nil
}

func (s *ScheduleService) ListStudent(ctx *fiber.Ctx, params dto.ListScheduleRequest) (dto.ListScheduleTeacherResponse, error) {
	var result dto.ListScheduleTeacherResponse
	var data models.Teacher

	if err := s.db.Preload("Schedules", func(db *gorm.DB) *gorm.DB {
		return db.Where("date >= ? AND date <= ?", params.StartDate, params.EndDate)
	}).
		Where("nik = ?", params.TeacherNik).
		Find(&data).Error; err != nil {
		s.log.Error(err.Error())
		ctx.Status(fiber.StatusInternalServerError)
		return result, errors.New("terjadi kesalahan sistem")
	}
	if data.ID == 0 {
		s.log.Error("guru tidak ditemukan")
		ctx.Status(fiber.StatusNotFound)
		return result, errors.New("guru tidak ditemukan")
	}

	jadwal := []dto.JadwalTeacher{}
	for _, v := range data.Schedules {
		jadwal = append(jadwal, dto.JadwalTeacher{
			Date:        v.Date.Format("2006-01-02"),
			ClassName:   v.ClassName,
			SubjectCode: v.SubjectCode,
			JamKe:       v.JamKe,
			TimeStart:   v.TimeStart,
			TimeEnd:     v.TimeEnd,
		})
	}

	result = dto.ListScheduleTeacherResponse{
		TeacherName: data.Name,
		Periode: dto.Periode{
			StartDate: params.StartDate,
			EndDate:   params.EndDate,
		},
		TotalJp: utils.HitungTotalJP(data.Schedules),
		Jadwal:  jadwal,
	}

	return result, nil
}
