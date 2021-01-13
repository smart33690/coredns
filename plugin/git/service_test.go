package git

import (
	"fmt"
	"testing"
	"time"
)

func TestServices(t *testing.T) {
	repo := &Repo{URL: "git@github.com", Interval: time.Millisecond * 100}

	start(repo)
	if len(Services.services) != 1 {
		t.Errorf("Expected 1 service, found %v", len(Services.services))
	}

	Services.Stop(repo.URL, 1)
	if len(Services.services) != 0 {
		t.Errorf("Expected 1 service, found %v", len(Services.services))
	}

	repos := make([]*Repo, 5)
	for i := 0; i < 5; i++ {
		repos[i] = &Repo{URL: fmt.Sprintf("test%v", i), Interval: time.Second * 2}
		start(repos[i])
		if len(Services.services) != i+1 {
			t.Errorf("Expected %v service(s), found %v", i+1, len(Services.services))
		}
	}

	time.Sleep(time.Microsecond * 500)
	Services.Stop(string(repos[0].URL), 1)
	if len(Services.services) != 4 {
		t.Errorf("Expected %v service(s), found %v", 4, len(Services.services))
	}

	repo = &Repo{URL: "git@github.com", Interval: time.Second}
	start(repo)
	if len(Services.services) != 5 {
		t.Errorf("Expected %v service(s), found %v", 5, len(Services.services))
	}

	repo = &Repo{URL: "git@github.com", Interval: time.Second * 2}
	start(repo)
	if len(Services.services) != 6 {
		t.Errorf("Expected %v service(s), found %v", 6, len(Services.services))
	}

	time.Sleep(time.Microsecond * 500)
	Services.Stop(string(repo.URL), -1)
	if len(Services.services) != 4 {
		t.Errorf("Expected %v service(s), found %v", 4, len(Services.services))
	}

	for _, repo := range repos {
		Services.Stop(string(repo.URL), -1)
	}
	if len(Services.services) != 0 {
		t.Errorf("Expected %v service(s), found %v", 0, len(Services.services))
	}

	repo.Interval = 0
	start(repo)
	if len(Services.services) != 0 {
		t.Errorf("Expected %v service(s), found %v", 0, len(Services.services))
	}
}
