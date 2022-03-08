// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
)

const (
	countStmt  = "SELECT COUNT(*) FROM process_indicators"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM process_indicators WHERE Id = $1)"

	getStmt     = "SELECT serialized FROM process_indicators WHERE Id = $1"
	deleteStmt  = "DELETE FROM process_indicators WHERE Id = $1"
	walkStmt    = "SELECT serialized FROM process_indicators"
	getIDsStmt  = "SELECT Id FROM process_indicators"
	getManyStmt = "SELECT serialized FROM process_indicators WHERE Id = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM process_indicators WHERE Id = ANY($1::text[])"
)

var (
	log = logging.LoggerForModule()

	table = "process_indicators"
)

type Store interface {
	Count() (int, error)
	Exists(id string) (bool, error)
	Get(id string) (*storage.ProcessIndicator, bool, error)
	Upsert(obj *storage.ProcessIndicator) error
	UpsertMany(objs []*storage.ProcessIndicator) error
	Delete(id string) error
	GetIDs() ([]string, error)
	GetMany(ids []string) ([]*storage.ProcessIndicator, []int, error)
	DeleteMany(ids []string) error

	Walk(fn func(obj *storage.ProcessIndicator) error) error

	AckKeysIndexed(keys ...string) error
	GetKeysToIndex() ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableProcessIndicators(db *pgxpool.Pool) {
	table := `
create table if not exists process_indicators (
    Id varchar,
    DeploymentId varchar,
    ContainerName varchar,
    PodId varchar,
    PodUid varchar,
    Signal_Id varchar,
    Signal_ContainerId varchar,
    Signal_Time timestamp,
    Signal_Name varchar,
    Signal_Args varchar,
    Signal_ExecFilePath varchar,
    Signal_Pid numeric,
    Signal_Uid numeric,
    Signal_Gid numeric,
    Signal_Lineage text[],
    Signal_Scraped bool,
    ClusterId varchar,
    Namespace varchar,
    ContainerStartTime timestamp,
    ImageId varchar,
    serialized bytea,
    PRIMARY KEY(Id)
)
`

	_, err := db.Exec(context.Background(), table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{}
	for _, index := range indexes {
		if _, err := db.Exec(context.Background(), index); err != nil {
			panic(err)
		}
	}

	createTableProcessIndicatorsLineageInfo(db)
}

func createTableProcessIndicatorsLineageInfo(db *pgxpool.Pool) {
	table := `
create table if not exists process_indicators_LineageInfo (
    process_indicators_Id varchar,
    idx numeric,
    ParentUid numeric,
    ParentExecFilePath varchar,
    PRIMARY KEY(process_indicators_Id, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (process_indicators_Id) REFERENCES process_indicators(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(context.Background(), table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{

		"create index if not exists processIndicatorsLineageInfo_idx on process_indicators_LineageInfo using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(context.Background(), index); err != nil {
			panic(err)
		}
	}

}

func insertIntoProcessIndicators(tx pgx.Tx, obj *storage.ProcessIndicator) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start

		obj.GetId(),

		obj.GetDeploymentId(),

		obj.GetContainerName(),

		obj.GetPodId(),

		obj.GetPodUid(),

		obj.GetSignal().GetId(),

		obj.GetSignal().GetContainerId(),

		pgutils.NilOrStringTimestamp(obj.GetSignal().GetTime()),

		obj.GetSignal().GetName(),

		obj.GetSignal().GetArgs(),

		obj.GetSignal().GetExecFilePath(),

		obj.GetSignal().GetPid(),

		obj.GetSignal().GetUid(),

		obj.GetSignal().GetGid(),

		obj.GetSignal().GetLineage(),

		obj.GetSignal().GetScraped(),

		obj.GetClusterId(),

		obj.GetNamespace(),

		pgutils.NilOrStringTimestamp(obj.GetContainerStartTime()),

		obj.GetImageId(),

		serialized,
	}

	finalStr := "INSERT INTO process_indicators (Id, DeploymentId, ContainerName, PodId, PodUid, Signal_Id, Signal_ContainerId, Signal_Time, Signal_Name, Signal_Args, Signal_ExecFilePath, Signal_Pid, Signal_Uid, Signal_Gid, Signal_Lineage, Signal_Scraped, ClusterId, Namespace, ContainerStartTime, ImageId, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, DeploymentId = EXCLUDED.DeploymentId, ContainerName = EXCLUDED.ContainerName, PodId = EXCLUDED.PodId, PodUid = EXCLUDED.PodUid, Signal_Id = EXCLUDED.Signal_Id, Signal_ContainerId = EXCLUDED.Signal_ContainerId, Signal_Time = EXCLUDED.Signal_Time, Signal_Name = EXCLUDED.Signal_Name, Signal_Args = EXCLUDED.Signal_Args, Signal_ExecFilePath = EXCLUDED.Signal_ExecFilePath, Signal_Pid = EXCLUDED.Signal_Pid, Signal_Uid = EXCLUDED.Signal_Uid, Signal_Gid = EXCLUDED.Signal_Gid, Signal_Lineage = EXCLUDED.Signal_Lineage, Signal_Scraped = EXCLUDED.Signal_Scraped, ClusterId = EXCLUDED.ClusterId, Namespace = EXCLUDED.Namespace, ContainerStartTime = EXCLUDED.ContainerStartTime, ImageId = EXCLUDED.ImageId, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(context.Background(), finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetSignal().GetLineageInfo() {
		if err := insertIntoProcessIndicatorsLineageInfo(tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from process_indicators_LineageInfo where process_indicators_Id = $1 AND idx >= $2"
	_, err = tx.Exec(context.Background(), query, obj.GetId(), len(obj.GetSignal().GetLineageInfo()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoProcessIndicatorsLineageInfo(tx pgx.Tx, obj *storage.ProcessSignal_LineageInfo, process_indicators_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start

		process_indicators_Id,

		idx,

		obj.GetParentUid(),

		obj.GetParentExecFilePath(),
	}

	finalStr := "INSERT INTO process_indicators_LineageInfo (process_indicators_Id, idx, ParentUid, ParentExecFilePath) VALUES($1, $2, $3, $4) ON CONFLICT(process_indicators_Id, idx) DO UPDATE SET process_indicators_Id = EXCLUDED.process_indicators_Id, idx = EXCLUDED.idx, ParentUid = EXCLUDED.ParentUid, ParentExecFilePath = EXCLUDED.ParentExecFilePath"
	_, err := tx.Exec(context.Background(), finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

// New returns a new Store instance using the provided sql instance.
func New(db *pgxpool.Pool) Store {
	globaldb.RegisterTable(table, "ProcessIndicator")

	createTableProcessIndicators(db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) upsert(objs ...*storage.ProcessIndicator) error {
	conn, release := s.acquireConn(ops.Get, "ProcessIndicator")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(context.Background())
		if err != nil {
			return err
		}

		if err := insertIntoProcessIndicators(tx, obj); err != nil {
			if err := tx.Rollback(context.Background()); err != nil {
				return err
			}
			return err
		}
		if err := tx.Commit(context.Background()); err != nil {
			return err
		}
	}
	return nil
}

func (s *storeImpl) Upsert(obj *storage.ProcessIndicator) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "ProcessIndicator")

	return s.upsert(obj)
}

func (s *storeImpl) UpsertMany(objs []*storage.ProcessIndicator) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "ProcessIndicator")

	return s.upsert(objs...)
}

// Count returns the number of objects in the store
func (s *storeImpl) Count() (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "ProcessIndicator")

	row := s.db.QueryRow(context.Background(), countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "ProcessIndicator")

	row := s.db.QueryRow(context.Background(), existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(id string) (*storage.ProcessIndicator, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "ProcessIndicator")

	conn, release := s.acquireConn(ops.Get, "ProcessIndicator")
	defer release()

	row := conn.QueryRow(context.Background(), getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.ProcessIndicator
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(op ops.Op, typ string) (*pgxpool.Conn, func()) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(context.Background())
	if err != nil {
		panic(err)
	}
	return conn, conn.Release
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(id string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "ProcessIndicator")

	conn, release := s.acquireConn(ops.Remove, "ProcessIndicator")
	defer release()

	if _, err := conn.Exec(context.Background(), deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs() ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.ProcessIndicatorIDs")

	rows, err := s.db.Query(context.Background(), getIDsStmt)
	if err != nil {
		return nil, pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// GetMany returns the objects specified by the IDs or the index in the missing indices slice
func (s *storeImpl) GetMany(ids []string) ([]*storage.ProcessIndicator, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "ProcessIndicator")

	conn, release := s.acquireConn(ops.GetMany, "ProcessIndicator")
	defer release()

	rows, err := conn.Query(context.Background(), getManyStmt, ids)
	if err != nil {
		if err == pgx.ErrNoRows {
			missingIndices := make([]int, 0, len(ids))
			for i := range ids {
				missingIndices = append(missingIndices, i)
			}
			return nil, missingIndices, nil
		}
		return nil, nil, err
	}
	defer rows.Close()
	elems := make([]*storage.ProcessIndicator, 0, len(ids))
	foundSet := make(map[string]struct{})
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		var msg storage.ProcessIndicator
		if err := proto.Unmarshal(data, &msg); err != nil {
			return nil, nil, err
		}
		foundSet[msg.GetId()] = struct{}{}
		elems = append(elems, &msg)
	}
	missingIndices := make([]int, 0, len(ids)-len(foundSet))
	for i, id := range ids {
		if _, ok := foundSet[id]; !ok {
			missingIndices = append(missingIndices, i)
		}
	}
	return elems, missingIndices, nil
}

// Delete removes the specified IDs from the store
func (s *storeImpl) DeleteMany(ids []string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "ProcessIndicator")

	conn, release := s.acquireConn(ops.RemoveMany, "ProcessIndicator")
	defer release()
	if _, err := conn.Exec(context.Background(), deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(fn func(obj *storage.ProcessIndicator) error) error {
	rows, err := s.db.Query(context.Background(), walkStmt)
	if err != nil {
		return pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		var msg storage.ProcessIndicator
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

//// Used for testing

func dropTableProcessIndicators(db *pgxpool.Pool) {
	_, _ = db.Exec(context.Background(), "DROP TABLE process_indicators CASCADE")
	dropTableProcessIndicatorsLineageInfo(db)

}

func dropTableProcessIndicatorsLineageInfo(db *pgxpool.Pool) {
	_, _ = db.Exec(context.Background(), "DROP TABLE process_indicators_LineageInfo CASCADE")

}

func Destroy(db *pgxpool.Pool) {
	dropTableProcessIndicators(db)
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex() ([]string, error) {
	return nil, nil
}
