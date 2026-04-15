package handler

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func extractObjectIDFromPath(path string, basePath string) (primitive.ObjectID, bool, error) {
	trimmedPath := strings.TrimSuffix(path, "/")
	trimmedBasePath := strings.TrimSuffix(basePath, "/")

	if trimmedPath == trimmedBasePath {
		return primitive.NilObjectID, false, nil
	}

	prefix := trimmedBasePath + "/"
	if !strings.HasPrefix(trimmedPath, prefix) {
		return primitive.NilObjectID, false, nil
	}

	idPart := strings.TrimPrefix(trimmedPath, prefix)
	if strings.Contains(idPart, "/") || idPart == "" {
		return primitive.NilObjectID, false, nil
	}

	id, err := primitive.ObjectIDFromHex(idPart)
	if err != nil {
		return primitive.NilObjectID, false, err
	}

	return id, true, nil
}
