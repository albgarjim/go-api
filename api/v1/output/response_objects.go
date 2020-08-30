package output

type ItemDataRDB struct {
	StringField string   `json:"stringField",rethinkdb:"stringField",omitempty`
	ListField   []string `json:"listField",rethinkdb:"listField,omitempty`
	IntField    int      `json:"intField",rethinkdb:"intField",omitempty`
	BoolField   bool     `json:"boolField",rethinkdb:"boolField",omitempty`

	ID string `json:"id",rethinkdb:"id",omitempty`
}
