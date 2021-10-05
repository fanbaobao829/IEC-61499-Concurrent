package functionblock

type FbInputEventInterface struct {
	Name      string
	DataLink  []FbInputEventInterface
	EventLink []FbEventLinked
}

type FbOutputEventInterface struct {
	Name      string
	DataLink  []FbOutputEventInterface
	EventLink []FbEventLinked
}

type FbInputDataInterface struct {
	Name     string
	Value    interface{}
	DataLink []FbDataLinked
}

type FbOutputDataInterface struct {
	Name     string
	Value    interface{}
	DataLink []FbDataLinked
}

type FbInfo struct {
	FbName      string
	FbIndex     int
	FbCache     int
	FbThreshold int
	FbTtl       int
	FbLast      int
	EventIn     []FbInputEventInterface
	EventOut    []FbOutputEventInterface
	DataIn      []FbInputDataInterface
	DataOut     []FbOutputDataInterface
}
