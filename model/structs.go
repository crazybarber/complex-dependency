package model

type ImplementationStatus uint16

const (
	Planned ImplementationStatus = iota
	Defined
	Implemented
	Tested
	InUse
	Decommisioned
)

type SourceSystem struct {
	Id          uint32
	Name        string
	Description string
}

type EventType struct {
	Id             uint64
	Name           string
	RecommendedUse string
	Status         ImplementationStatus
}

type EventTypeImplementation struct {
	PreciseUse string
	Status     ImplementationStatus

	EventTypeId uint64 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	EventType   *EventType

	SourceSystemId uint32 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	SourceSystem   *SourceSystem
}

type Field struct {
	Id             uint64
	Name           string
	Type           string
	RecommendedUse string
	Status         ImplementationStatus
}

type FieldImplementation struct {
	PreciseUse string
	Status     ImplementationStatus

	FieldId uint64 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	Field   *Field

	SourceSystemId uint32 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	SourceSystem   *SourceSystem
}

type RestrictedValue struct {
	Value       string `sql:",pk"`
	Description string

	FieldImplementationId uint64 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	FieldImplementation   *FieldImplementation
}
