package gonlinesim

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTariffsBalance(t *testing.T) {
	Convey("GetTariffs method", t, func() {
		client := NewClient("")
		httpmock.ActivateNonDefault(client.client.GetClient())
		defer httpmock.Deactivate()
		url := "https://onlinesim.io/api/getBalance.php?apikey="
		fixture := struct {
			Result string `json:"response"`
		}{"ERROR_NO_KEY"}
		responder, err := httpmock.NewJsonResponder(http.StatusOK, fixture)
		So(err, ShouldBeNil)
		httpmock.RegisterResponder(http.MethodGet, url, responder)
		resp, err := client.GetTariffs(context.Background())
		So(err, ShouldNotBeNil)
		So(resp, ShouldBeNil)
	})
}
