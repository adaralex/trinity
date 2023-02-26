package graph

import (
	"github.com/adaralex/trinity/graph/db"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB     *db.SecurityDatabase
	Logger *zap.SugaredLogger
}
