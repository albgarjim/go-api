package input

type ItemData struct {
	StringField string   `json:"stringField", rethinkdb:"stringField"`
	ListField   []string `json:"listField", rethinkdb:"listField"`
	IntField    int      `json:"intField", rethinkdb:"intField"`
	BoolField   bool     `json:"boolField", rethinkdb:"boolField"`

	ID string `json:"id", rethinkdb:"id"`
}
