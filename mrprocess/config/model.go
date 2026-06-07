package config

import (
	"time"
)

type (
	// MessageCollector - настройки многопоточного сервиса сбора и пакетной обработки сообщений (PUSH-модель).
	MessageCollector struct {
		ReadyTimeout   time.Duration `yaml:"ready_timeout"`
		FlushPeriod    time.Duration `yaml:"flush_period"`
		HandlerTimeout time.Duration `yaml:"handler_timeout"`
		BatchSize      uint32        `yaml:"batch_size"`
		WorkersCount   uint8         `yaml:"workers_count"`
	}

	// MessageProcessor - настройки многопоточного сервиса обработки сообщений (PULL-модель).
	MessageProcessor struct {
		ReadyTimeout         time.Duration `yaml:"ready_timeout"`
		ReadPeriod           time.Duration `yaml:"read_period"`
		ConsumerReadTimeout  time.Duration `yaml:"consumer_read_timeout"`
		ConsumerWriteTimeout time.Duration `yaml:"consumer_write_timeout"`
		HandlerTimeout       time.Duration `yaml:"handler_timeout"`
		QueueSize            uint16        `yaml:"queue_size"`
		WorkersCount         uint8         `yaml:"workers_count"`
		NotificationChannel  string        `yaml:"notification_channel,omitempty"`
	}

	// SchedulerTask - настройки многопоточного сервиса запуска задач по расписанию.
	SchedulerTask struct {
		Period              time.Duration `yaml:"period"`
		Timeout             time.Duration `yaml:"timeout"`
		NotificationChannel string        `yaml:"notification_channel,omitempty"`
	}
)
