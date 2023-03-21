package types

const (
	defaultKeyEvent  = "Event"
	defaultKeySite   = "Site"
	defaultKeyDate   = "Date"
	defaultKeyRound  = "Round"
	defaultKeyWhite  = "White"
	defaultKeyBlack  = "Black"
	defaultKeyResult = "Result"

	metadataNotAvailable = "N/A"
)

type Metadata map[string]string

func getMetadataKey(m Metadata, k string) (v string) {
	var ok bool

	if v, ok = m[k]; !ok {
		return metadataNotAvailable
	}

	return
}

func (m Metadata) Event() (evt string) {
	return getMetadataKey(m, defaultKeyEvent)
}

func (m Metadata) Site() (evt string) {
	return getMetadataKey(m, defaultKeySite)
}

func (m Metadata) Date() (evt string) {
	return getMetadataKey(m, defaultKeyDate)
}

func (m Metadata) Round() (evt string) {
	return getMetadataKey(m, defaultKeyRound)
}

func (m Metadata) White() (evt string) {
	return getMetadataKey(m, defaultKeyWhite)
}

func (m Metadata) Black() (evt string) {
	return getMetadataKey(m, defaultKeyBlack)
}

func (m Metadata) Result() (evt string) {
	return getMetadataKey(m, defaultKeyResult)
}
