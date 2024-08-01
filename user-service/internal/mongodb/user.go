package mongodb

import (
	"UserService/internal/models"
	pb "UserService/userproto"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserMongoDb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUser() (*UserMongoDb, error) {
	clientOptions := options.Client().ApplyURI("mongodb://user-mongodb:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database("User").Collection("users")
	return &UserMongoDb{client: client, collection: collection}, nil
}

func (u *UserMongoDb) AddUser(ctx context.Context, username, email, password, name, address string) error {
	res := u.collection.FindOne(ctx, bson.M{"email": email})
	if res.Err() == nil {
		return fmt.Errorf("email already registered")
	}
	if res.Err() != mongo.ErrNoDocuments {
		return res.Err()
	}

	doc := bson.M{
		"user_name": username,
		"email":     email,
		"password":  password,
		"profile": bson.M{
			"name":    name,
			"address": address,
		},
	}

	_, err := u.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserMongoDb) GetUserByEmail(ctx context.Context, email string) (*pb.User, error) {
	res := u.collection.FindOne(ctx, bson.M{"email": email})

	var user pb.User
	err := res.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserMongoDb) GetUserId(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
    objectId, err := primitive.ObjectIDFromHex(req.UserId)
    if err != nil {
        return nil, err
    }

    res := u.collection.FindOne(ctx, bson.M{"_id": objectId})

    var user pb.User
    var mongoUser models.User
    err = res.Decode(&mongoUser)
    if err != nil {
        return nil, fmt.Errorf("decode error: %v", err)
    }
    user.UserId = mongoUser.ID
    user.UserName = mongoUser.UserName
    user.Email = mongoUser.Email
    user.Password = mongoUser.Password
    user.Profile = &pb.Profile{
        Name:    mongoUser.Profile.Name,
        Address: mongoUser.Profile.Address,
    }

    return &user, nil
}

func (u *UserMongoDb) GetAllUser(ctx context.Context) (*pb.ListUser, error) {
	cursor, err := u.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*pb.User

	for cursor.Next(ctx) {
		var user pb.User
		var mongoUser models.User
		err := cursor.Decode(&mongoUser)
		if err != nil {
			return nil, err
		}

		user.UserId = mongoUser.ID
		user.UserName = mongoUser.UserName
		user.Email = mongoUser.Email
		user.Password = mongoUser.Password
		user.Profile = &pb.Profile{
			Name:    mongoUser.Profile.Name,
			Address: mongoUser.Profile.Address,
		}

		users = append(users, &user)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListUser{User: users}, nil
}

func (u *UserMongoDb) UpdateUser(ctx context.Context, req *pb.Request) (*pb.UpdateResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"user_name": req.UserName,
			"profile": bson.M{
				"name":    req.Profile.Name,
				"address": req.Profile.Address,
			},
		},
	}
	_, err = u.collection.UpdateOne(ctx, bson.M{"_id": objectId}, update)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{Message: "User updated successfully"}, nil
}

func (u *UserMongoDb) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.UpdateResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"email": req.NewEmail,
		},
	}

	_, err = u.collection.UpdateOne(ctx, bson.M{"_id": objectId, "email": req.OldEmail}, update)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{Message: "Email updated successfully"}, nil
}

func (u *UserMongoDb) UpdatePassword(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		fmt.Println("ObjectId error:", err)
		return nil, err
	}

	var user struct {
		Password string `bson:"password"`
	}
	filter := bson.M{"_id": objectId}
	err = u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))
	if err != nil {
		fmt.Println("oldpassword error:", err)
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Hashed password error:", err)
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"password": string(hashedPassword),
		},
	}

	_, err = u.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pb.UpdateResponse{Message: "Password updated successfully"}, nil
}

func (u *UserMongoDb) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.UpdateResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, err
	}

	res := u.collection.FindOne(ctx, bson.M{"_id": objectId})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, res.Err()
	}

	_, err = u.collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{Message: "User deleted successfully"}, nil
}
