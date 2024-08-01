package userservice

import (
	"UserService/internal/jwt"
	"UserService/internal/mongodb"
	pkg "UserService/internal/pkg/email"
	redisdb "UserService/internal/redis"
	pb "UserService/userproto"
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type UserService struct {
	*pb.UnimplementedUserServiceServer
	mongodb *mongodb.UserMongoDb
	redis   *redisdb.RedisClient
}

func NewUserService(mongodb *mongodb.UserMongoDb, redis *redisdb.RedisClient) *UserService {
	return &UserService{mongodb: mongodb, redis: redis}
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
    if req == nil {
        return nil, fmt.Errorf("request cannot be nil")
    }

    bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    hashedPassword := string(bytes)

    token, err := jwt.GenerateJWTToken(req.Email)
    if err != nil {
        return nil, err
    }

    if req.Profile == nil {
        return nil, fmt.Errorf("profile cannot be nil")
    }

    code := 10000 + rand.Intn(90000)
    err = pkg.SendEmail(req.Email, pkg.SendClientCode(code, req.UserName))
    if err != nil {
        log.Println("ERROR: sending email to user !!", err)
    }

    userData := map[string]interface{}{
        "userName":       req.UserName,
        "email":          req.Email,
        "password":       hashedPassword,
        "profileName":    req.Profile.Name,
        "profileAddress": req.Profile.Address,
        "code":           code,
    }

    if u.redis == nil {
        return nil, fmt.Errorf("redis client is not initialized")
    }

    err = u.redis.SetHash(req.Email, userData)
    if err != nil {
        return nil, fmt.Errorf("failed to save user data in Redis: %v", err)
    }

    return &pb.UserResponse{Message: "User created successfully", Token: token}, nil
}


func (u *UserService) VerifyCode(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	user, err := u.redis.VerifyEmail(ctx, req.Email, req.Password)
	if err != nil {
		return nil, fmt.Errorf("error code")
	}

	err = u.mongodb.AddUser(ctx, user.UserName, user.Email, user.Password, user.Profile.Name, user.Profile.Address)
	if err != nil {
		return nil, err
	}

	return &pb.Res{Message: "verify successfully"}, nil
}

func (u *UserService) GetByIdUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	userData, err := u.redis.GetHash(req.UserId)
	if err == nil && len(userData) > 0 {
		user := &pb.User{
			UserName: userData["userName"],
			Email:    userData["email"],
			Profile: &pb.Profile{
				Name:    userData["profileName"],
				Address: userData["profileAddress"],
			},
		}
		return user, nil
	}

	user, err := u.mongodb.GetUserId(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetUsers(ctx context.Context, req *pb.Empty) (*pb.ListUser, error) {
	users, err := u.mongodb.GetAllUser(ctx)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}

	// Redis'ga saqlash
	for _, user := range users.User {
		userData := map[string]interface{}{
			"userName":       user.UserName,
			"email":          user.Email,
			"password":       user.Password,
			"profileName":    user.Profile.Name,
			"profileAddress": user.Profile.Address,
		}
		err := u.redis.SetHash(user.Email, userData)
		if err != nil {
			log.Printf("failed to save user data in Redis: %v", err)
		}
	}

	return users, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.Request) (*pb.UpdateResponse, error) {
	resp, err := u.mongodb.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UserService) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.UpdateResponse, error) {
	resp, err := u.mongodb.UpdateEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	// Redis'dagi eski ma'lumotlarni o'chirish va yangi ma'lumotlarni qo'shish
	err = u.redis.Delete(req.OldEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to delete old user data in Redis: %v", err)
	}

	userData := map[string]interface{}{
		"email": req.NewEmail,
	}
	err = u.redis.SetHash(req.NewEmail, userData)
	if err != nil {
		return nil, fmt.Errorf("failed to save new user data in Redis: %v", err)
	}

	return resp, nil
}

func (u *UserService) UpdatePassword(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateResponse, error) {
	resp, err := u.mongodb.UpdatePassword(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.UpdateResponse, error) {
	resp, err := u.mongodb.DeleteUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserResponse, error) {
	user, err := u.mongodb.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	token, err := jwt.GenerateJWTToken(user.Email)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Message: "Login successful", Token: token}, nil
}
