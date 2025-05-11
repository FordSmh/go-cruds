package repo

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TokenRepository interface {
	Save(ctx *gin.Context, token string) error
	Find(ctx *gin.Context, token string) (string, error)
	Delete(ctx *gin.Context, token string) error
}

type TokenRepo struct {
	collection *mongo.Collection
}

func NewTokenRepository(db *mongo.Database) TokenRepository {
	return &TokenRepo{collection: db.Collection("jwt_tokens")}
}

func (r *TokenRepo) Save(ctx *gin.Context, token string) error {
	_, err := r.collection.InsertOne(ctx, bson.M{
		"token":     token,
		"username":  ctx.Param("username"),
		"createdAt": time.Now(),
		//"expiresAt": expirationTime,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *TokenRepo) Find(ctx *gin.Context, token string) (string, error) {
	filter := bson.M{"token": token}

	var result bson.M
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (r *TokenRepo) Delete(ctx *gin.Context, token string) error {
	filter := bson.M{"token": token}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
