package validation

import "go.mongodb.org/mongo-driver/bson/primitive"

func IsValidObjectID(s string) string {
	objID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		newObj := primitive.NewObjectID()
		return newObj.Hex()
	}
	return objID.Hex()
}
