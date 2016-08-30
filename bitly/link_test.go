package bitly

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestLink_Lookup_single(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/link/lookup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, buildJSONRes(`{
			"link_lookup": [
				{
					"aggregate_link": "http://bit.ly/2V6CFi",
					"url": "http://www.google.com/"
				}
			]
		}`, 200, "OK"))
	})

	links, err := client.Link.Lookup("http://www.google.com/")
	if err != nil {
		t.Fatalf("Link.Lookup returned error: %v", err)
	}

	want := LinkLookupRes{AggregateLink: "http://bit.ly/2V6CFi", URL: "http://www.google.com/"}
	if !reflect.DeepEqual(links[0], want) {
		t.Errorf("Link.Lookup returned %+v, want %+v", links[0], want)
	}
}

func TestLink_Lookup_multiple(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/link/lookup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, buildJSONRes(`{
			"link_lookup": [
				{
					"aggregate_link": "http://bit.ly/2V6CFi",
					"url": "http://www.google.com/"
				},
				{
					"aggregate_link": "http://bit.ly/4VGeu",
					"url": "http://www.facebook.com/"
				}
			]
		}`, 200, "OK"))
	})

	links, err := client.Link.Lookup("http://www.google.com/", "http://www.facebook.com/")
	if err != nil {
		t.Fatalf("Link.Lookup returned error: %v", err)
	}

	want := []LinkLookupRes{
		LinkLookupRes{AggregateLink: "http://bit.ly/2V6CFi", URL: "http://www.google.com/"},
		LinkLookupRes{AggregateLink: "http://bit.ly/4VGeu", URL: "http://www.facebook.com/"},
	}
	if !reflect.DeepEqual(links, want) {
		t.Errorf("Link.Lookup returned %#v, want %#v", links, want)
	}
}
