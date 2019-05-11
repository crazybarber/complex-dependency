package model

type ImplementationStatus uint16

const (
	Planned ImplementationStatus = iota // TODO: Docu
	Defined
	Implemented
	Tested
	InUse
	Decommisioned
)


//Definition of a source system that produces data feeding the wide-row.
type SourceSystem struct {
	Id          uint32
	Name        string // Arbitrary name/tag for the source system
	Description string // What the source system represents, it's job/competence
}

// An event type is a data structure consisting of a subset of fields available in the wide-row structure.
// This structure represents the canonical definition of the event type.
type EventType struct {
	Id             uint64
	Name           string // The name of a event_type according to the wide-row design
	RecommendedUse string // The explanation of how to use this event type in terms of its universal business definition
	Status         ImplementationStatus //What is the status of the event type on a receiving point (data lake)
}

// Definition of an EventType in the context of a given source system
type EventTypeImplementation struct {
	PreciseUse string //The explanation of how to use this event type in terms of its business definition according to a certain source system
	Status     ImplementationStatus //What is the status of the event type on a sending point (source system)

	EventTypeId uint64 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"` // Reference to the canonical definition of an event type
	EventType   *EventType

	SourceSystemId uint32 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"` // Reference to a source system that the implementation refers to
	SourceSystem   *SourceSystem
}

// Definition of a field/attribute within the wide-row structure.
// This structure represents the canonical definition of the field.
type Field struct {
	Id             uint64
	Name           string // The name of a field according to the wide-row design
	Type           string // Data type of the field
	RecommendedUse string // The explanation of how to use this field in terms of its universal business definition
	Status         ImplementationStatus //What is the status of the field on a receiving point (data lake)
}

// Definition of a Field in the context of a given source system
type FieldImplementation struct {
	PreciseUse string //The explanation of how to use this field in terms of its business definition according to a certain source system
	Status     ImplementationStatus //What is the status of the field on a sending point (source system)

	FieldId uint64 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	Field   *Field

	SourceSystemId uint32 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`
	SourceSystem   *SourceSystem
}

// Definition of a value from a restricted list that applies for a certain field (and source system)
type RestrictedValue struct {
	Value       string `sql:",pk"`
	Description string // What this exact value represents business-wise

	FieldImplementationId uint64 `sql:"on_delete:RESTRICT, on_update: CASCADE, pk"`// Field (in context of a source system) that the value refers to
	FieldImplementation   *FieldImplementation
}
