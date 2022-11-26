package utils


import (
	"main/db"
    "database/sql"
    "encoding/json"
	"github.com/dgrijalva/jwt-go"
	"main/models"
    "strings"
	"net/http"
    "fmt"
    "time"
    "reflect"
    "math/rand"
    "github.com/go-sql-driver/mysql"
    

)

// Additional scan types returned by the MySQL driver. I haven't looked at
// what PostgreSQL does.

type jsonNullInt64 struct {
    sql.NullInt64
}

func (v jsonNullInt64) MarshalJSON() ([]byte, error) {
    if !v.Valid {
        return json.Marshal(nil)
    }
    return json.Marshal(v.Int64)
}

type jsonNullFloat64 struct {
    sql.NullFloat64
}

func (v jsonNullFloat64) MarshalJSON() ([]byte, error) {
    if !v.Valid {
        return json.Marshal(nil)
    }
    return json.Marshal(v.Float64)
}

type jsonNullTime struct {
    mysql.NullTime
}

func (v jsonNullTime) MarshalJSON() ([]byte, error) {
    if !v.Valid {
        return json.Marshal(nil)
    }
    return json.Marshal(v.Time)
}

var jsonNullInt64Type = reflect.TypeOf(jsonNullInt64{})
var jsonNullFloat64Type = reflect.TypeOf(jsonNullFloat64{})
var jsonNullTimeType = reflect.TypeOf(jsonNullTime{})
var nullInt64Type = reflect.TypeOf(sql.NullInt64{})
var nullFloat64Type = reflect.TypeOf(sql.NullFloat64{})
var nullTimeType = reflect.TypeOf(mysql.NullTime{})

// SQLToJSON takes an SQL result and converts it to a nice JSON form. It also
// handles possibly-null values nicely. See https://stackoverflow.com/a/52572145/265521
func SQLToJSON(rows *sql.Rows) ([]byte, error) {
    columns, err := rows.Columns()
    if err != nil {
        return nil, fmt.Errorf("Column error: %v", err)
    }

    tt, err := rows.ColumnTypes()
    if err != nil {
        return nil, fmt.Errorf("Column type error: %v", err)
    }

    types := make([]reflect.Type, len(tt))
    for i, tp := range tt {
        st := tp.ScanType()
        if st == nil {
            return nil, fmt.Errorf("Scantype is null for column: %v", err)
        }
        switch st {
        case nullInt64Type:
            types[i] = jsonNullInt64Type
        case nullFloat64Type:
            types[i] = jsonNullFloat64Type
        case nullTimeType:
            types[i] = jsonNullTimeType
        default:
            types[i] = st
        }
    }

    values := make([]interface{}, len(tt))
    data := make(map[string][]interface{})

    for rows.Next() {
        for i := range values {
            values[i] = reflect.New(types[i]).Interface()
        }
        err = rows.Scan(values...)
        if err != nil {
            return nil, fmt.Errorf("Failed to scan values: %v", err)
        }
        for i, v := range values {
            data[columns[i]] = append(data[columns[i]], v)
        }
    }

    return json.Marshal(data)
}

// generate secret key
func GenerateSecretKey() string {
    key := ""
    x1 := rand.NewSource(time.Now().UnixNano())
    y1 := rand.New(x1)
    for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
        fmt.Println(i)
        asciiValue := y1.Intn(90-65) + 65
        character := rune(asciiValue)
        key = key + string(character)
    }
    return key
}

// Verify Function to check if token is still defined
func CheckAuth(header string, w http.ResponseWriter, r *http.Request) (string, string,bool) {
    // var jwtKey = []byte("secret_key")

    // getting token and user guid
    separateHeaders := strings.Fields(header)
    if len(separateHeaders) < 3 {
        return "", "", false
    }
    token := separateHeaders[1]
    userGuid := separateHeaders[2]

    fmt.Println("token: ", token)
    fmt.Println("userguid: ", userGuid)

    // Search Secret key
    DB := db.ConnectDB()
	rows, queryerr:= DB.Query("SELECT secret_key, user_guid, role FROM user_details WHERE user_guid = ?",userGuid)
	if queryerr != nil {
		fmt.Println("Error:", queryerr)
	}

	var myuser models.User
	for rows.Next() {

        newerr := rows.Scan(&myuser.Secret_Key, &myuser.User_Guid, &myuser.Role)
        if newerr != nil {
            fmt.Println("err:", newerr) 
        }
    }
    
    jwtKey := []byte (myuser.Secret_Key)
    // ----------------

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
            fmt.Println("Error in Signature")
            fmt.Println(err)
            return "", "", false
		}
        fmt.Println("Some other error")
        fmt.Println(err)
        return "", "", false
	}

	if !tkn.Valid {
        fmt.Println("Token not valid")
        return "", "", false
	}

    return myuser.User_Guid, myuser.Role, true
	// w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
}

//  Get UserGuid
func ValidateRole(role string, rolesToCheck[] string) bool {
    if len(role) > 0 && len(rolesToCheck) > 0 {
        for i := 0; i < len(rolesToCheck); i++ {
            if rolesToCheck[i] == role {
                return true
            }
        }
    }
    return false
}