package parsers

import (
	"testing"

	parser_application "github.com/steve-care-software/steve/applications/languages/parsers"
	"github.com/steve-care-software/steve/domain/languages/scripts"
)

func TestParser_program_Success(t *testing.T) {
	root := "program"
	script := `
		-> uint firstVariable:first_input
		-> list<int> myIntList:numbers_input
		<- uint secondVariable:second_input
		<- string last:last_keyname?

		string myVariable := "some value"
		uint myUintValue
		if (first < 56.87) {
			value := 56
		}

		for keyname, value := myIntList {
			if value <= 0 {
				continue
			}

			last = keyname
			if value > 5 {
				break
			}
		}

		secondVariable++
		log(secondVariable)
		break
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Program); ok {
		if casted == nil {
			t.Errorf("the program was expected to NOT be nil")
		}

		if !casted.HasVariables() {
			t.Errorf("the program was expecting header variables")
		}

		headerVariables := casted.Variables().All()
		if len(headerVariables) != 4 {
			t.Errorf("the program was expected %d headerVariables, %d returned", 4, len(headerVariables))
		}

		instructions := casted.Instructions().All()
		if len(instructions) != 7 {
			t.Errorf("the program was expected %d instructions, %d returned", 7, len(instructions))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_ifOperation_withoutInstruction_Success(t *testing.T) {
	root := "ifDeclaration"
	script := `
		if (myVariable < 45) {

		}
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(IfDeclaration); ok {
		if casted == nil {
			t.Errorf("the ifDeclaration was expected to NOT be nil")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assignment_withMultiple_Success(t *testing.T) {
	root := "assignment"
	script := `
		string str, myInt, _ := "my value", 34, 45
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Assignment); ok {
		if casted == nil {
			t.Errorf("the assignment was expected to NOT be nil")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assignment_withTwo_Success(t *testing.T) {
	root := "assignment"
	script := `
		str, myInt = "my value", 34
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Assignment); ok {
		if casted == nil {
			t.Errorf("the assignment was expected to NOT be nil")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assignment_single_Success(t *testing.T) {
	root := "assignment"
	script := `
		string myVariable := "my value"
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Assignment); ok {
		if casted == nil {
			t.Errorf("the assignment was expected to NOT be nil")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assigneeName_isSingle_Success(t *testing.T) {
	root := "assigneeNameFirsts"
	script := `
		myStr, int myInt, _
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(AssigneeNameFirsts); ok {
		list := casted.All()
		if len(list) != 3 {
			t.Errorf("%d elements were expected, %d returned", 3, len(list))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assigneeName_isAssigneeName_Success(t *testing.T) {
	root := "assigneeNameFirst"
	script := `
		myStr
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(AssigneeNameFirst); ok {
		if !casted.IsAssigneeName() {
			t.Errorf("the AssigneeNameFirst was expected to be an AssigneeName")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assigneeName_isVariableDeclaration_Success(t *testing.T) {
	root := "assigneeNameFirst"
	script := `
		string myStr
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(AssigneeNameFirst); ok {
		if !casted.IsDeclaration() {
			t.Errorf("the AssigneeNameFirst was expected to be a VariableDeclaration")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assigneeNames_withSingle_Success(t *testing.T) {
	root := "assigneeNames"
	script := `
		single
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(AssigneeNames); ok {
		list := casted.All()
		if len(list) != 1 {
			t.Errorf("%d elements were expected, %d returned", 1, len(list))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assigneeNames_withMultiple_Success(t *testing.T) {
	root := "assigneeNames"
	script := `
		_, second, _, fourth, _
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(AssigneeNames); ok {
		list := casted.All()
		if len(list) != 5 {
			t.Errorf("%d elements were expected, %d returned", 5, len(list))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assigneeName_isSkip_Success(t *testing.T) {
	root := "assigneeName"
	script := `
		_
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(AssigneeName); ok {
		if !casted.IsSkip() {
			t.Errorf("the assigneeName was expected to be skip")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assigneeName_Success(t *testing.T) {
	root := "assigneeName"
	script := `
		myVariable
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(AssigneeName); ok {
		if casted.Variable() != "myVariable" {
			t.Errorf("the variableName was expected to be '%s', '%s' returned", "myVariable", casted.Variable())
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_slice_Success(t *testing.T) {
	root := "slice"
	script := `
		[34,45,65,67][1:]
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Slice); ok {
		if casted == nil {
			t.Errorf("the Slice was expected to be valid")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_singleAssignable_withExtract_Success(t *testing.T) {
	root := "singleAssignable"
	script := `
		[first, 3, 4]...
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(SingleAssignable); ok {
		if !casted.IsExtract() {
			t.Errorf("the SingleAssignable was expected to contain an extract")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_singleAssignable_withOperation_Success(t *testing.T) {
	root := "singleAssignable"
	script := `
		(4566 <= myVariable)
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(SingleAssignable); ok {
		if !casted.IsOperation() {
			t.Errorf("the SingleAssignable was expected to contain an Operation")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_singleAssignable_withDeclarableValue_Success(t *testing.T) {
	root := "singleAssignable"
	script := `
		456
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(SingleAssignable); ok {
		if !casted.IsDeclarableValue() {
			t.Errorf("the SingleAssignable was expected to contain a declarable Value")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_singleAssignable_withListFetchElement_Success(t *testing.T) {
	root := "singleAssignable"
	script := `
		input[0]
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(SingleAssignable); ok {
		if !casted.IsListFetchElement() {
			t.Errorf("the singleAssignable was expected to contain a ListFetchElement instance")
		}
		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_listAssignableToAssignables_Success(t *testing.T) {
	root := "listAssignableToAssignables"
	script := `
		[
			[34, 45, 54, 34],
			[21, 23, 45, 67]
		]...
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(TotalListAssignable); ok {
		if !casted.IsListValue() {
			t.Errorf("the TotalListAssignable was expected to contain a ListValue")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_totalListValue_withVariableName_Success(t *testing.T) {
	root := "totalListValue"
	script := `
		myVariable
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(TotalListValue); ok {
		if !casted.IsVariableName() {
			t.Errorf("the TotalListValue was expected to contain a variableName")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_totalListValue_withTotalListAssignable_Success(t *testing.T) {
	root := "totalListValue"
	script := `
		append myList, 45
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(TotalListValue); ok {
		if !casted.IsAssignable() {
			t.Errorf("the TotalListValue was expected to contain a TotalListAssignable")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_totalListAssignable_withAppend_Success(t *testing.T) {
	root := "totalListAssignable"
	script := `
		append myList, 45
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(TotalListAssignable); ok {
		if !casted.IsAppend() {
			t.Errorf("the TotalListAssignable was expected to contain an Append")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_totalListAssignable_withSlice_Success(t *testing.T) {
	root := "totalListAssignable"
	script := `
		[1, 2, myVariable][1:]
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(TotalListAssignable); ok {
		if !casted.IsSlice() {
			t.Errorf("the TotalListAssignable was expected to contain a Slice")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_totalListAssignable_withListValue_Success(t *testing.T) {
	root := "totalListAssignable"
	script := `
		[1, 2, myVariable]
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(TotalListAssignable); ok {
		if !casted.IsListValue() {
			t.Errorf("the TotalListAssignable was expected to contain a ListValue")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_listValue_withAssignables_Success(t *testing.T) {
	root := "listValue"
	script := `
		[1, 2, myVariable]
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(ListValue); ok {
		if !casted.IsAssignables() {
			t.Errorf("the ListValue was expected to contain an Assignables")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_listValue_withVariableName_Success(t *testing.T) {
	root := "listValue"
	script := `
		myVariable
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(ListValue); ok {
		if !casted.IsVariableName() {
			t.Errorf("the ListValue was expected to contain a variableName")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_listAssignable_withMultiple_Success(t *testing.T) {
	root := "listAssignable"
	script := `
		[1, 2, 3]
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Assignables); ok {
		list := casted.All()
		if len(list) != 3 {
			t.Errorf("%d Assignable instances were expected, %d returned", 3, len(list))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_listAssignable_isEmpty_Success(t *testing.T) {
	root := "listAssignable"
	script := `
		[]
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Assignables); ok {
		list := casted.All()
		if len(list) != 0 {
			t.Errorf("%d Assignable instances were expected, %d returned", 0, len(list))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assignables_withMultiple_Success(t *testing.T) {
	root := "assignables"
	script := `
		34, 22, 21
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Assignables); ok {
		list := casted.All()
		if len(list) != 3 {
			t.Errorf("%d Assignable instances were expected, %d returned", 3, len(list))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_assignables_withSingleSuccess(t *testing.T) {
	root := "assignables"
	script := `
		34
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Assignables); ok {
		list := casted.All()
		if len(list) != 1 {
			t.Errorf("%d Assignable instances were expected, %d returned", 1, len(list))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_sliceDelimiter_withAmount_Success(t *testing.T) {
	root := "sliceDelimiter"
	script := `
		1:amount
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(SliceDelimiter); ok {
		index := casted.Index()
		if index.Value().Element().(uint) != 1 {
			t.Errorf("the index was expected to be %d, %d returned", 1, index.Value().Element().(uint))
		}

		if !casted.HasAmount() {
			t.Errorf("the sliceDelimiter was expecting an amount")
		}

		amount := casted.Amount()
		if amount.VariableName() != "amount" {
			t.Errorf("the amount variableName was expected to be '%s', '%s' returned", "amount", amount.VariableName())
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_sliceDelimiter_Success(t *testing.T) {
	root := "sliceDelimiter"
	script := `
		1:
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(SliceDelimiter); ok {
		index := casted.Index()
		if index.Value().Element().(uint) != 1 {
			t.Errorf("the index was expected to be %d, %d returned", 1, index.Value().Element().(uint))
		}

		if casted.HasAmount() {
			t.Errorf("the sliceDelimiter was NOT expecting an amount")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_indexValue_withVariableName_Success(t *testing.T) {
	root := "indexValue"
	script := `
		myVariale
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(IndexValue); ok {
		if !casted.IsVariableName() {
			t.Errorf("the indexValue was expecting a variableName instance")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_indexValue_withValue_Success(t *testing.T) {
	root := "indexValue"
	script := `
		56
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(IndexValue); ok {
		if !casted.IsValue() {
			t.Errorf("the indexValue was expecting a Value instance")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operation_withFuncCal_Success(t *testing.T) {
	root := "operation"
	script := `
		isNull(myVariable)
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Operation); ok {
		if !casted.Element().IsFuncCall() {
			t.Errorf("the operation's element was expecting a funcCall instance")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operationWithSuffix_withSuffix_Success(t *testing.T) {
	root := "operationWithSuffix"
	script := `
		first && (second < 34)
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Operation); ok {
		if !casted.HasSuffix() {
			t.Errorf("the operation was expecting a suffix")
		}
		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operationWithSuffix_Success(t *testing.T) {
	root := "operationWithSuffix"
	script := `
		true
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Operation); ok {
		if casted.HasSuffix() {
			t.Errorf("the operation was not expecting a suffix")
		}
		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operationSuffix_Success(t *testing.T) {
	root := "operationSuffix"
	script := `
		&& true
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(OperationSuffix); ok {
		operator := casted.Operator()
		if operator&OpLogAnd == 0 {
			t.Errorf("the operator was expected to contain an and flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operationElement_withComputableValue_Success(t *testing.T) {
	root := "operationElement"
	script := `
		34
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(OperationElement); ok {
		if !casted.IsComputableValue() {
			t.Errorf("the operationElement was expecting a computableValue instance")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operationElement_withBoolValue_Success(t *testing.T) {
	root := "operationElement"
	script := `
		true
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(OperationElement); ok {
		if !casted.IsBoolValue() {
			t.Errorf("the operationElement was expecting a boolValue instance")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operator_withLogical_Success(t *testing.T) {
	root := "operator"
	script := `
		||
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpLogical == 0 {
			t.Errorf("the operator was expected to contain the logical flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operator_withRelational_Success(t *testing.T) {
	root := "operator"
	script := `
		!=
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpRelational == 0 {
			t.Errorf("the operator was expected to contain the relational flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operator_withArithmetic_Success(t *testing.T) {
	root := "operator"
	script := `
		/
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpArithmetic == 0 {
			t.Errorf("the operator was expected to contain the arithmetic flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorArithmetic_isDivide_Success(t *testing.T) {
	root := "operatorArithmetic"
	script := `
		/
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpAriDivide == 0 {
			t.Errorf("the operator was expected to contain divide")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorArithmetic_isMultiply_Success(t *testing.T) {
	root := "operatorArithmetic"
	script := `
		*
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpAriMultiply == 0 {
			t.Errorf("the operator was expected to contain multiply")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorArithmetic_isMinus_Success(t *testing.T) {
	root := "operatorArithmetic"
	script := `
		-
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpAriMinus == 0 {
			t.Errorf("the operator was expected to contain minus")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorArithmetic_isPlus_Success(t *testing.T) {
	root := "operatorArithmetic"
	script := `
		+
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpAriPlus == 0 {
			t.Errorf("the operator was expected to contain plus")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorRelational_isEqual_Success(t *testing.T) {
	root := "operatorRelational"
	script := `
		==
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpRelational == 0 {
			t.Errorf("the operator was expected to contain relational")
		}

		if casted&OpRelNot != 0 {
			t.Errorf("the operator was expected to NOT contain not")
		}

		if casted&OpRelEqual == 0 {
			t.Errorf("the operator was expected to contain equal")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorRelational_isNotEqual_Success(t *testing.T) {
	root := "operatorRelational"
	script := `
		!=
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpRelational == 0 {
			t.Errorf("the operator was expected to contain relational")
		}

		if casted&OpRelNot == 0 {
			t.Errorf("the operator was expected to contain not")
		}

		if casted&OpRelEqual == 0 {
			t.Errorf("the operator was expected to contain equal")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorRelational_isGreaterThan_Success(t *testing.T) {
	root := "operatorRelational"
	script := `
		>
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpRelational == 0 {
			t.Errorf("the operator was expected to contain relational")
		}

		if casted&OpRelGreaterThan == 0 {
			t.Errorf("the operator was expected to contain greaterThan")
		}

		if casted&OpRelEqual != 0 {
			t.Errorf("the operator was expected to NOT be equal")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorRelational_isGreaterThan_orEqual_Success(t *testing.T) {
	root := "operatorRelational"
	script := `
		>=
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpRelational == 0 {
			t.Errorf("the operator was expected to contain relational")
		}

		if casted&OpRelGreaterThan == 0 {
			t.Errorf("the operator was expected to contain greaterThan")
		}

		if casted&OpRelEqual == 0 {
			t.Errorf("the operator was expected to contain equal")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorRelational_isLessThan_Success(t *testing.T) {
	root := "operatorRelational"
	script := `
		<
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpRelational == 0 {
			t.Errorf("the operator was expected to contain relational")
		}

		if casted&OpRelLessThan == 0 {
			t.Errorf("the operator was expected to contain lessThan")
		}

		if casted&OpRelEqual != 0 {
			t.Errorf("the operator was expected to NOT be equal")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorRelational_isLessThan_orEqual_Success(t *testing.T) {
	root := "operatorRelational"
	script := `
		<=
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpRelational == 0 {
			t.Errorf("the operator was expected to contain relational")
		}

		if casted&OpRelLessThan == 0 {
			t.Errorf("the operator was expected to contain lessThan")
		}

		if casted&OpRelEqual == 0 {
			t.Errorf("the operator was expected to contain equal")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorLogical_isOr_Success(t *testing.T) {
	root := "operatorLogical"
	script := `
		||
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpLogical == 0 {
			t.Errorf("the operator was expected to contain logical")
		}

		if casted&OpLogAnd != 0 {
			t.Errorf("the operator was expected to NOT be an and")
		}

		if casted&OpLogOr == 0 {
			t.Errorf("the operator was expected to contain an or")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_operatorLogical_isAnd_Success(t *testing.T) {
	root := "operatorLogical"
	script := `
		&&
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint16); ok {
		if casted&OpLogical == 0 {
			t.Errorf("the operator was expected to contain logical")
		}

		if casted&OpLogAnd == 0 {
			t.Errorf("the operator was expected to contain an and")
		}

		if casted&OpLogOr != 0 {
			t.Errorf("the operator was expected to NOT be an or")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_intIncrement_isMinus_Success(t *testing.T) {
	root := "intIncrement"
	script := `
		cpt--
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(IntIncrement); ok {
		name := casted.VariableName()
		if name != "cpt" {
			t.Errorf("the operation member was expected to contain the '%s' variable, '%s' returned", "cpt", name)
		}

		if casted.Increment().Element().(int) != -1 {
			t.Errorf("the suffix was expected to contain %d, %d returned", -1, casted.Increment().Element().(int))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_intIncrement_isPlus_Success(t *testing.T) {
	root := "intIncrement"
	script := `
		cpt++
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(IntIncrement); ok {
		name := casted.VariableName()
		if name != "cpt" {
			t.Errorf("the operation member was expected to contain the '%s' variable, '%s' returned", "cpt", name)
		}

		if casted.Increment().Element().(int) != 1 {
			t.Errorf("the suffix was expected to contain %d, %d returned", 1, casted.Increment().Element().(int))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_declarableValue_withPrimitiveValue_Success(t *testing.T) {
	root := "declarableValue"
	script := `
		"works"
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		str := casted.Element().(string)
		if str != "works" {
			t.Errorf("the value was expected to be '%s', '%s' returned", "works", str)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_declarableValue_withComputableValue_Success(t *testing.T) {
	root := "declarableValue"
	script := `
		-56
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(int)
		if val != -56 {
			t.Errorf("the value was expected to be %d, %d returned", -56, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_computableValue_withFloat_Success(t *testing.T) {
	root := "computableValue"
	script := `
		3456.98
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(float64)
		if val != 3456.98 {
			t.Errorf("the value was expected to be %f, %f returned", 3456.98, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_computableValue_withUint_Success(t *testing.T) {
	root := "computableValue"
	script := `
		3456
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(uint)
		if val != 3456 {
			t.Errorf("the value was expected to be %d, %d returned", 3456, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_primitiveValue_withBool_Success(t *testing.T) {
	root := "primitiveValue"
	script := `
		false
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(bool)
		if val {
			t.Errorf("the value was expected to be %t, %t returned", false, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_primitiveValue_withString_Success(t *testing.T) {
	root := "primitiveValue"
	script := `
		" -> this is some value "
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(string)
		if val != " -> this is some value " {
			t.Errorf("the value was expected to be '%s', '%s' returned", " -> this is some value ", val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_stringValue_Success(t *testing.T) {
	root := "stringValue"
	script := `
		" -> this is some value "
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(string)
		if val != " -> this is some value " {
			t.Errorf("the value was expected to be '%s', '%s' returned", " -> this is some value ", val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_boolValue_isFalse_Success(t *testing.T) {
	root := "boolValue"
	script := `
		false
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(bool)
		if val {
			t.Errorf("the returned value was expected to be %t, %t returned", false, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_boolValue_isTrue_Success(t *testing.T) {
	root := "boolValue"
	script := `
		true
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(bool)
		if !val {
			t.Errorf("the returned value was expected to be %t, %t returned", true, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_intValue_isPositive_Success(t *testing.T) {
	root := "intValue"
	script := `
		2376
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(uint)
		if val != 2376 {
			t.Errorf("the value was expected to be %d, %d returned", 2376, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_intValue_isNegative_Success(t *testing.T) {
	root := "intValue"
	script := `
		-56
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(int)
		if val != -56 {
			t.Errorf("the value was expected to be %d, %d returned", -56, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_floatValue_isPositive_Success(t *testing.T) {
	root := "floatValue"
	script := `
		42322.00
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(float64)
		if val != 42322.00 {
			t.Errorf("the value was expected to be %f, %f returned", 42322.00, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_floatValue_isNegative_Success(t *testing.T) {
	root := "floatValue"
	script := `
		-56.98
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Value); ok {
		val := casted.Element().(float64)
		if val != -56.98 {
			t.Errorf("the value was expected to be %f, %f returned", -56.98, val)
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_variableDeclarations_withMultiple_Success(t *testing.T) {
	root := "variableDeclarations"
	script := `
		string myVariable,
		list<int> myInt,
		uint cpt
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(VariableDeclarations); ok {
		list := casted.All()
		if len(list) != 3 {
			t.Errorf("%d elements were expected, %d returned", 3, len(list))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_variableDeclarations_withSingle_Success(t *testing.T) {
	root := "variableDeclarations"
	script := `
		string myVariable
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(VariableDeclarations); ok {
		list := casted.All()
		if len(list) != 1 {
			t.Errorf("%d elements were expected, %d returned", 1, len(list))
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_variableDeclaration_Success(t *testing.T) {
	root := "variableDeclaration"
	script := `
		list<uint> myList
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(VariableDeclaration); ok {
		if casted.Name() != "myList" {
			t.Errorf("the variableName was expected to be '%s', '%s' returned", "myList", casted.Name())
		}

		typ := casted.Type()
		if typ.Dept() != 1 {
			t.Errorf("the dept was expected to be %d, %d returned", 1, typ.Dept())
		}

		flag := typ.Flag()
		if flag&TypeIntegerFlag == 0 {
			t.Errorf("the flag was expected to contain an int flag")
		}

		if flag&TypeUnsignedFlag == 0 {
			t.Errorf("the flag was expected to contain an unsigned flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_declarativeType_withList_withDeptFive_Success(t *testing.T) {
	root := "declarativeType"
	script := `
		list<list<list<list<list<uint>>>>>
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Type); ok {
		dept := casted.Dept()
		if dept != 5 {
			t.Errorf("the dept was expected to be %d, %d returned", 5, dept)
		}

		flag := casted.Flag()
		if flag&TypeIntegerFlag == 0 {
			t.Errorf("the flag was expected to contain an int flag")
		}

		if flag&TypeUnsignedFlag == 0 {
			t.Errorf("the flag was expected to contain an unsigned flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_declarativeType_withList_withDeptOne_Success(t *testing.T) {
	root := "declarativeType"
	script := `
		list<uint>
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Type); ok {
		dept := casted.Dept()
		if dept != 1 {
			t.Errorf("the dept was expected to be %d, %d returned", 1, dept)
		}

		flag := casted.Flag()
		if flag&TypeIntegerFlag == 0 {
			t.Errorf("the flag was expected to contain an int flag")
		}

		if flag&TypeUnsignedFlag == 0 {
			t.Errorf("the flag was expected to contain an unsigned flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_declarativeType_withComputableType_Success(t *testing.T) {
	root := "declarativeType"
	script := `
		uint
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Type); ok {
		dept := casted.Dept()
		if dept != 0 {
			t.Errorf("the dept was expected to be %d, %d returned", 0, dept)
		}

		flag := casted.Flag()
		if flag&TypeIntegerFlag == 0 {
			t.Errorf("the flag was expected to contain an int flag")
		}

		if flag&TypeUnsignedFlag == 0 {
			t.Errorf("the flag was expected to contain an unsigned flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_declarativeType_withPrimitiveType_Success(t *testing.T) {
	root := "declarativeType"
	script := `
		bool
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(Type); ok {
		dept := casted.Dept()
		if dept != 0 {
			t.Errorf("the dept was expected to be %d, %d returned", 0, dept)
		}

		flag := casted.Flag()
		if flag&TypeStringFlag != 0 {
			t.Errorf("the flag was expected to NOT contain a string flag")
		}

		if flag&TypeBoolFlag == 0 {
			t.Errorf("the flag was expected to contain a bool flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_primitiveType_withBool_Success(t *testing.T) {
	root := "primitiveType"
	script := `
		bool
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint8); ok {
		if casted&TypeStringFlag != 0 {
			t.Errorf("the flag was expected to NOT contain a string flag")
		}

		if casted&TypeBoolFlag == 0 {
			t.Errorf("the flag was expected to contain a bool flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_primitiveType_withString_Success(t *testing.T) {
	root := "primitiveType"
	script := `
		string
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint8); ok {
		if casted&TypeStringFlag == 0 {
			t.Errorf("the flag was expected to contain a string flag")
		}

		if casted&TypeBoolFlag != 0 {
			t.Errorf("the flag was expected to NOT contain a bool flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_computableType_withFloat_Success(t *testing.T) {
	root := "computableType"
	script := `
		float
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint8); ok {
		if casted&TypeFloatFlag == 0 {
			t.Errorf("the flag was expected to contain a float flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_computableType_withInt_Success(t *testing.T) {
	root := "computableType"
	script := `
		int
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint8); ok {
		if casted&TypeIntegerFlag == 0 {
			t.Errorf("the flag was expected to contain an int flag")
		}

		if casted&TypeUnsignedFlag != 0 {
			t.Errorf("the flag was expected to NOT contain an int flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_computableType_withUint_Success(t *testing.T) {
	root := "computableType"
	script := `
		uint
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if casted, ok := retIns.(uint8); ok {
		if casted&TypeIntegerFlag == 0 {
			t.Errorf("the flag was expected to contain an int flag")
		}

		if casted&TypeUnsignedFlag == 0 {
			t.Errorf("the flag was expected to contain an int flag")
		}

		return
	}

	t.Errorf("the returned type could not be casted properly")
}

func TestParser_variableName_Success(t *testing.T) {
	root := "variableName"
	script := `
		myVariable
	`

	baseDir := "./../../resources"
	langDir := "base"
	instance, mistake := scripts.ExecuteLexerForTests(root, script, baseDir, langDir)
	if mistake != nil {
		t.Errorf("the mistake was expected to be nil")
		return
	}

	events := NewEvents(64)
	parser := parser_application.NewApplication()
	retIns, err := parser.Execute(instance, events)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retIns.(string) != "myVariable" {
		t.Errorf("the returned value was expected to contain '%s', '%s' returned", "myVariable", retIns.(string))
		return
	}
}
