package main

type Service interface {
	SearchVacancy(query string) ([]Vacancy, error)
	GetVacancy(id string) (*Vacancy, error)
	ListVacancies() ([]Vacancy, error)
	DeleteVacancy(id string) error

	ListSearchHistory() ([]SearchHistory, error)
	DeleteSearchHistory(id int) error
}

type MyService struct {
	repo Repository
}

func NewMyService(repo Repository) *MyService {
	return &MyService{repo: repo}
}

func (s *MyService) SearchVacancy(query string) ([]Vacancy, error) {
	vacancies, err := s.repo.SearchVacancy(query)
	if err != nil {
		return nil, err
	}

	return vacancies, nil
}

func (s *MyService) GetVacancy(id string) (*Vacancy, error) {
	vacancy, err := s.repo.GetVacancy(id)
	if err != nil {
		return nil, err
	}

	return vacancy, nil
}

func (s *MyService) ListVacancies() ([]Vacancy, error) {
	vacancies, err := s.repo.ListVacancies()
	if err != nil {
		return nil, err
	}

	return vacancies, nil
}

func (s *MyService) DeleteVacancy(id string) error {
	err := s.repo.DeleteVacancy(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *MyService) ListSearchHistory() ([]SearchHistory, error) {
	history, err := s.repo.ListSearchHistory()
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (s *MyService) DeleteSearchHistory(id int) error {
	err := s.repo.DeleteSearchHistory(id)
	if err != nil {
		return err
	}

	return nil
}
