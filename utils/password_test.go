package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CompareHashAndPass(hashedPassword, password)
	require.NoError(t, err)

	wrong := RandomString(6)
	err = CompareHashAndPass(hashedPassword, wrong)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	newHash, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, newHash)
	require.NotEqual(t, newHash, hashedPassword)

}
