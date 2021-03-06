package codelab

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ubuntu/tutorial-deployment/paths"
)

func TestDiscover(t *testing.T) {
	testCases := []struct {
		tutorialPaths []string
		expected      []string
		wantErr       bool
	}{
		{[]string{}, nil, false},
		{[]string{"/doesnt/exist"}, nil, true},
		{[]string{"testdata/nothing"}, nil, false},
		{[]string{"testdata/flat"}, []string{"testdata/flat/tut1.md", "testdata/flat/tut2.md"}, false},
		{[]string{"testdata/nested"}, []string{"testdata/nested/subdir1/subsub/tut1.md", "testdata/nested/subdir1/subsub/tut2.md", "testdata/nested/subdir2/tut1.md", "testdata/nested/subdir2/tut2.md"}, false},
		{[]string{"testdata/flat", "testdata/flat2"}, []string{"testdata/flat/tut1.md", "testdata/flat/tut2.md", "testdata/flat2/tut1.md"}, false},
		{[]string{"testdata/withgdoc"}, []string{"gdoc:mytut1", "gdoc:mytut2"}, false},
		{[]string{"testdata/withgdocduplicate"}, []string{"gdoc:mytut1", "gdoc:mytut2"}, false},
		{[]string{"testdata/withignored"}, []string{"testdata/withignored/tut1.md"}, false},
		{[]string{"testdata/flat", "testdata/withgdoc", "testdata/withignored"}, []string{"testdata/flat/tut1.md", "testdata/flat/tut2.md", "gdoc:mytut1", "gdoc:mytut2", "testdata/withignored/tut1.md"}, false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("scanning %s", tc.tutorialPaths), func(t *testing.T) {
			// Setup/Teardown
			p, teardown := paths.MockPath()
			defer teardown()
			p.TutorialInputs = tc.tutorialPaths

			// Test
			tutorials, err := Discover()

			if (err != nil) != tc.wantErr {
				t.Errorf("Discover() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(tutorials, tc.expected) {
				t.Errorf("got %+v; want %+v", tutorials, tc.expected)
			}
		})
	}
}
