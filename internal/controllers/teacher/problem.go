package teacher

import (
	"net/http"
	"strconv"

	"github.com/MNU/exam-go"
	"github.com/labstack/echo"
)

// ProblemController is
type ProblemController struct {
	problemSvc goexam.ProblemService
	courseSvc  goexam.CourseService
}

// NewProblemController is
func NewProblemController(problemSvc goexam.ProblemService, courseSvc goexam.CourseService) *ProblemController {
	return &ProblemController{
		problemSvc,
		courseSvc,
	}
}

// Create is
func (p *ProblemController) Create(ctx echo.Context) error {
	problem := new(goexam.Problem)
	err := ctx.Bind(problem)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "参数错误")
	}

	_, err = p.courseSvc.Get(problem.CourseID)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "course not found")
	}

	err = p.problemSvc.Create(problem)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

// Update is
func (p *ProblemController) Update(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "ID不存在")
	}

	problem := new(goexam.Problem)
	err = ctx.Bind(problem)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "参数错误")
	}
	problem.ID = uint(id)
	err = p.problemSvc.Update(problem)
	return ctx.JSON(http.StatusOK, "成功")
}

// Delele is
func (p *ProblemController) Delele(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	err = p.problemSvc.Delete(uint(id))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	return ctx.JSON(http.StatusOK, "成功")
}

// Get is ...
func (p *ProblemController) Get(ctx echo.Context) error {
	_id := ctx.Param("id")
	id, err := strconv.ParseUint(_id, 10, 64)
	problem, err := p.problemSvc.Get(uint(id))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "sevide")
	}
	return ctx.JSON(http.StatusOK, problem)
}

// GetList is
func (p *ProblemController) GetList(ctx echo.Context) error {
	filter := new(goexam.ProblemFilter)
	err := ctx.Bind(filter)
	filter.LoadDefault()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "参数错误")
	}
	problemList, err := p.problemSvc.GetList(filter)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "服务错误")
	}
	return ctx.JSON(http.StatusOK, problemList)
}
