package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var InternalServerError = status.Errorf(codes.Internal, "Internal server error")
