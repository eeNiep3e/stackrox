// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type NotificationSchedulesStoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestNotificationSchedulesStore(t *testing.T) {
	suite.Run(t, new(NotificationSchedulesStoreSuite))
}

func (s *NotificationSchedulesStoreSuite) SetupTest() {
	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *NotificationSchedulesStoreSuite) TearDownTest() {
	s.testDB.Teardown(s.T())
}

func (s *NotificationSchedulesStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	notificationSchedule := &storage.NotificationSchedule{}
	s.NoError(testutils.FullInit(notificationSchedule, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundNotificationSchedule, exists, err := store.Get(ctx)
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNotificationSchedule)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, notificationSchedule))
	foundNotificationSchedule, exists, err = store.Get(ctx)
	s.NoError(err)
	s.True(exists)
	s.Equal(notificationSchedule, foundNotificationSchedule)

	foundNotificationSchedule, exists, err = store.Get(ctx)
	s.NoError(err)
	s.True(exists)
	s.Equal(notificationSchedule, foundNotificationSchedule)

	s.NoError(store.Delete(ctx))
	foundNotificationSchedule, exists, err = store.Get(ctx)
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNotificationSchedule)

	s.ErrorIs(store.Delete(withNoAccessCtx), sac.ErrResourceAccessDenied)

	notificationSchedule = &storage.NotificationSchedule{}
	s.NoError(testutils.FullInit(notificationSchedule, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))
	s.NoError(store.Upsert(ctx, notificationSchedule))

	foundNotificationSchedule, exists, err = store.Get(ctx)
	s.NoError(err)
	s.True(exists)
	s.Equal(notificationSchedule, foundNotificationSchedule)

	notificationSchedule = &storage.NotificationSchedule{}
	s.NoError(testutils.FullInit(notificationSchedule, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))
	s.NoError(store.Upsert(ctx, notificationSchedule))

	foundNotificationSchedule, exists, err = store.Get(ctx)
	s.NoError(err)
	s.True(exists)
	s.Equal(notificationSchedule, foundNotificationSchedule)
}
