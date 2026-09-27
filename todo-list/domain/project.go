package domain

type ProjectID string

func (p ProjectID) String() string {
	return string(p)
}

type Project struct {
	id    ProjectID
	name  string
	tasks []TaskID
}

func NewProject(id ProjectID, name string) *Project {
	return &Project{
		id:    id,
		name:  name,
		tasks: make([]TaskID, 0),
	}
}

func NewProjectID(id string) ProjectID {
	return ProjectID(id)
}

func (p *Project) AddTask(task *Task) error {
	if len(p.tasks) >= 5 {
		return ErrProjectFull
	}
	if err := task.AssignToProject(p.id); err != nil {
		return err
	}
	p.tasks = append(p.tasks, task.ID())

	return nil
}

func (p *Project) RemoveTask(id TaskID) {
	for i := range p.tasks {
		if p.tasks[i] == id {
			p.tasks = append(p.tasks[:i], p.tasks[i+1:]...)
			break
		}
	}
}
