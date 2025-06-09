package movement_test

import (
	"log/slog"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMovement(t *testing.T) {
	handler := slog.NewTextHandler(GinkgoWriter, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	slog.SetDefault(slog.New(handler))
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movement Suite")
}
