package test

import (
	"github.com/pangxieke/portray/config"
	"github.com/pangxieke/portray/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	theFixtures *fixtures
)

func Init() {
	var err error
	if err := config.Init("../"); err != nil {
		panic(err)
	}

	theFixtures, err = newFixtures()
	if err != nil {
		panic(err)
	}

	if err := model.InitForTest(theFixtures.DB); err != nil {
		panic(err)
	}

	//if err := model.Migrate(); err != nil {
	//	panic(err)
	//}

	return
}

/*	options:
	fixtures, bool, default: true
*/
func Prepare(t *testing.T, options ...map[string]interface{}) *assert.Assertions {
	prepareFixtures := true
	if len(options) > 0 {
		if v, ok := options[0]["fixture"]; ok {
			prepareFixtures = v.(bool)
		}
	}

	if prepareFixtures {
		theFixtures.Prepare(t)
	}
	return assert.New(t)
}
