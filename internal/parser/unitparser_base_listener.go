// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // UnitParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseUnitParserListener is a complete listener for a parse tree produced by UnitParser.
type BaseUnitParserListener struct{}

var _ UnitParserListener = &BaseUnitParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUnitParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUnitParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUnitParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUnitParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseUnitParserListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseUnitParserListener) ExitUnit(ctx *UnitContext) {}

// EnterSection is called when production section is entered.
func (s *BaseUnitParserListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BaseUnitParserListener) ExitSection(ctx *SectionContext) {}

// EnterSectionHeader is called when production sectionHeader is entered.
func (s *BaseUnitParserListener) EnterSectionHeader(ctx *SectionHeaderContext) {}

// ExitSectionHeader is called when production sectionHeader is exited.
func (s *BaseUnitParserListener) ExitSectionHeader(ctx *SectionHeaderContext) {}

// EnterSectionName is called when production sectionName is entered.
func (s *BaseUnitParserListener) EnterSectionName(ctx *SectionNameContext) {}

// ExitSectionName is called when production sectionName is exited.
func (s *BaseUnitParserListener) ExitSectionName(ctx *SectionNameContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseUnitParserListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseUnitParserListener) ExitFields(ctx *FieldsContext) {}

// EnterField is called when production field is entered.
func (s *BaseUnitParserListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseUnitParserListener) ExitField(ctx *FieldContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseUnitParserListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseUnitParserListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterValues is called when production values is entered.
func (s *BaseUnitParserListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseUnitParserListener) ExitValues(ctx *ValuesContext) {}

// EnterPlainValues is called when production plainValues is entered.
func (s *BaseUnitParserListener) EnterPlainValues(ctx *PlainValuesContext) {}

// ExitPlainValues is called when production plainValues is exited.
func (s *BaseUnitParserListener) ExitPlainValues(ctx *PlainValuesContext) {}

// EnterPlainValue is called when production plainValue is entered.
func (s *BaseUnitParserListener) EnterPlainValue(ctx *PlainValueContext) {}

// ExitPlainValue is called when production plainValue is exited.
func (s *BaseUnitParserListener) ExitPlainValue(ctx *PlainValueContext) {}

// EnterQuotedValues is called when production quotedValues is entered.
func (s *BaseUnitParserListener) EnterQuotedValues(ctx *QuotedValuesContext) {}

// ExitQuotedValues is called when production quotedValues is exited.
func (s *BaseUnitParserListener) ExitQuotedValues(ctx *QuotedValuesContext) {}

// EnterQuotedValue is called when production quotedValue is entered.
func (s *BaseUnitParserListener) EnterQuotedValue(ctx *QuotedValueContext) {}

// ExitQuotedValue is called when production quotedValue is exited.
func (s *BaseUnitParserListener) ExitQuotedValue(ctx *QuotedValueContext) {}
