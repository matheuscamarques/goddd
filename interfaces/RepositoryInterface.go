package interfaces

// make a repository pattern
//
// https://en.wikipedia.org/wiki/Repository_pattern
//
// https://en.wikipedia.org/wiki/Domain-driven_design
//
// https://en.wikipedia.org/wiki/Domain-driven_design#In_software_development
//

type RepositoryInterface interface {
	// method find by id
	Find(id int) (interface{}, error)
	// method find all
	FindAll() ([]interface{}, error)
	// method create
	Create(interface{}) (interface{}, error)
	// method update
	Update(interface{}) (interface{}, error)
	// method delete
	Delete(interface{}) (interface{}, error)
}
