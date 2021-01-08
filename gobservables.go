package gobservables

import "fmt"

// Observable documentation
type Observable struct {
	Observers []Observer
}

// Subscribe documentation
func (o *Observable) Subscribe(callback func(o Observer, payload interface{})) Observer {
	observer := Observer{ID: len(o.Observers), OnDispatch: callback}
	o.Observers = append(o.Observers, observer)
	return observer
}

// Unsubscribe documentation
func (o *Observable) Unsubscribe(observer Observer, callback func(payload interface{})) {
	fmt.Println("observers before unsubscribe:", len(o.Observers))
	for index, tempObserver := range o.Observers {
		if tempObserver.ID == observer.ID {
			firstHalf := o.Observers[:index]
			secondHalf := o.Observers[index+1:]
			o.Observers = append(firstHalf, secondHalf...)
		}
	}
	callback(len(o.Observers))
}

// Dispatch documentation
func (o Observable) Dispatch(payload interface{}) {
	for _, observer := range o.Observers {
		observer.OnDispatch(observer, payload)
	}
}

// Observer documentation
type Observer struct {
	ID         int
	OnDispatch func(o Observer, payload interface{})
}
