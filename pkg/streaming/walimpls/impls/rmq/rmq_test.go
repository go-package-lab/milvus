package rmq

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/milvus-io/milvus/pkg/v2/mq/mqimpl/rocksmq/server"
	"github.com/milvus-io/milvus/pkg/v2/streaming/util/message"
	"github.com/milvus-io/milvus/pkg/v2/streaming/walimpls"
	"github.com/milvus-io/milvus/pkg/v2/streaming/walimpls/registry"
	"github.com/milvus-io/milvus/pkg/v2/util/paramtable"
)

func TestMain(m *testing.M) {
	paramtable.Init()
	tmpPath, err := os.MkdirTemp("", "rocksdb_test")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpPath)
	server.InitRocksMQ(tmpPath)
	defer server.CloseRocksMQ()
	m.Run()
}

func TestRegistry(t *testing.T) {
	registeredB := registry.MustGetBuilder(WALName)
	assert.NotNil(t, registeredB)
	assert.Equal(t, WALName, registeredB.Name())

	id, err := message.UnmarshalMessageID(WALName, rmqID(1).Marshal())
	assert.NoError(t, err)
	assert.True(t, id.EQ(rmqID(1)))
}

func TestWAL(t *testing.T) {
	walimpls.NewWALImplsTestFramework(t, 1000, &builderImpl{}).Run()
}
