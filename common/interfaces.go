package common

type PersistenceAdapter interface {
	Creator
	Updater
	Deleter
}

type Creator interface {
	Create(msg Message) (id string, err error)
}

type Updater interface {
	Update(id string, msg Message) error
}

type Deleter interface {
	Delete(id string) error
}

type Searcher interface {
	SearchByUser(user string) ([]Message, error)
	SearchByOrganization(organization string) ([]Message, error)
	SearchByProject(organization string) ([]Message, error)
}

type Getter interface {
	GetById(id string) (Message, error)
	Get(beginning int, amount int) ([]Message, error)
}
