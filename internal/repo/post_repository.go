package repo

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mindlab/internal/model"
)

type PostRepository interface {
	Index(ctx *gin.Context) ([]*model.Post, error)
	Create(ctx *gin.Context, post *model.Post) (*model.Post, error)
	Show(ctx *gin.Context, id string) (*model.Post, error)
	Update(ctx *gin.Context, post *model.Post) (*model.Post, error)
	Delete(ctx *gin.Context, id string) error
}

type postRepo struct {
	collection *mongo.Collection
}

func NewPostRepository(db *mongo.Database) PostRepository {
	return &postRepo{collection: db.Collection("posts")}
}

func (r *postRepo) Index(ctx *gin.Context) ([]*model.Post, error) {
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var posts []*model.Post
	for cur.Next(ctx) {
		var u model.Post
		if err := cur.Decode(&u); err != nil {
			return nil, err
		}
		posts = append(posts, &u)
	}

	return posts, cur.Err()
}

func (r *postRepo) Create(ctx *gin.Context, post *model.Post) (*model.Post, error) {
	result, err := r.collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		post.ID = oid.Hex()
	} else {
		return nil, fmt.Errorf("failed to convert InsertedID to ObjectID")
	}
	logrus.Info("Success insert post", result)
	return post, err
}

func (r *postRepo) Show(ctx *gin.Context, id string) (*model.Post, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	filter := bson.M{"_id": oid}

	post := &model.Post{}

	err = r.collection.FindOne(ctx, filter).Decode(post)

	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, fmt.Errorf("no post found with ID: %s", id)
		}
		return nil, err
	}

	return post, nil
}

func (r *postRepo) Update(ctx *gin.Context, post *model.Post) (*model.Post, error) {
	oid, err := primitive.ObjectIDFromHex(post.ID)

	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}
	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"title": post.Title,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no document found with ID %s", post.ID)
	}

	return post, nil
}

func (r *postRepo) Delete(ctx *gin.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("invalid ID format: %w", err)
	}

	filter := bson.M{"_id": oid}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting post %s: %w", id, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %s", id)
	}

	return nil
}
