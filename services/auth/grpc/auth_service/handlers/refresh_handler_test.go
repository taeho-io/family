package handlers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/auth/config"
	mockToken "github.com/taeho-io/family/services/auth/mocks/token"
	"github.com/taeho-io/family/services/auth/token"
)

func TestRefreshHandler(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSvc := token.New(cfg)
	refreshToken, _ := tokenSvc.NewRefreshToken(testAccountId)
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}
	res, err := Refresh(cfg, tokenSvc)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRefreshHandler_Error_InvalidRefreshToken(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSvc := token.New(cfg)
	req := &auth.RefreshRequest{
		RefreshToken: "invalid_token",
	}
	res, err := Refresh(cfg, tokenSvc)(ctx, req)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestRefreshHandler_NewAccessToken_Error(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	refreshToken, _ := token.New(cfg).NewRefreshToken(testAccountId)
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}
	tokenSvc := new(mockToken.IFace)
	tokenSvc.On("NewAccessToken", testAccountId).Return("", errors.New("failed"))
	_, err := Refresh(cfg, tokenSvc)(ctx, req)
	assert.NotNil(t, err)
}
