package logic

func (s *TodoServiceImpl) AddUserTodos(ID string) error {

	err := s.Repo.AddTodo(ID)
	if err != nil {
		return err
	}
	return nil
}
