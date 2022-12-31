// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // UnitParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// UnitParserListener is a complete listener for a parse tree produced by UnitParser.
type UnitParserListener interface {
	antlr.ParseTreeListener

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterSectionHeader is called when entering the sectionHeader production.
	EnterSectionHeader(c *SectionHeaderContext)

	// EnterSectionName is called when entering the sectionName production.
	EnterSectionName(c *SectionNameContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// EnterPlainValues is called when entering the plainValues production.
	EnterPlainValues(c *PlainValuesContext)

	// EnterPlainValue is called when entering the plainValue production.
	EnterPlainValue(c *PlainValueContext)

	// EnterQuotedValues is called when entering the quotedValues production.
	EnterQuotedValues(c *QuotedValuesContext)

	// EnterQuotedValue is called when entering the quotedValue production.
	EnterQuotedValue(c *QuotedValueContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitSectionHeader is called when exiting the sectionHeader production.
	ExitSectionHeader(c *SectionHeaderContext)

	// ExitSectionName is called when exiting the sectionName production.
	ExitSectionName(c *SectionNameContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)

	// ExitPlainValues is called when exiting the plainValues production.
	ExitPlainValues(c *PlainValuesContext)

	// ExitPlainValue is called when exiting the plainValue production.
	ExitPlainValue(c *PlainValueContext)

	// ExitQuotedValues is called when exiting the quotedValues production.
	ExitQuotedValues(c *QuotedValuesContext)

	// ExitQuotedValue is called when exiting the quotedValue production.
	ExitQuotedValue(c *QuotedValueContext)
}
