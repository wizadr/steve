package parsers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/steve-care-software/steve/domain/languages/lexers/roots"
	"github.com/steve-care-software/steve/applications/languages/parsers"
)

func newEvents(bitsize uint) map[string]parsers.EventFunc {
	programBuilder := NewProgramBuilder()
	headerVariablesBuilder := NewHeaderVariablesBuilder()
	headerVariableBuilder := NewHeaderVariableBuilder()
	instructionsBuilder := NewInstructionsBuilder()
	instructionBuilder := NewInstructionBuilder()
	voidFuncCallBuilder := NewVoidFuncCallBuilder()
	directionBuilder := NewDirectionBuilder()
	forDeclarationBuilder := NewForDeclarationBuilder()
	forStatementBuilder := NewForStatementBuilder()
	ifDeclarationBuilder := NewIfDeclarationBuilder()
	assignmentBuilder := NewAssignmentBuilder()
	assigneeBuilder := NewAssigneeBuilder()
	assigneeNameFirstsBuilder := NewAssigneeNameFirstsBuilder()
	assigneeNameFirstBuilder := NewAssigneeNameFirstBuilder()
	assigneeNamesBuilder := NewAssigneeNamesBuilder()
	assigneeNameBuilder := NewAssigneeNameBuilder()
	assignablesBuilder := NewAssignablesBuilder()
	assignableBuilder := NewAssignableBuilder()
	singleAssignableBuilder := NewSingleAssignableBuilder()
	listFetchElementBuilder := NewListFetchElementBuilder()
	totalListValueBuilder := NewTotalListValueBuilder()
	totalListAssignableBuilder := NewTotalListAssignableBuilder()
	appendBuilder := NewAppendBuilder()
	listValueBuilder := NewListValueBuilder()
	sliceBuilder := NewSliceBuilder()
	sliceDelimiterBuilder := NewSliceDelimiterBuilder()
	indexValueBuilder := NewIndexValueBuilder()
	operationFuncCallBuilder := NewOperationFuncCallBuilder()
	operationSuffixBuilder := NewOperationSuffixBuilder()
	operationBuilder := NewOperationBuilder()
	operationElementBuilder := NewOperatonElementBuilder()
	intIncrementBuilder := NewIntIncrementBuilder()
	valueBuilder := NewValueBuilder()
	variableDeclarationsBuilder := NewVariableDeclarationsBuilder()
	variableDeclarationBuilder := NewVariableDeclarationBuilder()
	typeBuilder := NewTypeBuilder()
	return map[string]parsers.EventFunc{
		"program": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			instructionsResult, err := manager.Trigger("instructions", root)
			if err != nil {
				return nil, err
			}

			builder := programBuilder.Create().WithInstructions(instructionsResult.First().(Instructions))
			if nodes.Exists("headerVariables") {
				result, err := manager.Trigger("headerVariables", root)
				if err != nil {
					return nil, err
				}

				builder.WithVariables(result.First().(HeaderVariables))
			}

			return builder.Now()
		},
		"headerVariables": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			result, err := manager.Trigger("headerVariable", root)
			if err != nil {
				return nil, err
			}

			list := []HeaderVariable{}
			resultList := result.List()
			for _, oneHeaderVariable := range resultList {
				list = append(list, oneHeaderVariable.(HeaderVariable))
			}

			return headerVariablesBuilder.Create().WithVariables(list).Now()
		},
		"headerVariable": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			isInputResult, err := manager.Trigger("inputOutput", root)
			if err != nil {
				return nil, err
			}

			declarationResult, err := manager.Trigger("variableDeclaration", root)
			if err != nil {
				return nil, err
			}

			keynameResult, err := manager.Trigger("KEYNAME_PATTERN", root)
			if err != nil {
				return nil, err
			}

			builder := headerVariableBuilder.Create().WithDeclaration(declarationResult.First().(VariableDeclaration)).WithKeyname(keynameResult.First().(string))
			if isInputResult.First().(bool) {
				builder.IsInput()
			}

			if !root.Nodes().Exists("QUESTION_MARK") {
				builder.IsMandatory()
			}

			return builder.Now()
		},
		"inputOutput": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			if nodes.Exists("RIGHT_ARROW") {
				return true, nil
			}

			return false, nil
		},
		"instructions": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			result, err := manager.Trigger("instruction", root)
			if err != nil {
				return nil, err
			}

			list := []Instruction{}
			resultList := result.List()
			for _, oneInstruction := range resultList {
				list = append(list, oneInstruction.(Instruction))
			}

			return instructionsBuilder.Create().WithInstructions(list).Now()
		},
		"instruction": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := instructionBuilder.Create()
			if nodes.Exists("assignment") {
				result, err := manager.Trigger("assignment", root)
				if err != nil {
					return nil, err
				}

				builder.WithAssignment(result.First().(Assignment))
			}

			if nodes.Exists("variableDeclaration") {
				result, err := manager.Trigger("variableDeclaration", root)
				if err != nil {
					return nil, err
				}

				builder.WithVariableDeclaration(result.First().(VariableDeclaration))
			}

			if nodes.Exists("ifDeclaration") {
				result, err := manager.Trigger("ifDeclaration", root)
				if err != nil {
					return nil, err
				}

				builder.WithIfDeclaration(result.First().(IfDeclaration))
			}

			if nodes.Exists("forDeclaration") {
				result, err := manager.Trigger("forDeclaration", root)
				if err != nil {
					return nil, err
				}

				builder.WithForDeclaration(result.First().(ForDeclaration))
			}

			if nodes.Exists("voidFuncCalls") {
				result, err := manager.Trigger("voidFuncCalls", root)
				if err != nil {
					return nil, err
				}

				builder.WithFuncCall(result.First().(VoidFuncCall))
			}

			if nodes.Exists("direction") {
				result, err := manager.Trigger("direction", root)
				if err != nil {
					return nil, err
				}

				builder.WithDirection(result.First().(Direction))
			}

			if nodes.Exists("intIncrement") {
				result, err := manager.Trigger("intIncrement", root)
				if err != nil {
					return nil, err
				}

				builder.WithIntIncrement(result.First().(IntIncrement))
			}

			return builder.Now()
		},
		"voidFuncCalls": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := voidFuncCallBuilder.Create()
			if nodes.Exists("logFuncCall") {
				result, err := manager.Trigger("logFuncCall", root)
				if err != nil {
					return nil, err
				}

				builder.WithLog(result.First().(Assignable))
			}

			return builder.Now()
		},
		"logFuncCall": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("assignable", root)
		},
		"direction": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := directionBuilder.Create()
			if nodes.Exists("BREAK") {
				builder.IsBreak()
			}

			if nodes.Exists("CONTINUE") {
				builder.IsContinue()
			}

			return builder.Now()
		},
		"forDeclaration": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			statementResult, err := manager.Trigger("forStatement", root)
			if err != nil {
				return nil, err
			}

			builder := forDeclarationBuilder.Create().WithStatement(statementResult.First().(ForStatement))
			if root.Nodes().Exists("instructions") {
				result, err := manager.Trigger("instructions", root)
				if err != nil {
					return nil, err
				}

				builder.WithInstructions(result.First().(Instructions))
			}

			return builder.Now()
		},
		"forStatement": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			if nodes.Exists("forStatement") {
				return manager.Trigger("forStatement", root)
			}

			builder := forStatementBuilder.Create()
			if nodes.Exists("assigneeNames") {
				result, err := manager.Trigger("assigneeNames", root)
				if err != nil {
					return nil, err
				}

				list := result.First().(AssigneeNames).All()
				builder.WithIndex(list[0].(AssigneeName)).WithValue(list[1].(AssigneeName))
			}

			if nodes.Exists("totalListValue") {
				result, err := manager.Trigger("totalListValue", root)
				if err != nil {
					return nil, err
				}

				builder.WithIterable(result.First().(TotalListValue))
			}

			return builder.Now()
		},
		"ifDeclaration": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := ifDeclarationBuilder.Create()
			if nodes.Exists("ifOperation") {
				result, err := manager.Trigger("ifOperation", root)
				if err != nil {
					return nil, err
				}

				builder.WithCondition(result.First().(Operation))
			}

			if nodes.Exists("instructions") {
				result, err := manager.Trigger("instructions", root)
				if err != nil {
					return nil, err
				}

				builder.WithInstructions(result.First().(Instructions))
			}

			return builder.Now()
		},
		"ifOperation": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("operation", root)
		},
		"assignment": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := assignmentBuilder.Create()
			if nodes.Exists("assignee") {
				result, err := manager.Trigger("assignee", root)
				if err != nil {
					return nil, err
				}

				builder.WithAssignee(result.First().(Assignee))
			}

			if nodes.Exists("assignables") {
				result, err := manager.Trigger("assignables", root)
				if err != nil {
					return nil, err
				}

				builder.WithAssignables(result.First().(Assignables))
			}

			return builder.Now()
		},
		"assignee": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := assigneeBuilder.Create()
			if nodes.Exists("assigneeNameFirsts") {
				result, err := manager.Trigger("assigneeNameFirsts", root)
				if err != nil {
					return nil, err
				}

				builder.WithFirst(result.First().(AssigneeNameFirsts))
			}

			if nodes.Exists("assigneeNames") {
				result, err := manager.Trigger("assigneeNames", root)
				if err != nil {
					return nil, err
				}

				builder.WithName(result.First().(AssigneeNames))
			}

			return builder.Now()
		},
		"assigneeNameFirsts": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			builder := assigneeNameFirstsBuilder.Create()
			assigneeNameFirstResult, err := manager.Trigger("assigneeNameFirst", root)
			if err != nil {
				return nil, err
			}

			list := []AssigneeNameFirst{
				assigneeNameFirstResult.First().(AssigneeNameFirst),
			}

			if root.Nodes().Exists("commaAssigneeFirst") {
				result, err := manager.Trigger("commaAssigneeFirst", root)
				if err != nil {
					return nil, err
				}

				resultList := result.List()
				for _, oneAssigneeFirst := range resultList {
					list = append(list, oneAssigneeFirst.(AssigneeNameFirst))
				}
			}

			return builder.WithNames(list).Now()
		},
		"commaAssigneeFirst": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("assigneeNameFirst", root)
		},
		"assigneeNameFirst": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := assigneeNameFirstBuilder.Create()
			if nodes.Exists("variableDeclaration") {
				result, err := manager.Trigger("variableDeclaration", root)
				if err != nil {
					return nil, err
				}

				builder.WithDeclaration(result.First().(VariableDeclaration))
			}

			if nodes.Exists("assigneeName") {
				result, err := manager.Trigger("assigneeName", root)
				if err != nil {
					return nil, err
				}

				builder.WithAssigneeName(result.First().(AssigneeName))
			}

			return builder.Now()
		},
		"assigneeNames": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			builder := assigneeNamesBuilder.Create()
			assigneeNameResult, err := manager.Trigger("assigneeName", root)
			if err != nil {
				return nil, err
			}

			list := []AssigneeName{
				assigneeNameResult.First().(AssigneeName),
			}

			if root.Nodes().Exists("commaAssigneeName") {
				result, err := manager.Trigger("commaAssigneeName", root)
				if err != nil {
					return nil, err
				}

				resultList := result.List()
				for _, oneAssigneeName := range resultList {
					list = append(list, oneAssigneeName.(AssigneeName))
				}
			}

			return builder.WithNames(list).Now()
		},
		"commaAssigneeName": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("assigneeName", root)
		},
		"assigneeName": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := assigneeNameBuilder.Create()
			if nodes.Exists("variableName") {
				result, err := manager.Trigger("variableName", root)
				if err != nil {
					return nil, err
				}

				builder.WithVariable(result.First().(string))
			}

			if nodes.Exists("UNDERSCORE") {
				builder.IsSkip()
			}

			return builder.Now()
		},
		"slice": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			listResult, err := manager.Trigger("listValue", root)
			if err != nil {
				return nil, err
			}

			delimiterResult, err := manager.Trigger("sliceDelimiter", root)
			if err != nil {
				return nil, err
			}

			return sliceBuilder.Create().
				WithList(listResult.First().(ListValue)).
				WithDelimiter(delimiterResult.First().(SliceDelimiter)).
				Now()
		},
		"assignables": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			assignableResult, err := manager.Trigger("assignable", root)
			if err != nil {
				return nil, err
			}

			list := []Assignable{
				assignableResult.First().(Assignable),
			}

			if root.Nodes().Exists("commaAssignable") {
				results, err := manager.Trigger("commaAssignable", root)
				if err != nil {
					return nil, err
				}

				resultList := results.List()
				for _, oneAssignable := range resultList {
					list = append(list, oneAssignable.(Assignable))
				}
			}

			return assignablesBuilder.Create().WithAssignables(list).Now()
		},
		"commaAssignable": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("assignable", root)
		},
		"assignable": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := assignableBuilder.Create()
			if nodes.Exists("totalListAssignable") {
				result, err := manager.Trigger("totalListAssignable", root)
				if err != nil {
					return nil, err
				}

				builder.WithList(result.First().(TotalListAssignable))
			}

			if nodes.Exists("singleAssignable") {
				result, err := manager.Trigger("singleAssignable", root)
				if err != nil {
					return nil, err
				}

				builder.WithSingle(result.First().(SingleAssignable))
			}

			return builder.Now()
		},
		"singleAssignable": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := singleAssignableBuilder.Create()
			if nodes.Exists("declarableValue") {
				result, err := manager.Trigger("declarableValue", root)
				if err != nil {
					return nil, err
				}

				builder.WithDeclarableValue(result.First().(Value))
			}

			if nodes.Exists("operation") {
				result, err := manager.Trigger("operation", root)
				if err != nil {
					return nil, err
				}

				builder.WithOperation(result.First().(Operation))
			}

			if nodes.Exists("listAssignableToAssignables") {
				result, err := manager.Trigger("listAssignableToAssignables", root)
				if err != nil {
					return nil, err
				}

				builder.WithExtract(result.First().(TotalListAssignable))
			}

			if nodes.Exists("listFetchElement") {
				result, err := manager.Trigger("listFetchElement", root)
				if err != nil {
					return nil, err
				}

				builder.WithListFetchElement(result.First().(ListFetchElement))
			}

			return builder.Now()
		},
		"listFetchElement": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			listResult, err := manager.Trigger("totalListValue", root)
			if err != nil {
				return nil, err
			}

			indexValueResult, err := manager.Trigger("listElementFetcher", root)
			if err != nil {
				return nil, err
			}

			return listFetchElementBuilder.Create().
				WithList(listResult.First().(TotalListValue)).
				WithIndex(indexValueResult.First().(IndexValue)).
				Now()
		},
		"listElementFetcher": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("indexValue", root)
		},
		"listAssignableToAssignables": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("totalListAssignable", root)
		},
		"totalListValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := totalListValueBuilder.Create()
			if nodes.Exists("totalListAssignable") {
				result, err := manager.Trigger("totalListAssignable", root)
				if err != nil {
					return nil, err
				}

				builder.WithAssignable(result.First().(TotalListAssignable))
			}

			if nodes.Exists("variableName") {
				result, err := manager.Trigger("variableName", root)
				if err != nil {
					return nil, err
				}

				builder.WithVariableName(result.First().(string))
			}

			return builder.Now()
		},
		"totalListAssignable": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := totalListAssignableBuilder.Create()
			if nodes.Exists("listValue") {
				result, err := manager.Trigger("listValue", root)
				if err != nil {
					return nil, err
				}

				builder.WithListValue(result.First().(ListValue))
			}

			if nodes.Exists("slice") {
				result, err := manager.Trigger("slice", root)
				if err != nil {
					return nil, err
				}

				builder.WithSlice(result.First().(Slice))
			}

			if nodes.Exists("listAppend") {
				result, err := manager.Trigger("listAppend", root)
				if err != nil {
					return nil, err
				}

				builder.WithAppend(result.First().(Append))
			}

			return builder.Now()
		},
		"listAppend": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			listResult, err := manager.Trigger("listValue", root)
			if err != nil {
				return nil, err
			}

			assignableResult, err := manager.Trigger("singleAssignable", root)
			if err != nil {
				return nil, err
			}

			return appendBuilder.Create().
				WithList(listResult.First().(ListValue)).
				WithAssignable(assignableResult.First().(SingleAssignable)).
				Now()
		},
		"listValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := listValueBuilder.Create()
			if nodes.Exists("listAssignable") {
				result, err := manager.Trigger("listAssignable", root)
				if err != nil {
					return nil, err
				}

				builder.WithAssignables(result.First().(Assignables))
			}

			if nodes.Exists("variableName") {
				result, err := manager.Trigger("variableName", root)
				if err != nil {
					return nil, err
				}

				builder.WithVariableName(result.First().(string))
			}

			return builder.Now()
		},
		"listAssignable": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			if root.Nodes().Exists("assignables") {
				return manager.Trigger("assignables", root)
			}

			return assignablesBuilder.Create().Now()
		},
		"sliceDelimiter": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			builder := sliceDelimiterBuilder.Create()
			indexValuesResult, err := manager.Trigger("indexValue", root)
			if err != nil {
				return nil, err
			}

			indexValues := indexValuesResult.List()
			if len(indexValues) > 1 {
				builder.WithAmount(indexValues[1].(IndexValue))
			}

			return builder.WithIndex(indexValues[0].(IndexValue)).Now()
		},
		"indexValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := indexValueBuilder.Create()
			if nodes.Exists("INT_PATTERN") {
				intResult, err := manager.Trigger("INT_PATTERN", root)
				if err != nil {
					return nil, err
				}

				intValue, err := strconv.ParseUint(intResult.First().(string), 10, int(bitsize))
				if err != nil {
					return nil, err
				}

				value, err := valueBuilder.Create().WithUint(uint(intValue)).Now()
				if err != nil {
					return nil, err
				}

				builder.WithValue(value)
			}

			if nodes.Exists("variableName") {
				result, err := manager.Trigger("variableName", root)
				if err != nil {
					return nil, err
				}

				builder.WithVariableName(result.First().(string))
			}

			return builder.Now()
		},
		"operation": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			if nodes.Exists("operationWithSuffix") {
				return manager.Trigger("operationWithSuffix", root)
			}

			if nodes.Exists("operationFuncCalls") {
				fnCallResult, err := manager.Trigger("operationFuncCalls", root)
				if err != nil {
					return nil, err
				}

				fnCall := fnCallResult.First().(OperationFuncCall)
				element, err := operationElementBuilder.Create().WithFuncCall(fnCall).Now()
				if err != nil {
					return nil, err
				}

				return operationBuilder.Create().WithElement(element).Now()
			}

			return manager.Trigger("operation", root)
		},
		"operationFuncCalls": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := operationFuncCallBuilder.Create()
			if nodes.Exists("isNullFuncCall") {
				result, err := manager.Trigger("isNullFuncCall", root)
				if err != nil {
					return nil, err
				}

				builder.WithNull(result.First().(Operation))
			}

			return builder.Now()
		},
		"isNullFuncCall": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("operation", root)
		},
		"operationWithSuffix": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := operationBuilder.Create()
			if nodes.Exists("operationElement") {
				elResult, err := manager.Trigger("operationElement", root)
				if err != nil {
					return nil, err
				}

				builder.WithElement(elResult.First().(OperationElement))
			}

			if nodes.Exists("operationSuffix") {
				result, err := manager.Trigger("operationSuffix", root)
				if err != nil {
					return nil, err
				}

				builder.WithSuffix(result.First().(OperationSuffix))
			}

			return builder.Now()
		},
		"operationSuffix": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := operationSuffixBuilder.Create()
			if nodes.Exists("operator") {
				result, err := manager.Trigger("operator", root)
				if err != nil {
					return nil, err
				}

				builder.WithOperator(result.First().(uint16))
			}

			if nodes.Exists("operationElement") {
				result, err := manager.Trigger("operationElement", root)
				if err != nil {
					return nil, err
				}

				builder.WithElement(result.First().(OperationElement))
			}

			return builder.Now()
		},
		"operationElement": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := operationElementBuilder.Create()
			if nodes.Exists("boolValue") {
				result, err := manager.Trigger("boolValue", root)
				if err != nil {
					return nil, err
				}

				casted := result.First().(Value)
				flag := casted.Type().Flag()
				if flag&TypeBoolFlag == 0 {
					return nil, errors.New("the boolValue was expected to contain a boolean value")
				}

				builder.WithBoolValue(casted.Element().(bool))
			}

			if nodes.Exists("variableName") {
				varResult, err := manager.Trigger("variableName", root)
				if err != nil {
					return nil, err
				}

				builder.WithVariableName(varResult.First().(string))
			}

			if nodes.Exists("computableValue") {
				valueResult, err := manager.Trigger("computableValue", root)
				if err != nil {
					return nil, err
				}

				builder.WithComputableValue(valueResult.First().(Value))
			}

			if nodes.Exists("operation") {
				operationResult, err := manager.Trigger("operation", root)
				if err != nil {
					return nil, err
				}

				builder.WithOperation(operationResult.First().(Operation))
			}

			return builder.Now()
		},
		"operator": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			if nodes.Exists("operatorArithmetic") {
				return manager.Trigger("operatorArithmetic", root)
			}

			if nodes.Exists("operatorRelational") {
				return manager.Trigger("operatorRelational", root)
			}

			return manager.Trigger("operatorLogical", root)
		},
		"operatorArithmetic": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			flag := uint16(OpArithmetic)
			nodes := root.Nodes()
			if nodes.Exists("PLUS") {
				flag |= OpAriPlus
			}

			if nodes.Exists("MINUS") {
				flag |= OpAriMinus
			}

			if nodes.Exists("STAR") {
				flag |= OpAriMultiply
			}

			if nodes.Exists("DIV") {
				flag |= OpAriDivide
			}

			return flag, nil
		},
		"operatorRelational": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			flag := uint16(OpRelational)
			nodes := root.Nodes()
			if nodes.Exists("EQUAL") {
				flag |= OpRelEqual
			}

			if nodes.Exists("EXCLAMATION") {
				flag |= OpRelNot
			}

			if nodes.Exists("LESS") {
				flag |= OpRelLessThan
			}

			if nodes.Exists("MORE") {
				flag |= OpRelGreaterThan
			}

			return flag, nil
		},
		"operatorLogical": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			flag := uint16(OpLogical)
			nodes := root.Nodes()
			if nodes.Exists("AND") {
				flag |= OpLogAnd
			}

			if nodes.Exists("PIPE") {
				flag |= OpLogOr
			}

			return flag, nil
		},
		"intIncrement": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			variableResult, err := manager.Trigger("variableName", root)
			if err != nil {
				return nil, err
			}

			variableName := variableResult.First().(string)
			builder := intIncrementBuilder.Create().WithVariableName(variableName)
			incrResult, err := manager.Trigger("intIncrementValue", root)
			if err != nil {
				return nil, err
			}

			increment := incrResult.First().(Value)
			return builder.WithIncrement(increment).Now()
		},
		"intIncrementValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := valueBuilder.Create()
			if nodes.Exists("PLUS") {
				builder.WithInt(1)
			}

			if nodes.Exists("MINUS") {
				builder.WithInt(-1)
			}

			return builder.Now()
		},
		"declarableValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			if nodes.Exists("computableValue") {
				return manager.Trigger("computableValue", root)
			}

			return manager.Trigger("primitiveValue", root)
		},
		"computableValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			if nodes.Exists("floatValue") {
				return manager.Trigger("floatValue", root)
			}

			return manager.Trigger("intValue", root)
		},
		"primitiveValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			if nodes.Exists("boolValue") {
				return manager.Trigger("boolValue", root)
			}

			return manager.Trigger("stringValue", root)
		},
		"stringValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			result, err := manager.Trigger("EVERYTHING_EXCEPT_QUOTATION", root)
			if err != nil {
				return nil, err
			}

			str := result.First().(string)
			return valueBuilder.Create().WithString(str).Now()
		},
		"boolValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := valueBuilder.Create()
			if nodes.Exists("TRUE") {
				builder.WithBool(true)
			}

			if nodes.Exists("FALSE") {
				builder.WithBool(false)
			}

			return builder.Now()
		},
		"intValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			builder := valueBuilder.Create()
			result, err := manager.Trigger("INT_PATTERN", root)
			if err != nil {
				return nil, err
			}

			intValue := result.First().(string)
			if root.Nodes().Exists("MINUS") {
				str := fmt.Sprintf("-%s", intValue)
				intValue, err := strconv.ParseInt(str, 10, int(bitsize))
				if err != nil {
					return nil, err
				}

				return builder.WithInt(int(intValue)).Now()
			}

			uintValue, err := strconv.ParseUint(intValue, 10, int(bitsize))
			if err != nil {
				return nil, err
			}

			return builder.WithUint(uint(uintValue)).Now()
		},
		"floatValue": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			intResult, err := manager.Trigger("INT_PATTERN", root)
			if err != nil {
				return nil, err
			}

			intProperties := intResult.List()
			if len(intProperties) != 2 {
				str := fmt.Sprintf("%d INT_PATTERN were expected, %d returned", 2, len(intProperties))
				return nil, errors.New(str)
			}

			prefix := ""
			if root.Nodes().Exists("MINUS") {
				prefix = "-"
			}

			str := fmt.Sprintf("%s%s.%s", prefix, intProperties[0], intProperties[1])
			val, err := strconv.ParseFloat(str, int(bitsize))
			if err != nil {
				return nil, err
			}

			return valueBuilder.Create().WithFloat(val).Now()
		},
		"variableDeclarations": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			variableDeclaration, err := manager.Trigger("variableDeclaration", root)
			if err != nil {
				return nil, err
			}

			list := []VariableDeclaration{
				variableDeclaration.First().(VariableDeclaration),
			}

			if root.Nodes().Exists("commaVariableDeclaration") {
				result, err := manager.Trigger("commaVariableDeclaration", root)
				if err != nil {
					return nil, err
				}

				results := result.List()
				for _, oneVariableDeclaration := range results {
					list = append(list, oneVariableDeclaration.(VariableDeclaration))
				}
			}

			return variableDeclarationsBuilder.Create().WithDeclarations(list).Now()
		},
		"commaVariableDeclaration": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("variableDeclaration", root)
		},
		"variableDeclaration": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			declType, err := manager.Trigger("declarativeType", root)
			if err != nil {
				return nil, err
			}

			variable, err := manager.Trigger("variableName", root)
			if err != nil {
				return nil, err
			}

			return variableDeclarationBuilder.Create().WithType(declType.First().(Type)).WithName(variable.First().(string)).Now()
		},
		"listType": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			return manager.Trigger("declarativeType", root)
		},
		"declarativeType": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			builder := typeBuilder.Create()
			if nodes.Exists("primitiveType") {
				result, err := manager.Trigger("primitiveType", root)
				if err != nil {
					return nil, err
				}

				if casted, ok := result.First().(uint8); ok {
					builder.WithFlag(casted)
				}
			}

			if nodes.Exists("computableType") {
				result, err := manager.Trigger("computableType", root)
				if err != nil {
					return nil, err
				}

				if casted, ok := result.First().(uint8); ok {
					builder.WithFlag(casted)
				}
			}

			if nodes.Exists("listType") {
				result, err := manager.Trigger("listType", root)
				if err != nil {
					return nil, err
				}

				if casted, ok := result.First().(Type); ok {
					flag := casted.Flag()
					dept := casted.Dept() + 1
					builder.WithFlag(flag).WithDept(dept)
				}
			}

			return builder.Now()
		},
		"primitiveType": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			flag := uint8(0)
			nodes := root.Nodes()
			if nodes.Exists("BOOL") {
				flag |= TypeBoolFlag
			}

			if nodes.Exists("STRING") {
				flag |= TypeStringFlag
			}

			return flag, nil
		},
		"computableType": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			flag := uint8(0)
			nodes := root.Nodes()
			if nodes.Exists("FLOAT") {
				flag |= TypeFloatFlag
			}

			if nodes.Exists("UINT") {
				flag |= TypeIntegerFlag | TypeUnsignedFlag
			}

			if nodes.Exists("INT") {
				flag |= TypeIntegerFlag
			}

			return flag, nil
		},
		"variableName": func(root roots.Root, manager parsers.Manager) (interface{}, error) {
			nodes := root.Nodes()
			firstLetter, err := nodes.Fetch("FIRST_LETTER")
			if err != nil {
				return nil, err
			}

			remainingLetters, err := nodes.Fetch("MIN_MAJ_LETTERS")
			if err != nil {
				return nil, err
			}

			return strings.Join([]string{
				firstLetter.Content(),
				remainingLetters.Content(),
			}, ""), nil
		},
	}
}
