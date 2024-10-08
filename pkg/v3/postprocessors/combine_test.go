package postprocessors

import (
	"context"
	"fmt"
	"testing"

	ocr2keepers "github.com/goplugin/plugin-common/pkg/types/automation"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCombinedPostprocessor(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		one := new(MockPostProcessor)
		two := new(MockPostProcessor)
		tre := new(MockPostProcessor)

		cmb := NewCombinedPostprocessor(one, two, tre)
		rst := []ocr2keepers.CheckResult{{UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{1}), WorkID: "0x1", Retryable: true}}
		p := []ocr2keepers.UpkeepPayload{{UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{1}), WorkID: "0x1"}}

		one.On("PostProcess", ctx, rst, p).Return(nil)
		two.On("PostProcess", ctx, rst, p).Return(nil)
		tre.On("PostProcess", ctx, rst, p).Return(nil)

		assert.NoError(t, cmb.PostProcess(ctx, rst, p), "no error expected from combined post processing")
	})

	t.Run("with errors", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		one := new(MockPostProcessor)
		two := new(MockPostProcessor)
		tre := new(MockPostProcessor)

		cmb := NewCombinedPostprocessor(one, two, tre)
		rst := []ocr2keepers.CheckResult{{UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{1}), WorkID: "0x1", Retryable: true}}
		p := []ocr2keepers.UpkeepPayload{{UpkeepID: ocr2keepers.UpkeepIdentifier([32]byte{1}), WorkID: "0x1"}}

		one.On("PostProcess", ctx, rst, p).Return(nil)
		two.On("PostProcess", ctx, rst, p).Return(fmt.Errorf("error"))
		tre.On("PostProcess", ctx, rst, p).Return(fmt.Errorf("error"))

		assert.Error(t, cmb.PostProcess(ctx, rst, p), "error expected from combined post processing")
	})
}

type MockPostProcessor struct {
	mock.Mock
}

func (_m *MockPostProcessor) PostProcess(ctx context.Context, r []ocr2keepers.CheckResult, p []ocr2keepers.UpkeepPayload) error {
	ret := _m.Called(ctx, r, p)
	return ret.Error(0)
}
