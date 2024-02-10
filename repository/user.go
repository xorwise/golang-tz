package repository

import (
	"context"

	"github.com/xorwise/golang-tz/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	collection := r.db.Collection(string(domain.UserCollection))
	_, err := collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}

func (r *userRepository) InsertOrUpdate(ctx context.Context, user *domain.User) error {
	collection := r.db.Collection(string(domain.UserCollection))
	_, err := collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user}, options.Update().SetUpsert(true))
	return err
}

func (r *userRepository) GetByRefresh(ctx context.Context, refreshToken string) (domain.User, error) {
	collection := r.db.Collection(string(domain.UserCollection))
	var user domain.User
	err := collection.FindOne(ctx, bson.M{"refresh_token": refreshToken}).Decode(&user)
	return user, err
}
