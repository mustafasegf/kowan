package links

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"github.com/mustafasegf/go-shortener/entity"
)

type Service struct {
	repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetLinkByURL(shortUrl string) (result *entity.CreateLinkRequest, err error) {
	longURL, err := s.repo.RedisGetLinkByURL(shortUrl)
	if err != redis.Nil {
		result = &entity.CreateLinkRequest{}
		result.LongUrl = longURL
		result.ShortUrl = shortUrl
		err = nil
		return
	}

	data, err := s.repo.GetLinkByURL(shortUrl)
	if err != nil {
		return
	}

	err = s.repo.RedisSetURL(shortUrl, data.LongUrl)
	if err != redis.Nil {
		log.Print(err)
		err = nil
	}
	result = &entity.CreateLinkRequest{}
	err = copier.Copy(&result, data)

	return
}

func (s *Service) InsertURL(req entity.CreateLinkRequest) (err error) {
	err = s.repo.InsertURL(req.ShortUrl, req.LongUrl)
	s.repo.RedisSetURL(req.ShortUrl, req.LongUrl)
	return
}
