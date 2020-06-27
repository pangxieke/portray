package model

func MgoAdd(collection string, data interface{}) (err error) {
	ss := Mgo.session.Copy()
	defer ss.Close()

	err = ss.DB(Mgo.db).C(collection).Insert(data)
	return
}
