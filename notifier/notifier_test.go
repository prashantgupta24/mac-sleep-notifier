package notifier

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestNotifier struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestNotifier))
}

//Run once before each test
func (suite *TestNotifier) SetupTest() {
	instance = nil
}

func (suite *TestNotifier) TestNotNil() {
	t := suite.T()
	notifierInstance := GetInstance()
	assert.NotNil(t, notifierInstance, "instance should not be nil")
}

func (suite *TestNotifier) TestType() {
	t := suite.T()
	notifierInstance := GetInstance()
	assert.IsType(t, &Notifier{}, notifierInstance, "instance should be of notifier type")
}

func (suite *TestNotifier) TestReturnValue() {
	t := suite.T()
	notifierInstance := GetInstance()
	notifierCh := notifierInstance.Start()
	channel := make(chan *Activity)

	assert.NotNil(t, notifierCh, "channel returned should not be nil")
	assert.IsType(t, channel, notifierCh, "returned value should be of channel type")
}

func (suite *TestNotifier) TestStartAndStop() {
	t := suite.T()
	notifierInstance := GetInstance()
	notifierInstance.Start()

	assert.True(t, notifierInstance.isStatusRunning(), "notifier should be running after start")

	notifierInstance.Quit()
	time.Sleep(time.Second) //wait for quit
	assert.False(t, notifierInstance.isStatusRunning(), "notifier should not be running after quit")
}
