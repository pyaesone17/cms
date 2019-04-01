package datastore

type PostDB interface {
	Create()
}

type postdb struct {
}

func NewPostDataStore() PostDB {
	return postdb{}
}

func (db postdb) Create() {

}
