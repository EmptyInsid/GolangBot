package object

type DB interface {
	SaveFB(tableName string, obj *Object) error
	SaveLink(tableName string, obj *Object) error
	ShowAll(tableName string, userID int) (*[]Object, error)
	ShowOne(tableName string, userName string) (**Object, error)
	NoteDone(tableName string, obj *Object) error
	Remove(tableName string, obj *Object) error
	PickRandom(tableName string, obj *Object) error
	IsExists(tableName string, obj *Object) (bool, error)
}

type Object struct {
	Id      int
	Name    string
	Author  string
	Genre   string
	Year    int
	Adviser string
	Done    bool
	UserId  int
}
