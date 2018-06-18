package db

import arango "github.com/diegogub/aranGO"

func GetSessionDB() *arango.Session {

	//Connection ArangoDB
	s, err := arango.Connect("http://localhost:8529","root", "parangaturimicuaro", false)

	if err != nil {
		panic(err)
	}
	return s
}