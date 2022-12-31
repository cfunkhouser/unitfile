// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // UnitParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type UnitParser struct {
	*antlr.BaseParser
}

var unitparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func unitparserParserInit() {
	staticData := &unitparserParserStaticData
	staticData.literalNames = []string{
		"", "'['", "", "'='", "", "", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "LBRACKET", "RBRACKET", "ASSIGN", "DQUOTE", "SQUOTE", "EOL", "ESCAPE",
		"WHITESPACE", "IDENTIFIER", "COMMENT", "CONTINUATION", "WHITESPACE_IN_ASSIGN",
		"VALUE",
	}
	staticData.ruleNames = []string{
		"unit", "section", "sectionHeader", "sectionName", "fields", "field",
		"fieldName", "values", "plainValues", "plainValue", "quotedValues",
		"quotedValue",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 93, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 5,
		4, 44, 8, 4, 10, 4, 12, 4, 47, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 53,
		8, 5, 11, 5, 12, 5, 54, 1, 5, 3, 5, 58, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 5, 7, 65, 8, 7, 10, 7, 12, 7, 68, 9, 7, 1, 8, 1, 8, 5, 8, 72, 8, 8,
		10, 8, 12, 8, 75, 9, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 4, 10, 87, 8, 10, 11, 10, 12, 10, 88, 1, 11, 1, 11,
		1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 0, 90, 0,
		27, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 35, 1, 0, 0, 0, 6, 40, 1, 0, 0, 0,
		8, 45, 1, 0, 0, 0, 10, 48, 1, 0, 0, 0, 12, 59, 1, 0, 0, 0, 14, 66, 1, 0,
		0, 0, 16, 69, 1, 0, 0, 0, 18, 76, 1, 0, 0, 0, 20, 86, 1, 0, 0, 0, 22, 90,
		1, 0, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0,
		27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 27, 1,
		0, 0, 0, 30, 31, 5, 0, 0, 1, 31, 1, 1, 0, 0, 0, 32, 33, 3, 4, 2, 0, 33,
		34, 3, 8, 4, 0, 34, 3, 1, 0, 0, 0, 35, 36, 5, 1, 0, 0, 36, 37, 3, 6, 3,
		0, 37, 38, 5, 2, 0, 0, 38, 39, 5, 6, 0, 0, 39, 5, 1, 0, 0, 0, 40, 41, 5,
		9, 0, 0, 41, 7, 1, 0, 0, 0, 42, 44, 3, 10, 5, 0, 43, 42, 1, 0, 0, 0, 44,
		47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 9, 1, 0, 0,
		0, 47, 45, 1, 0, 0, 0, 48, 49, 3, 12, 6, 0, 49, 50, 5, 3, 0, 0, 50, 57,
		3, 14, 7, 0, 51, 53, 5, 6, 0, 0, 52, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 58, 5,
		0, 0, 1, 57, 52, 1, 0, 0, 0, 57, 56, 1, 0, 0, 0, 58, 11, 1, 0, 0, 0, 59,
		60, 5, 9, 0, 0, 60, 13, 1, 0, 0, 0, 61, 65, 3, 16, 8, 0, 62, 65, 3, 20,
		10, 0, 63, 65, 5, 11, 0, 0, 64, 61, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64,
		63, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0,
		0, 67, 15, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 73, 3, 18, 9, 0, 70, 72,
		3, 18, 9, 0, 71, 70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0,
		73, 74, 1, 0, 0, 0, 74, 17, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5,
		13, 0, 0, 77, 19, 1, 0, 0, 0, 78, 79, 5, 4, 0, 0, 79, 80, 3, 22, 11, 0,
		80, 81, 5, 4, 0, 0, 81, 87, 1, 0, 0, 0, 82, 83, 5, 5, 0, 0, 83, 84, 3,
		22, 11, 0, 84, 85, 5, 5, 0, 0, 85, 87, 1, 0, 0, 0, 86, 78, 1, 0, 0, 0,
		86, 82, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1,
		0, 0, 0, 89, 21, 1, 0, 0, 0, 90, 91, 5, 13, 0, 0, 91, 23, 1, 0, 0, 0, 9,
		27, 45, 54, 57, 64, 66, 73, 86, 88,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// UnitParserInit initializes any static state used to implement UnitParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewUnitParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func UnitParserInit() {
	staticData := &unitparserParserStaticData
	staticData.once.Do(unitparserParserInit)
}

// NewUnitParser produces a new parser instance for the optional input antlr.TokenStream.
func NewUnitParser(input antlr.TokenStream) *UnitParser {
	UnitParserInit()
	this := new(UnitParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &unitparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// UnitParser tokens.
const (
	UnitParserEOF                  = antlr.TokenEOF
	UnitParserLBRACKET             = 1
	UnitParserRBRACKET             = 2
	UnitParserASSIGN               = 3
	UnitParserDQUOTE               = 4
	UnitParserSQUOTE               = 5
	UnitParserEOL                  = 6
	UnitParserESCAPE               = 7
	UnitParserWHITESPACE           = 8
	UnitParserIDENTIFIER           = 9
	UnitParserCOMMENT              = 10
	UnitParserCONTINUATION         = 11
	UnitParserWHITESPACE_IN_ASSIGN = 12
	UnitParserVALUE                = 13
)

// UnitParser rules.
const (
	UnitParserRULE_unit          = 0
	UnitParserRULE_section       = 1
	UnitParserRULE_sectionHeader = 2
	UnitParserRULE_sectionName   = 3
	UnitParserRULE_fields        = 4
	UnitParserRULE_field         = 5
	UnitParserRULE_fieldName     = 6
	UnitParserRULE_values        = 7
	UnitParserRULE_plainValues   = 8
	UnitParserRULE_plainValue    = 9
	UnitParserRULE_quotedValues  = 10
	UnitParserRULE_quotedValue   = 11
)

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(UnitParserEOF, 0)
}

func (s *UnitContext) AllSection() []ISectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISectionContext); ok {
			len++
		}
	}

	tst := make([]ISectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISectionContext); ok {
			tst[i] = t.(ISectionContext)
			i++
		}
	}

	return tst
}

func (s *UnitContext) Section(i int) ISectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (p *UnitParser) Unit() (localctx IUnitContext) {
	this := p
	_ = this

	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UnitParserRULE_unit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UnitParserLBRACKET {
		{
			p.SetState(24)
			p.Section()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
		p.Match(UnitParserEOF)
	}

	return localctx
}

// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) SectionHeader() ISectionHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectionHeaderContext)
}

func (s *SectionContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterSection(s)
	}
}

func (s *SectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitSection(s)
	}
}

func (p *UnitParser) Section() (localctx ISectionContext) {
	this := p
	_ = this

	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, UnitParserRULE_section)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.SectionHeader()
	}
	{
		p.SetState(33)
		p.Fields()
	}

	return localctx
}

// ISectionHeaderContext is an interface to support dynamic dispatch.
type ISectionHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionHeaderContext differentiates from other interfaces.
	IsSectionHeaderContext()
}

type SectionHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionHeaderContext() *SectionHeaderContext {
	var p = new(SectionHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_sectionHeader
	return p
}

func (*SectionHeaderContext) IsSectionHeaderContext() {}

func NewSectionHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionHeaderContext {
	var p = new(SectionHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_sectionHeader

	return p
}

func (s *SectionHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionHeaderContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(UnitParserLBRACKET, 0)
}

func (s *SectionHeaderContext) SectionName() ISectionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISectionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISectionNameContext)
}

func (s *SectionHeaderContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(UnitParserRBRACKET, 0)
}

func (s *SectionHeaderContext) EOL() antlr.TerminalNode {
	return s.GetToken(UnitParserEOL, 0)
}

func (s *SectionHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterSectionHeader(s)
	}
}

func (s *SectionHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitSectionHeader(s)
	}
}

func (p *UnitParser) SectionHeader() (localctx ISectionHeaderContext) {
	this := p
	_ = this

	localctx = NewSectionHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, UnitParserRULE_sectionHeader)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(UnitParserLBRACKET)
	}
	{
		p.SetState(36)
		p.SectionName()
	}
	{
		p.SetState(37)
		p.Match(UnitParserRBRACKET)
	}
	{
		p.SetState(38)
		p.Match(UnitParserEOL)
	}

	return localctx
}

// ISectionNameContext is an interface to support dynamic dispatch.
type ISectionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionNameContext differentiates from other interfaces.
	IsSectionNameContext()
}

type SectionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionNameContext() *SectionNameContext {
	var p = new(SectionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_sectionName
	return p
}

func (*SectionNameContext) IsSectionNameContext() {}

func NewSectionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionNameContext {
	var p = new(SectionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_sectionName

	return p
}

func (s *SectionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(UnitParserIDENTIFIER, 0)
}

func (s *SectionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterSectionName(s)
	}
}

func (s *SectionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitSectionName(s)
	}
}

func (p *UnitParser) SectionName() (localctx ISectionNameContext) {
	this := p
	_ = this

	localctx = NewSectionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, UnitParserRULE_sectionName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(UnitParserIDENTIFIER)
	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *UnitParser) Fields() (localctx IFieldsContext) {
	this := p
	_ = this

	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, UnitParserRULE_fields)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UnitParserIDENTIFIER {
		{
			p.SetState(42)
			p.Field()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(UnitParserASSIGN, 0)
}

func (s *FieldContext) Values() IValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *FieldContext) EOF() antlr.TerminalNode {
	return s.GetToken(UnitParserEOF, 0)
}

func (s *FieldContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(UnitParserEOL)
}

func (s *FieldContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(UnitParserEOL, i)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *UnitParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, UnitParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.FieldName()
	}
	{
		p.SetState(49)
		p.Match(UnitParserASSIGN)
	}
	{
		p.SetState(50)
		p.Values()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UnitParserEOL:
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == UnitParserEOL {
			{
				p.SetState(51)
				p.Match(UnitParserEOL)
			}

			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case UnitParserEOF:
		{
			p.SetState(56)
			p.Match(UnitParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(UnitParserIDENTIFIER, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *UnitParser) FieldName() (localctx IFieldNameContext) {
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, UnitParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(UnitParserIDENTIFIER)
	}

	return localctx
}

// IValuesContext is an interface to support dynamic dispatch.
type IValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValuesContext differentiates from other interfaces.
	IsValuesContext()
}

type ValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesContext() *ValuesContext {
	var p = new(ValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_values
	return p
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_values

	return p
}

func (s *ValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesContext) AllPlainValues() []IPlainValuesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPlainValuesContext); ok {
			len++
		}
	}

	tst := make([]IPlainValuesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPlainValuesContext); ok {
			tst[i] = t.(IPlainValuesContext)
			i++
		}
	}

	return tst
}

func (s *ValuesContext) PlainValues(i int) IPlainValuesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlainValuesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPlainValuesContext)
}

func (s *ValuesContext) AllQuotedValues() []IQuotedValuesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQuotedValuesContext); ok {
			len++
		}
	}

	tst := make([]IQuotedValuesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQuotedValuesContext); ok {
			tst[i] = t.(IQuotedValuesContext)
			i++
		}
	}

	return tst
}

func (s *ValuesContext) QuotedValues(i int) IQuotedValuesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedValuesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedValuesContext)
}

func (s *ValuesContext) AllCONTINUATION() []antlr.TerminalNode {
	return s.GetTokens(UnitParserCONTINUATION)
}

func (s *ValuesContext) CONTINUATION(i int) antlr.TerminalNode {
	return s.GetToken(UnitParserCONTINUATION, i)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *UnitParser) Values() (localctx IValuesContext) {
	this := p
	_ = this

	localctx = NewValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, UnitParserRULE_values)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&10288) != 0 {
		p.SetState(64)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case UnitParserVALUE:
			{
				p.SetState(61)
				p.PlainValues()
			}

		case UnitParserDQUOTE, UnitParserSQUOTE:
			{
				p.SetState(62)
				p.QuotedValues()
			}

		case UnitParserCONTINUATION:
			{
				p.SetState(63)
				p.Match(UnitParserCONTINUATION)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPlainValuesContext is an interface to support dynamic dispatch.
type IPlainValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlainValuesContext differentiates from other interfaces.
	IsPlainValuesContext()
}

type PlainValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlainValuesContext() *PlainValuesContext {
	var p = new(PlainValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_plainValues
	return p
}

func (*PlainValuesContext) IsPlainValuesContext() {}

func NewPlainValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlainValuesContext {
	var p = new(PlainValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_plainValues

	return p
}

func (s *PlainValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *PlainValuesContext) AllPlainValue() []IPlainValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPlainValueContext); ok {
			len++
		}
	}

	tst := make([]IPlainValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPlainValueContext); ok {
			tst[i] = t.(IPlainValueContext)
			i++
		}
	}

	return tst
}

func (s *PlainValuesContext) PlainValue(i int) IPlainValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlainValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPlainValueContext)
}

func (s *PlainValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlainValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterPlainValues(s)
	}
}

func (s *PlainValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitPlainValues(s)
	}
}

func (p *UnitParser) PlainValues() (localctx IPlainValuesContext) {
	this := p
	_ = this

	localctx = NewPlainValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, UnitParserRULE_plainValues)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.PlainValue()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(70)
				p.PlainValue()
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IPlainValueContext is an interface to support dynamic dispatch.
type IPlainValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlainValueContext differentiates from other interfaces.
	IsPlainValueContext()
}

type PlainValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlainValueContext() *PlainValueContext {
	var p = new(PlainValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_plainValue
	return p
}

func (*PlainValueContext) IsPlainValueContext() {}

func NewPlainValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlainValueContext {
	var p = new(PlainValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_plainValue

	return p
}

func (s *PlainValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PlainValueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(UnitParserVALUE, 0)
}

func (s *PlainValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlainValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterPlainValue(s)
	}
}

func (s *PlainValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitPlainValue(s)
	}
}

func (p *UnitParser) PlainValue() (localctx IPlainValueContext) {
	this := p
	_ = this

	localctx = NewPlainValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, UnitParserRULE_plainValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(UnitParserVALUE)
	}

	return localctx
}

// IQuotedValuesContext is an interface to support dynamic dispatch.
type IQuotedValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedValuesContext differentiates from other interfaces.
	IsQuotedValuesContext()
}

type QuotedValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedValuesContext() *QuotedValuesContext {
	var p = new(QuotedValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_quotedValues
	return p
}

func (*QuotedValuesContext) IsQuotedValuesContext() {}

func NewQuotedValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedValuesContext {
	var p = new(QuotedValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_quotedValues

	return p
}

func (s *QuotedValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedValuesContext) AllDQUOTE() []antlr.TerminalNode {
	return s.GetTokens(UnitParserDQUOTE)
}

func (s *QuotedValuesContext) DQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(UnitParserDQUOTE, i)
}

func (s *QuotedValuesContext) AllQuotedValue() []IQuotedValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQuotedValueContext); ok {
			len++
		}
	}

	tst := make([]IQuotedValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQuotedValueContext); ok {
			tst[i] = t.(IQuotedValueContext)
			i++
		}
	}

	return tst
}

func (s *QuotedValuesContext) QuotedValue(i int) IQuotedValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedValueContext)
}

func (s *QuotedValuesContext) AllSQUOTE() []antlr.TerminalNode {
	return s.GetTokens(UnitParserSQUOTE)
}

func (s *QuotedValuesContext) SQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(UnitParserSQUOTE, i)
}

func (s *QuotedValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterQuotedValues(s)
	}
}

func (s *QuotedValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitQuotedValues(s)
	}
}

func (p *UnitParser) QuotedValues() (localctx IQuotedValuesContext) {
	this := p
	_ = this

	localctx = NewQuotedValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, UnitParserRULE_quotedValues)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(86)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case UnitParserDQUOTE:
				{
					p.SetState(78)
					p.Match(UnitParserDQUOTE)
				}
				{
					p.SetState(79)
					p.QuotedValue()
				}
				{
					p.SetState(80)
					p.Match(UnitParserDQUOTE)
				}

			case UnitParserSQUOTE:
				{
					p.SetState(82)
					p.Match(UnitParserSQUOTE)
				}
				{
					p.SetState(83)
					p.QuotedValue()
				}
				{
					p.SetState(84)
					p.Match(UnitParserSQUOTE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IQuotedValueContext is an interface to support dynamic dispatch.
type IQuotedValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedValueContext differentiates from other interfaces.
	IsQuotedValueContext()
}

type QuotedValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedValueContext() *QuotedValueContext {
	var p = new(QuotedValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnitParserRULE_quotedValue
	return p
}

func (*QuotedValueContext) IsQuotedValueContext() {}

func NewQuotedValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedValueContext {
	var p = new(QuotedValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnitParserRULE_quotedValue

	return p
}

func (s *QuotedValueContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedValueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(UnitParserVALUE, 0)
}

func (s *QuotedValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.EnterQuotedValue(s)
	}
}

func (s *QuotedValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnitParserListener); ok {
		listenerT.ExitQuotedValue(s)
	}
}

func (p *UnitParser) QuotedValue() (localctx IQuotedValueContext) {
	this := p
	_ = this

	localctx = NewQuotedValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, UnitParserRULE_quotedValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(UnitParserVALUE)
	}

	return localctx
}
