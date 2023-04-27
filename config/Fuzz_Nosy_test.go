package config

import (
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_Config_Consensus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		c := NewConfig(path)
		c.Consensus()
	})
}

func Fuzz_Nosy_Config_Create__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, selectedExecution string, selectedConsensus string) {
		c := NewConfig(path)
		c.Create(selectedExecution, selectedConsensus)
	})
}

func Fuzz_Nosy_Config_Execution__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		c := NewConfig(path)
		c.Execution()
	})
}

func Fuzz_Nosy_Config_Exists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		c := NewConfig(path)
		c.Exists()
	})
}

func Fuzz_Nosy_Config_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		c := NewConfig(path)
		c.Read()
	})
}

func Fuzz_Nosy_Config_WriteConsensus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, selectedConsensus string) {
		c := NewConfig(path)
		c.WriteConsensus(selectedConsensus)
	})
}

func Fuzz_Nosy_Config_WriteExecution__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, selectedExecution string) {
		c := NewConfig(path)
		c.WriteExecution(selectedExecution)
	})
}

func Fuzz_Nosy_parsePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		parsePath(path)
	})
}
