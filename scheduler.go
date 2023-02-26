package main

import (
	"context"
	"github.com/adaralex/trinity/graph/db"
	"go.uber.org/zap"
	"time"
)

type ProjectSchedules struct {
	Name     string
	Analysis map[string]int
}

type SchedulerHandler struct {
	Projects map[string]ProjectSchedules
	Context  context.Context
	Database *db.SecurityDatabase
	Logger   *zap.SugaredLogger
}

func (h *SchedulerHandler) Init() {
	h.Context = context.Background()
	h.Logger.Infow("starting scheduler handler", "time", time.Now().UnixMilli())
}

func (h *SchedulerHandler) AddProjectAnalysisSchedule(id string, analysis string) error {
	_, err := h.Database.GetProject(id)
	if err != nil {
		h.Logger.Errorw("error retrieving project", "project", id, "time", time.Now().UnixMilli())
		return err
	}

	return nil
}

func (h *SchedulerHandler) ClearProjectAnalysisSchedules(id string) {
	delete(h.Projects, id)
}
