package register

import (
	"github.com/gofiber/fiber/v2"
	"time"

	"github.com/go-co-op/gocron"
)

func (r *Runner) Inject(routes []fiber.Route) {
	//r.scheduleTask()
	r.RegisterSAUser()
	r.RegisterSystemApi(routes)
}

func (r *Runner) scheduleTask() {
	s := gocron.NewScheduler(time.Local)
	_, err := s.Every(5).Second().Do(r.ScheduleTask1)
	if err != nil {
		r.Logger.Error(err.Error())
	}
	s.StartAsync()
}
