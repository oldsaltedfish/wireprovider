package example

import (
	"database/sql"
	"fmt"
	"testing"
)

func Test_wireDB(t *testing.T) {
	tests := []struct {
		name    string
		want    *sql.DB
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := wireDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("wireDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			a := ""
			err = db.QueryRow("select 1").Scan(&a)
			if err != nil {
				t.Errorf("wireDB() error = %v", err)
				return
			}
			fmt.Println(a)
		})
	}
}
