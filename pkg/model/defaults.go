package model

import (
	"time"
)

const (
	DefaultFormat          = FormatCustom
	DefaultQuality         = QualityHigh
	DefaultPageSize        = 5
	DefaultUpdatePeriod    = 1 * time.Hour
	DefaultLogMaxSize      = 50 // megabytes
	DefaultLogMaxAge       = 30 // days
	DefaultLogMaxBackups   = 7
	PathRegex              = `^[A-Za-z0-9]+$`
	OPML                   = true
	DefaultYouTubeDLFormat = `bestaudio[ext=m4a]`
	DefaultExtension       = `m4a`
	DefaultMinDuration     = 300
	DefaultKeepLast        = 5
)
