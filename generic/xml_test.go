package generic

import (
	"encoding/xml"
	"strings"
	"testing"
)

var (
	TestXml = []byte(`
<config version="10.2.0" urldb="paloaltonetworks" detail-version="10.2.4">
    <normal> some text </normal>
    <person name="jane"></person>
    <unique uuid="1234-56-789"></unique>
    <nested>
        <and unknown="foobar"></and>
    </nested>
</config>`)
)

func TestUnmarshalDoesNotReturnError(t *testing.T) {
	var x Xml
	if err := xml.Unmarshal(TestXml, &x); err != nil {
		t.Fatalf("Error in unmarshal: %s", err)
	}
}

func TestUnmarshalSavesName(t *testing.T) {
	var x Xml
	if err := xml.Unmarshal(TestXml, &x); err != nil {
		t.Fatalf("Error in unmarshal: %s", err)
	}
	if len(x.Nodes) == 0 {
		t.Fatalf("no nodes present")
	}
	for _, node := range x.Nodes {
		if node.XMLName.Local == "person" {
			if node.Name == nil {
				t.Fatalf("config > person.name is nil")
			}
			if *node.Name != "jane" {
				t.Fatalf("config > person.name = %q", *node.Name)
			}
			return
		}
	}

	t.Fatalf("could not find config > person")
}

func TestUnmarshalSavesUuid(t *testing.T) {
	var x Xml
	if err := xml.Unmarshal(TestXml, &x); err != nil {
		t.Fatalf("Error in unmarshal: %s", err)
	}
	if len(x.Nodes) == 0 {
		t.Fatalf("no nodes present")
	}
	for _, node := range x.Nodes {
		if node.XMLName.Local == "unique" {
			if node.Uuid == nil {
				t.Fatalf("config > unique.uuid is nil")
			}
			if *node.Uuid != "1234-56-789" {
				t.Fatalf("config > unique.uuid = %q", *node.Uuid)
			}
			return
		}
	}

	t.Fatalf("could not find config > unique")
}

func TestUnmarshalSavesText(t *testing.T) {
	var x Xml
	if err := xml.Unmarshal(TestXml, &x); err != nil {
		t.Fatalf("Error in unmarshal: %s", err)
	}
	if len(x.Nodes) == 0 {
		t.Fatalf("no nodes present")
	}
	for _, node := range x.Nodes {
		if node.XMLName.Local == "normal" {
			if len(node.Text) == 0 {
				t.Fatalf("config > normal.text is empty")
			}
			if string(node.Text) != " some text " {
				t.Fatalf("config > normal.text = %q", *node.Uuid)
			}
			return
		}
	}

	t.Fatalf("could not find config > normal")
}

func TestMarshalDoesNotContainThings(t *testing.T) {
	var x Xml
	if err := xml.Unmarshal(TestXml, &x); err != nil {
		t.Fatalf("Error in unmarshal: %s", err)
	}

	b, err := xml.Marshal(x)
	if err != nil {
		t.Fatalf("Error in marshal: %s", err)
	}

	ans := string(b)

	missing := []string{"10.2.0", "foobar", "Node"}

	for _, chk := range missing {
		if strings.Contains(ans, chk) {
			t.Fatalf("marshalled bytes includes %q: %s", chk, ans)
		}
	}
}
