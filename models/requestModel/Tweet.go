package requestModel

type Tweet struct {
	Text      string             `bson:"text" json:"text,omitempty"`
}
