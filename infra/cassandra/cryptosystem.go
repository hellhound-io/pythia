package cassandra
/*
import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"gitlab.com/consensys-hellhound/pythia/log"

	"gitlab.com/consensys-hellhound/pythia/domain/cryptosystem"
)

const (
	keyspace          = "hellhound"
	cryptoSystemTable = "pythia_cryptosystem"
)

type crypoStystemRepository struct {
	cassandraHost string
	session       *gocql.Session
}

func (r crypoStystemRepository) Store(cryptosystem cryptosystem.CryptoSystem) (c cryptosystem.CryptoSystem, err error) {
	stmt, names := qb.Insert(cryptoSystemTable).
		Columns("id", "name", "type", "sub_type", "stage").
		ToCql()
	err = gocqlx.Query(r.session.Query(stmt), names).BindStruct(&cryptosystem).ExecRelease()
	c = cryptosystem
	return
}

func (r crypoStystemRepository) Find(id string) (c *cryptosystem.CryptoSystem, err error) {
	stmt, names := qb.Select(cryptoSystemTable).
		Where(qb.Eq("id")).
		ToCql()
	var foundCryptoSystem cryptosystem.CryptoSystem
	q := gocqlx.Query(r.session.Query(stmt), names).BindMap(qb.M{
		"id": id,
	})
	if err = q.GetRelease(&foundCryptoSystem); err != nil {
		return
	}
	c = &foundCryptoSystem
	return
}

func (r crypoStystemRepository) FindByType(cryptoSystemType cryptosystem.Type) (c []cryptosystem.CryptoSystem, err error) {
	stmt, names := qb.Select(cryptoSystemTable).
		Where(qb.Eq("type")).
		AllowFiltering().
		ToCql()
	q := gocqlx.Query(r.session.Query(stmt), names).BindMap(qb.M{
		"type": string(cryptoSystemType),
	})
	if err = q.SelectRelease(&c); err != nil {
		return
	}
	return
}

func (r crypoStystemRepository) FindAll() (c []cryptosystem.CryptoSystem, err error) {
	stmt, names := qb.Select(cryptoSystemTable).
		ToCql()
	q := gocqlx.Query(r.session.Query(stmt), names)
	if err = q.SelectRelease(&c); err != nil {
		return
	}
	return
}

func NewCryptoSystemRepository(cassandraHost string) cryptosystem.Repository {
	c := &crypoStystemRepository{
		cassandraHost: cassandraHost,
	}
	go c.connect()
	return c
}

func (r *crypoStystemRepository) connect() {
	log.Logger.Debugf("initializing cassandra connection")
	cluster := gocql.NewCluster(r.cassandraHost)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Logger.Errorf("cannot start cassandra session : %s", err.Error())
	} else {
		r.session = session
		log.Logger.Debugln("connected to cassandra cluster")
	}
}
*/