package functionblock

type FbEventLinked struct {
	FbLinkedInedx               int
	FbLinkedInputEventInterface []FbInputEventInterface
}

type FbDataLinked struct {
	FbLinkedInedx              int
	FbLinkedInputDataInterface []FbInputDataInterface
}
