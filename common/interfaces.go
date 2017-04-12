package common

type PersistenceAdapter interface {
	Persister
	Updater
	Deleter
}

type Persister interface {
	Persist(msg Message) (id string, err error)
}

type Updater interface {
	Update(msg Message) error
}

type Deleter interface {
	Delete(id string) error
}

type Searcher interface {
	SearchByUser(user string) []Message
	SearchByOrganization(organization string) []Message
	SearchByProject(organization string) []Message
}

