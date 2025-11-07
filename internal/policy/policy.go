package policy

import (
	"github.com/graphql-go/graphql"
	"github.com/joepk90/graphql-auth/internal/auth"
	log "github.com/sirupsen/logrus"
)

const (
	CreatePostRequest = "is_authorized_to_create_post"
	ReadPostRequest   = "is_authorized_to_read_post"
	UpdatePostRequest = "is_authorized_to_update_post"
	DeletePostRequest = "is_authorized_to_delete_post"
)

const (
	CreatePostErrorMsg = "Principal not authorized to get create post resource"
	ReadPostErrorMsg   = "Principal not authorized to get read post resource"
	UpdatePostErrorMsg = "Principal not authorized to get update post resource"
	DeletePostErrorMsg = "Principal not authorized to get delete post resource"
)

func (s *Service) IsAuthorizedToCreatePost(p graphql.ResolveParams) (interface{}, error) {
	_, err := s.authorize(p.Context, auth.CanCreatePost())
	if err != nil {
		s.metrics.ObserveQueryError(queryType, CreatePostRequest)

		errMsg := CreatePostErrorMsg
		log.WithError(err).Error(errMsg)
		return MapAuthenticationResponse(&errMsg), err
	}

	s.metrics.ObserveQuery(queryType, CreatePostRequest)
	return MapAuthenticationResponse(nil), nil
}

func (s *Service) IsAuthorizedToReadPost(p graphql.ResolveParams) (interface{}, error) {
	_, err := s.authorize(p.Context, auth.CanReadPost("*"))
	if err != nil {
		s.metrics.ObserveQueryError(queryType, ReadPostRequest)

		errorMsg := ReadPostErrorMsg
		log.WithError(err).Error(errorMsg)
		return MapAuthenticationResponse(&errorMsg), err
	}

	s.metrics.ObserveQuery(queryType, ReadPostRequest)
	return MapAuthenticationResponse(nil), nil
}

func (s *Service) IsAuthorizedToUpdatePost(p graphql.ResolveParams) (interface{}, error) {
	_, err := s.authorize(p.Context, auth.CanUpdatePost("*"))
	if err != nil {
		s.metrics.ObserveQueryError(queryType, UpdatePostRequest)

		errorMsg := UpdatePostErrorMsg
		log.WithError(err).Error(errorMsg)
		return MapAuthenticationResponse(&errorMsg), err
	}

	s.metrics.ObserveQuery(queryType, UpdatePostRequest)
	return MapAuthenticationResponse(nil), nil
}

func (s *Service) IsAuthorizedToDeletePost(p graphql.ResolveParams) (interface{}, error) {
	_, err := s.authorize(p.Context, auth.CanDeletePost("*"))
	if err != nil {
		s.metrics.ObserveQueryError(queryType, DeletePostRequest)

		errorMsg := DeletePostErrorMsg
		log.WithError(err).Error(errorMsg)
		return MapAuthenticationResponse(&errorMsg), err
	}

	s.metrics.ObserveQuery(queryType, DeletePostRequest)
	return MapAuthenticationResponse(nil), nil
}
