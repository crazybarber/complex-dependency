package repository

type ImplementationStatus int32

const (
	Waiting ImplementationStatus = iota
	Ready
	Done
)

type SourceSystem struct {
	Id          int32
	Name        string
	Description string
}

type EventType struct {
	Id             int64
	Name           string
	RecommendedUse string
	Status         ImplementationStatus
}

type EventTypeImplementation struct {
	PreciseUse string
	Status     ImplementationStatus

	EventTypeId int64 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	EventType   *EventType

	SourceSystemId int32 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	SourceSystem   *SourceSystem
}

type Field struct {
	Id             int64
	Name           string
	Type           string
	RecommendedUse string
	Status         ImplementationStatus
}

type FieldImplementation struct {
	PreciseUse string
	Status     ImplementationStatus

	FieldId int64 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	Field   *Field

	SourceSystemId int32 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	SourceSystem   *SourceSystem
}

type RestrictedValue struct {
	Value       string `sql:",pk"`
	Description string

	FieldImplementationId int64 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	FieldImplementation   *FieldImplementation
}
