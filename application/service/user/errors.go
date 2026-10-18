package user

import apperrors "Momentum/application/errors"

var ErrNotFound = apperrors.New(apperrors.CodeNotFound, 404, "user not found")
