package parsers

import (
	"github.com/steve-care-software/steve/applications/languages/parsers"
)

const (
	// TypeUnsignedFlag representsthe unsigned flag
	TypeUnsignedFlag uint8 = 1 << iota

	// TypeIntegerFlag represents the integer flag
	TypeIntegerFlag

	// TypeFloatFlag represents the float flag
	TypeFloatFlag

	// TypeStringFlag represents the string flag
	TypeStringFlag

	// TypeBoolFlag represents the bool flag
	TypeBoolFlag
)

const (
	// OpArithmetic represents the arithmetic flag
	OpArithmetic uint16 = 1 << iota

	// OpAriPlus represents the plus flag
	OpAriPlus

	// OpAriMinus represents the minus flag
	OpAriMinus

	// OpAriMultiply represents the multiply flag
	OpAriMultiply

	// OpAriDivide represents the divide flag
	OpAriDivide

	// OpLogical represents the logical flag
	OpLogical

	// OpLogAnd represents the and flag
	OpLogAnd

	// OpLogOr represents the or flag
	OpLogOr

	// OpRelational represents the relational flag
	OpRelational

	// OpRelNot represents the not flag
	OpRelNot

	// OpRelEqual represents the equal flag
	OpRelEqual

	// OpRelLessThan represents the lessThan flag
	OpRelLessThan

	// OpRelGreaterThan represents the greaterThan flag
	OpRelGreaterThan
)

// NewEvents create the event maps
func NewEvents(bitsize uint) map[string][]parsers.EventFunc {
	out := map[string][]parsers.EventFunc{}
	mp := newEvents(bitsize)
	for name, evtFunc := range mp {
		out[name] = []parsers.EventFunc{
			evtFunc,
		}
	}

	return out
}

// NewProgramBuilder creates a new program builder
func NewProgramBuilder() ProgramBuilder {
	return createProgramBuilder()
}

// NewHeaderVariablesBuilder creates a new headerVariables builder
func NewHeaderVariablesBuilder() HeaderVariablesBuilder {
	return createHeaderVariablesBuilder()
}

// NewHeaderVariableBuilder creates a new headerVariable builder
func NewHeaderVariableBuilder() HeaderVariableBuilder {
	return createHeaderVariableBuilder()
}

// NewInstructionsBuilder creates a new instructions builder
func NewInstructionsBuilder() InstructionsBuilder {
	return createInstructionsBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewDirectionBuilder creates a new direction builder
func NewDirectionBuilder() DirectionBuilder {
	return createDirectionBuilder()
}

// NewVoidFuncCallBuilder creates a new voidFuncCall builder
func NewVoidFuncCallBuilder() VoidFuncCallBuilder {
	return createVoidFuncCallBuilder()
}

// NewForDeclarationBuilder creates a new forDeclaration builder
func NewForDeclarationBuilder() ForDeclarationBuilder {
	return createForDeclarationBuilder()
}

// NewForStatementBuilder creates a new forStatement builder
func NewForStatementBuilder() ForStatementBuilder {
	return createForStatementBuilder()
}

// NewIfDeclarationBuilder creates a new ifDeclaration builder
func NewIfDeclarationBuilder() IfDeclarationBuilder {
	return createIfDeclarationBuilder()
}

// NewIntIncrementBuilder creates a new intIncrementBuilder
func NewIntIncrementBuilder() IntIncrementBuilder {
	return createIntIncrementBuilder()
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	return createAssignmentBuilder()
}

// NewAssigneeBuilder creates a new assignee builder
func NewAssigneeBuilder() AssigneeBuilder {
	return createAssigneeBuilder()
}

// NewAssigneeNameFirstsBuilder creates a new assignee name firsts builder
func NewAssigneeNameFirstsBuilder() AssigneeNameFirstsBuilder {
	return createAssigneeNameFirstsBuilder()
}

// NewAssigneeNameFirstBuilder creates a new assigneeNameFirst builder
func NewAssigneeNameFirstBuilder() AssigneeNameFirstBuilder {
	return createAssigneeNameFirstBuilder()
}

// NewAssigneeNamesBuilder creates a new assignee names builder
func NewAssigneeNamesBuilder() AssigneeNamesBuilder {
	return createAssigneeNamesBuilder()
}

// NewAssigneeNameBuilder creates a new assigneeName builder
func NewAssigneeNameBuilder() AssigneeNameBuilder {
	return createAssigneeNameBuilder()
}

// NewAssignablesBuilder creates a new assignables builder
func NewAssignablesBuilder() AssignablesBuilder {
	return createAssignablesBuilder()
}

// NewAssignableBuilder creates a new assignable builder
func NewAssignableBuilder() AssignableBuilder {
	return createAssignableBuilder()
}

// NewSingleAssignableBuilder creates a new singleAssignable builder
func NewSingleAssignableBuilder() SingleAssignableBuilder {
	return createSingleAssignableBuilder()
}

// NewListFetchElementBuilder creates a new listFetchElement builder
func NewListFetchElementBuilder() ListFetchElementBuilder {
	return createListfetchElementBuilder()
}

// NewTotalListValueBuilder creates a new totalListValue builder
func NewTotalListValueBuilder() TotalListValueBuilder {
	return createTotalListValueBuilder()
}

// NewTotalListAssignableBuilder creates a new totalListAssignable builder
func NewTotalListAssignableBuilder() TotalListAssignableBuilder {
	return createTotalListAssignableBuilder()
}

// NewListValueBuilder creates a new listValue builder
func NewListValueBuilder() ListValueBuilder {
	return createListValueBuilder()
}

// NewAppendBuilder creates a new append builder
func NewAppendBuilder() AppendBuilder {
	return createAppendBuilder()
}

// NewSliceBuilder creates a new slice builder
func NewSliceBuilder() SliceBuilder {
	return createSliceBuilder()
}

// NewSliceDelimiterBuilder creates a new slice delimiter builder
func NewSliceDelimiterBuilder() SliceDelimiterBuilder {
	return createSliceDelimiterBuilder()
}

// NewIndexValueBuilder creates a new indexValue builder
func NewIndexValueBuilder() IndexValueBuilder {
	return createIndexValueBuilder()
}

// NewOperationBuilder creates a new operation builder
func NewOperationBuilder() OperationBuilder {
	return createOperationBuilder()
}

// NewOperationSuffixBuilder creates a new operationSuffix builder
func NewOperationSuffixBuilder() OperationSuffixBuilder {
	return createOperationSuffixBuilder()
}

// NewOperatonElementBuilder creates a new operation element builder
func NewOperatonElementBuilder() OperationElementBuilder {
	return createOperationElementBuilder()
}

// NewOperationFuncCallBuilder creates a new operationFuncCall builder
func NewOperationFuncCallBuilder() OperationFuncCallBuilder {
	return createOperationFuncCallBuilder()
}

// NewVariableDeclarationsBuilder creates a new variable declarations builder
func NewVariableDeclarationsBuilder() VariableDeclarationsBuilder {
	return createVariableDeclarationsBuilder()
}

// NewVariableDeclarationBuilder creates a new variable declaration builder
func NewVariableDeclarationBuilder() VariableDeclarationBuilder {
	return createVariableDeclarationBuilder()
}

// NewValueBuilder creates a new value builder instance
func NewValueBuilder() ValueBuilder {
	typeBuilder := NewTypeBuilder()
	return createValueBuilder(typeBuilder)
}

// NewTypeBuilder creates a new type builder instance
func NewTypeBuilder() TypeBuilder {
	return createTypeBuilder()
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithInstructions(instructions Instructions) ProgramBuilder
	WithVariables(variables HeaderVariables) ProgramBuilder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Instructions() Instructions
	HasVariables() bool
	Variables() HeaderVariables
}

// HeaderVariablesBuilder represents an header variables builder
type HeaderVariablesBuilder interface {
	Create() HeaderVariablesBuilder
	WithVariables(variables []HeaderVariable) HeaderVariablesBuilder
	Now() (HeaderVariables, error)
}

// HeaderVariables represents header variables
type HeaderVariables interface {
	All() []HeaderVariable
}

// HeaderVariableBuilder represents an header variable builder
type HeaderVariableBuilder interface {
	Create() HeaderVariableBuilder
	WithDeclaration(declaration VariableDeclaration) HeaderVariableBuilder
	WithKeyname(keyname string) HeaderVariableBuilder
	IsMandatory() HeaderVariableBuilder
	IsInput() HeaderVariableBuilder
	Now() (HeaderVariable, error)
}

// HeaderVariable represents an header variable
type HeaderVariable interface {
	IsMandatory() bool
	IsInput() bool
	Declaration() VariableDeclaration
	Keyname() string
}

// InstructionsBuilder represents an instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithInstructions(instructions []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	All() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	WithVariableDeclaration(variableDeclaration VariableDeclaration) InstructionBuilder
	WithIfDeclaration(ifDeclaration IfDeclaration) InstructionBuilder
	WithForDeclaration(forDeclaration ForDeclaration) InstructionBuilder
	WithFuncCall(funcCall VoidFuncCall) InstructionBuilder
	WithDirection(direction Direction) InstructionBuilder
	WithIntIncrement(intIncrement IntIncrement) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsAssignment() bool
	Assignment() Assignment
	IsVariableDeclaration() bool
	VariableDeclaration() VariableDeclaration
	IsIfDeclaration() bool
	IfDeclaration() IfDeclaration
	IsForDeclaration() bool
	ForDeclaration() ForDeclaration
	IsFuncCall() bool
	FuncCall() VoidFuncCall
	IsDirection() bool
	Direction() Direction
	IsIntIncrement() bool
	IntIncrement() IntIncrement
}

// DirectionBuilder represents a direction builder
type DirectionBuilder interface {
	Create() DirectionBuilder
	IsBreak() DirectionBuilder
	IsContinue() DirectionBuilder
	Now() (Direction, error)
}

// Direction represents a direction
type Direction interface {
	IsBreak() bool
	IsContinue() bool
}

// VoidFuncCallBuilder represents a VoidFuncCall builder
type VoidFuncCallBuilder interface {
	Create() VoidFuncCallBuilder
	WithLog(log Assignable) VoidFuncCallBuilder
	Now() (VoidFuncCall, error)
}

// VoidFuncCall represents a void func call
type VoidFuncCall interface {
	IsLog() bool
	Log() Assignable
}

// ForDeclarationBuilder represents a forDeclaration builder
type ForDeclarationBuilder interface {
	Create() ForDeclarationBuilder
	WithStatement(statement ForStatement) ForDeclarationBuilder
	WithInstructions(instructions Instructions) ForDeclarationBuilder
	Now() (ForDeclaration, error)
}

// ForDeclaration represents a for declaration
type ForDeclaration interface {
	Statement() ForStatement
	HasInstructions() bool
	Instructions() Instructions
}

// ForStatementBuilder represents a forStatement builder
type ForStatementBuilder interface {
	Create() ForStatementBuilder
	WithIndex(index AssigneeName) ForStatementBuilder
	WithValue(value AssigneeName) ForStatementBuilder
	WithIterable(iterable TotalListValue) ForStatementBuilder
	Now() (ForStatement, error)
}

// ForStatement represents a for statement
type ForStatement interface {
	Index() AssigneeName
	Value() AssigneeName
	Iterable() TotalListValue
}

// IfDeclarationBuilder represents an ifDeclaration builder
type IfDeclarationBuilder interface {
	Create() IfDeclarationBuilder
	WithCondition(condition Operation) IfDeclarationBuilder
	WithInstructions(instructions Instructions) IfDeclarationBuilder
	Now() (IfDeclaration, error)
}

// IfDeclaration represents an if declaration
type IfDeclaration interface {
	Condition() Operation
	HasInstructions() bool
	Instructions() Instructions
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithAssignee(assignee Assignee) AssignmentBuilder
	WithAssignables(assignables Assignables) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Assignee() Assignee
	Assignables() Assignables
}

// IntIncrementBuilder represents an int increment builder
type IntIncrementBuilder interface {
	Create() IntIncrementBuilder
	WithVariableName(variableName string) IntIncrementBuilder
	WithIncrement(increment Value) IntIncrementBuilder
	Now() (IntIncrement, error)
}

// IntIncrement represents an int increment
type IntIncrement interface {
	VariableName() string
	Increment() Value
}

// AssigneeBuilder represents an assignee builder
type AssigneeBuilder interface {
	Create() AssigneeBuilder
	WithFirst(first AssigneeNameFirsts) AssigneeBuilder
	WithName(name AssigneeNames) AssigneeBuilder
	Now() (Assignee, error)
}

// Assignee represents an assignee
type Assignee interface {
	IsFirst() bool
	First() AssigneeNameFirsts
	IsName() bool
	Name() AssigneeNames
}

// AssigneeNameFirstBuilder represents an assigneeNameFirst builder
type AssigneeNameFirstBuilder interface {
	Create() AssigneeNameFirstBuilder
	WithDeclaration(declaration VariableDeclaration) AssigneeNameFirstBuilder
	WithAssigneeName(name AssigneeName) AssigneeNameFirstBuilder
	Now() (AssigneeNameFirst, error)
}

// AssigneeNameFirstsBuilder represents an assigneeNameFirsts builder
type AssigneeNameFirstsBuilder interface {
	Create() AssigneeNameFirstsBuilder
	WithNames(names []AssigneeNameFirst) AssigneeNameFirstsBuilder
	Now() (AssigneeNameFirsts, error)
}

// AssigneeNameFirsts represents assignee name firsts
type AssigneeNameFirsts interface {
	All() []AssigneeNameFirst
}

// AssigneeNameFirst represents an assignee name first
type AssigneeNameFirst interface {
	IsDeclaration() bool
	Declaration() VariableDeclaration
	IsAssigneeName() bool
	AssigneeName() AssigneeName
}

// AssigneeNameBuilder represents an assignee name builder
type AssigneeNameBuilder interface {
	Create() AssigneeNameBuilder
	WithVariable(variable string) AssigneeNameBuilder
	IsSkip() AssigneeNameBuilder
	Now() (AssigneeName, error)
}

// AssigneeNamesBuilder represents an assigneeNames builder
type AssigneeNamesBuilder interface {
	Create() AssigneeNamesBuilder
	WithNames(names []AssigneeName) AssigneeNamesBuilder
	Now() (AssigneeNames, error)
}

// AssigneeNames represents assignee names
type AssigneeNames interface {
	All() []AssigneeName
}

// AssigneeName represents an assignee name
type AssigneeName interface {
	IsSkip() bool
	IsVariable() bool
	Variable() string
}

// AssignablesBuilder represents an assignales builder
type AssignablesBuilder interface {
	Create() AssignablesBuilder
	WithAssignables(assignables []Assignable) AssignablesBuilder
	Now() (Assignables, error)
}

// Assignables represents assignables
type Assignables interface {
	IsEmpty() bool
	All() []Assignable
}

// AssignableBuilder represents an assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithList(list TotalListAssignable) AssignableBuilder
	WithSingle(single SingleAssignable) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsList() bool
	List() TotalListAssignable
	IsSingle() bool
	Single() SingleAssignable
}

// SingleAssignableBuilder represents a single assignable builder
type SingleAssignableBuilder interface {
	Create() SingleAssignableBuilder
	WithDeclarableValue(declarableValue Value) SingleAssignableBuilder
	WithOperation(operation Operation) SingleAssignableBuilder
	WithExtract(extract TotalListAssignable) SingleAssignableBuilder
	WithListFetchElement(listFetchElement ListFetchElement) SingleAssignableBuilder
	Now() (SingleAssignable, error)
}

// SingleAssignable represents a single assignable
type SingleAssignable interface {
	IsDeclarableValue() bool
	DeclarableValue() Value
	IsOperation() bool
	Operation() Operation
	IsExtract() bool
	Extract() TotalListAssignable
	IsListFetchElement() bool
	ListFetchElement() ListFetchElement
}

// ListFetchElementBuilder represents a listFetchElement builder
type ListFetchElementBuilder interface {
	Create() ListFetchElementBuilder
	WithList(list TotalListValue) ListFetchElementBuilder
	WithIndex(index IndexValue) ListFetchElementBuilder
	Now() (ListFetchElement, error)
}

// ListFetchElement represents a listFetchElement
type ListFetchElement interface {
	List() TotalListValue
	Index() IndexValue
}

// TotalListValueBuilder represents a total list value builder
type TotalListValueBuilder interface {
	Create() TotalListValueBuilder
	WithAssignable(assignable TotalListAssignable) TotalListValueBuilder
	WithVariableName(variableName string) TotalListValueBuilder
	Now() (TotalListValue, error)
}

// TotalListValue represents a total list value
type TotalListValue interface {
	IsAssignable() bool
	Assignable() TotalListAssignable
	IsVariableName() bool
	VariableName() string
}

// TotalListAssignableBuilder represents a total list assignable builder
type TotalListAssignableBuilder interface {
	Create() TotalListAssignableBuilder
	WithListValue(listValue ListValue) TotalListAssignableBuilder
	WithSlice(slice Slice) TotalListAssignableBuilder
	WithAppend(apnd Append) TotalListAssignableBuilder
	Now() (TotalListAssignable, error)
}

// TotalListAssignable represents a total list assignable
type TotalListAssignable interface {
	IsListValue() bool
	ListValue() ListValue
	IsSlice() bool
	Slice() Slice
	IsAppend() bool
	Append() Append
}

// ListValueBuilder represents a list value builder
type ListValueBuilder interface {
	Create() ListValueBuilder
	WithAssignables(assignables Assignables) ListValueBuilder
	WithVariableName(variableName string) ListValueBuilder
	Now() (ListValue, error)
}

// ListValue represents a list value
type ListValue interface {
	IsAssignables() bool
	Assignables() Assignables
	IsVariableName() bool
	VariableName() string
}

// AppendBuilder represents an append builder
type AppendBuilder interface {
	Create() AppendBuilder
	WithList(list ListValue) AppendBuilder
	WithAssignable(assignable SingleAssignable) AppendBuilder
	Now() (Append, error)
}

// Append represents an append
type Append interface {
	List() ListValue
	Assignable() SingleAssignable
}

// SliceBuilder represents a slice builder
type SliceBuilder interface {
	Create() SliceBuilder
	WithList(list ListValue) SliceBuilder
	WithDelimiter(delimiter SliceDelimiter) SliceBuilder
	Now() (Slice, error)
}

// Slice represents a slice
type Slice interface {
	List() ListValue
	Delimiter() SliceDelimiter
}

// SliceDelimiterBuilder represents a slice delimiter builder
type SliceDelimiterBuilder interface {
	Create() SliceDelimiterBuilder
	WithIndex(index IndexValue) SliceDelimiterBuilder
	WithAmount(amount IndexValue) SliceDelimiterBuilder
	Now() (SliceDelimiter, error)
}

// SliceDelimiter represents a sub list delimiter
type SliceDelimiter interface {
	Index() IndexValue
	HasAmount() bool
	Amount() IndexValue
}

// IndexValueBuilder represents an index value builder
type IndexValueBuilder interface {
	Create() IndexValueBuilder
	WithValue(value Value) IndexValueBuilder
	WithVariableName(variableName string) IndexValueBuilder
	Now() (IndexValue, error)
}

// IndexValue represents an index value
type IndexValue interface {
	IsValue() bool
	Value() Value
	IsVariableName() bool
	VariableName() string
}

// OperationBuilder represents an operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithElement(element OperationElement) OperationBuilder
	WithSuffix(suffix OperationSuffix) OperationBuilder
	Now() (Operation, error)
}

// Operation represents an operation
type Operation interface {
	Optimize() Operation
	Element() OperationElement
	HasSuffix() bool
	Suffix() OperationSuffix
}

// OperationSuffixBuilder represents an operation suffix builder
type OperationSuffixBuilder interface {
	Create() OperationSuffixBuilder
	WithOperator(operator uint16) OperationSuffixBuilder
	WithElement(element OperationElement) OperationSuffixBuilder
	Now() (OperationSuffix, error)
}

// OperationSuffix represents an operation suffix
type OperationSuffix interface {
	Operator() uint16
	Element() OperationElement
}

// OperationElementBuilder represents an operation element builder
type OperationElementBuilder interface {
	Create() OperationElementBuilder
	WithBoolValue(boolValue bool) OperationElementBuilder
	WithVariableName(variableName string) OperationElementBuilder
	WithComputableValue(computableValue Value) OperationElementBuilder
	WithOperation(operation Operation) OperationElementBuilder
	WithFuncCall(funcCall OperationFuncCall) OperationElementBuilder
	Now() (OperationElement, error)
}

// OperationElement represents an operation element
type OperationElement interface {
	IsBoolValue() bool
	BoolValue() *bool
	IsVariableName() bool
	VariableName() *string
	IsComputableValue() bool
	ComputableValue() Value
	IsOperation() bool
	Operation() Operation
	IsFuncCall() bool
	FuncCall() OperationFuncCall
}

// OperationFuncCallBuilder represents an operation func call builder
type OperationFuncCallBuilder interface {
	Create() OperationFuncCallBuilder
	WithNull(null Operation) OperationFuncCallBuilder
	Now() (OperationFuncCall, error)
}

// OperationFuncCall represents a func call that can be called inside an operation
type OperationFuncCall interface {
	IsNull() bool
	Null() Operation
}

// VariableDeclarationsBuilder represents variableDeclarations builder
type VariableDeclarationsBuilder interface {
	Create() VariableDeclarationsBuilder
	WithDeclarations(declarations []VariableDeclaration) VariableDeclarationsBuilder
	Now() (VariableDeclarations, error)
}

// VariableDeclarations represents variable declarations
type VariableDeclarations interface {
	All() []VariableDeclaration
}

// VariableDeclarationBuilder represents a variable declaration builder
type VariableDeclarationBuilder interface {
	Create() VariableDeclarationBuilder
	WithType(typ Type) VariableDeclarationBuilder
	WithName(name string) VariableDeclarationBuilder
	Now() (VariableDeclaration, error)
}

// VariableDeclaration represents a variable declaration
type VariableDeclaration interface {
	Type() Type
	Name() string
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithType(typ Type) ValueBuilder
	WithElement(element interface{}) ValueBuilder
	WithUint(uintValue uint) ValueBuilder
	WithInt(intValue int) ValueBuilder
	WithFloat(floatValue float64) ValueBuilder
	WithBool(boolValue bool) ValueBuilder
	WithString(str string) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Type() Type
	Element() interface{}
}

// TypeBuilder represents a type builder
type TypeBuilder interface {
	Create() TypeBuilder
	WithDept(dept uint) TypeBuilder
	WithFlag(flag uint8) TypeBuilder
	Now() (Type, error)
}

// Type represents a type
type Type interface {
	Dept() uint
	Flag() uint8
	Compare(input Type) error
}
