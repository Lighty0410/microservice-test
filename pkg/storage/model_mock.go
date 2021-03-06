// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package storage

import (
	"sync"
	"time"
)

var (
	lockServiceMockAddEkadashi   sync.RWMutex
	lockServiceMockAddSession    sync.RWMutex
	lockServiceMockAddUser       sync.RWMutex
	lockServiceMockGetSession    sync.RWMutex
	lockServiceMockGetUser       sync.RWMutex
	lockServiceMockNextEkadashi  sync.RWMutex
	lockServiceMockUpdateSession sync.RWMutex
)

// Ensure, that ServiceMock does implement Service.
// If this is not the case, regenerate this file with moq.
var _ Service = &ServiceMock{}

// ServiceMock is a mock implementation of Service.
//
//     func TestSomethingThatUsesService(t *testing.T) {
//
//         // make and configure a mocked Service
//         mockedService := &ServiceMock{
//             AddEkadashiFunc: func(in1 *Ekadashi) error {
// 	               panic("mock out the AddEkadashi method")
//             },
//             AddSessionFunc: func(in1 *Session) error {
// 	               panic("mock out the AddSession method")
//             },
//             AddUserFunc: func(in1 *User) error {
// 	               panic("mock out the AddUser method")
//             },
//             GetSessionFunc: func(hash string) (*Session, error) {
// 	               panic("mock out the GetSession method")
//             },
//             GetUserFunc: func(username string) (User, error) {
// 	               panic("mock out the GetUser method")
//             },
//             NextEkadashiFunc: func(in1 time.Time) (*Ekadashi, error) {
// 	               panic("mock out the NextEkadashi method")
//             },
//             UpdateSessionFunc: func(in1 *Session) error {
// 	               panic("mock out the UpdateSession method")
//             },
//         }
//
//         // use mockedService in code that requires Service
//         // and then make assertions.
//
//     }
type ServiceMock struct {
	// AddEkadashiFunc mocks the AddEkadashi method.
	AddEkadashiFunc func(in1 *Ekadashi) error

	// AddSessionFunc mocks the AddSession method.
	AddSessionFunc func(in1 *Session) error

	// AddUserFunc mocks the AddUser method.
	AddUserFunc func(in1 *User) error

	// GetSessionFunc mocks the GetSession method.
	GetSessionFunc func(hash string) (*Session, error)

	// GetUserFunc mocks the GetUser method.
	GetUserFunc func(username string) (User, error)

	// NextEkadashiFunc mocks the NextEkadashi method.
	NextEkadashiFunc func(in1 time.Time) (*Ekadashi, error)

	// UpdateSessionFunc mocks the UpdateSession method.
	UpdateSessionFunc func(in1 *Session) error

	// calls tracks calls to the methods.
	calls struct {
		// AddEkadashi holds details about calls to the AddEkadashi method.
		AddEkadashi []struct {
			// In1 is the in1 argument value.
			In1 *Ekadashi
		}
		// AddSession holds details about calls to the AddSession method.
		AddSession []struct {
			// In1 is the in1 argument value.
			In1 *Session
		}
		// AddUser holds details about calls to the AddUser method.
		AddUser []struct {
			// In1 is the in1 argument value.
			In1 *User
		}
		// GetSession holds details about calls to the GetSession method.
		GetSession []struct {
			// Hash is the hash argument value.
			Hash string
		}
		// GetUser holds details about calls to the GetUser method.
		GetUser []struct {
			// Username is the username argument value.
			Username string
		}
		// NextEkadashi holds details about calls to the NextEkadashi method.
		NextEkadashi []struct {
			// In1 is the in1 argument value.
			In1 time.Time
		}
		// UpdateSession holds details about calls to the UpdateSession method.
		UpdateSession []struct {
			// In1 is the in1 argument value.
			In1 *Session
		}
	}
}

// AddEkadashi calls AddEkadashiFunc.
func (mock *ServiceMock) AddEkadashi(in1 *Ekadashi) error {
	if mock.AddEkadashiFunc == nil {
		panic("ServiceMock.AddEkadashiFunc: method is nil but Service.AddEkadashi was just called")
	}
	callInfo := struct {
		In1 *Ekadashi
	}{
		In1: in1,
	}
	lockServiceMockAddEkadashi.Lock()
	mock.calls.AddEkadashi = append(mock.calls.AddEkadashi, callInfo)
	lockServiceMockAddEkadashi.Unlock()
	return mock.AddEkadashiFunc(in1)
}

// AddEkadashiCalls gets all the calls that were made to AddEkadashi.
// Check the length with:
//     len(mockedService.AddEkadashiCalls())
func (mock *ServiceMock) AddEkadashiCalls() []struct {
	In1 *Ekadashi
} {
	var calls []struct {
		In1 *Ekadashi
	}
	lockServiceMockAddEkadashi.RLock()
	calls = mock.calls.AddEkadashi
	lockServiceMockAddEkadashi.RUnlock()
	return calls
}

// AddSession calls AddSessionFunc.
func (mock *ServiceMock) AddSession(in1 *Session) error {
	if mock.AddSessionFunc == nil {
		panic("ServiceMock.AddSessionFunc: method is nil but Service.AddSession was just called")
	}
	callInfo := struct {
		In1 *Session
	}{
		In1: in1,
	}
	lockServiceMockAddSession.Lock()
	mock.calls.AddSession = append(mock.calls.AddSession, callInfo)
	lockServiceMockAddSession.Unlock()
	return mock.AddSessionFunc(in1)
}

// AddSessionCalls gets all the calls that were made to AddSession.
// Check the length with:
//     len(mockedService.AddSessionCalls())
func (mock *ServiceMock) AddSessionCalls() []struct {
	In1 *Session
} {
	var calls []struct {
		In1 *Session
	}
	lockServiceMockAddSession.RLock()
	calls = mock.calls.AddSession
	lockServiceMockAddSession.RUnlock()
	return calls
}

// AddUser calls AddUserFunc.
func (mock *ServiceMock) AddUser(in1 *User) error {
	if mock.AddUserFunc == nil {
		panic("ServiceMock.AddUserFunc: method is nil but Service.AddUser was just called")
	}
	callInfo := struct {
		In1 *User
	}{
		In1: in1,
	}
	lockServiceMockAddUser.Lock()
	mock.calls.AddUser = append(mock.calls.AddUser, callInfo)
	lockServiceMockAddUser.Unlock()
	return mock.AddUserFunc(in1)
}

// AddUserCalls gets all the calls that were made to AddUser.
// Check the length with:
//     len(mockedService.AddUserCalls())
func (mock *ServiceMock) AddUserCalls() []struct {
	In1 *User
} {
	var calls []struct {
		In1 *User
	}
	lockServiceMockAddUser.RLock()
	calls = mock.calls.AddUser
	lockServiceMockAddUser.RUnlock()
	return calls
}

// GetSession calls GetSessionFunc.
func (mock *ServiceMock) GetSession(hash string) (*Session, error) {
	if mock.GetSessionFunc == nil {
		panic("ServiceMock.GetSessionFunc: method is nil but Service.GetSession was just called")
	}
	callInfo := struct {
		Hash string
	}{
		Hash: hash,
	}
	lockServiceMockGetSession.Lock()
	mock.calls.GetSession = append(mock.calls.GetSession, callInfo)
	lockServiceMockGetSession.Unlock()
	return mock.GetSessionFunc(hash)
}

// GetSessionCalls gets all the calls that were made to GetSession.
// Check the length with:
//     len(mockedService.GetSessionCalls())
func (mock *ServiceMock) GetSessionCalls() []struct {
	Hash string
} {
	var calls []struct {
		Hash string
	}
	lockServiceMockGetSession.RLock()
	calls = mock.calls.GetSession
	lockServiceMockGetSession.RUnlock()
	return calls
}

// GetUser calls GetUserFunc.
func (mock *ServiceMock) GetUser(username string) (User, error) {
	if mock.GetUserFunc == nil {
		panic("ServiceMock.GetUserFunc: method is nil but Service.GetUser was just called")
	}
	callInfo := struct {
		Username string
	}{
		Username: username,
	}
	lockServiceMockGetUser.Lock()
	mock.calls.GetUser = append(mock.calls.GetUser, callInfo)
	lockServiceMockGetUser.Unlock()
	return mock.GetUserFunc(username)
}

// GetUserCalls gets all the calls that were made to GetUser.
// Check the length with:
//     len(mockedService.GetUserCalls())
func (mock *ServiceMock) GetUserCalls() []struct {
	Username string
} {
	var calls []struct {
		Username string
	}
	lockServiceMockGetUser.RLock()
	calls = mock.calls.GetUser
	lockServiceMockGetUser.RUnlock()
	return calls
}

// NextEkadashi calls NextEkadashiFunc.
func (mock *ServiceMock) NextEkadashi(in1 time.Time) (*Ekadashi, error) {
	if mock.NextEkadashiFunc == nil {
		panic("ServiceMock.NextEkadashiFunc: method is nil but Service.NextEkadashi was just called")
	}
	callInfo := struct {
		In1 time.Time
	}{
		In1: in1,
	}
	lockServiceMockNextEkadashi.Lock()
	mock.calls.NextEkadashi = append(mock.calls.NextEkadashi, callInfo)
	lockServiceMockNextEkadashi.Unlock()
	return mock.NextEkadashiFunc(in1)
}

// NextEkadashiCalls gets all the calls that were made to NextEkadashi.
// Check the length with:
//     len(mockedService.NextEkadashiCalls())
func (mock *ServiceMock) NextEkadashiCalls() []struct {
	In1 time.Time
} {
	var calls []struct {
		In1 time.Time
	}
	lockServiceMockNextEkadashi.RLock()
	calls = mock.calls.NextEkadashi
	lockServiceMockNextEkadashi.RUnlock()
	return calls
}

// UpdateSession calls UpdateSessionFunc.
func (mock *ServiceMock) UpdateSession(in1 *Session) error {
	if mock.UpdateSessionFunc == nil {
		panic("ServiceMock.UpdateSessionFunc: method is nil but Service.UpdateSession was just called")
	}
	callInfo := struct {
		In1 *Session
	}{
		In1: in1,
	}
	lockServiceMockUpdateSession.Lock()
	mock.calls.UpdateSession = append(mock.calls.UpdateSession, callInfo)
	lockServiceMockUpdateSession.Unlock()
	return mock.UpdateSessionFunc(in1)
}

// UpdateSessionCalls gets all the calls that were made to UpdateSession.
// Check the length with:
//     len(mockedService.UpdateSessionCalls())
func (mock *ServiceMock) UpdateSessionCalls() []struct {
	In1 *Session
} {
	var calls []struct {
		In1 *Session
	}
	lockServiceMockUpdateSession.RLock()
	calls = mock.calls.UpdateSession
	lockServiceMockUpdateSession.RUnlock()
	return calls
}
