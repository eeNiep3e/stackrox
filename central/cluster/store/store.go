package store

import (
	"time"

	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/bolthelper"
	"github.com/boltdb/bolt"
)

const (
	clusterBucket       = "clusters"
	clusterStatusBucket = "clusters_status"
)

// Store provides storage functionality for alerts.
//go:generate mockery -name=Store
type Store interface {
	GetCluster(id string) (*v1.Cluster, bool, error)
	GetClusters() ([]*v1.Cluster, error)
	CountClusters() (int, error)
	AddCluster(cluster *v1.Cluster) (string, error)
	UpdateCluster(cluster *v1.Cluster) error
	RemoveCluster(id string) error
	UpdateClusterContactTime(id string, t time.Time) error
}

// New returns a new Store instance using the provided bolt DB instance.
func New(db *bolt.DB) Store {
	bolthelper.RegisterBucket(db, clusterBucket)
	bolthelper.RegisterBucket(db, clusterStatusBucket)
	return &storeImpl{
		DB: db,
	}
}
