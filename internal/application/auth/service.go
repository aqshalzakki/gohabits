package auth

import (
	"errors"
	"gohabits/internal/domain/user"
	"gohabits/internal/domain/user/dto"
	"gohabits/internal/infra"
	"gohabits/internal/repository"
	"gohabits/internal/shared"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Service struct {
	db          *gorm.DB
	config      *infra.Config
	redisClient *infra.RedisClient
	userRepo    user.Repository
	jwtManager  *JWTManager
}

func NewAuthService(ctr *infra.Container, jwtManager *JWTManager) *Service {
	return &Service{
		db:          ctr.DB,
		config:      ctr.Config,
		redisClient: ctr.Redis,
		userRepo:    repository.NewUserRepository(ctr.DB),
		jwtManager:  jwtManager,
	}
}

func (a *Service) RegisterUser(ctx *fiber.Ctx, req dto.RegisterRequest) (user.User, dto.GenerateTokenResponse, error) {
	existingUser, err := a.userRepo.FindByEmail(ctx.Context(), req.Email)

	if existingUser != nil {
		// User with this email already exists
		return user.User{}, dto.GenerateTokenResponse{}, errors.New("user with this email already exists")
	}

	hashedPassword, err := shared.HashPassword(req.Password)
	if err != nil {
		return user.User{}, dto.GenerateTokenResponse{}, err
	}

	newUser := user.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Points:   0,
	}

	if err := a.db.Create(&newUser).Error; err != nil {
		return user.User{}, dto.GenerateTokenResponse{}, err
	}

	//generate token jwt
	token, err := a.jwtManager.GenerateToken(newUser.ID, newUser.Email)
	if err != nil {
		return user.User{}, dto.GenerateTokenResponse{}, err
	}

	return newUser, token, nil
}

func (a *Service) Login(ctx *fiber.Ctx, req dto.LoginRequest) (dto.LoginResponse, error) {
	foundUser, err := a.userRepo.FindByEmail(ctx.Context(), req.Email)
	if err != nil {
		return dto.LoginResponse{}, errors.New("user not found")
	}

	if isValid := shared.CheckPasswordHash(req.Password, foundUser.Password); isValid == false {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := a.jwtManager.GenerateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		AccessToken: token,
	}, nil
}

func (a *Service) Me(ctx *fiber.Ctx, userID uint) (dto.MeResponse, error) {
	foundUser, err := a.userRepo.FindByID(ctx.Context(), userID)
	if err != nil {
		return dto.MeResponse{}, errors.New("user not found")
	}

	return dto.MeResponse{
		Id:       foundUser.ID,
		Username: foundUser.Username,
		Email:    foundUser.Email,
		Points:   foundUser.Points,
	}, nil
}
