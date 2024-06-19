package filejanitor

import (
	"sort"
	"strings"
	"time"

	"github.com/toxyl/flo"
	"github.com/toxyl/scheduler"
)

type FileJanitor struct {
	config  *Config
	stopFns []func()
}

func (fj *FileJanitor) stop() {
	for _, fn := range fj.stopFns {
		fn()
	}
}

func (fj *FileJanitor) remove(retentionPeriod time.Duration, files []*flo.FileObj) []error {
	errors := []error{}
	threshold := time.Now().Add(-1 * retentionPeriod)
	for _, f := range files {
		if f.OlderThan(threshold) {
			if err := f.Remove(); err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

func (fj *FileJanitor) start(errorHandler func(errors []error)) (stop func()) {
	for _, p := range fj.config.Policies {
		fj.stopFns = append(fj.stopFns,
			scheduler.Run(
				p.ScanEvery, 0,
				func() (stop bool) {
					files := []*flo.FileObj{}
					flo.Dir(p.Path).Each(
						func(f *flo.FileObj) {
							if p.Extension == "" || (p.Extension != "" && strings.HasSuffix(strings.ToLower(f.Path()), "."+strings.ToLower(p.Extension))) {
								files = append(files, f)
							}
						},
						nil,
					)
					l := len(files)
					if l == 0 || l < int(p.KeepLast) {
						return false // nothing to remove
					}
					sort.Slice(files, func(i, j int) bool {
						return files[j].OlderThan(files[i].LastModified())
					})
					errors := fj.remove(p.RetentionPeriod, files[int(p.KeepLast):])
					if len(errors) > 0 && errorHandler != nil {
						errorHandler(errors)
					}
					return false
				},
				nil,
			),
		)
	}
	return fj.stop
}

func Run(config *Config, errorHandler func(errors []error)) (stop func()) {
	fj := &FileJanitor{
		config: config,
	}
	return fj.start(errorHandler)
}
