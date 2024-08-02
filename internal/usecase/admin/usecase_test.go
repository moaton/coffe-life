package admin

import (
	"coffe-life/config"
	"coffe-life/internal/dto"
	"coffe-life/internal/repository"
	"coffe-life/pkg/gorm/postgres"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
	pd "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

//go:generate mockgen -source=../../interfaces/usecase.go -destination=../../../mocks/usecase.go -package=mocks

var (
	errReturned = errors.New("failed")
)

func TestAdminSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
	suite.Run(t, new(UserSuite))
}

type AuthSuite struct {
	suite.Suite

	ctx          context.Context
	db           *sql.DB
	mock         sqlmock.Sqlmock
	deps         Dependencies
	req          dto.LoginRequest
	hashPassword []byte
}

func (s *AuthSuite) SetupTest() {
	s.ctx = context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		s.T().Errorf("failed to create sqlmock: %v", err)
	}
	s.db = db
	s.mock = mock

	cfg := &pd.Config{
		Conn: db,
	}

	gorm, err := gorm.Open(pd.Dialector{Config: cfg}, &gorm.Config{})
	if err != nil {
		s.T().Fatalf("failed to init db session: %v", err)
	}

	g := postgres.Gorm{DB: gorm}

	s.deps = Dependencies{
		Repository: repository.New(&g),
		JwtToken: config.JwtToken{
			JwtTokenSecret: "test",
			TokenTimeLimit: time.Hour,
		},
	}
	hash, err := bcrypt.GenerateFromPassword([]byte("test"), 10)
	if err != nil {
		s.T().Fatalf("failed to generate hash: %v", err)
	}
	s.hashPassword = hash
}

func (s *AuthSuite) TearDownTest() {
	s.db.Close()
}

func (s *AuthSuite) getLoginOkMocks(mock sqlmock.Sqlmock) {
	s.req = dto.LoginRequest{
		Username: "test",
		Password: "test",
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"first_name",
		"last_name",
		"password",
		"username",
		"is_first",
	}).AddRow(
		"e7de712a-06a0-45cd-ae62-1ba2efce9cb1",
		"test",
		"test",
		s.hashPassword,
		"test",
		false,
	)

	mock.ExpectQuery(`SELECT \* FROM "users"*`).
		WithArgs("test", 1).
		WillReturnRows(rows)
}

func (s *AuthSuite) getLoginErrMocks(mock sqlmock.Sqlmock) {
	s.req = dto.LoginRequest{
		Username: "test",
		Password: "test",
	}

	mock.ExpectQuery(`SELECT \* FROM "users"*`).
		WithArgs("test", 1).
		WillReturnError(errReturned)
}

func (s *AuthSuite) TestAuthSuite() {
	admin := New(s.deps)
	s.T().Run("login: ok", func(t *testing.T) {
		s.getLoginOkMocks(s.mock)

		resp, err := admin.Login(s.ctx, s.req)

		s.Assert().Equal(nil, err)
		s.Assert().True(resp.Token != "")
	})
	s.T().Run("login: err", func(t *testing.T) {
		s.getLoginErrMocks(s.mock)

		resp, err := admin.Login(s.ctx, s.req)

		var expectedResp *dto.AuthResponse
		s.Assert().Equal(fmt.Errorf("failed to login: %w", errReturned), err)
		s.Assert().Equal(expectedResp, resp)
	})
}

//	User

type UserSuite struct {
	suite.Suite

	ctx  context.Context
	db   *sql.DB
	mock sqlmock.Sqlmock
	deps Dependencies
	req  dto.CreateUserRequest
}

func (s *UserSuite) SetupTest() {
	s.ctx = context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		s.T().Errorf("failed to create sqlmock: %v", err)
	}
	s.db = db
	s.mock = mock

	cfg := &pd.Config{
		Conn: db,
	}

	gorm, err := gorm.Open(pd.Dialector{Config: cfg}, &gorm.Config{})
	if err != nil {
		s.T().Fatalf("failed to init db session: %v", err)
	}

	g := postgres.Gorm{DB: gorm}

	s.deps = Dependencies{
		Repository: repository.New(&g),
		JwtToken: config.JwtToken{
			JwtTokenSecret: "test",
			TokenTimeLimit: time.Hour,
		},
	}
}

func (s *UserSuite) TearDownTest() {
	s.db.Close()
}

func (s *UserSuite) getCreateUserOkMocks(mock sqlmock.Sqlmock) {
	s.req = dto.CreateUserRequest{
		FirstName: "test",
		LastName:  "test",
		Username:  "test",
		Password:  "test",
	}
	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id"}).AddRow("")
	mock.ExpectQuery(`INSERT INTO "users"*`).
		WithArgs("test", "test", "test", "test", false).
		WillReturnRows(rows)

	mock.ExpectCommit()
}

func (s *UserSuite) getCreateUserErrMocks(mock sqlmock.Sqlmock) {
	s.req = dto.CreateUserRequest{
		FirstName: "test",
		LastName:  "test",
		Username:  "test",
		Password:  "test",
	}
	mock.ExpectBegin()

	mock.ExpectQuery(`INSERT INTO "users"*`).
		WithArgs("test", "test", "test", "test", false).
		WillReturnError(errReturned)

	mock.ExpectRollback()
}

func (s *UserSuite) TestCreateUser() {
	admin := New(s.deps)
	s.T().Run("create user: ok", func(t *testing.T) {
		s.getCreateUserOkMocks(s.mock)

		err := admin.CreateUser(s.ctx, s.req)

		s.Assert().Equal(nil, err)
	})
	s.T().Run("create user: err", func(t *testing.T) {
		s.getCreateUserErrMocks(s.mock)

		err := admin.CreateUser(s.ctx, s.req)

		s.Assert().Equal(fmt.Errorf("failed to create user: %w", errReturned), err)
	})
}
