package templates

const DbServiceTpl = `
func Connect{{ .Name }}DB() (*gorm.DB, error) {
	return database.Connect({{ databaseName . }})
}

// The same as Connect, but panics if there is an error.
func MustConnect{{ .Name }}DB() *gorm.DB {
	db, err := Connect{{ .Name }}DB()

	if err != nil {
		panic(err)
	}

	return db
}
`
