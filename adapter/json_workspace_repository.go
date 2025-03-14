package adapter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/mshanah/lcnc-domain/domain"
	"github.com/mshanah/lcnc-port/port"
)

type JSONWorkspaceRepository struct {
	filePath string
}

func NewJSONWorkspaceRepository(filePath string) port.WorkspaceRepository {
	return &JSONWorkspaceRepository{filePath: filePath}
}

func (r *JSONWorkspaceRepository) Save(workspace *domain.Workspace) error {
	workspaces, err := r.FindAll()
	if err != nil {
		return err
	}

	for _, ws := range workspaces {
		if ws.ID == workspace.ID {
			return errors.New("workspace already exists")
		}
	}

	workspaces = append(workspaces, workspace)
	return r.writeToFile(workspaces)
}

func (r *JSONWorkspaceRepository) FindByID(id string) (*domain.Workspace, error) {
	workspaces, err := r.FindAll()
	if err != nil {
		return nil, err
	}

	for _, ws := range workspaces {
		if ws.ID == id {
			return ws, nil
		}
	}

	return nil, errors.New("workspace not found")
}

func (r *JSONWorkspaceRepository) FindAll() ([]*domain.Workspace, error) {
	data, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []*domain.Workspace{}, nil
		}
		return nil, err
	}

	var workspaces []*domain.Workspace
	err = json.Unmarshal(data, &workspaces)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (r *JSONWorkspaceRepository) Delete(id string) error {
	workspaces, err := r.FindAll()
	if err != nil {
		return err
	}

	updatedWorkspaces := []*domain.Workspace{}
	for _, ws := range workspaces {
		if ws.ID != id {
			updatedWorkspaces = append(updatedWorkspaces, ws)
		}
	}

	if len(updatedWorkspaces) == len(workspaces) {
		return errors.New("workspace not found")
	}

	return r.writeToFile(updatedWorkspaces)
}

func (r *JSONWorkspaceRepository) writeToFile(workspaces []*domain.Workspace) error {
	data, err := json.MarshalIndent(workspaces, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(r.filePath, data, 0644)
}
