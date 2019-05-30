package tests

import (
	"fmt"
	"github.com/google/uuid"
	. "github.com/sub0zero/go-sqlbuilder/sqlbuilder"
	"github.com/sub0zero/go-sqlbuilder/tests/.test_files/dvd_rental/test_sample/model"
	. "github.com/sub0zero/go-sqlbuilder/tests/.test_files/dvd_rental/test_sample/table"
	"gotest.tools/assert"
	"testing"
)

func TestAllTypesSelect(t *testing.T) {

	dest := []model.AllTypes{}

	err := AllTypes.SELECT(AllTypes.AllColumns).Query(db, &dest)

	fmt.Println(err)
	assert.NilError(t, err)

	assert.Equal(t, len(dest), 2)

	assert.DeepEqual(t, dest[0], allTypesRow0)
	assert.DeepEqual(t, dest[1], allTypesRow1)
}

func TestStringOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Text.EQ(AllTypes.Character),
		AllTypes.Text.EQ(String("Text")),
		AllTypes.Text.NOT_EQ(AllTypes.CharacterVaryingPtr),
		AllTypes.Text.NOT_EQ(String("Text")),
		AllTypes.Text.GT(AllTypes.Text),
		AllTypes.Text.GT(String("Text")),
		AllTypes.Text.GT_EQ(AllTypes.TextPtr),
		AllTypes.Text.GT_EQ(String("Text")),
		AllTypes.Text.LT(AllTypes.Character),
		AllTypes.Text.LT(String("Text")),
		AllTypes.Text.LT_EQ(AllTypes.CharacterVaryingPtr),
		AllTypes.Text.LT_EQ(String("Text")),
	)

	fmt.Println(query.DebugSql())

	err := query.Query(db, &struct{}{})

	assert.NilError(t, err)
}

func TestExpressionOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Integer.IS_NULL(),
		AllTypes.Timestamp.IS_NOT_NULL(),
	)

	fmt.Println(query.DebugSql())

	err := query.Query(db, &struct{}{})

	assert.NilError(t, err)
}

func TestBoolOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Boolean.EQ(AllTypes.BooleanPtr),
		AllTypes.Boolean.EQ(Bool(true)),
		AllTypes.Boolean.NOT_EQ(AllTypes.BooleanPtr),
		AllTypes.Boolean.NOT_EQ(Bool(false)),
		AllTypes.Boolean.IS_DISTINCT_FROM(AllTypes.BooleanPtr),
		AllTypes.Boolean.IS_DISTINCT_FROM(Bool(true)),
		AllTypes.Boolean.IS_NOT_DISTINCT_FROM(AllTypes.BooleanPtr),
		AllTypes.Boolean.IS_NOT_DISTINCT_FROM(Bool(true)),
		AllTypes.Boolean.IS_TRUE(),
		AllTypes.Boolean.IS_NOT_TRUE(),
		AllTypes.Boolean.IS_NOT_FALSE(),
		AllTypes.Boolean.IS_UNKNOWN(),
		AllTypes.Boolean.IS_NOT_UNKNOWN(),

		AllTypes.Boolean.AND(AllTypes.Boolean).EQ(AllTypes.Boolean.AND(AllTypes.Boolean)),
		AllTypes.Boolean.OR(AllTypes.Boolean).EQ(AllTypes.Boolean.AND(AllTypes.Boolean)),
	)

	fmt.Println(query.DebugSql())

	err := query.Query(db, &struct{}{})

	assert.NilError(t, err)
}

func TestNumericOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Numeric.EQ(AllTypes.Numeric),
		AllTypes.Decimal.EQ(Int(12)),
		AllTypes.Real.EQ(Float(12.12)),
		AllTypes.Smallint.NOT_EQ(AllTypes.Real),
		AllTypes.Integer.NOT_EQ(Int(12)),
		AllTypes.Bigint.NOT_EQ(Float(12)),
		AllTypes.Numeric.IS_DISTINCT_FROM(AllTypes.Numeric),
		AllTypes.Decimal.IS_DISTINCT_FROM(Int(12)),
		AllTypes.Real.IS_DISTINCT_FROM(Float(12.12)),
		AllTypes.Numeric.IS_NOT_DISTINCT_FROM(AllTypes.Numeric),
		AllTypes.Decimal.IS_NOT_DISTINCT_FROM(Int(12)),
		AllTypes.Real.IS_NOT_DISTINCT_FROM(Float(12.12)),
		AllTypes.Numeric.LT(AllTypes.Integer),
		AllTypes.Numeric.LT(Int(124)),
		AllTypes.Numeric.LT(Float(34.56)),
		AllTypes.Smallint.LT_EQ(AllTypes.Numeric),
		AllTypes.Integer.LT_EQ(Int(45)),
		AllTypes.Bigint.LT_EQ(Float(65)),
		AllTypes.Numeric.GT(AllTypes.Smallint),
		AllTypes.Numeric.GT(Int(124)),
		AllTypes.Numeric.GT(Float(34.56)),
		AllTypes.Smallint.GT_EQ(AllTypes.Numeric),
		AllTypes.Integer.GT_EQ(Int(45)),
		AllTypes.Bigint.GT_EQ(Float(65)),
	)

	fmt.Println(query.DebugSql())

	err := query.Query(db, &struct{}{})

	assert.NilError(t, err)
}

func TestTimeOperators(t *testing.T) {
	query := AllTypes.SELECT(
		AllTypes.Time.EQ(AllTypes.Time),
		AllTypes.Time.EQ(Time(23, 6, 6, 1)),
		AllTypes.Timez.EQ(AllTypes.TimezPtr),
		AllTypes.Timez.EQ(Timez(23, 6, 6, 222, +200)),
		AllTypes.Timestamp.EQ(AllTypes.TimestampPtr),
		AllTypes.Timestamp.EQ(Timestamp(2010, 10, 21, 15, 30, 12, 333)),
		AllTypes.Timestampz.EQ(AllTypes.TimestampzPtr),
		AllTypes.Timestampz.EQ(Timestampz(2010, 10, 21, 15, 30, 12, 444, 0)),
		AllTypes.Date.EQ(AllTypes.DatePtr),
		AllTypes.Date.EQ(Date(2010, 12, 3)),

		AllTypes.Time.NOT_EQ(AllTypes.Time),
		AllTypes.Time.NOT_EQ(Time(23, 6, 6, 10)),
		AllTypes.Timez.NOT_EQ(AllTypes.TimezPtr),
		AllTypes.Timez.NOT_EQ(Timez(23, 6, 6, 555, +200)),
		AllTypes.Timestamp.NOT_EQ(AllTypes.TimestampPtr),
		AllTypes.Timestamp.NOT_EQ(Timestamp(2010, 10, 21, 15, 30, 12, 666)),
		AllTypes.Timestampz.NOT_EQ(AllTypes.TimestampzPtr),
		AllTypes.Timestampz.NOT_EQ(Timestampz(2010, 10, 21, 15, 30, 12, 777, 0)),
		AllTypes.Date.NOT_EQ(AllTypes.DatePtr),
		AllTypes.Date.NOT_EQ(Date(2010, 12, 3)),

		AllTypes.Time.IS_DISTINCT_FROM(AllTypes.Time),
		AllTypes.Time.IS_DISTINCT_FROM(Time(23, 6, 6, 100)),

		AllTypes.Time.IS_NOT_DISTINCT_FROM(AllTypes.Time),
		AllTypes.Time.IS_NOT_DISTINCT_FROM(Time(23, 6, 6, 200)),

		AllTypes.Time.LT(AllTypes.Time),
		AllTypes.Time.LT(Time(23, 6, 6, 22)),

		AllTypes.Time.LT_EQ(AllTypes.Time),
		AllTypes.Time.LT_EQ(Time(23, 6, 6, 33)),

		AllTypes.Time.GT(AllTypes.Time),
		AllTypes.Time.GT(Time(23, 6, 6, 0)),

		AllTypes.Time.GT_EQ(AllTypes.Time),
		AllTypes.Time.GT_EQ(Time(23, 6, 6, 1)),
	)

	fmt.Println(query.DebugSql())

	err := query.Query(db, &struct{}{})

	assert.NilError(t, err)
}

var allTypesRow0 = model.AllTypes{
	SmallintPtr:        int16Ptr(1),
	Smallint:           1,
	IntegerPtr:         int32Ptr(300),
	Integer:            300,
	BigintPtr:          int64Ptr(50000),
	Bigint:             5000,
	DecimalPtr:         float64Ptr(11.44),
	Decimal:            11.44,
	NumericPtr:         float64Ptr(55.77),
	Numeric:            55.77,
	RealPtr:            float32Ptr(99.1),
	Real:               99.1,
	DoublePrecisionPtr: float64Ptr(11111111.22),
	DoublePrecision:    11111111.22,
	Smallserial:        1,
	Serial:             1,
	Bigserial:          1,
	//MoneyPtr: nil,
	//Money:
	CharacterVaryingPtr:  stringPtr("ABBA"),
	CharacterVarying:     "ABBA",
	CharacterPtr:         stringPtr("JOHN                                                                            "),
	Character:            "JOHN                                                                            ",
	TextPtr:              stringPtr("Some text"),
	Text:                 "Some text",
	ByteaPtr:             []byte("bytea"),
	Bytea:                []byte("bytea"),
	TimestampzPtr:        timestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0),
	Timestampz:           *timestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0),
	TimestampPtr:         timestampWithoutTimeZone("1999-01-08 04:05:06", 0),
	Timestamp:            *timestampWithoutTimeZone("1999-01-08 04:05:06", 0),
	DatePtr:              timestampWithoutTimeZone("1999-01-08 00:00:00", 0),
	Date:                 *timestampWithoutTimeZone("1999-01-08 00:00:00", 0),
	TimezPtr:             timeWithTimeZone("04:05:06 -0800"),
	Timez:                *timeWithTimeZone("04:05:06 -0800"),
	TimePtr:              timeWithoutTimeZone("04:05:06"),
	Time:                 *timeWithoutTimeZone("04:05:06"),
	IntervalPtr:          stringPtr("3 days 04:05:06"),
	Interval:             "3 days 04:05:06",
	BooleanPtr:           boolPtr(true),
	Boolean:              false,
	PointPtr:             stringPtr("(2,3)"),
	BitPtr:               stringPtr("101"),
	Bit:                  "101",
	BitVaryingPtr:        stringPtr("101111"),
	BitVarying:           "101111",
	TsvectorPtr:          stringPtr("'supernova':1"),
	Tsvector:             "'supernova':1",
	UUIDPtr:              uuidPtr("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
	UUID:                 uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
	XMLPtr:               stringPtr("<Sub>abc</Sub>"),
	XML:                  "<Sub>abc</Sub>",
	JSONPtr:              stringPtr(`{"a": 1, "b": 3}`),
	JSON:                 `{"a": 1, "b": 3}`,
	JsonbPtr:             stringPtr(`{"a": 1, "b": 3}`),
	Jsonb:                `{"a": 1, "b": 3}`,
	IntegerArrayPtr:      stringPtr("{1,2,3}"),
	IntegerArray:         "{1,2,3}",
	TextArrayPtr:         stringPtr("{breakfast,consulting}"),
	TextArray:            "{breakfast,consulting}",
	JsonbArray:           `{"{\"a\": 1, \"b\": 2}","{\"a\": 3, \"b\": 4}"}`,
	TextMultiDimArrayPtr: stringPtr("{{meeting,lunch},{training,presentation}}"),
	TextMultiDimArray:    "{{meeting,lunch},{training,presentation}}",
}

var allTypesRow1 = model.AllTypes{
	SmallintPtr:        nil,
	Smallint:           1,
	IntegerPtr:         nil,
	Integer:            300,
	BigintPtr:          nil,
	Bigint:             5000,
	DecimalPtr:         nil,
	Decimal:            11.44,
	NumericPtr:         nil,
	Numeric:            55.77,
	RealPtr:            nil,
	Real:               99.1,
	DoublePrecisionPtr: nil,
	DoublePrecision:    11111111.22,
	Smallserial:        2,
	Serial:             2,
	Bigserial:          2,
	//MoneyPtr: nil,
	//Money:
	CharacterVaryingPtr:  nil,
	CharacterVarying:     "ABBA",
	CharacterPtr:         nil,
	Character:            "JOHN                                                                            ",
	TextPtr:              nil,
	Text:                 "Some text",
	ByteaPtr:             nil,
	Bytea:                []byte("bytea"),
	TimestampzPtr:        nil,
	Timestampz:           *timestampWithTimeZone("1999-01-08 13:05:06 +0100 CET", 0),
	TimestampPtr:         nil,
	Timestamp:            *timestampWithoutTimeZone("1999-01-08 04:05:06", 0),
	DatePtr:              nil,
	Date:                 *timestampWithoutTimeZone("1999-01-08 00:00:00", 0),
	TimezPtr:             nil,
	Timez:                *timeWithTimeZone("04:05:06 -0800"),
	TimePtr:              nil,
	Time:                 *timeWithoutTimeZone("04:05:06"),
	IntervalPtr:          nil,
	Interval:             "3 days 04:05:06",
	BooleanPtr:           nil,
	Boolean:              false,
	PointPtr:             nil,
	BitPtr:               nil,
	Bit:                  "101",
	BitVaryingPtr:        nil,
	BitVarying:           "101111",
	TsvectorPtr:          nil,
	Tsvector:             "'supernova':1",
	UUIDPtr:              nil,
	UUID:                 uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
	XMLPtr:               nil,
	XML:                  "<Sub>abc</Sub>",
	JSONPtr:              nil,
	JSON:                 `{"a": 1, "b": 3}`,
	JsonbPtr:             nil,
	Jsonb:                `{"a": 1, "b": 3}`,
	IntegerArrayPtr:      nil,
	IntegerArray:         "{1,2,3}",
	TextArrayPtr:         nil,
	TextArray:            "{breakfast,consulting}",
	JsonbArray:           `{"{\"a\": 1, \"b\": 2}","{\"a\": 3, \"b\": 4}"}`,
	TextMultiDimArrayPtr: nil,
	TextMultiDimArray:    "{{meeting,lunch},{training,presentation}}",
}
