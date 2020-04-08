package examples

import (
	"encoding/json"
	"testing"
)

func TestJsonGoparsec(t *testing.T) {
	bts, err := jsonParser().Parse(testData)
	if err != nil {
		t.Errorf("\"%s\" %s", testData[:bts], err)
	}
}

func BenchmarkJSONSTD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Valid(testData)
	}
}

func BenchmarkJSONGOP(b *testing.B) {
	p := jsonParser()
	for i := 0; i < b.N; i++ {
		p.Parse(testData)
	}
}

var testData = []byte(`[
  {
    "_id": "5e7fc3b1ee8a1e0968b99110",
    "index": 0,
    "guid": "20395076-7407-4f13-85be-8a9383b1e85b",
    "isActive": false,
    "balance": "$3,930.34",
    "picture": "http://placehold.it/32x32",
    "age": 32,
    "eyeColor": "blue",
    "name": "Melody Hendricks",
    "gender": "female",
    "company": "QUOTEZART",
    "email": "melodyhendricks@quotezart.com",
    "phone": "+1 (943) 403-3668",
    "address": "163 Richards Street, Disautel, South Carolina, 9399",
    "about": "Sint enim excepteur eu ea ea labore. Dolore officia esse amet sunt sunt elit duis ad ea amet aute ex adipisicing nostrud. Anim sit eiusmod est ex nulla aliqua. Officia tempor consectetur ipsum ipsum.\r\n",
    "registered": "2015-04-29T10:55:15 -03:00",
    "latitude": 64.009952,
    "longitude": 75.835803,
    "tags": [
      "quis",
      "officia",
      "nulla",
      "in",
      "laboris",
      "nulla",
      "qui"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Hardin Hamilton"
      },
      {
        "id": 1,
        "name": "Cruz Kirkland"
      },
      {
        "id": 2,
        "name": "Amy Gallagher"
      }
    ],
    "greeting": "Hello, Melody Hendricks! You have 2 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5e7fc3b19a38cf97b90047be",
    "index": 1,
    "guid": "f7fd3cd9-99b9-4b14-8f94-97724d886bbb",
    "isActive": false,
    "balance": "$2,849.35",
    "picture": "http://placehold.it/32x32",
    "age": 20,
    "eyeColor": "blue",
    "name": "Mccoy Winters",
    "gender": "male",
    "company": "JASPER",
    "email": "mccoywinters@jasper.com",
    "phone": "+1 (874) 455-3112",
    "address": "794 Main Street, Muir, Kansas, 8383",
    "about": "Id fugiat magna labore veniam ullamco do excepteur Lorem labore proident et voluptate nostrud. Anim labore excepteur ut eu excepteur nulla ullamco magna. Irure reprehenderit ad aliqua magna amet sint. Proident eiusmod magna duis do non id duis esse consectetur veniam mollit id.\r\n",
    "registered": "2014-02-07T10:39:31 -03:00",
    "latitude": -51.269066,
    "longitude": 178.892274,
    "tags": [
      "esse",
      "incididunt",
      "minim",
      "magna",
      "ullamco",
      "adipisicing",
      "id"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Margie Beach"
      },
      {
        "id": 1,
        "name": "Gutierrez Cantrell"
      },
      {
        "id": 2,
        "name": "Rodgers Crosby"
      }
    ],
    "greeting": "Hello, Mccoy Winters! You have 9 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5e7fc3b11c7b0d6c6274c25b",
    "index": 2,
    "guid": "3472bcaa-2e6e-4363-a73a-c9c26dcfc4d6",
    "isActive": false,
    "balance": "$1,897.79",
    "picture": "http://placehold.it/32x32",
    "age": 25,
    "eyeColor": "green",
    "name": "Kaufman Yates",
    "gender": "male",
    "company": "TERAPRENE",
    "email": "kaufmanyates@teraprene.com",
    "phone": "+1 (808) 511-2891",
    "address": "319 Banner Avenue, Irwin, American Samoa, 8268",
    "about": "Reprehenderit eiusmod aliqua labore aliqua duis irure mollit. Commodo id officia culpa irure. Aliquip reprehenderit velit incididunt ut culpa officia in. Enim mollit tempor ad in elit nostrud et esse nisi commodo deserunt qui aliqua. Excepteur commodo occaecat pariatur officia duis labore adipisicing. Pariatur sunt commodo irure cupidatat Lorem magna ad consequat enim laborum.\r\n",
    "registered": "2016-01-17T12:23:42 -03:00",
    "latitude": 10.571119,
    "longitude": -159.763252,
    "tags": [
      "magna",
      "anim",
      "elit",
      "nostrud",
      "et",
      "ut",
      "magna"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Herman Jennings"
      },
      {
        "id": 1,
        "name": "Valeria Griffin"
      },
      {
        "id": 2,
        "name": "Castro Knowles"
      }
    ],
    "greeting": "Hello, Kaufman Yates! You have 8 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5e7fc3b14691f8fb0f819738",
    "index": 3,
    "guid": "4d992275-7928-4aed-8241-cf3f376720b4",
    "isActive": false,
    "balance": "$3,192.87",
    "picture": "http://placehold.it/32x32",
    "age": 21,
    "eyeColor": "blue",
    "name": "Macias Thornton",
    "gender": "male",
    "company": "BUZZNESS",
    "email": "maciasthornton@buzzness.com",
    "phone": "+1 (925) 571-3972",
    "address": "587 Atlantic Avenue, Belgreen, Wisconsin, 5113",
    "about": "Cupidatat tempor nostrud dolor aute exercitation Lorem. Dolor ipsum incididunt magna deserunt qui tempor. Laboris nisi sunt officia voluptate est incididunt fugiat. Irure aliqua occaecat occaecat dolor deserunt ex labore et officia cupidatat non duis reprehenderit. Nulla reprehenderit irure enim nostrud mollit anim non non aliqua officia non dolor ea consectetur. Exercitation est ea labore eu velit Lorem non eiusmod do in tempor eiusmod occaecat. Qui ad occaecat velit pariatur aute.\r\n",
    "registered": "2018-05-22T12:17:59 -03:00",
    "latitude": 40.683081,
    "longitude": -48.243942,
    "tags": [
      "proident",
      "commodo",
      "ullamco",
      "in",
      "exercitation",
      "labore",
      "excepteur"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Brandie Washington"
      },
      {
        "id": 1,
        "name": "Harvey Terry"
      },
      {
        "id": 2,
        "name": "Robles Stephenson"
      }
    ],
    "greeting": "Hello, Macias Thornton! You have 3 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5e7fc3b14eb2d0043274ae89",
    "index": 4,
    "guid": "c37ab2ce-4475-4dca-9460-1401db20f4a1",
    "isActive": true,
    "balance": "$1,675.69",
    "picture": "http://placehold.it/32x32",
    "age": 39,
    "eyeColor": "brown",
    "name": "Hansen Davidson",
    "gender": "male",
    "company": "HYDROCOM",
    "email": "hansendavidson@hydrocom.com",
    "phone": "+1 (892) 553-2900",
    "address": "401 Irving Street, Leola, Ohio, 904",
    "about": "Laborum nisi ut ipsum commodo consectetur irure dolor aliqua ea sit exercitation proident exercitation ea. Pariatur laboris esse est proident ipsum adipisicing excepteur dolore magna minim pariatur aliqua proident pariatur. Anim esse ullamco cillum reprehenderit. Nostrud duis tempor reprehenderit sunt sint id sunt exercitation enim proident aliqua est voluptate. Mollit Lorem duis cupidatat dolore dolor deserunt velit labore pariatur esse cillum ea et nostrud. Quis quis incididunt ullamco ex. Irure occaecat mollit laboris sunt mollit est amet.\r\n",
    "registered": "2016-09-17T03:00:23 -03:00",
    "latitude": 81.027656,
    "longitude": 64.38203,
    "tags": [
      "cupidatat",
      "exercitation",
      "consequat",
      "culpa",
      "eiusmod",
      "id",
      "ullamco"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Nieves Porter"
      },
      {
        "id": 1,
        "name": "Sheppard Pollard"
      },
      {
        "id": 2,
        "name": "Jodi Mathis"
      }
    ],
    "greeting": "Hello, Hansen Davidson! You have 1 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "5e7fc3b1dd50a755af109f62",
    "index": 5,
    "guid": "aecb8359-c452-4d74-b5f9-dfdaefb6e145",
    "isActive": true,
    "balance": "$1,732.19",
    "picture": "http://placehold.it/32x32",
    "age": 23,
    "eyeColor": "brown",
    "name": "Fleming Burton",
    "gender": "male",
    "company": "OCEANICA",
    "email": "flemingburton@oceanica.com",
    "phone": "+1 (913) 447-3805",
    "address": "304 Jerome Street, Carbonville, South Dakota, 4269",
    "about": "Incididunt do commodo irure do cillum cupidatat. Commodo esse ea aute nostrud dolore dolor culpa cupidatat. Aliqua sunt ad reprehenderit nulla ipsum eu esse labore sint reprehenderit labore irure in. Incididunt id pariatur amet labore Lorem incididunt minim tempor anim. Ex mollit ut ipsum deserunt et. Ea elit ex ut ullamco exercitation voluptate sit nostrud aliquip voluptate ipsum sunt. Sint in laborum nulla Lorem enim irure.\r\n",
    "registered": "2016-04-10T10:11:55 -03:00",
    "latitude": 74.199343,
    "longitude": -170.721298,
    "tags": [
      "anim",
      "laborum",
      "excepteur",
      "fugiat",
      "labore",
      "nulla",
      "esse"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Trina Black"
      },
      {
        "id": 1,
        "name": "Maryann Dyer"
      },
      {
        "id": 2,
        "name": "Rasmussen Scott"
      }
    ],
    "greeting": "Hello, Fleming Burton! You have 6 unread messages.",
    "favoriteFruit": "strawberry"
  }
]`)
