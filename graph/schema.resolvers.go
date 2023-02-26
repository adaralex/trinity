package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"time"

	"github.com/adaralex/trinity/graph/model"
)

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, project model.ProjectInput) (*model.Project, error) {
	err := r.DB.CreateProject(project.Name)
	if err != nil {
		r.Logger.Errorw("graphql createproject error", "time", time.Now().UnixMilli(), "project", project.Name)
		return nil, err
	} else {
		_, err := r.DB.GetProject(project.Name)
		r.Logger.Infow("graphql createproject success", "time", time.Now().UnixMilli(), "project", project.Name)
		return nil, err
	}
}

// CreateScanner is the resolver for the createScanner field.
func (r *mutationResolver) CreateScanner(ctx context.Context, scanner model.ScannerInput) (*model.Scanner, error) {
	panic(fmt.Errorf("not implemented: CreateScanner - createScanner"))
}

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - CreateUser"))
}

// Project is the resolver for the project field.
func (r *queryResolver) Project(ctx context.Context, id string) (*model.Project, error) {
	p, err := r.DB.GetProject(id)
	r.Logger.Infow("graphql project call", "time", time.Now().UnixMilli(), "project", id)
	if err == nil {
		result := ProjectToGQL(p)
		return &result, nil
	}
	return nil, err
}

// Scanner is the resolver for the scanner field.
func (r *queryResolver) Scanner(ctx context.Context, id string) (*model.Scanner, error) {
	panic(fmt.Errorf("not implemented: Scanner - scanner"))
}

// Scanners is the resolver for the scanners field.
func (r *queryResolver) Scanners(ctx context.Context, id *string) ([]*model.Scanner, error) {
	panic(fmt.Errorf("not implemented: Scanners - scanners"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
