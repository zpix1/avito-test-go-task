package service

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

func normalizePercent(percent uint32) (uint32, error) {
	if percent > 100 {
		return 0, fmt.Errorf("can't use %d as percent", percent)
	}
	return uint32((math.MaxUint32 * float64(percent)) / 100.0), nil
}

func (s *Service) CreateSlug(slugName string, autoAddPercent uint32) (int, error) {
	normalizedPercent, err := normalizePercent(autoAddPercent)
	logrus.Warn("converted ", autoAddPercent, " to ", normalizedPercent)
	if err != nil {
		return 0, err
	}
	return s.repository.CreateSlug(slugName, normalizedPercent)
}

func (s *Service) DeleteSlug(slugName string) error {
	return s.repository.DeleteSlug(slugName)
}

func (s *Service) UpdateUserSlugs(userId int, addSlugNames []string, deleteSlugNames []string, ttl uint64) error {
	if ttl == 0 {
		return s.repository.UpdateUserSlugs(userId, addSlugNames, deleteSlugNames, time.Time{})
	}
	validUntil := time.Now().Add(time.Duration(ttl) * time.Second)
	return s.repository.UpdateUserSlugs(userId, addSlugNames, deleteSlugNames, validUntil)
}

func (s *Service) autoAddSlugs(userId int) error {
	autoAddSlugNames, err := s.repository.GetAutoAddSlugs(userId)
	logrus.Warn("autoAddSlugnames", autoAddSlugNames)
	if err != nil {
		return err
	}

	var toAddSlugNames []string
	for _, sn := range autoAddSlugNames {
		hash := md5.New()
		_, err := hash.Write([]byte(fmt.Sprintf("%s%d", sn.SlugName, userId)))
		if err != nil {
			return err
		}
		value := binary.BigEndian.Uint32(hash.Sum(nil))
		// so user won and his userId is chosen to have slugName
		logrus.Warn("value vs auto add weight ", value, " ", sn.AutoAddWeight, ", slug=", sn.SlugName)
		if value < sn.AutoAddWeight {
			toAddSlugNames = append(toAddSlugNames, sn.SlugName)
		}
	}
	logrus.Warn(userId, " ", toAddSlugNames, " kek")
	return s.repository.UpdateUserSlugs(userId, toAddSlugNames, nil, time.Time{})
}

func (s *Service) GetUserSlugs(userId int) ([]string, error) {
	err := s.autoAddSlugs(userId)
	if err != nil {
		logrus.Warn("failed to autoAddSlugs ", err)
		return nil, err
	}
	return s.repository.GetUserSlugs(userId)
}
