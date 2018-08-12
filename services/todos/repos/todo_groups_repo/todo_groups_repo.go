package todo_groups_repo

import (
	"fmt"
	"strings"

	"github.com/guregu/dynamo"

	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/base/aws/dynamodb/table"
	"github.com/taeho-io/family/services/todos/config"
	"github.com/taeho-io/family/services/todos/models"
)

const (
	todoGroupIDFieldKey    = "todo_group_id"
	orderedTodoIDsFieldKey = "ordered_todo_ids"
)

type Table struct {
	table.IFace

	todoGroupsTable *dynamo.Table
}

func New(ddb dynamodb.IFace, cfg config.IFace) *Table {
	fullTableName := fullTableName(cfg)
	todoGroupsTable := ddb.DB().Table(fullTableName)

	return &Table{
		todoGroupsTable: &todoGroupsTable,
	}
}

func NewMock() *Table {
	ddb := dynamodb.NewMock()
	cfg := config.NewMock()

	return New(ddb, cfg)
}

func fullTableName(cfg config.IFace) string {
	prefix := cfg.Prefix()
	tableName := cfg.Settings().DynamodbTodoGroupsTableName
	return strings.Join([]string{prefix, tableName}, "-")
}

func (t *Table) Table() *dynamo.Table {
	return t.todoGroupsTable
}

func (t *Table) GetByID(todoGroupID string) (*models.TodoGroup, error) {
	var todoGroup models.TodoGroup

	if err := t.Table().Get(todoGroupIDFieldKey, todoGroupID).One(&todoGroup); err != nil {
		return nil, err
	}

	return &todoGroup, nil
}

func (t *Table) ListByIDs(todoGroupIDs []string) ([]*models.TodoGroup, error) {
	var todoGroups []*models.TodoGroup

	// TODO: make it concurrent.
	for _, todoGroupID := range todoGroupIDs {
		todoGroup, err := t.GetByID(todoGroupID)
		if err != nil {
			return nil, err
		}
		todoGroups = append(todoGroups, todoGroup)
	}

	return todoGroups, nil
}

func (t *Table) validateTodoGroupInput(todoGroup *models.TodoGroup) error {
	if todoGroup.TodoGroupID == "" {
		return fmt.Errorf("invalid todo_group_id")
	}

	return nil
}

func (t *Table) Put(todoGroup *models.TodoGroup) error {
	if err := t.validateTodoGroupInput(todoGroup); err != nil {
		return err
	}

	return t.Table().
		Put(todoGroup).
		Run()
}

func (t *Table) ReorderTodos(todoGroupID string, orderedTodoIDs []string) (*models.TodoGroup, error) {
	var todoGroup models.TodoGroup

	err := t.Table().
		Update(orderedTodoIDsFieldKey, orderedTodoIDs).
		Value(&todoGroup)
	if err != nil {
		return nil, err
	}

	return &todoGroup, nil
}

func (t *Table) DeleteByID(todoGroupID string) error {
	return t.Table().
		Delete(todoGroupIDFieldKey, todoGroupID).
		Run()
}

func (t *Table) AppendTodoID(todoGroupID string, todoID string) (*models.TodoGroup, error) {
	var todoGroup models.TodoGroup

	err := t.Table().
		Update(todoGroupIDFieldKey, todoGroupID).
		Append(orderedTodoIDsFieldKey, todoID).
		Value(&todoGroup)
	if err != nil {
		return nil, err
	}

	return &todoGroup, nil
}

func (t *Table) DeleteTodoID(todoGroupID string, todoID string) (*models.TodoGroup, error) {
	var todoGroup models.TodoGroup

	err := t.Table().
		Update(todoGroupIDFieldKey, todoGroupID).
		DeleteStringsFromSet(orderedTodoIDsFieldKey, todoID).
		Value(&todoGroup)
	if err != nil {
		return nil, err
	}

	return &todoGroup, nil
}
