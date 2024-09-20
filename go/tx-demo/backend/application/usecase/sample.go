package usecase

import "backend/domain/repository"

type Sample interface {
	UpdateTitleAndName(userId, itemId, name, title string) error
}

type sample struct {
	userRepository repository.User
	itemRepository repository.Item
}

func NewSample(userRepository repository.User, itemRepository repository.Item) Sample {
	return &sample{
		userRepository: userRepository,
		itemRepository: itemRepository,
	}
}

func (s *sample) UpdateTitleAndName(userId, itemId, name, title string) error {
	return s.userRepository.Transaction(func() error {
		err := s.userRepository.UpdateName(userId, name)
		if err != nil {
			return err
		}

		err = s.itemRepository.UpdateTitle(itemId, title)
		if err != nil {
			return err
		}

		return nil
	})
}
