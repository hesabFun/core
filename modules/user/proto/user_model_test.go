package userpb

import (
	"context"
	"testing"
	"time"

	"elbix.dev/engine/pkg/token/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"elbix.dev/engine/pkg/mockery"
)

func compareUser(t *testing.T, u1, u2 *User) {
	assert.Equal(t, u1.Id, u2.Id)
	assert.Equal(t, u1.Email, u2.Email)
	// assert.Equal(t, u1.Password, u2.Password)
	assert.Equal(t, u1.Status, u2.Status)
}

func TestManager_RegisterUser(t *testing.T) {
	ctx := context.Background()
	defer mockery.Start(ctx, t)()

	m := NewManager()
	u, err := m.RegisterUser(ctx, "valid@gmail.com", "name", "bita123")
	require.NoError(t, err)
	assert.True(t, u.VerifyPassword("bita123"))

	u1, err := m.GetUserByPrimary(ctx, u.Id)
	require.NoError(t, err)
	compareUser(t, u, u1)

	// duplicate email
	u, err = m.RegisterUser(ctx, "valid@gmail.com", "name", "bita123")
	assert.Nil(t, u)
	assert.Error(t, err)
}

func TestManager_FindUserByEmail(t *testing.T) {
	ctx := context.Background()
	defer mockery.Start(ctx, t)()

	m := NewManager()
	u, err := m.RegisterUser(ctx, "valid@gmail.com", "name", "bita123")
	require.NoError(t, err)

	u1, err := m.FindUserByEmail(ctx, "valid@gmail.com")
	require.NoError(t, err)

	assert.Equal(t, u, u1)

	u2, err := m.FindUserByEmail(ctx, "notvalid@gmail.com")
	assert.Error(t, err)
	assert.Nil(t, u2)
}

func TestManager_FindUserByEmailPassword(t *testing.T) {
	ctx := context.Background()
	defer mockery.Start(ctx, t)()

	m := NewManager()
	u, err := m.RegisterUser(ctx, "valid@gmail.com", "name", "bita123")
	require.NoError(t, err)

	u1, err := m.FindUserByEmailPassword(ctx, "valid@gmail.com", "bita123")
	require.NoError(t, err)

	assert.Equal(t, u, u1)

	u2, err := m.FindUserByEmailPassword(ctx, "valid@gmail.com", "NOPASSS")
	require.Error(t, err)
	require.Nil(t, u2)

	u2, err = m.FindUserByEmailPassword(ctx, "notvalid@gmail.com", "NOPASSS")
	require.Error(t, err)
	require.Nil(t, u2)

	u1.Status = UserStatus_USER_STATUS_BANNED
	err = m.ChangePassword(ctx, u1, "newpass")
	require.NoError(t, err)

	u1, err = m.FindUserByEmailPassword(ctx, "valid@gmail.com", "newpass")
	require.Error(t, err)
	require.Nil(t, u1)
}

func TestManager_FindUserByIndirectToken(t *testing.T) {
	ctx := context.Background()
	defer mockery.Start(ctx, t)()

	m := NewManager()
	u, err := m.RegisterUser(ctx, "valid@gmail.com", "name", "bita123")
	require.NoError(t, err)

	token := m.CreateToken(ctx, u, time.Hour)

	u1, err := m.FindUserByIndirectToken(ctx, token)
	require.NoError(t, err)
	compareUser(t, u, u1)

	m.DeleteToken(ctx, token)

	u1, err = m.FindUserByIndirectToken(ctx, token)
	require.Error(t, err)
	assert.Nil(t, u1)
}

func TestManager_ForgottenToken(t *testing.T) {
	ctx := context.Background()
	defer mockery.Start(ctx, t)()

	m := NewManager()
	u, err := m.RegisterUser(ctx, "valid@gmail.com", "name", "bita123")
	require.NoError(t, err)

	require.Error(t, m.VerifyForgottenToken(ctx, u, "RANDOOOOM"))

	token, tm, err := m.CreateForgottenToken(ctx, u)
	require.NoError(t, err)

	token2, tm2, err := m.CreateForgottenToken(ctx, u)
	require.NoError(t, err)
	assert.Equal(t, token, token2)
	assert.True(t, tm >= tm2)

	require.Error(t, m.VerifyForgottenToken(ctx, u, "RANDOOOOM"))
	require.NoError(t, m.VerifyForgottenToken(ctx, u, token))
}

func TestManager_ReloadUser(t *testing.T) {
	ctx := context.Background()
	defer mockery.Start(ctx, t)()

	m := NewManager()
	u, err := m.RegisterUser(ctx, "valid@gmail.com", "name", "bita123")
	require.NoError(t, err)
	assert.True(t, u.VerifyPassword("bita123"))

	u2 := &User{Id: u.Id}
	require.NoError(t, m.ReloadUser(ctx, u2))

	compareUser(t, u, u2)
}

func init() {
	SetProvider(mock.NewMockStorage())
}
