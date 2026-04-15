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

func CreateUser(user model.User) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, user)
	if err == nil {
		user.ID = result.InsertedID.(primitive.ObjectID)
	}
	return user, err
}

func GetUserByID(id primitive.ObjectID) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user model.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func GetUserByEmail(email string) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user model.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}

// FindBySteamID busca um usuário pelo steamid64 no banco.
// Retorna mongo.ErrNoDocuments se o usuário ainda não existe —
// o chamador usa isso pra decidir se cria ou atualiza.
func FindBySteamID(steamID string) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := collection.FindOne(ctx, bson.M{"steam_id": steamID}).Decode(&user)
	return user, err
}

// UpsertSteamUser cria ou atualiza um usuário baseado no steamid.
//
// "Upsert" = update + insert:
//   - Se já existe um documento com esse steam_id → atualiza nome, avatar e último login
//   - Se não existe → cria um documento novo com todos os campos
//
// Isso resolve o caso de login repetido: o usuário sempre tem os dados
// mais recentes do perfil Steam sem duplicar registros no banco.
func UpsertSteamUser(user model.User) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()

	// upsert: true → cria o documento se não encontrar
	// setOnInsert → campos que só são definidos na criação (ID e CreatedAt)
	// $set → campos que são sempre atualizados
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"steam_id": user.SteamID}
	update := bson.M{
		"$set": bson.M{
			"username":      user.Username,
			"avatar_url":    user.AvatarURL,
			"profile_url":   user.ProfileURL,
			"updated_at":    now,
			"last_login_at": now,
		},
		"$setOnInsert": bson.M{
			"_id":        primitive.NewObjectID(),
			"steam_id":   user.SteamID,
			"created_at": now,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return model.User{}, err
	}

	// Busca o documento atualizado/criado pra retornar com o ID correto.
	return FindBySteamID(user.SteamID)
}

func UpdateUser(id primitive.ObjectID, user model.User) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user.UpdatedAt = time.Now()
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	return user, err
}

func DeleteUser(id primitive.ObjectID) error {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func ListUsers() ([]model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	var users []model.User
	err = cursor.All(ctx, &users)
	return users, err
}
