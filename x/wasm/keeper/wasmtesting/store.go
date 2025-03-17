package wasmtesting

import (
	"io"

	storetypes "cosmossdk.io/store/types"
)

// MockCommitMultiStore mock with a CacheMultiStore to capture commits
type MockCommitMultiStore struct {
	storetypes.CommitMultiStore
	Committed []bool
}

func (m *MockCommitMultiStore) CacheMultiStore() storetypes.CacheMultiStore {
	m.Committed = append(m.Committed, false)
	return &mockCMS{m, &m.Committed[len(m.Committed)-1]}
}

type mockCMS struct {
	storetypes.CommitMultiStore
	committed *bool
}

// CacheMultiStore implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).CacheMultiStore of mockCMS.CommitMultiStore.
func (m *mockCMS) CacheMultiStore() storetypes.CacheMultiStore {
	panic("unimplemented")
}

// CacheMultiStoreWithVersion implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).CacheMultiStoreWithVersion of mockCMS.CommitMultiStore.
func (m *mockCMS) CacheMultiStoreWithVersion(version int64) (storetypes.CacheMultiStore, error) {
	panic("unimplemented")
}

// CacheWrap implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).CacheWrap of mockCMS.CommitMultiStore.
func (m *mockCMS) CacheWrap() storetypes.CacheWrap {
	panic("unimplemented")
}

// CacheWrapWithTrace implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).CacheWrapWithTrace of mockCMS.CommitMultiStore.
func (m *mockCMS) CacheWrapWithTrace(w io.Writer, tc storetypes.TraceContext) storetypes.CacheWrap {
	panic("unimplemented")
}

// Copy implements types.CacheMultiStore.
func (m *mockCMS) Copy() storetypes.CacheMultiStore {
	panic("unimplemented")
}

// GetKVStore implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).GetKVStore of mockCMS.CommitMultiStore.
func (m *mockCMS) GetKVStore(storetypes.StoreKey) storetypes.KVStore {
	panic("unimplemented")
}

// GetStore implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).GetStore of mockCMS.CommitMultiStore.
func (m *mockCMS) GetStore(storetypes.StoreKey) storetypes.Store {
	panic("unimplemented")
}

// GetStoreType implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).GetStoreType of mockCMS.CommitMultiStore.
func (m *mockCMS) GetStoreType() storetypes.StoreType {
	panic("unimplemented")
}

// LatestVersion implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).LatestVersion of mockCMS.CommitMultiStore.
func (m *mockCMS) LatestVersion() int64 {
	panic("unimplemented")
}

// SetTracer implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).SetTracer of mockCMS.CommitMultiStore.
func (m *mockCMS) SetTracer(w io.Writer) storetypes.MultiStore {
	panic("unimplemented")
}

// SetTracingContext implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).SetTracingContext of mockCMS.CommitMultiStore.
func (m *mockCMS) SetTracingContext(storetypes.TraceContext) storetypes.MultiStore {
	panic("unimplemented")
}

// TracingEnabled implements types.CacheMultiStore.
// Subtle: this method shadows the method (CommitMultiStore).TracingEnabled of mockCMS.CommitMultiStore.
func (m *mockCMS) TracingEnabled() bool {
	panic("unimplemented")
}

func (m *mockCMS) Write() {
	*m.committed = true
}
