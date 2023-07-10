package service

// import (
// 	"errors"
// 	"kmipn-2023/model"
// 	repo "kmipn-2023/repository"
// 	"time"

// 	"github.com/golang-jwt/jwt"
// )

// type SellerService interface {
// 	Register(seller *model.Seller) (model.Seller, error)
// 	Login(seller *model.Seller) (token *string, err error)
// }

// type sellerService struct {
// 	sellerRepo   repo.SellerRepository
// 	sessionsRepo repo.SessionRepository
// }

// func NewSellerService(sellerRepository repo.SellerRepository, sessionsRepo repo.SessionRepository) SellerService {
// 	return &sellerService{sellerRepo: sellerRepository, sessionsRepo: sessionsRepo}
// }

// func (s *sellerService) Register(seller *model.Seller) (model.Seller, error) {
// 	dbSeller, err := s.sellerRepo.GetSellerByEmail(seller.Email)
// 	if err != nil {
// 		return *seller, err
// 	}

// 	if dbSeller.Email != "" || dbSeller.ID != 0 {
// 		return *seller, errors.New("email already exists")
// 	}

// 	seller.CreatedAt = time.Now()

// 	newSeller, err := s.sellerRepo.CreateSeller(*seller)
// 	if err != nil {
// 		return *seller, err
// 	}

// 	return newSeller, nil
// }

// func (s *sellerService) Login(seller *model.Seller) (token *string, err error) {
// 	dbSeller, err := s.sellerRepo.GetSellerByEmail(seller.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if dbSeller.Email == "" || dbSeller.ID == 0 {
// 		return nil, errors.New("seller not found")
// 	}

// 	if seller.Password != dbSeller.Password {
// 		return nil, errors.New("wrong email or password")
// 	}

// 	expirationTime := time.Now().Add(20 * time.Minute)
// 	claims := &model.Claims{
// 		Email: dbSeller.Email,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := t.SignedString(model.JwtKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	session := model.Session{
// 		Token:  tokenString,
// 		Email:  seller.Email,
// 		Expiry: expirationTime,
// 	}

// 	_, err = s.sessionsRepo.SessionAvailEmail(session.Email)
// 	if err != nil {
// 		err = s.sessionsRepo.AddSessions(session)
// 	} else {
// 		err = s.sessionsRepo.UpdateSessions(session)
// 	}

// 	return &tokenString, nil
// }
