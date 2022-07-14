package model

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUID_Valid(t *testing.T) {
	tests := []struct {
		name    string
		exp     time.Time
		wantErr bool
	}{
		{"happy case", time.Now().Add(3 * time.Minute), false},
		{"Not valid", time.Now().Add(-3 * time.Minute), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := UID{Exp: tt.exp}
			if err := p.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("UID.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnmarshalJSON(t *testing.T) {
	jsonBytes, _ := base64.RawStdEncoding.DecodeString("eyJhdF9oYXNoIjoiRTE1ZGF3WUlXR0I3UjRyTUF3dC1zUSIsImF1ZCI6WyJ0aWtpLXRtcy1hcHAiXSwiYXV0aF90aW1lIjoxNjIwNjE2OTk1LCJlbWFpbCI6InNvbWVvbmVJblRtc0B0aWtpLnZuIiwiZXhwIjoxNjIwOTgwNDk0LCJpYXQiOjE2MjA5NzY4OTQsImlzcyI6Imh0dHBzOi8vdGVzdC50YWxhLnh5ei8iLCJqdGkiOiI0ZjI2ZWI2ZC04NmYwLTQyY2QtYTZmMC1hYmRhMzIyNmJjOWMiLCJuYW1lIjoic29tZW9uZSBpbiB0bXMiLCJub25jZSI6IiIsInJhdCI6MTYyMDYxNjk5Mywic2lkIjoiODQ0OGNjNzAtNjhhYy00MDNlLWEwY2UtNTQxOGYxOTBmZjFjIiwic3ViIjoiY3VzdG9tZXJzOjEwMDA5MDM5MiIsInRybiI6ImFjY291bnRzOjEwMDA2MzEifQ")
	uid := UID{}
	err := json.Unmarshal(jsonBytes, &uid)
	if err != nil {
		t.Fatalf("Should unmarshall success but got err %v", err)
	}

	exp, err := time.Parse(time.RFC3339, "2021-05-14T15:21:34+07:00")
	if err != nil {
		panic(err)
	}

	exp = exp.In(time.UTC)
	uid.Exp = uid.Exp.In(time.UTC)

	wantedUID := UID{
		ID:    "accounts:1000631",
		SUB:   "customers:100090392",
		Name:  "someone in tms",
		Exp:   exp,
		Email: "someoneInTms@tiki.vn",
	}

	if !reflect.DeepEqual(uid, wantedUID) {
		t.Fatalf("Unmarshal() %v, wantUid %v", uid, wantedUID)
	}
}

func TestUnmarshalJSON_err(t *testing.T) {
	jsonBytes := "{\"exp\":\"asd\"}"

	uid := UID{}
	err := json.Unmarshal([]byte(jsonBytes), &uid)
	if err == nil {
		t.Fatalf("Unmarshal() should get error")
	}
}

func TestUnmarshalJSON_trn(t *testing.T) {
	testTrns := []string{"trn_id", ""}

	for _, trn := range testTrns {
		t.Run("trn: "+trn, func(t *testing.T) {
			maps := map[string]interface{}{
				"email": "someoneInTiki@tiki.vn",
				"sub":   "sub_id",
				"exp":   1623471572,
				"name":  "someone in tiki",
			}

			if trn != "" {
				maps["trn"] = "trn_id"
			}

			jsonBytes, err := json.Marshal(maps)
			if err != nil {
				panic(err)
			}

			uid := UID{}
			if err = json.Unmarshal(jsonBytes, &uid); err != nil {
				t.Fatalf("Unmarshal() should success, but got err %v", err)
			}

			exp, err := time.Parse(time.RFC3339, "2021-06-12T11:19:32+07:00")
			if err != nil {
				panic(err)
			}

			exp = exp.In(time.UTC)
			uid.Exp = uid.Exp.In(time.UTC)

			wantedUID := UID{
				SUB:   "sub_id",
				Name:  "someone in tiki",
				Email: "someoneInTiki@tiki.vn",
				Exp:   exp,
			}

			if trn == "" {
				wantedUID.ID = "sub_id"
			} else {
				wantedUID.ID = trn
			}

			if !reflect.DeepEqual(uid, wantedUID) {
				t.Fatalf("Unmarshal() got uid %v, wantedUid %v", uid, wantedUID)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	exp := time.Date(2017, time.April, 1, 0, 0, 0, 0, time.UTC).Round(1 * time.Second)
	uid := &UID{
		ID:    "ID",
		SUB:   "hal",
		Name:  "name",
		Email: "tiki@gmail.com",
		Exp:   exp,
	}

	bytes, err := json.Marshal(uid)
	if err != nil {
		t.Fatalf("Should marshall JSON successfully, but got err %v", err)
	}

	fields := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &fields); err != nil {
		t.Fatalf("Should marshall JSON successfully, but got err %v", err)
	}

	assert.Equal(t, 5, len(fields))
	assert.Equal(t, "ID", fields["id"])
	assert.Equal(t, "hal", fields["sub"])
	assert.Equal(t, "name", fields["name"])
	assert.Equal(t, "tiki@gmail.com", fields["email"])
	assert.Equal(t, exp.Unix(), int64(fields["exp"].(float64)))
}
