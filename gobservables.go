package gobservables

// Observable that dispatches events to Observers.
type Observable struct {
	Observers []Observer
}

// Subscribe adds a new Observer to the list of Observers and executes a callback.
func (o *Observable) Subscribe(callback func(o Observer, payload interface{})) Observer {
	observer := Observer{ID: len(o.Observers), observableRef: o, OnDispatch: callback}
	o.Observers = append(o.Observers, observer)
	return observer
}

// Unsubscribe removes the Observer from the list of Observers and executes a callback.
func (o *Observable) Unsubscribe(observer Observer, callback func()) {
	for index, tempObserver := range o.Observers {
		if tempObserver.ID == observer.ID {
			firstHalf := o.Observers[:index]
			secondHalf := o.Observers[index+1:]
			o.Observers = append(firstHalf, secondHalf...)
		}
	}
	callback()
}

// Dispatch events to a list of Observers.
func (o Observable) Dispatch(payload interface{}) {
	for _, observer := range o.Observers {
		observer.OnDispatch(observer, payload)
	}
}

// Observer to catch events dispatched by the Observable.
type Observer struct {
	ID            int
	observableRef *Observable
	OnDispatch    func(o Observer, payload interface{})
}

// Unsubscribe removes the observer from the observable's list of observers and executes a callback.
func (o Observer) Unsubscribe(callback func()) {
	o.observableRef.Unsubscribe(o, callback)
}
