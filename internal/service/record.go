package service

import (
	"github.com/matisiekpl/electrocardiogram-server/internal/dto"
	"github.com/matisiekpl/electrocardiogram-server/internal/model"
	"github.com/matisiekpl/electrocardiogram-server/internal/repository"
	"github.com/sirupsen/logrus"
	"go.bug.st/serial"
	"strings"
	"time"
)

type RecordService interface {
	ListRecords(filter dto.Filter) ([]model.Record, error)
	Analyze(filter dto.Filter) (dto.Analysis, error)
	Channel() chan model.Record
	Connect()
	Clean() error
}

type recordService struct {
	recordRepository repository.RecordRepository
	channel          chan model.Record
}

func newRecordService(recordRepository repository.RecordRepository) RecordService {
	return &recordService{
		recordRepository: recordRepository,
		channel:          make(chan model.Record, 5),
	}
}

func (r recordService) ListRecords(filter dto.Filter) ([]model.Record, error) {
	return r.recordRepository.Filter(filter.Start, filter.End)
}

func (r recordService) Analyze(filter dto.Filter) (dto.Analysis, error) {
	//TODO implement me
	panic("implement me")
}

func (r recordService) Channel() chan model.Record {
	return r.channel
}

func (r recordService) Save(value int64) error {
	logrus.Infof("saving value: %d", value)
	record := model.Record{
		Value:     value,
		Timestamp: time.Now().UnixMilli(),
	}
	err := r.recordRepository.Insert(&record)
	if err != nil {
		logrus.Errorf("failed to save record with value %d", value)
		return err
	}
	for len(r.channel) > 0 {
		<-r.channel
	}
	r.channel <- record
	return nil
}

func (r recordService) Connect() {
	ports, err := serial.GetPortsList()
	if err != nil {
		logrus.Error("could not list serial ports: %w", err)
		return
	}
	var usbPorts []string
	for _, p := range ports {
		logrus.Infof("found serial port: %s", p)
		if strings.Contains(p, "USB") {
			usbPorts = append(usbPorts, p)
		}
	}

	delta := 1
	current := 100

	for {
		if current > 200 {
			delta = -1
		}
		if current < 100 {
			delta = 1
		}
		current += delta
		r.Save(int64(current))
		time.Sleep(8 * time.Millisecond)
	}

	//if len(usbPorts) == 0 {
	//	logrus.Error("no usb serial ports found")
	//	return
	//}
	//mode := &serial.Mode{
	//	BaudRate: 115200,
	//}
	//port, err := serial.Open(usbPorts[0], mode)
	//if err != nil {
	//	logrus.Error("cannot open serial port: %w", err)
	//}
	//
	//buff := make([]byte, 100)
	//for {
	//	n, err := port.Read(buff)
	//	if err != nil {
	//		logrus.Error("cannot serial read: %w", err)
	//		break
	//	}
	//	if n == 0 {
	//		logrus.Info("end of serial")
	//		break
	//	}
	//	line := string(buff[:n])
	//	logrus.Debugf("serial: %s", line)
	//}
}

func (r recordService) Clean() error {
	logrus.Info("cleaning up database")
	return r.recordRepository.PurgeOlderThan(time.Now().Add(-10 * time.Second))
}
