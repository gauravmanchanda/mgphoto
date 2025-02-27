package main

import (
	"time"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

type ProgressBar struct {
	progress *mpb.Progress
	bar      *mpb.Bar
	start    time.Time
}

// NewProgressBar - get new progress bar
func NewProgressBar(total int) *ProgressBar {

	progressBar := new(ProgressBar)

	progressBar.progress = mpb.New(mpb.WithWidth(64), mpb.WithRefreshRate(180*time.Millisecond))

	progressBar.bar = progressBar.progress.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(decor.Elapsed(decor.ET_STYLE_MMSS)),
	)

	progressBar.start = time.Now()

	return progressBar
}

func (p *ProgressBar) increment() {
	p.bar.IncrBy(1, time.Since(p.start))
}

func (p *ProgressBar) wait() {
	p.progress.Wait()
}
