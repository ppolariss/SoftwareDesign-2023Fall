package command

import e "design/myError"

type observer interface {
	update(command Command) error
}

var observers []observer

func notifyObserver(command Command) error {
	for _, o := range observers {
		err := o.update(command)
		if err != nil {
			return e.NewMyError("notifyObserver error")
		}
	}
	return nil
}

func registerObserver(o observer) {
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
