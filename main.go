package zs

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"

    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
)

func init() {
    // Serve static files from "static" directory.
    http.Handle("/static/", http.FileServer(http.Dir(".")))
    // Redirect home URL to search page
    http.HandleFunc("/", redirectToSearchPageHandler)
    // Zillow search API endpoint
    http.HandleFunc("/api/search", searchHandler)
}

func redirectToSearchPageHandler(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/static/search.html", http.StatusMovedPermanently)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    u, err := url.Parse(r.RequestURI)
    m, _ := url.ParseQuery(u.RawQuery)

    address := m["address"][0]
    citystatezip := m["citystatezip"][0]

    apiKey := "X1-ZWz1dyb53fdhjf_6jziz"
    url := fmt.Sprintf("http://www.zillow.com/webservice/GetSearchResults.htm?zws-id=%v&address=%v&citystatezip=%v",
        apiKey, url.QueryEscape(address), url.QueryEscape(citystatezip))

    // Fetch search result from Zillow service
    ctx := appengine.NewContext(r)
    client := urlfetch.Client(ctx)
    resp, err := client.Get(url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    bytes, err := ioutil.ReadAll(resp.Body)

    // Parse XML response to data object
    var srs ZSearchResults
    xml.Unmarshal(bytes, &srs)

    // Serialize search result data object to JSON
    // and send to client
    jsonBytes, err := json.Marshal(srs)
    fmt.Fprint(w, string(jsonBytes))
}
