package xmlapi

// CdataText is for getting CDATA contents of XML docs.
type CdataText struct {
    Text string `xml:",cdata"`
}
