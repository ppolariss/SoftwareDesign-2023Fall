package interfaces

import (
	"errors"
)

type observer interface {
	Update(command Command) error
}

var observers []observer

func NotifyObserver(command Command) error {
	for _, o := range observers {
		err := o.Update(command)
		if err != nil {
			return errors.New("notifyObserver error")
		}
	}
	return nil
}

func RegisterObserver(o observer) {
	observers = append(observers, o)
}

func removeObserver(o observer) {
	for i, ob := range observers {
		if ob == o {
			observers = append(observers[:i], observers[i+1:]...)
			break
		}
	}
}
