package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codershubham350/bookings/internal/models"
)

// type postData struct {
// 	key   string
// 	value string
// }

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"gq", "/generals-quarters", "GET", http.StatusOK},
	{"ms", "/majors-suite", "GET", http.StatusOK},
	{"sa", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
	// {"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	// {"post-search-avail", "/search-availability", "POST", []postData{
	// 	{key: "start", value: "2022-10-06"},
	// 	{key: "end", value: "2022-10-07"},
	// }, http.StatusOK},
	// {"post-search-avail-json", "/search-availability-json", "POST", []postData{
	// 	{key: "start", value: "2022-10-06"},
	// 	{key: "end", value: "2022-10-07"},
	// }, http.StatusOK},
	// {"make-reservation-post", "/make-reservation", "POST", []postData{
	// 	{key: "first_name", value: "John"},
	// 	{key: "last_name", value: "Smith"},
	// 	{key: "email", value: "smith@mail.com"},
	// 	{key: "phone", value: "555-356-4521"},
	// }, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		//	if e.method == "GET" {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
		//	}
		// else {

		// 	values := url.Values{}
		// 	for _, x := range e.params {
		// 		values.Add(x.key, x.value)
		// 	}
		// 	resp, err := ts.Client().PostForm(ts.URL+e.url, values)
		// 	if err != nil {
		// 		t.Log(err)
		// 		t.Fatal(err)
		// 	}

		// 	if resp.StatusCode != e.expectedStatusCode {
		// 		t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		// 	}
		// }
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomId: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: got %d wanted %d", rr.Code, http.StatusOK)
	}

	// test case where reservation is not in session (reset everything)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test with non-existent room
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomId = 100
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

// postReservationTests is the test data for hte PostReservation handler test
// var postReservationTests = []struct {
// 	name                 string
// 	postedData           url.Values
// 	expectedResponseCode int
// 	expectedLocation     string
// 	expectedHTML         string
// }{
// 	{
// 		name: "valid-data",
// 		postedData: url.Values{
// 			"start_date": {"2050-01-01"},
// 			"end_date":   {"2050-01-02"},
// 			"first_name": {"John"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 			"room_id":    {"1"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedHTML:         "",
// 		expectedLocation:     "/reservation-summary",
// 	},
// 	{
// 		name:                 "missing-post-body",
// 		postedData:           nil,
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedHTML:         "",
// 		expectedLocation:     "/",
// 	},
// 	{
// 		name: "invalid-start-date",
// 		postedData: url.Values{
// 			"start_date": {"invalid"},
// 			"end_date":   {"2050-01-02"},
// 			"first_name": {"John"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 			"room_id":    {"1"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedHTML:         "",
// 		expectedLocation:     "/",
// 	},
// 	{
// 		name: "invalid-end-date",
// 		postedData: url.Values{
// 			"start_date": {"2050-01-01"},
// 			"end_date":   {"end"},
// 			"first_name": {"John"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 			"room_id":    {"1"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedHTML:         "",
// 		expectedLocation:     "/",
// 	},
// 	{
// 		name: "invalid-room-id",
// 		postedData: url.Values{
// 			"start_date": {"2050-01-01"},
// 			"end_date":   {"2050-01-02"},
// 			"first_name": {"John"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 			"room_id":    {"invalid"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedHTML:         "",
// 		expectedLocation:     "/",
// 	},
// 	{
// 		name: "invalid-data",
// 		postedData: url.Values{
// 			"start_date": {"2050-01-01"},
// 			"end_date":   {"2050-01-02"},
// 			"first_name": {"J"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 			"room_id":    {"1"},
// 		},
// 		expectedResponseCode: http.StatusOK,
// 		expectedHTML:         `action="/make-reservation"`,
// 		expectedLocation:     "",
// 	},
// 	{
// 		name: "database-insert-fails-reservation",
// 		postedData: url.Values{
// 			"start_date": {"2050-01-01"},
// 			"end_date":   {"2050-01-02"},
// 			"first_name": {"John"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 			"room_id":    {"2"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedHTML:         "",
// 		expectedLocation:     "/",
// 	},
// 	{
// 		name: "database-insert-fails-restriction",
// 		postedData: url.Values{
// 			"start_date": {"2050-01-01"},
// 			"end_date":   {"2050-01-02"},
// 			"first_name": {"John"},
// 			"last_name":  {"Smith"},
// 			"email":      {"john@smith.com"},
// 			"phone":      {"555-555-5555"},
// 			"room_id":    {"1000"},
// 		},
// 		expectedResponseCode: http.StatusSeeOther,
// 		expectedHTML:         "",
// 		expectedLocation:     "/",
// 	},
// }

// func TestRepository_PostReservation(t *testing.T) {
// 	reqBody := "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=569-985-4521")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// postedData := url.Values{}
// postedData.Add("start_date", "2050-01-01")
// postedData.Add("end_date", "2050-01-02")
// postedData.Add("first_name", "John")
// postedData.Add("last_name", "Smith")
// postedData.Add("email", "john@smith.com")
// postedData.Add("phone", "589-654-2365")
// postedData.Add("room_id", "1")

// 	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
// 	ctx := getCtx(req)
// 	req = req.WithContext(ctx)

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	rr := httptest.NewRecorder()

// 	handler := http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusSeeOther {
// 		t.Errorf("PostReservation handler returned wrong response code: got %d wanted %d", rr.Code, http.StatusSeeOther)
// 	}

// 	// test for missing post body
// 	req, _ = http.NewRequest("POST", "/make-reservation", nil)
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler returned wrong response code for missing post body: got %d wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for invalid start date
// 	reqBody = "start_date=invalid"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=569-985-4521")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler returned wrong response code for invalid start date: got %d wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for invalid end date
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=invalid")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=569-985-4521")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler returned wrong response code for invalid end date: got %d wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for invalid room id
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=569-985-4521")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=invalid")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler returned wrong response code for invalid room id: got %d wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for invalid data
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=J")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=569-985-4521")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler returned wrong response code for invalid room id: got %d wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for failure to insert reservation into database
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=569-985-4521")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler failed when trying to fail inserting reservation: got %d wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// test for failure to insert restriction into database
// 	reqBody = "start_date=2050-01-01"
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "first_name=John")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Smith")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=john@smith.com")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=569-985-4521")
// 	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1000")

// 	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
// 	ctx = getCtx(req)
// 	req = req.WithContext(ctx)
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	rr = httptest.NewRecorder()

// 	handler = http.HandlerFunc(Repo.PostReservation)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusTemporaryRedirect {
// 		t.Errorf("PostReservation handler failed when trying to fail inserting reservation: got %d wanted %d", rr.Code, http.StatusTemporaryRedirect)
// 	}

// 	// from git source
// 	// for _, e := range postReservationTests {
// 	// 	var req *http.Request
// 	// 	if e.postedData != nil {
// 	// 		req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(e.postedData.Encode()))
// 	// 	} else {
// 	// 		req, _ = http.NewRequest("POST", "/make-reservation", nil)

// 	// 	}
// 	// 	ctx := getCtx(req)
// 	// 	req = req.WithContext(ctx)
// 	// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	// 	rr := httptest.NewRecorder()

// 	// 	handler := http.HandlerFunc(Repo.PostReservation)

// 	// 	handler.ServeHTTP(rr, req)

// 	// 	if rr.Code != e.expectedResponseCode {
// 	// 		t.Errorf("%s returned wrong response code: got %d, wanted %d", e.name, rr.Code, e.expectedResponseCode)
// 	// 	}

// 	// 	if e.expectedLocation != "" {
// 	// 		// get the URL from test
// 	// 		actualLoc, _ := rr.Result().Location()
// 	// 		if actualLoc.String() != e.expectedLocation {
// 	// 			t.Errorf("failed %s: expected location %s, but got location %s", e.name, e.expectedLocation, actualLoc.String())
// 	// 		}
// 	// 	}

// 	// 	if e.expectedHTML != "" {
// 	// 		// read the response body into a string
// 	// 		html := rr.Body.String()
// 	// 		if !strings.Contains(html, e.expectedHTML) {
// 	// 			t.Errorf("failed %s: expected to find %s but did not", e.name, e.expectedHTML)
// 	// 		}
// 	// 	}

// 	// }
// }

func TestRepository_AvailabilityJSON(t *testing.T) {
	// first case - rooms are not available
	reqBody := "start_date=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2050-01-02")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	// create request
	req, _ := http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get context with session
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make handler handlerFunc
	handler := http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recorder
	rr := httptest.NewRecorder()

	// make request to out handler
	handler.ServeHTTP(rr, req)

	var j jsonResponse
	err := json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}
	// since we specified a start date > 2049-12-31, we expect no availability
	if j.OK {
		t.Error("Got availability when none was expected in AvailabilityJSON")
	}

	// create our request body
	reqBody = "start=2040-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2040-01-02")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	// create our request
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get the context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// create our response recorder, which satisfies the requirements
	// for http.ResponseWriter
	rr = httptest.NewRecorder()

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.AvailabilityJSON)

	// make the request to our handler
	handler.ServeHTTP(rr, req)

	// this time we want to parse JSON and get the expected response
	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json!")
	}

	// since we specified a start date < 2049-12-31, we expect availability
	if !j.OK {
		t.Error("Got no availability when some was expected in AvailabilityJSON")
	}

	// create our request
	req, _ = http.NewRequest("POST", "/search-availability-json", nil)

	// get the context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// create our response recorder, which satisfies the requirements
	// for http.ResponseWriter
	rr = httptest.NewRecorder()

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.AvailabilityJSON)

	// make the request to our handler
	handler.ServeHTTP(rr, req)

	// this time we want to parse JSON and get the expected response
	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json!")
	}

	// since we specified a start date < 2049-12-31, we expect availability
	if j.OK || j.Message != "Internal server error" {
		t.Error("Got availability when request body was empty")
	}
	// create our request body
	reqBody = "start=2060-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2060-01-02")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get the context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// create our response recorder, which satisfies the requirements
	// for http.ResponseWriter
	rr = httptest.NewRecorder()

	// make our handler a http.HandlerFunc
	handler = http.HandlerFunc(Repo.AvailabilityJSON)

	// make the request to our handler
	handler.ServeHTTP(rr, req)

	// this time we want to parse JSON and get the expected response
	err = json.Unmarshal([]byte(rr.Body.String()), &j)
	if err != nil {
		t.Error("failed to parse json!")
	}

	// since we specified a start date < 2049-12-31, we expect availability
	if j.OK || j.Message != "Error querying database" {
		t.Error("Got availability when simulating database error")
	}
}

func TestRepository_ReservationSummary(t *testing.T) {
	reservation := models.Reservation{
		RoomId: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/reservation-summary", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.ReservationSummary)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("ReservationSummary handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}
	req, _ = http.NewRequest("GET", "/reservation-summary", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.ReservationSummary)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("ReservationSummary handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}
}

func TestRepository_ChooseRoom(t *testing.T) {
	reservation := models.Reservation{
		RoomId: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/choose-room/1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	// set the RequestURI on the request so that we can grab the ID
	// from the URL
	req.RequestURI = "/choose-room/1"

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.ChooseRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	req, _ = http.NewRequest("GET", "/choose-room/1", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.RequestURI = "/choose-room/1"

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.ChooseRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	req, _ = http.NewRequest("GET", "/choose-room/fish", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.RequestURI = "/choose-room/fish"

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.ChooseRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_BookRoom(t *testing.T) {

	reservation := models.Reservation{
		RoomId: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/book-room?s=2050-01-01&e=2050-01-02&id=1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.BookRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	req, _ = http.NewRequest("GET", "/book-room?s=2040-01-01&e=2040-01-02&id=4", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.BookRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
