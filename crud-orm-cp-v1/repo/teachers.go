package repo

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	if result := t.db.Create(&data); result.Error != nil {
		return fmt.Errorf("Error INSERT Teacher")
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	garis, err := t.db.Table("teachers").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	var listGuru []model.Teacher
	for garis.Next() {
		t.db.ScanRows(garis, &listGuru)
	}
	return listGuru, nil // TODO: replace this
}

func (t TeacherRepo) Update(id uint, name string) error {
	if err := t.db.Table("teachers").Where("id = ?", id).Update("name", name).Error; err != nil {
		return fmt.Errorf("Error UPDATE Teacher")
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error {
	guru := model.Teacher{}
	if result := t.db.Where("id = ?", id).Delete(&guru); result.Error != nil {
		return fmt.Errorf("Error Delete Teacher")
	}
	return nil // TODO: replace this
}
