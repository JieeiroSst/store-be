package ratelimmit

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	Rate               int64
	MaxTokens          int64
	CurrentTokens      int64
	LastRefillTimstamp time.Time
	Mutex              sync.Mutex
}

func NewTokenBucket(rate int64, maxtokens int64) *TokenBucket {
	return &TokenBucket{
		Rate:      rate,
		MaxTokens: maxtokens,
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	end := time.Since(tb.LastRefillTimstamp)

	tokenTobeAdded := (end.Nanoseconds() * tb.Rate) / 1000000000
	tb.CurrentTokens = int64(math.Min(float64(tb.CurrentTokens+tokenTobeAdded), float64(tb.MaxTokens)))
	tb.LastRefillTimstamp = now
}

func (tb *TokenBucket) IsRequestAllowed(tokens int64) bool {
	tb.Mutex.Lock()
	defer tb.Mutex.Unlock()
	tb.refill()
	if tb.CurrentTokens >= tokens {
		tb.CurrentTokens = tb.CurrentTokens - tokens
		return true 
	}
	return false
}


var clientBucketMap = make(map[string]* TokenBucket) 

type Rule struct {
	MaxTokens int64
	Rate int64
}

var rulesMap = map[string]Rule {
	"gen-user": {MaxTokens: 1, Rate: 5},
}

func GetBucket(identifer string, userType string) *TokenBucket {
	if clientBucketMap[identifer]== nil {
		clientBucketMap[identifer] = NewTokenBucket(rulesMap[userType].MaxTokens, rulesMap[userType].Rate)
	}
	return clientBucketMap[identifer]
}
