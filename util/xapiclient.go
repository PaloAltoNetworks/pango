package util

import (
	"time"

	"github.com/PaloAltoNetworks/pango/plugin"
	"github.com/PaloAltoNetworks/pango/version"
)

// XapiClient is the interface that describes an pango.Client.
type XapiClient interface {
	String() string
	Versioning() version.Number
	Plugins() []plugin.Info
	LogAction(string, ...interface{})
	LogQuery(string, ...interface{})
	LogOp(string, ...interface{})
	LogUid(string, ...interface{})
	LogLog(string, ...interface{})
	Op(interface{}, string, interface{}, interface{}) ([]byte, error)
	Show(interface{}, interface{}, interface{}) ([]byte, error)
	Get(interface{}, interface{}, interface{}) ([]byte, error)
	Delete(interface{}, interface{}, interface{}) ([]byte, error)
	Set(interface{}, interface{}, interface{}, interface{}) ([]byte, error)
	Edit(interface{}, interface{}, interface{}, interface{}) ([]byte, error)
	Move(interface{}, string, string, interface{}, interface{}) ([]byte, error)
	Log(string, string, string, string, int, int, interface{}, interface{}) ([]byte, error)
	Export(string, interface{}, interface{}) (string, []byte, error)
	Uid(interface{}, string, interface{}, interface{}) ([]byte, error)
	EntryListUsing(Retriever, []string) ([]string, error)
	MemberListUsing(Retriever, []string) ([]string, error)
	RequestPasswordHash(string) (string, error)
	VsysImport(string, string, string, string, []string) error
	VsysUnimport(string, string, string, []string) error
	WaitForJob(uint, time.Duration, interface{}) error
	WaitForLogs(uint, time.Duration, time.Duration, interface{}) ([]byte, error)
	Commit(interface{}, string, interface{}) (uint, []byte, error)
	Clock() (time.Time, error)
	PositionFirstEntity(int, string, string, []string, []string) error
}
