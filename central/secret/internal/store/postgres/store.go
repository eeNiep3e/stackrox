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
	countStmt  = "SELECT COUNT(*) FROM secrets"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM secrets WHERE Id = $1)"

	getStmt     = "SELECT serialized FROM secrets WHERE Id = $1"
	deleteStmt  = "DELETE FROM secrets WHERE Id = $1"
	walkStmt    = "SELECT serialized FROM secrets"
	getIDsStmt  = "SELECT Id FROM secrets"
	getManyStmt = "SELECT serialized FROM secrets WHERE Id = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM secrets WHERE Id = ANY($1::text[])"
)

var (
	log = logging.LoggerForModule()

	table = "secrets"
)

type Store interface {
	Count() (int, error)
	Exists(id string) (bool, error)
	Get(id string) (*storage.Secret, bool, error)
	Upsert(obj *storage.Secret) error
	UpsertMany(objs []*storage.Secret) error
	Delete(id string) error
	GetIDs() ([]string, error)
	GetMany(ids []string) ([]*storage.Secret, []int, error)
	DeleteMany(ids []string) error

	Walk(fn func(obj *storage.Secret) error) error

	AckKeysIndexed(keys ...string) error
	GetKeysToIndex() ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableSecrets(db *pgxpool.Pool) {
	table := `
create table if not exists secrets (
    Id varchar,
    Name varchar,
    ClusterId varchar,
    ClusterName varchar,
    Namespace varchar,
    Type varchar,
    Labels jsonb,
    Annotations jsonb,
    CreatedAt timestamp,
    Relationship_Id varchar,
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

	createTableSecretsFiles(db)
	createTableSecretsContainerRelationships(db)
	createTableSecretsDeploymentRelationships(db)
}

func createTableSecretsFiles(db *pgxpool.Pool) {
	table := `
create table if not exists secrets_Files (
    secrets_Id varchar,
    idx numeric,
    Name varchar,
    Type integer,
    Cert_Subject_CommonName varchar,
    Cert_Subject_Country varchar,
    Cert_Subject_Organization varchar,
    Cert_Subject_OrganizationUnit varchar,
    Cert_Subject_Locality varchar,
    Cert_Subject_Province varchar,
    Cert_Subject_StreetAddress varchar,
    Cert_Subject_PostalCode varchar,
    Cert_Subject_Names text[],
    Cert_Issuer_CommonName varchar,
    Cert_Issuer_Country varchar,
    Cert_Issuer_Organization varchar,
    Cert_Issuer_OrganizationUnit varchar,
    Cert_Issuer_Locality varchar,
    Cert_Issuer_Province varchar,
    Cert_Issuer_StreetAddress varchar,
    Cert_Issuer_PostalCode varchar,
    Cert_Issuer_Names text[],
    Cert_Sans text[],
    Cert_StartDate timestamp,
    Cert_EndDate timestamp,
    Cert_Algorithm varchar,
    PRIMARY KEY(secrets_Id, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (secrets_Id) REFERENCES secrets(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(context.Background(), table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{

		"create index if not exists secretsFiles_idx on secrets_Files using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(context.Background(), index); err != nil {
			panic(err)
		}
	}

	createTableSecretsFilesRegistries(db)
}

func createTableSecretsFilesRegistries(db *pgxpool.Pool) {
	table := `
create table if not exists secrets_Files_Registries (
    secrets_Id varchar,
    secrets_Files_idx numeric,
    idx numeric,
    Name varchar,
    Username varchar,
    PRIMARY KEY(secrets_Id, secrets_Files_idx, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (secrets_Id, secrets_Files_idx) REFERENCES secrets_Files(secrets_Id, idx) ON DELETE CASCADE
)
`

	_, err := db.Exec(context.Background(), table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{

		"create index if not exists secretsFilesRegistries_idx on secrets_Files_Registries using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(context.Background(), index); err != nil {
			panic(err)
		}
	}

}

func createTableSecretsContainerRelationships(db *pgxpool.Pool) {
	table := `
create table if not exists secrets_ContainerRelationships (
    secrets_Id varchar,
    idx numeric,
    Id varchar,
    Path varchar,
    PRIMARY KEY(secrets_Id, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (secrets_Id) REFERENCES secrets(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(context.Background(), table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{

		"create index if not exists secretsContainerRelationships_idx on secrets_ContainerRelationships using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(context.Background(), index); err != nil {
			panic(err)
		}
	}

}

func createTableSecretsDeploymentRelationships(db *pgxpool.Pool) {
	table := `
create table if not exists secrets_DeploymentRelationships (
    secrets_Id varchar,
    idx numeric,
    Id varchar,
    Name varchar,
    PRIMARY KEY(secrets_Id, idx),
    CONSTRAINT fk_parent_table FOREIGN KEY (secrets_Id) REFERENCES secrets(Id) ON DELETE CASCADE
)
`

	_, err := db.Exec(context.Background(), table)
	if err != nil {
		panic("error creating table: " + table)
	}

	indexes := []string{

		"create index if not exists secretsDeploymentRelationships_idx on secrets_DeploymentRelationships using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(context.Background(), index); err != nil {
			panic(err)
		}
	}

}

func insertIntoSecrets(tx pgx.Tx, obj *storage.Secret) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start

		obj.GetId(),

		obj.GetName(),

		obj.GetClusterId(),

		obj.GetClusterName(),

		obj.GetNamespace(),

		obj.GetType(),

		obj.GetLabels(),

		obj.GetAnnotations(),

		pgutils.NilOrStringTimestamp(obj.GetCreatedAt()),

		obj.GetRelationship().GetId(),

		serialized,
	}

	finalStr := "INSERT INTO secrets (Id, Name, ClusterId, ClusterName, Namespace, Type, Labels, Annotations, CreatedAt, Relationship_Id, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, ClusterId = EXCLUDED.ClusterId, ClusterName = EXCLUDED.ClusterName, Namespace = EXCLUDED.Namespace, Type = EXCLUDED.Type, Labels = EXCLUDED.Labels, Annotations = EXCLUDED.Annotations, CreatedAt = EXCLUDED.CreatedAt, Relationship_Id = EXCLUDED.Relationship_Id, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(context.Background(), finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetFiles() {
		if err := insertIntoSecretsFiles(tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from secrets_Files where secrets_Id = $1 AND idx >= $2"
	_, err = tx.Exec(context.Background(), query, obj.GetId(), len(obj.GetFiles()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetRelationship().GetContainerRelationships() {
		if err := insertIntoSecretsContainerRelationships(tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from secrets_ContainerRelationships where secrets_Id = $1 AND idx >= $2"
	_, err = tx.Exec(context.Background(), query, obj.GetId(), len(obj.GetRelationship().GetContainerRelationships()))
	if err != nil {
		return err
	}
	for childIdx, child := range obj.GetRelationship().GetDeploymentRelationships() {
		if err := insertIntoSecretsDeploymentRelationships(tx, child, obj.GetId(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from secrets_DeploymentRelationships where secrets_Id = $1 AND idx >= $2"
	_, err = tx.Exec(context.Background(), query, obj.GetId(), len(obj.GetRelationship().GetDeploymentRelationships()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoSecretsFiles(tx pgx.Tx, obj *storage.SecretDataFile, secrets_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start

		secrets_Id,

		idx,

		obj.GetName(),

		obj.GetType(),

		obj.GetCert().GetSubject().GetCommonName(),

		obj.GetCert().GetSubject().GetCountry(),

		obj.GetCert().GetSubject().GetOrganization(),

		obj.GetCert().GetSubject().GetOrganizationUnit(),

		obj.GetCert().GetSubject().GetLocality(),

		obj.GetCert().GetSubject().GetProvince(),

		obj.GetCert().GetSubject().GetStreetAddress(),

		obj.GetCert().GetSubject().GetPostalCode(),

		obj.GetCert().GetSubject().GetNames(),

		obj.GetCert().GetIssuer().GetCommonName(),

		obj.GetCert().GetIssuer().GetCountry(),

		obj.GetCert().GetIssuer().GetOrganization(),

		obj.GetCert().GetIssuer().GetOrganizationUnit(),

		obj.GetCert().GetIssuer().GetLocality(),

		obj.GetCert().GetIssuer().GetProvince(),

		obj.GetCert().GetIssuer().GetStreetAddress(),

		obj.GetCert().GetIssuer().GetPostalCode(),

		obj.GetCert().GetIssuer().GetNames(),

		obj.GetCert().GetSans(),

		pgutils.NilOrStringTimestamp(obj.GetCert().GetStartDate()),

		pgutils.NilOrStringTimestamp(obj.GetCert().GetEndDate()),

		obj.GetCert().GetAlgorithm(),
	}

	finalStr := "INSERT INTO secrets_Files (secrets_Id, idx, Name, Type, Cert_Subject_CommonName, Cert_Subject_Country, Cert_Subject_Organization, Cert_Subject_OrganizationUnit, Cert_Subject_Locality, Cert_Subject_Province, Cert_Subject_StreetAddress, Cert_Subject_PostalCode, Cert_Subject_Names, Cert_Issuer_CommonName, Cert_Issuer_Country, Cert_Issuer_Organization, Cert_Issuer_OrganizationUnit, Cert_Issuer_Locality, Cert_Issuer_Province, Cert_Issuer_StreetAddress, Cert_Issuer_PostalCode, Cert_Issuer_Names, Cert_Sans, Cert_StartDate, Cert_EndDate, Cert_Algorithm) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26) ON CONFLICT(secrets_Id, idx) DO UPDATE SET secrets_Id = EXCLUDED.secrets_Id, idx = EXCLUDED.idx, Name = EXCLUDED.Name, Type = EXCLUDED.Type, Cert_Subject_CommonName = EXCLUDED.Cert_Subject_CommonName, Cert_Subject_Country = EXCLUDED.Cert_Subject_Country, Cert_Subject_Organization = EXCLUDED.Cert_Subject_Organization, Cert_Subject_OrganizationUnit = EXCLUDED.Cert_Subject_OrganizationUnit, Cert_Subject_Locality = EXCLUDED.Cert_Subject_Locality, Cert_Subject_Province = EXCLUDED.Cert_Subject_Province, Cert_Subject_StreetAddress = EXCLUDED.Cert_Subject_StreetAddress, Cert_Subject_PostalCode = EXCLUDED.Cert_Subject_PostalCode, Cert_Subject_Names = EXCLUDED.Cert_Subject_Names, Cert_Issuer_CommonName = EXCLUDED.Cert_Issuer_CommonName, Cert_Issuer_Country = EXCLUDED.Cert_Issuer_Country, Cert_Issuer_Organization = EXCLUDED.Cert_Issuer_Organization, Cert_Issuer_OrganizationUnit = EXCLUDED.Cert_Issuer_OrganizationUnit, Cert_Issuer_Locality = EXCLUDED.Cert_Issuer_Locality, Cert_Issuer_Province = EXCLUDED.Cert_Issuer_Province, Cert_Issuer_StreetAddress = EXCLUDED.Cert_Issuer_StreetAddress, Cert_Issuer_PostalCode = EXCLUDED.Cert_Issuer_PostalCode, Cert_Issuer_Names = EXCLUDED.Cert_Issuer_Names, Cert_Sans = EXCLUDED.Cert_Sans, Cert_StartDate = EXCLUDED.Cert_StartDate, Cert_EndDate = EXCLUDED.Cert_EndDate, Cert_Algorithm = EXCLUDED.Cert_Algorithm"
	_, err := tx.Exec(context.Background(), finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetImagePullSecret().GetRegistries() {
		if err := insertIntoSecretsFilesRegistries(tx, child, secrets_Id, idx, childIdx); err != nil {
			return err
		}
	}

	query = "delete from secrets_Files_Registries where secrets_Id = $1 AND secrets_Files_idx = $2 AND idx >= $3"
	_, err = tx.Exec(context.Background(), query, secrets_Id, idx, len(obj.GetImagePullSecret().GetRegistries()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoSecretsFilesRegistries(tx pgx.Tx, obj *storage.ImagePullSecret_Registry, secrets_Id string, secrets_Files_idx int, idx int) error {

	values := []interface{}{
		// parent primary keys start

		secrets_Id,

		secrets_Files_idx,

		idx,

		obj.GetName(),

		obj.GetUsername(),
	}

	finalStr := "INSERT INTO secrets_Files_Registries (secrets_Id, secrets_Files_idx, idx, Name, Username) VALUES($1, $2, $3, $4, $5) ON CONFLICT(secrets_Id, secrets_Files_idx, idx) DO UPDATE SET secrets_Id = EXCLUDED.secrets_Id, secrets_Files_idx = EXCLUDED.secrets_Files_idx, idx = EXCLUDED.idx, Name = EXCLUDED.Name, Username = EXCLUDED.Username"
	_, err := tx.Exec(context.Background(), finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoSecretsContainerRelationships(tx pgx.Tx, obj *storage.SecretContainerRelationship, secrets_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start

		secrets_Id,

		idx,

		obj.GetId(),

		obj.GetPath(),
	}

	finalStr := "INSERT INTO secrets_ContainerRelationships (secrets_Id, idx, Id, Path) VALUES($1, $2, $3, $4) ON CONFLICT(secrets_Id, idx) DO UPDATE SET secrets_Id = EXCLUDED.secrets_Id, idx = EXCLUDED.idx, Id = EXCLUDED.Id, Path = EXCLUDED.Path"
	_, err := tx.Exec(context.Background(), finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func insertIntoSecretsDeploymentRelationships(tx pgx.Tx, obj *storage.SecretDeploymentRelationship, secrets_Id string, idx int) error {

	values := []interface{}{
		// parent primary keys start

		secrets_Id,

		idx,

		obj.GetId(),

		obj.GetName(),
	}

	finalStr := "INSERT INTO secrets_DeploymentRelationships (secrets_Id, idx, Id, Name) VALUES($1, $2, $3, $4) ON CONFLICT(secrets_Id, idx) DO UPDATE SET secrets_Id = EXCLUDED.secrets_Id, idx = EXCLUDED.idx, Id = EXCLUDED.Id, Name = EXCLUDED.Name"
	_, err := tx.Exec(context.Background(), finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

// New returns a new Store instance using the provided sql instance.
func New(db *pgxpool.Pool) Store {
	globaldb.RegisterTable(table, "Secret")

	createTableSecrets(db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) upsert(objs ...*storage.Secret) error {
	conn, release := s.acquireConn(ops.Get, "Secret")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(context.Background())
		if err != nil {
			return err
		}

		if err := insertIntoSecrets(tx, obj); err != nil {
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

func (s *storeImpl) Upsert(obj *storage.Secret) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "Secret")

	return s.upsert(obj)
}

func (s *storeImpl) UpsertMany(objs []*storage.Secret) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "Secret")

	return s.upsert(objs...)
}

// Count returns the number of objects in the store
func (s *storeImpl) Count() (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "Secret")

	row := s.db.QueryRow(context.Background(), countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "Secret")

	row := s.db.QueryRow(context.Background(), existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(id string) (*storage.Secret, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "Secret")

	conn, release := s.acquireConn(ops.Get, "Secret")
	defer release()

	row := conn.QueryRow(context.Background(), getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.Secret
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
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "Secret")

	conn, release := s.acquireConn(ops.Remove, "Secret")
	defer release()

	if _, err := conn.Exec(context.Background(), deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs() ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.SecretIDs")

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
func (s *storeImpl) GetMany(ids []string) ([]*storage.Secret, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "Secret")

	conn, release := s.acquireConn(ops.GetMany, "Secret")
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
	elems := make([]*storage.Secret, 0, len(ids))
	foundSet := make(map[string]struct{})
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		var msg storage.Secret
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
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "Secret")

	conn, release := s.acquireConn(ops.RemoveMany, "Secret")
	defer release()
	if _, err := conn.Exec(context.Background(), deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(fn func(obj *storage.Secret) error) error {
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
		var msg storage.Secret
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

func dropTableSecrets(db *pgxpool.Pool) {
	_, _ = db.Exec(context.Background(), "DROP TABLE secrets CASCADE")
	dropTableSecretsFiles(db)
	dropTableSecretsContainerRelationships(db)
	dropTableSecretsDeploymentRelationships(db)

}

func dropTableSecretsFiles(db *pgxpool.Pool) {
	_, _ = db.Exec(context.Background(), "DROP TABLE secrets_Files CASCADE")
	dropTableSecretsFilesRegistries(db)

}

func dropTableSecretsFilesRegistries(db *pgxpool.Pool) {
	_, _ = db.Exec(context.Background(), "DROP TABLE secrets_Files_Registries CASCADE")

}

func dropTableSecretsContainerRelationships(db *pgxpool.Pool) {
	_, _ = db.Exec(context.Background(), "DROP TABLE secrets_ContainerRelationships CASCADE")

}

func dropTableSecretsDeploymentRelationships(db *pgxpool.Pool) {
	_, _ = db.Exec(context.Background(), "DROP TABLE secrets_DeploymentRelationships CASCADE")

}

func Destroy(db *pgxpool.Pool) {
	dropTableSecrets(db)
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
