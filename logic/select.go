package logic

import "todolist/domain"

func (s *TodoServiceImpl) GetUserTodos(ID string) (domain.Todo, error) {

	todo, err := s.Repo.GetTodos(ID)
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}
