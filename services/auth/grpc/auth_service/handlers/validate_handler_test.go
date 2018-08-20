package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/services/auth/config"
	"github.com/taeho-io/family/services/auth/token"
)

func TestValidateHandler(t *testing.T) {
	ctx := context.Background()
	settings := config.NewSettings()
	cfg := config.New(settings)
	tokenSrv := token.New(cfg)
	accessToken, _ := tokenSrv.NewAccessToken(testAccountId)
	req := &auth.ValidateRequest{
		AccessToken: accessToken,
	}
	res, err := Validate(tokenSrv)(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}