package handler

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sig-agro/services/user-service/internal/repository"

	"google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

	pb "github.com/sig-agro/api/proto/user"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo      *repository.Repository
	jwtSecret string
}

func NewUserService(repo *repository.Repository, jwtSecret string) *UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Register request for email: %s\n", req.Email)

	// Hash password
	hash := sha256.Sum256([]byte(req.Password))
	passwordHash := fmt.Sprintf("%x", hash)

	userID, err := s.repo.CreateUser(ctx, req.Email, passwordHash, req.FullName, req.Phone)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return &pb.RegisterResponse{Message: "Error registering user"}, err
	}

	// Add default role
	s.repo.AddUserRole(ctx, userID, "user")

	return &pb.RegisterResponse{
		UserId:   userID,
		Email:    req.Email,
		FullName: req.FullName,
		Message:  "User registered successfully",
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
    log.Printf("Login request for email: %s\n", req.Email)

    userID, storedHash, err := s.repo.GetUserByEmail(ctx, req.Email)
    if err != nil {
        return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
    }

    // Verify password
    hash := sha256.Sum256([]byte(req.Password))
    passwordHash := fmt.Sprintf("%x", hash)

    if passwordHash != storedHash {
        return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
    }

    // Get user roles
    roles, err := s.repo.GetUserRoles(ctx, userID)
    if err != nil {
        log.Printf("Error fetching roles: %v\n", err)
        roles = []string{"user"}
    }

    // Generate JWT token
    expiresAt := time.Now().Add(time.Hour).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "email":   req.Email,
        "roles":   roles,
        "exp":     expiresAt,
    })

    tokenString, err := token.SignedString([]byte(s.jwtSecret))
    if err != nil {
        return nil, status.Error(codes.Internal, "Error generating token")
    }

    return &pb.LoginResponse{
        UserId:    userID,
        Email:     req.Email,
        Token:     tokenString,
        ExpiresAt: expiresAt,
    }, nil
}

func (s *UserService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	log.Printf("Validating token\n")

	token, err := jwt.ParseWithClaims(req.Token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return &pb.ValidateTokenResponse{
			Valid:   false,
			Message: "Invalid token",
		}, nil
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int64(claims["user_id"].(float64))
	email := claims["email"].(string)

	roles := []string{}
	if rolesList, ok := claims["roles"].([]interface{}); ok {
		for _, role := range rolesList {
			roles = append(roles, role.(string))
		}
	}

	return &pb.ValidateTokenResponse{
		Valid:   true,
		UserId:  userID,
		Email:   email,
		Roles:   roles,
		Message: "Token is valid",
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("GetUser request for user ID: %d\n", req.UserId)

	email, fullName, phone, err := s.repo.GetUserByID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	roles, err := s.repo.GetUserRoles(ctx, req.UserId)
	if err != nil {
		log.Printf("Error fetching roles: %v\n", err)
		roles = []string{"user"}
	}

	return &pb.GetUserResponse{
		UserId:   req.UserId,
		Email:    email,
		FullName: fullName,
		Phone:    phone,
		Roles:    roles,
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	log.Printf("ListUsers request\n")

	users, err := s.repo.ListUsers(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	resp := &pb.ListUsersResponse{Total: int32(len(users))}
	for _, user := range users {
		pbUser := &pb.GetUserResponse{
			UserId:   user["id"].(int64),
			Email:    user["email"].(string),
			FullName: user["full_name"].(string),
			Phone:    user["phone"].(string),
		}
		resp.Users = append(resp.Users, pbUser)
	}

	return resp, nil
}
