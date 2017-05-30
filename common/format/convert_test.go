package format

import (
	"fmt"
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	type testParam struct {
		param    interface{}
		adt      int
		expected string
	}

	test := []testParam{}
	test = append(test, testParam{123, 0, "123"})
	test = append(test, testParam{123, 5, "123"})
	test = append(test, testParam{123456789012345678, 0, "123456789012345678"})
	test = append(test, testParam{123456789012345678, 2, "123456789012345678"})
	test = append(test, testParam{3.14159, 0, "3"})
	test = append(test, testParam{3.14159, 2, "3.14"})
	test = append(test, testParam{3.14159, 5, "3.14159"})
	test = append(test, testParam{float64(3.14159), 0, "3"})
	test = append(test, testParam{float64(3.14159), 2, "3.14"})
	test = append(test, testParam{float64(3.14159), 5, "3.14159"})
	test = append(test, testParam{byte('a'), 0, "a"})
	test = append(test, testParam{byte('a'), 2, "a"})
	test = append(test, testParam{[]byte("abc"), 0, "abc"})
	test = append(test, testParam{[]byte("abc"), 2, "abc"})
	test = append(test, testParam{"a0s$d8*09s^ki09", 0, "a0s$d8*09s^ki09"})
	test = append(test, testParam{"a0s$d8*09s^ki09", 5, "a0s$d8*09s^ki09"})

	for _, val := range test {
		actual := ToString(val.param, val.adt)
		if actual != val.expected {
			t.Error(
				"For:", fmt.Sprintf("%v", val.param),
				"with additional", val.adt,
				"expected:", val.expected,
				"actual:", actual,
			)
		}
	}
}

func TestDateFormat(t *testing.T) {
	type testParam struct {
		text     string
		expected string
	}
	test := []testParam{}
	test = append(test, testParam{"Jan 2, 2013 7:54", "2 Januari 2013"})
	test = append(test, testParam{"Feb 12, 2013 17:54", "12 Februari 2013"})
	test = append(test, testParam{"Mar 29, 2016 7:54", "29 Maret 2016"})
	test = append(test, testParam{"Apr 2, 2013 7:54", "2 April 2013"})
	test = append(test, testParam{"May 2, 2013 7:54", "2 Mei 2013"})
	test = append(test, testParam{"Jun 2, 2013 7:54", "2 Juni 2013"})
	test = append(test, testParam{"Jul 2, 2013 7:54", "2 Juli 2013"})
	test = append(test, testParam{"Aug 2, 2013 7:54", "2 Agustus 2013"})
	test = append(test, testParam{"Sep 2, 2013 7:54", "2 September 2013"})
	test = append(test, testParam{"Oct 2, 2013 7:54", "2 Oktober 2013"})
	test = append(test, testParam{"Nov 2, 2013 7:54", "2 November 2013"})
	test = append(test, testParam{"Dec 2, 2013 7:54", "2 Desember 2013"})

	for _, val := range test {
		ts, err := time.Parse("Jan 2, 2006 15:04", val.text)
		if err != nil {
			t.Error(
				"For:", val.text,
				"expected:", val.expected,
				"error:", err.Error(),
			)
		}
		actual := DateFormat(ts)
		if actual != val.expected {
			t.Error(
				"For:", val.text,
				"expected:", val.expected,
				"actual:", actual,
			)
		}
	}
}

func TestDateTimeFormat(t *testing.T) {
	type testParam struct {
		text     string
		expected string
	}
	test := []testParam{}
	test = append(test, testParam{"Jan 2, 2013 7:54", "2 Januari 2013, 07:54 WIB"})
	test = append(test, testParam{"Feb 12, 2013 17:54", "12 Februari 2013, 17:54 WIB"})
	test = append(test, testParam{"Mar 29, 2016 7:54", "29 Maret 2016, 07:54 WIB"})
	test = append(test, testParam{"Apr 2, 2013 7:54", "2 April 2013, 07:54 WIB"})
	test = append(test, testParam{"May 2, 2013 7:54", "2 Mei 2013, 07:54 WIB"})
	test = append(test, testParam{"Jun 2, 2013 7:54", "2 Juni 2013, 07:54 WIB"})
	test = append(test, testParam{"Jul 2, 2013 7:54", "2 Juli 2013, 07:54 WIB"})
	test = append(test, testParam{"Aug 2, 2013 7:54", "2 Agustus 2013, 07:54 WIB"})
	test = append(test, testParam{"Sep 2, 2013 7:54", "2 September 2013, 07:54 WIB"})
	test = append(test, testParam{"Oct 2, 2013 7:54", "2 Oktober 2013, 07:54 WIB"})
	test = append(test, testParam{"Nov 2, 2013 7:54", "2 November 2013, 07:54 WIB"})
	test = append(test, testParam{"Dec 2, 2013 7:54", "2 Desember 2013, 07:54 WIB"})

	for _, val := range test {
		ts, err := time.Parse("Jan 2, 2006 15:04", val.text)
		if err != nil {
			t.Error(
				"For:", val.text,
				"expected:", val.expected,
				"error:", err.Error(),
			)
		}
		actual := DateTimeFormat(ts)
		if actual != val.expected {
			t.Error(
				"For:", val.text,
				"expected:", val.expected,
				"actual:", actual,
			)
		}
	}
}

func TestLPad(t *testing.T) {
	type testParam struct {
		text     string
		count    int
		filler   string
		expected string
	}
	test := []testParam{}
	test = append(test, testParam{"", 12, "a", "aaaaaaaaaaaa"})
	test = append(test, testParam{"1234567890", 5, "a", "1234567890"})
	test = append(test, testParam{"123", 12, "a", "aaaaaaaaa123"})
	test = append(test, testParam{"12345", 5, "a", "12345"})

	for _, val := range test {
		actual := LPad(val.text, val.count, val.filler)
		if actual != val.expected {
			t.Error(
				"For:", val.text,
				"with count", val.count,
				"and filler", val.filler,
				"expected:", val.expected,
				"actual:", actual,
			)
		}
	}
}

func TestRPad(t *testing.T) {
	type testParam struct {
		text     string
		count    int
		filler   string
		expected string
	}
	test := []testParam{}
	test = append(test, testParam{"", 12, "a", "aaaaaaaaaaaa"})
	test = append(test, testParam{"1234567890", 5, "a", "1234567890"})
	test = append(test, testParam{"123", 12, "a", "123aaaaaaaaa"})
	test = append(test, testParam{"12345", 5, "a", "12345"})

	for _, val := range test {
		actual := RPad(val.text, val.count, val.filler)
		if actual != val.expected {
			t.Error(
				"For:", val.text,
				"with count", val.count,
				"and filler", val.filler,
				"expected:", val.expected,
				"actual:", actual,
			)
		}
	}
}
