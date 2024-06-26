package cafe

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	os.Setenv("FOO", "foo")
	os.Setenv("BAR", "10")
	os.Setenv("BAZ", "true")

	s := NewCafeSchema(Schema{
		"foo": String("FOO").Require(),
		"bar": Int("BAR").Require(),
		"baz": Bool("BAZ"),
	})

	err := s.Initialize()
	if err != nil {
		t.Error(err)
	}

	fooRes, err := s.GetString("foo")
	if fooRes != "foo" || err != nil {
		t.Error("expected foo to be foo")
	}

	barRes, err := s.GetInt("bar")
	if barRes != 10 || err != nil {
		t.Error("expected bar to be 10")
	}

	bazRes, err := s.GetBool("baz")
	if bazRes != true || err != nil {
		t.Error("expected baz to be true")
	}
}

func TestServerOptions(t *testing.T) {
	// seed
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "postgres")
	config, err := New(
		Schema{
			"PORT": Int(
				"SERVER_PORT",
			), // PORT is an integer that is set by the SERVER_PORT environment variable and is not required
			"DB_HOST": String("DB_HOST").Require(), // DB_HOST is a string that is required
			"DB_PORT": Int(
				"DB_PORT",
			).Require(),
			// DB_PORT is an integer that is required
			"DB_USER": String("DB_USER").Require(), // DB_USER is a string that is required
			"DB_PASSWORD": String(
				"DB_PASSWORD",
			).Require(),
			// DB_PASSWORD is a string that is required
			"DB_NAME": String("DB_NAME").Require(), // DB_NAME is a string that is required
		},
	)
	if err != nil {
		t.Error(err)
	}
	sPort, err := config.GetInt("PORT")
	if sPort != 8080 || err != nil {
		t.Error("expected PORT to be 8080")
	}
	dbHost, err := config.GetString("DB_HOST")
	if dbHost != "localhost" || err != nil {
		t.Error("expected DB_HOST to be localhost")
	}
	dbPort, err := config.GetInt("DB_PORT")
	if dbPort != 5432 || err != nil {
		t.Error("expected DB_PORT to be 5432")
	}
	dbUser, err := config.GetString("DB_USER")
	if dbUser != "postgres" || err != nil {
		t.Error("expected DB_USER to be postgres")
	}
	dbPassword, err := config.GetString("DB_PASSWORD")
	if dbPassword != "postgres" || err != nil {
		t.Error("expected DB_PASSWORD to be postgres")
	}
	dbName, err := config.GetString("DB_NAME")
	if dbName != "postgres" || err != nil {
		t.Error("expected DB_NAME to be postgres")
	}
}

func TestServerOptionsWithDefault(t *testing.T) {
	config, err := New(
		Schema{
			"PORT": Int(
				"SERVER_PORT_HTTP",
			).Require().
				Default(8080),
			// PORT is an integer that is set by the SERVER_PORT environment variable and is not required
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, config)
	sPort, err := config.GetInt("PORT")
	assert.NoError(t, err)
	assert.Equal(t, 8080, sPort)
}

func TestServerOptionsStringWithDefault(t *testing.T) {
	config, err := New(
		Schema{
			"PORT": String(
				"SERVER_PORT_HTTP",
			).Require().
				Default("8080"),
			// PORT is an integer that is set by the SERVER_PORT environment variable and is not required
		},
	)
	assert.NoError(t, err)
	assert.NotNil(t, config)
	sPort, err := config.GetString("PORT")
	assert.NoError(t, err)
	assert.Equal(t, "8080", sPort)
}

func TestServerOptionsStringWithMisMatchDefaultDefault(t *testing.T) {
	config, err := New(
		Schema{
			"PORT": String(
				"SERVER_PORT_HTTP",
			).Require().
				Default("8080"),
			// PORT is an integer that is set by the SERVER_PORT environment variable and is not required
		},
	)
	assert.NoError(t, err)
	_, err = config.GetInt("PORT")
	assert.Error(t, err)
}

func TestSubSchema(t *testing.T) {
	// seed
	os.Setenv("FOO", "foo")
	os.Setenv("BAR_BAZ", "10")
	os.Setenv("BAR_QUX", "true")

	s := NewCafeSchema(Schema{
		"foo": String("FOO").Require(),
		"bar": SubSchema("BAR", Schema{
			"baz": Int("BAR_BAZ").Require(),
			"qux": Bool("BAR_QUX").Require(),
		}),
	})

	err := s.Initialize()
	if err != nil {
		t.Fatal(err)
	}

	fooRes, err := s.GetString("foo")
	if fooRes != "foo" || err != nil {
		t.Fatalf("expected foo to be foo got %s", fooRes)
	}

	subSchema, err := s.GetSubSchema("bar")
	if err != nil {
		t.Fatal(err)
	}

	barBazRes, err := subSchema.GetInt("baz")
	if barBazRes != 10 || err != nil {
		t.Fatalf("expected bar.baz to be 10 got %d", barBazRes)
	}

	barQuxRes, err := subSchema.GetBool("qux")
	if barQuxRes != true || err != nil {
		t.Fatalf("expected bar.qux to be true got %t", barQuxRes)
	}
}

func TestPrintJSON(t *testing.T) {
	// seed
	os.Setenv("FOO", "foo")
	os.Setenv("BAR_BAZ", "10")
	os.Setenv("BAR_QUX", "true")

	s := NewCafeSchema(Schema{
		"foo": String("FOO").Require(),
		"bar": SubSchema("BAR", Schema{
			"baz": Int("BAR_BAZ").Require(),
			"qux": Bool("BAR_QUX").Require(),
		}),
	})

	err := s.Initialize()
	if err != nil {
		fmt.Println(s.JSON())
		t.Fatal(err)
	}
	fmt.Println(s.JSON())
}

func TestPrintEnv(t *testing.T) {
	// seed
	os.Setenv("FOO", "foo")
	os.Setenv("BAR_BAZ", "10")

	s := NewCafeSchema(Schema{
		"foo": String("FOOBAR").Require(),
		"bar": SubSchema("BAR", Schema{
			"baz": Int("BAR_BAZ#").Require(),
			"qux": Bool("BAR_QUX").Require(),
		}),
	})

	err := s.Initialize()
	if err != nil {
		fmt.Println(s.Env())
	}
}
