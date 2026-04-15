package repository

import (
	"context"
	"time"

	"github.com/Roger13san/games-review/backend/internal/database"
	"github.com/Roger13san/games-review/backend/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetLibraryItems retorna todos os jogos da biblioteca de um usuário específico,
// ordenados do mais recente pro mais antigo.
func GetLibraryItems(userID primitive.ObjectID) ([]model.LibraryItem, error) {
	collection := database.GetCollection("library")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{
		{Key: "created_at", Value: -1},
	})

	// Filtra apenas os jogos do usuário logado.
	cursor, err := collection.Find(ctx, bson.M{"user_id": userID}, findOptions)
	if err != nil {
		return nil, err
	}

	var items []model.LibraryItem
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func GetLibraryItemByID(id primitive.ObjectID) (model.LibraryItem, error) {
	collection := database.GetCollection("library")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var item model.LibraryItem
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return model.LibraryItem{}, err
	}

	return item, nil
}

// UpsertLibraryItem salva um jogo na biblioteca sem duplicar.
// O filtro usa user_id + platform_id — se o jogo já existe pra esse usuário,
// atualiza; se não existe, cria.
func UpsertLibraryItem(item model.LibraryItem) error {
	collection := database.GetCollection("library")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()
	opts := options.Update().SetUpsert(true)

	filter := bson.M{
		"user_id":     item.UserID,
		"platform_id": item.PlatformID,
	}
	update := bson.M{
		"$set": bson.M{
			"added_by":   item.AddedBy,
			"platform":   item.Platform,
			"updated_at": now,
		},
		"$setOnInsert": bson.M{
			"_id":        primitive.NewObjectID(),
			"user_id":    item.UserID,
			"platform_id": item.PlatformID,
			"created_at": now,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func CreateLibraryItem(item model.LibraryItem) (model.LibraryItem, error) {
	collection := database.GetCollection("library")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, item)
	if err != nil {
		return model.LibraryItem{}, err
	}

	return item, nil
}

func UpdateLibraryItem(id primitive.ObjectID, item model.LibraryItem) (model.LibraryItem, error) {
	collection := database.GetCollection("library")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	item.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": item},
	)
	if err != nil {
		return model.LibraryItem{}, err
	}

	return item, nil
}

func DeleteLibraryItem(id primitive.ObjectID) error {
	collection := database.GetCollection("library")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
