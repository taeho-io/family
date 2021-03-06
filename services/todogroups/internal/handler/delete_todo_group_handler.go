package handler

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todogroups/internal/repo"
)

type DeleteTodoGroupFunc func(
	ctx context.Context,
	req *todogroups.DeleteTodoGroupRequest,
) (*todogroups.DeleteTodoGroupResponse, error)

func DeleteTodoGroup(
	todoGroupsRepo repo.GroupsRepo,
	todoGroupPermitsRepo repo.PermitsRepo,
) DeleteTodoGroupFunc {
	return func(
		ctx context.Context,
		req *todogroups.DeleteTodoGroupRequest,
	) (*todogroups.DeleteTodoGroupResponse, error) {
		accountID := base.GetAccountIDFromContext(ctx)
		if err := base.HasPermissionByAccountID(ctx, accountID); err != nil {
			return nil, err
		}

		logger := ctxlogrus.Extract(ctx)

		if err := todoGroupsRepo.DeleteByID(req.TodoGroupId); err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupsRepo.DeleteByID",
				"req":  req,
			}).Error(err)

			return nil, err
		}

		if err := todoGroupPermitsRepo.Delete(accountID, req.TodoGroupId); err != nil {
			logger.WithFields(logrus.Fields{
				"what": "todoGroupPermitsRepo.Delete",
				"req":  req,
			}).Error(err)

			return nil, err
		}

		return &todogroups.DeleteTodoGroupResponse{}, nil
	}
}
