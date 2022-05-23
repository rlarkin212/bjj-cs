package scalar

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MarshalObjectId(id primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		res, err := id.MarshalText()
		if err != nil {
			fmt.Println(err)
		}

		w.Write(res)
	})
}

func UnmarshalObjectId(id string) (primitive.ObjectID, error) {
	res, err := primitive.ObjectIDFromHex(id)

	return res, err
}
