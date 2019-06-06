package main

import (
	"fmt"
	"os"
	"sync"
)

//START1 OMIT
type Repo struct {
	sync.RWMutex
	items map[string]int
}

var (
	r    *Repo
	once sync.Once
)

func GetInstance() *Repo {
	once.Do(func() {
		r = &Repo{
			items: make(map[string]int),
		}
	})
	return r
}

// END1 OMIT

//START2 OMIT
func (r *Repo) Set(key string, val int) {
	r.Lock()
	defer r.Unlock()
	r.items[key] = val
}

func (r *Repo) Get(key string) (int, error) {
	r.RLock()
	defer r.RUnlock()
	v, ok := r.items[key]
	if !ok {
		return 0, fmt.Errorf("The key %s is not present", key)
	}
	return v, nil
}

// END2 OMIT

//START3 OMIT
func main() {
	r := GetInstance()
	r.Set("jessy", 1)

	v, err := r.Get("jessy")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("jessy: %d\n", v)
}

//END3 OMIT
