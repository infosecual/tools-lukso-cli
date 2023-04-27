package main

import (
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/urfave/cli/v2"
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

func Fuzz_Nosy_ClientDependency_Download__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		var tag string
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var commitHash string
		fill_err = tp.Fill(&commitHash)
		if fill_err != nil {
			return
		}
		var isUpdate bool
		fill_err = tp.Fill(&isUpdate)
		if fill_err != nil {
			return
		}
		var permissions int
		fill_err = tp.Fill(&permissions)
		if fill_err != nil {
			return
		}
		if dependency == nil {
			return
		}

		dependency.Download(tag, commitHash, isUpdate, permissions)
	})
}

func Fuzz_Nosy_ClientDependency_Log__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		var logFilePath string
		fill_err = tp.Fill(&logFilePath)
		if fill_err != nil {
			return
		}
		if dependency == nil {
			return
		}

		dependency.Log(logFilePath)
	})
}

func Fuzz_Nosy_ClientDependency_ParseUrl__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		var tag string
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var commitHash string
		fill_err = tp.Fill(&commitHash)
		if fill_err != nil {
			return
		}
		if dependency == nil {
			return
		}

		dependency.ParseUrl(tag, commitHash)
	})
}

func Fuzz_Nosy_ClientDependency_PassStartFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if dependency == nil || ctx == nil {
			return
		}

		dependency.PassStartFlags(ctx)
	})
}

func Fuzz_Nosy_ClientDependency_ResolveBinaryPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		var tagName string
		fill_err = tp.Fill(&tagName)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		if dependency == nil {
			return
		}

		dependency.ResolveBinaryPath(tagName, datadir)
	})
}

func Fuzz_Nosy_ClientDependency_ResolveDirPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		var tagName string
		fill_err = tp.Fill(&tagName)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		if dependency == nil {
			return
		}

		dependency.ResolveDirPath(tagName, datadir)
	})
}

func Fuzz_Nosy_ClientDependency_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		var arguments []string
		fill_err = tp.Fill(&arguments)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if dependency == nil || ctx == nil {
			return
		}

		dependency.Start(arguments, ctx)
	})
}

func Fuzz_Nosy_ClientDependency_Stat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		if dependency == nil {
			return
		}

		dependency.Stat()
	})
}

func Fuzz_Nosy_ClientDependency_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		if dependency == nil {
			return
		}

		dependency.Stop()
	})
}

func Fuzz_Nosy_ClientDependency_createDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependency *ClientDependency
		fill_err = tp.Fill(&dependency)
		if fill_err != nil {
			return
		}
		if dependency == nil {
			return
		}

		dependency.createDir()
	})
}

func Fuzz_Nosy_boolToInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b bool) {
		boolToInt(b)
	})
}

func Fuzz_Nosy_fetchTag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, githubLocation string) {
		fetchTag(githubLocation)
	})
}

func Fuzz_Nosy_fetchTagAndCommitHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, githubLocation string) {
		fetchTagAndCommitHash(githubLocation)
	})
}

func Fuzz_Nosy_fileExists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		fileExists(path)
	})
}

func Fuzz_Nosy_flagFileExists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var flag string
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		flagFileExists(ctx, flag)
	})
}

func Fuzz_Nosy_getLastFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dir string, dependency string) {
		getLastFile(dir, dependency)
	})
}

func Fuzz_Nosy_isRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dependency string) {
		isRunning(dependency)
	})
}

func Fuzz_Nosy_logClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dependencyName string) {
		logClient(dependencyName)
	})
}

func Fuzz_Nosy_prepareGethStartFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		prepareGethStartFlags(ctx)
	})
}

func Fuzz_Nosy_prepareLogfileFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, logDir string, dependencyName string) {
		prepareLogfileFlag(logDir, dependencyName)
	})
}

func Fuzz_Nosy_preparePrysmStartFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		preparePrysmStartFlags(ctx)
	})
}

func Fuzz_Nosy_prepareTimestampedFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, logDir string, logFileName string) {
		prepareTimestampedFile(logDir, logFileName)
	})
}

func Fuzz_Nosy_prepareValidatorStartFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		prepareValidatorStartFlags(ctx)
	})
}

func Fuzz_Nosy_registerInputWithMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, message string) {
		registerInputWithMessage(message)
	})
}

func Fuzz_Nosy_removePrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, arg string, name string) {
		removePrefix(arg, name)
	})
}

// skipping Fuzz_Nosy_selectNetworkFor__ because parameters include func, chan, or unsupported interface: func(*github.com/urfave/cli/v2.Context) error

func Fuzz_Nosy_statClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dependencyName string, layer string) {
		statClient(dependencyName, layer)
	})
}

func Fuzz_Nosy_truncateFileFromDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		truncateFileFromDir(filePath)
	})
}
