package stats

import (
	"image/color"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserPin -
type UserPin struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	URL         string             `bson:"url"`
	Name        string             `bson:"name"`
	Label       string             `bson:"label"`
	Weight      int                `bson:"weight"`
	Hidden      bool               `bson:"hidden"`
	Description string             `bson:"description"`

	Size      int        `bson:"-"`
	Glow      bool       `bson:"glow"`
	GlowColor color.RGBA `bson:"glow_color"`
}
