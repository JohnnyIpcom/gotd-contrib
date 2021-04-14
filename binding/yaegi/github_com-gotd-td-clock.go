// Code generated by 'yaegi extract github.com/gotd/td/clock'. DO NOT EDIT.

package yaegi

import (
	"github.com/gotd/neo"
	"github.com/gotd/td/clock"
	"reflect"
	"time"
)

func init() {
	Symbols["github.com/gotd/td/clock"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"StopTimer": reflect.ValueOf(clock.StopTimer),
		"System":    reflect.ValueOf(&clock.System).Elem(),

		// type definitions
		"Clock":  reflect.ValueOf((*clock.Clock)(nil)),
		"Ticker": reflect.ValueOf((*clock.Ticker)(nil)),
		"Timer":  reflect.ValueOf((*clock.Timer)(nil)),

		// interface wrapper definitions
		"_Clock":  reflect.ValueOf((*_github_com_gotd_td_clock_Clock)(nil)),
		"_Ticker": reflect.ValueOf((*_github_com_gotd_td_clock_Ticker)(nil)),
		"_Timer":  reflect.ValueOf((*_github_com_gotd_td_clock_Timer)(nil)),
	}
}

// _github_com_gotd_td_clock_Clock is an interface wrapper for Clock type
type _github_com_gotd_td_clock_Clock struct {
	WNow    func() time.Time
	WTicker func(d time.Duration) neo.Ticker
	WTimer  func(d time.Duration) neo.Timer
}

func (W _github_com_gotd_td_clock_Clock) Now() time.Time                    { return W.WNow() }
func (W _github_com_gotd_td_clock_Clock) Ticker(d time.Duration) neo.Ticker { return W.WTicker(d) }
func (W _github_com_gotd_td_clock_Clock) Timer(d time.Duration) neo.Timer   { return W.WTimer(d) }

// _github_com_gotd_td_clock_Ticker is an interface wrapper for Ticker type
type _github_com_gotd_td_clock_Ticker struct {
	WC     func() <-chan time.Time
	WReset func(d time.Duration)
	WStop  func()
}

func (W _github_com_gotd_td_clock_Ticker) C() <-chan time.Time   { return W.WC() }
func (W _github_com_gotd_td_clock_Ticker) Reset(d time.Duration) { W.WReset(d) }
func (W _github_com_gotd_td_clock_Ticker) Stop()                 { W.WStop() }

// _github_com_gotd_td_clock_Timer is an interface wrapper for Timer type
type _github_com_gotd_td_clock_Timer struct {
	WC     func() <-chan time.Time
	WReset func(d time.Duration)
	WStop  func() bool
}

func (W _github_com_gotd_td_clock_Timer) C() <-chan time.Time   { return W.WC() }
func (W _github_com_gotd_td_clock_Timer) Reset(d time.Duration) { W.WReset(d) }
func (W _github_com_gotd_td_clock_Timer) Stop() bool            { return W.WStop() }
