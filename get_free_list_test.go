package gonlinesim

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	. "github.com/smartystreets/goconvey/convey"
)

const getFreeListSuccessResponse = `{
  "response": 1,
  "countries": [
    {
      "country": 7,
      "country_text": "Россия"
    },
    {
      "country": 380,
      "country_text": "Украина"
    },
    {
      "country": 77,
      "country_text": "Казахстан"
    },
    {
      "country": 44,
      "country_text": "Британия"
    },
    {
      "country": 371,
      "country_text": "Латвия"
    },
    {
      "country": 31,
      "country_text": "Нидерланды"
    },
    {
      "country": 34,
      "country_text": "Испания"
    },
    {
      "country": 372,
      "country_text": "Эстония"
    },
    {
      "country": 33,
      "country_text": "Франция"
    },
    {
      "country": 358,
      "country_text": "Финляндия"
    },
    {
      "country": 1,
      "country_text": "США"
    },
    {
      "country": 46,
      "country_text": "Швеция"
    },
    {
      "country": 39,
      "country_text": "Италия"
    },
    {
      "country": 351,
      "country_text": "Португалия"
    },
    {
      "country": 57,
      "country_text": "Колумбия"
    },
    {
      "country": 40,
      "country_text": "Румыния"
    },
    {
      "country": 212,
      "country_text": "Марокко"
    }
  ],
  "numbers": {
    "9099317316": {
      "country": 7,
      "data_humans": "6 дней назад",
      "full_number": "+79099317316",
      "is_archive": false
    },
    "9585027963": {
      "country": 7,
      "data_humans": "2 дня назад",
      "full_number": "+79585027963",
      "is_archive": false
    },
    "9915584205": {
      "country": 7,
      "data_humans": "2 дня назад",
      "full_number": "+79915584205",
      "is_archive": false
    },
    "9915584176": {
      "country": 7,
      "data_humans": "2 дня назад",
      "full_number": "+79915584176",
      "is_archive": false
    }
  },
  "messages": {
    "current_page": 1,
    "data": [
      {
        "text": "78123854000 received from onlinesim.io",
        "in_number": "***4000\n",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:44:21",
        "data_humans": "51 секунду назад"
      },
      {
        "text": "78123854000 received from onlinesim.io",
        "in_number": "***4000\n",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:43:47",
        "data_humans": "1 минуту назад"
      },
      {
        "text": "74957907277 received from onlinesim.io",
        "in_number": "***7277\n",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:36:44",
        "data_humans": "8 минут назад"
      },
      {
        "text": "74952235485 received from onlinesim.io",
        "in_number": "***5485\n",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:35:29",
        "data_humans": "9 минут назад"
      },
      {
        "text": "Ваш пароль 32166113 received from onlinesim.io",
        "in_number": "FCKRASNODAR",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:35:00",
        "data_humans": "10 минут назад"
      },
      {
        "text": "Введите 5407, чтобы подтвердить телефон в интернет-магазине е2е4 received from onlinesim.io",
        "in_number": "e2e4",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:31:13",
        "data_humans": "13 минут назад"
      },
      {
        "text": "79515984891 received from onlinesim.io",
        "in_number": "***4891\n",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:27:23",
        "data_humans": "17 минут назад"
      },
      {
        "text": "79622137000 received from onlinesim.io",
        "in_number": "***7000\n",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:24:37",
        "data_humans": "20 минут назад"
      },
      {
        "text": "79622137000 received from onlinesim.io",
        "in_number": "***7000\n",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:20:51",
        "data_humans": "24 минуты назад"
      },
      {
        "text": "79628078162 received from onlinesim.io",
        "in_number": "***8162\n",
        "my_number": 9099317316,
        "created_at": "2022-07-27 16:11:23",
        "data_humans": "33 минуты назад"
      }
    ],
    "first_page_url": "https://onlinesim.io/api/getFreeList?page=1",
    "from": 1,
    "last_page": 341,
    "last_page_url": "https://onlinesim.io/api/getFreeList?page=341",
    "links": [
      {
        "url": null,
        "label": "« Назад",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=1",
        "label": "1",
        "active": true
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=2",
        "label": "2",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=3",
        "label": "3",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=4",
        "label": "4",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=5",
        "label": "5",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=6",
        "label": "6",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=7",
        "label": "7",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=8",
        "label": "8",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=9",
        "label": "9",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=10",
        "label": "10",
        "active": false
      },
      {
        "url": null,
        "label": "...",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=340",
        "label": "340",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=341",
        "label": "341",
        "active": false
      },
      {
        "url": "https://onlinesim.io/api/getFreeList?page=2",
        "label": "Вперёд »",
        "active": false
      }
    ],
    "next_page_url": "https://onlinesim.io/api/getFreeList?page=2",
    "path": "https://onlinesim.io/api/getFreeList",
    "per_page": 10,
    "prev_page_url": null,
    "to": 10,
    "total": 3404,
    "number": "9099317316",
    "country": 7
  },
  "ignore": ""
}`

func TestFreeList(t *testing.T) {
	Convey("FreeList method", t, func() {
		client := NewClient("")
		httpmock.ActivateNonDefault(client.client.GetClient())
		defer httpmock.Deactivate()
		url := "https://onlinesim.io/api/getFreeList"
		responder := httpmock.NewStringResponder(http.StatusOK, getFreeListSuccessResponse)
		headers := http.Header{}
		headers.Set("Content-Type", "application/json")
		responder = responder.HeaderSet(headers)
		httpmock.RegisterResponder(http.MethodGet, url, responder)
		resp, err := client.GetFreeList(context.Background())
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.Response, ShouldEqual, 1)
		So(len(resp.Countries), ShouldEqual, 17)
	})
}
