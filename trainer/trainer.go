package trainer

import "golang-starter-pack/model"

type Store interface {
	CreateTrainer(*model.Trainer) error
	UpdateTrainer(*model.Trainer, []string) error
	DeleteTrainer(*model.Trainer) error
}
