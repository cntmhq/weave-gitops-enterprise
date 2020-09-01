package acceptance

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
)

var theT *testing.T
var webDriver *agouti.Page

var gitProvider string
var seleniumServiceUrl string
var wkpDashboardUrl string

const ARTEFACTS_BASE_DIR string = "/tmp/workspace/"
const SCREENSHOTS_DIR string = ARTEFACTS_BASE_DIR + "screenshots/"
const JUNIT_TEST_REPORT_FILE string = ARTEFACTS_BASE_DIR + "wkp_junit.xml"

const ASSERTION_DEFAULT_TIME_OUT time.Duration = 15 * time.Second //15 seconds

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func TakeScreenShot(name string) string {
	if webDriver != nil {
		filepath := path.Join(SCREENSHOTS_DIR, name+".png")
		webDriver.Screenshot(filepath)
		return filepath
	}
	return ""
}

func GomegaFail(message string, callerSkip ...int) {
	if webDriver != nil {
		filepath := TakeScreenShot(String(16)) //Save the screenshot of failure
		fmt.Printf("\033[1;34mFailure screenshot is saved in file %s\033[0m \n", filepath)
	}

	//Pass this down to the default handler for onward processing
	ginkgo.Fail(message, callerSkip...)
}

func TestAcceptance(t *testing.T) {

	theT = t //Save the testing instance for later use

	//Cleanup the workspace dir, it helps when running locally
	_ = os.RemoveAll(ARTEFACTS_BASE_DIR)
	_ = os.MkdirAll(SCREENSHOTS_DIR, 0700)

	RegisterFailHandler(Fail)

	//Intercept the assertiona Failure
	gomega.RegisterFailHandler(GomegaFail)

	//JUnit style test report
	junitReporter := reporters.NewJUnitReporter(JUNIT_TEST_REPORT_FILE)
	RunSpecsWithDefaultAndCustomReporters(t, "WKP Acceptance Suite", []Reporter{junitReporter})

}

var _ = BeforeSuite(func() {
	//Set the suite level defaults

	SetDefaultEventuallyTimeout(ASSERTION_DEFAULT_TIME_OUT) //Things are slow on WKP UI

	gitProvider = "github" //TODO - Read from config.yaml
	seleniumServiceUrl = "http://localhost:4444/wd/hub"
	wkpDashboardUrl = "http://localhost:8090/"

})

var _ = AfterSuite(func() {
	//Tear down the suite level setup
	if webDriver != nil {
		Expect(webDriver.Destroy()).To(Succeed())
	}
})