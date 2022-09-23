/*
 * Copyright © 2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package service

import (
	"context"

	"github.com/durudex/durudex-auth-service/internal/domain"
)

// User auth service interface.
type User interface {
	// User SignUp.
	SignUp(ctx context.Context, input domain.UserSignUpInput) (domain.UserTokens, error)
	// User SignIn.
	SignIn(ctx context.Context, input domain.UserSignInInput) (domain.UserTokens, error)
	// User SignOut.
	SignOut(ctx context.Context, input domain.UserSignOutInput) error
	// Refresh user token.
	RefreshToken(ctx context.Context, input domain.UserRefreshTokenInput) (string, error)
}

// User service structure.
type UserService struct{}

// Creating a new user service.
func NewUserService() *UserService {
	return &UserService{}
}

// User SignUp.
func (s *UserService) SignUp(ctx context.Context, input domain.UserSignUpInput) (domain.UserTokens, error) {
	return domain.UserTokens{}, nil
}

// User SignIn.
func (s *UserService) SignIn(ctx context.Context, input domain.UserSignInInput) (domain.UserTokens, error) {
	return domain.UserTokens{}, nil
}

// User SignOut.
func (s *UserService) SignOut(ctx context.Context, input domain.UserSignOutInput) error {
	return nil
}

// Refresh user token.
func (s *UserService) RefreshToken(ctx context.Context, input domain.UserRefreshTokenInput) (string, error) {
	return "", nil
}
