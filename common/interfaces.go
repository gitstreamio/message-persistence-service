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
	SearchByUser(user string) []Message
	SearchByOrganization(organization string) []Message
	SearchByProject(organization string) []Message
}

type Finder interface {
	FindById(id string) Message
}
