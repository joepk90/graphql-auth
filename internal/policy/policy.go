package policy

import (
	"github.com/graphql-go/graphql"
	"github.com/joepk90/graphql-auth/internal/auth"
	log "github.com/sirupsen/logrus"
)

const (
	IsAuthorizedToCreatePost = "is_authorized_to_create_post"
	IsAuthorizedToReadPost   = "is_authorized_to_read_post"
	IsAuthorizedToUpdatePost = "is_authorized_to_update_post"
	IsAuthorizedToDeletePost = "is_authorized_to_delete_post"
)

func (s *Service) IsAuthorizedToCreatePost(p graphql.ResolveParams) (interface{}, error) {
	_, err := s.authorize(p.Context, auth.CanCreatePost())
	if err != nil {
		s.metrics.ObserveQueryError(queryType, IsAuthorizedToCreatePost)

		errorMsg := "Principal not authorized to get create post resource"
		log.WithError(err).Error(errorMsg)
		return MapAuthenticationResponse(&errorMsg), err
	}

	s.metrics.ObserveQuery(queryType, IsAuthorizedToCreatePost)
	return MapAuthenticationResponse(nil), nil
}

func (s *Service) IsAuthorizedToReadPost(p graphql.ResolveParams) (interface{}, error) {
	_, err := s.authorize(p.Context, auth.CanReadPost("*"))
	if err != nil {
		s.metrics.ObserveQueryError(queryType, IsAuthorizedToReadPost)

		errorMsg := "Principal not authorized to get read post resource"
		log.WithError(err).Error(errorMsg)
		return MapAuthenticationResponse(&errorMsg), err
	}

	s.metrics.ObserveQuery(queryType, IsAuthorizedToReadPost)
	return MapAuthenticationResponse(nil), nil
}

func (s *Service) IsAuthorizedToUpdatePost(p graphql.ResolveParams) (interface{}, error) {
	_, err := s.authorize(p.Context, auth.CanUpdatePost("*"))
	if err != nil {
		s.metrics.ObserveQueryError(queryType, IsAuthorizedToUpdatePost)

		errorMsg := "Principal not authorized to get update post resource"
		log.WithError(err).Error(errorMsg)
		return MapAuthenticationResponse(&errorMsg), err
	}

	s.metrics.ObserveQuery(queryType, IsAuthorizedToUpdatePost)
	return MapAuthenticationResponse(nil), nil
}

func (s *Service) IsAuthorizedToDeletePost(p graphql.ResolveParams) (interface{}, error) {
	_, err := s.authorize(p.Context, auth.CanDeletePost("*"))
	if err != nil {
		s.metrics.ObserveQueryError(queryType, IsAuthorizedToDeletePost)

		errorMsg := "Principal not authorized to get delete post resource"
		log.WithError(err).Error(errorMsg)
		return MapAuthenticationResponse(&errorMsg), err
	}

	s.metrics.ObserveQuery(queryType, IsAuthorizedToDeletePost)
	return MapAuthenticationResponse(nil), nil
}
