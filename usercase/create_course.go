package usercase

import (
	"github.com/andreggpereira/go-kafka/entity"
	"github.com/google/uuid"
)

type CreateCouse struct {
	Repository entity.CourseRepository
}

func (c CreateCouse) Excute(input CreateCourseInputDto) (CreateCourseOutputDto, error) {

	course := entity.Course{}
	course.ID = uuid.New().String()
	course.Name = input.Name
	course.Description = input.Description
	course.Status = input.Status

	err := c.Repository.Insert(course)
	if err != nil {
		return CreateCourseOutputDto{}, err
	}

	return CreateCourseOutputDto{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		Status:      course.Status,
	}, nil
}
