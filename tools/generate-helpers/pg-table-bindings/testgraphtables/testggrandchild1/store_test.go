// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/protoassert"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stretchr/testify/suite"
)

type TestGGrandChild1StoreSuite struct {
	suite.Suite
	store  Store
	testDB *pgtest.TestPostgres
}

func TestTestGGrandChild1Store(t *testing.T) {
	suite.Run(t, new(TestGGrandChild1StoreSuite))
}

func (s *TestGGrandChild1StoreSuite) SetupSuite() {

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.DB)
}

func (s *TestGGrandChild1StoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE test_g_grand_child1 CASCADE")
	s.T().Log("test_g_grand_child1", tag)
	s.store = New(s.testDB.DB)
	s.NoError(err)
}

func (s *TestGGrandChild1StoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
}

func (s *TestGGrandChild1StoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	testGGrandChild1 := &storage.TestGGrandChild1{}
	s.NoError(testutils.FullInit(testGGrandChild1, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundTestGGrandChild1, exists, err := store.Get(ctx, testGGrandChild1.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestGGrandChild1)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, testGGrandChild1))
	foundTestGGrandChild1, exists, err = store.Get(ctx, testGGrandChild1.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), testGGrandChild1, foundTestGGrandChild1)

	testGGrandChild1Count, err := store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(1, testGGrandChild1Count)
	testGGrandChild1Count, err = store.Count(withNoAccessCtx, search.EmptyQuery())
	s.NoError(err)
	s.Zero(testGGrandChild1Count)

	testGGrandChild1Exists, err := store.Exists(ctx, testGGrandChild1.GetId())
	s.NoError(err)
	s.True(testGGrandChild1Exists)
	s.NoError(store.Upsert(ctx, testGGrandChild1))
	s.ErrorIs(store.Upsert(withNoAccessCtx, testGGrandChild1), sac.ErrResourceAccessDenied)

	foundTestGGrandChild1, exists, err = store.Get(ctx, testGGrandChild1.GetId())
	s.NoError(err)
	s.True(exists)
	protoassert.Equal(s.T(), testGGrandChild1, foundTestGGrandChild1)

	s.NoError(store.Delete(ctx, testGGrandChild1.GetId()))
	foundTestGGrandChild1, exists, err = store.Get(ctx, testGGrandChild1.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestGGrandChild1)
	s.NoError(store.Delete(withNoAccessCtx, testGGrandChild1.GetId()))

	var testGGrandChild1s []*storage.TestGGrandChild1
	var testGGrandChild1IDs []string
	for i := 0; i < 200; i++ {
		testGGrandChild1 := &storage.TestGGrandChild1{}
		s.NoError(testutils.FullInit(testGGrandChild1, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		testGGrandChild1s = append(testGGrandChild1s, testGGrandChild1)
		testGGrandChild1IDs = append(testGGrandChild1IDs, testGGrandChild1.GetId())
	}

	s.NoError(store.UpsertMany(ctx, testGGrandChild1s))

	testGGrandChild1Count, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(200, testGGrandChild1Count)

	s.NoError(store.DeleteMany(ctx, testGGrandChild1IDs))

	testGGrandChild1Count, err = store.Count(ctx, search.EmptyQuery())
	s.NoError(err)
	s.Equal(0, testGGrandChild1Count)
}
