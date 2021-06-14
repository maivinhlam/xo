package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"time"
)

// ABitOfEverything represents a row from 'a_bit_of_everything.a_bit_of_everything'.
type ABitOfEverything struct {
	ABigint                  int64           `json:"a_bigint"`                    // a_bigint
	ABigintNullable          sql.NullInt64   `json:"a_bigint_nullable"`           // a_bigint_nullable
	ABinary                  []byte          `json:"a_binary"`                    // a_binary
	ABinaryNullable          []byte          `json:"a_binary_nullable"`           // a_binary_nullable
	ABit                     bool            `json:"a_bit"`                       // a_bit
	ABitNullable             sql.NullBool    `json:"a_bit_nullable"`              // a_bit_nullable
	ABlob                    []byte          `json:"a_blob"`                      // a_blob
	ABlobNullable            []byte          `json:"a_blob_nullable"`             // a_blob_nullable
	AChar                    string          `json:"a_char"`                      // a_char
	ACharNullable            sql.NullString  `json:"a_char_nullable"`             // a_char_nullable
	ADate                    time.Time       `json:"a_date"`                      // a_date
	ADateNullable            sql.NullTime    `json:"a_date_nullable"`             // a_date_nullable
	ADatetime                time.Time       `json:"a_datetime"`                  // a_datetime
	ADatetimeNullable        sql.NullTime    `json:"a_datetime_nullable"`         // a_datetime_nullable
	ADec                     float64         `json:"a_dec"`                       // a_dec
	ADecNullable             sql.NullFloat64 `json:"a_dec_nullable"`              // a_dec_nullable
	AFixed                   float64         `json:"a_fixed"`                     // a_fixed
	AFixedNullable           sql.NullFloat64 `json:"a_fixed_nullable"`            // a_fixed_nullable
	ADecimal                 float64         `json:"a_decimal"`                   // a_decimal
	ADecimalNullable         sql.NullFloat64 `json:"a_decimal_nullable"`          // a_decimal_nullable
	ADoublePrecision         float64         `json:"a_double_precision"`          // a_double_precision
	ADoublePrecisionNullable sql.NullFloat64 `json:"a_double_precision_nullable"` // a_double_precision_nullable
	AEnum                    AEnum           `json:"a_enum"`                      // a_enum
	AEnumNullable            AEnumNullable   `json:"a_enum_nullable"`             // a_enum_nullable
	AFloat                   float32         `json:"a_float"`                     // a_float
	AFloatNullable           sql.NullFloat64 `json:"a_float_nullable"`            // a_float_nullable
	AInt                     int             `json:"a_int"`                       // a_int
	AIntNullable             sql.NullInt64   `json:"a_int_nullable"`              // a_int_nullable
	AInteger                 int             `json:"a_integer"`                   // a_integer
	AIntegerNullable         sql.NullInt64   `json:"a_integer_nullable"`          // a_integer_nullable
	AJSON                    string          `json:"a_json"`                      // a_json
	AJSONNullable            sql.NullString  `json:"a_json_nullable"`             // a_json_nullable
	ALongblob                []byte          `json:"a_longblob"`                  // a_longblob
	ALongblobNullable        []byte          `json:"a_longblob_nullable"`         // a_longblob_nullable
	ALongtext                string          `json:"a_longtext"`                  // a_longtext
	ALongtextNullable        sql.NullString  `json:"a_longtext_nullable"`         // a_longtext_nullable
	AMediumblob              []byte          `json:"a_mediumblob"`                // a_mediumblob
	AMediumblobNullable      []byte          `json:"a_mediumblob_nullable"`       // a_mediumblob_nullable
	AMediumint               int             `json:"a_mediumint"`                 // a_mediumint
	AMediumintNullable       sql.NullInt64   `json:"a_mediumint_nullable"`        // a_mediumint_nullable
	AMediumtext              string          `json:"a_mediumtext"`                // a_mediumtext
	AMediumtextNullable      sql.NullString  `json:"a_mediumtext_nullable"`       // a_mediumtext_nullable
	ANumeric                 float64         `json:"a_numeric"`                   // a_numeric
	ANumericNullable         sql.NullFloat64 `json:"a_numeric_nullable"`          // a_numeric_nullable
	AReal                    float32         `json:"a_real"`                      // a_real
	ARealNullable            sql.NullFloat64 `json:"a_real_nullable"`             // a_real_nullable
	ASet                     SetOneTwo       `json:"a_set"`                       // a_set
	ASetNullable             SetOneTwo       `json:"a_set_nullable"`              // a_set_nullable
	ASmallint                int16           `json:"a_smallint"`                  // a_smallint
	ASmallintNullable        sql.NullInt64   `json:"a_smallint_nullable"`         // a_smallint_nullable
	AText                    string          `json:"a_text"`                      // a_text
	ATextNullable            sql.NullString  `json:"a_text_nullable"`             // a_text_nullable
	ATime                    string          `json:"a_time"`                      // a_time
	ATimeNullable            sql.NullString  `json:"a_time_nullable"`             // a_time_nullable
	ATimestamp               time.Time       `json:"a_timestamp"`                 // a_timestamp
	ATimestampNullable       time.Time       `json:"a_timestamp_nullable"`        // a_timestamp_nullable
	ATinyblob                []byte          `json:"a_tinyblob"`                  // a_tinyblob
	ATinyblobNullable        []byte          `json:"a_tinyblob_nullable"`         // a_tinyblob_nullable
	ATinyint                 int8            `json:"a_tinyint"`                   // a_tinyint
	ATinyintNullable         sql.NullInt64   `json:"a_tinyint_nullable"`          // a_tinyint_nullable
	ATinytext                string          `json:"a_tinytext"`                  // a_tinytext
	ATinytextNullable        sql.NullString  `json:"a_tinytext_nullable"`         // a_tinytext_nullable
	AVarbinary               []byte          `json:"a_varbinary"`                 // a_varbinary
	AVarbinaryNullable       []byte          `json:"a_varbinary_nullable"`        // a_varbinary_nullable
	AVarchar                 string          `json:"a_varchar"`                   // a_varchar
	AVarcharNullable         sql.NullString  `json:"a_varchar_nullable"`          // a_varchar_nullable
	AYear                    int16           `json:"a_year"`                      // a_year
	AYearNullable            sql.NullInt64   `json:"a_year_nullable"`             // a_year_nullable

}
