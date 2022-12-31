// This file uses camelCased parser grammar rule names so that the ANTLR- generated Go actually
// looks like Go.
parser grammar UnitParser;

options {
	tokenVocab = UnitLexer;
}

unit: (section)* EOF;

section: sectionHeader fields;

sectionHeader: LBRACKET sectionName RBRACKET EOL;

sectionName: IDENTIFIER;

fields: (field)*;

field: fieldName ASSIGN values (EOL+ | EOF);

fieldName: IDENTIFIER;

values: (plainValues | quotedValues | CONTINUATION)*;

plainValues: plainValue (plainValue)*;

plainValue: VALUE;

quotedValues: (
		(DQUOTE quotedValue DQUOTE)
		| (SQUOTE quotedValue SQUOTE)
	)+;

quotedValue: VALUE;
