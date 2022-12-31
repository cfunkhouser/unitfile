lexer grammar UnitLexer;

fragment LETTER
  : [\p{L}] // Any Unicode Letter
  ;

fragment DIGIT
  : [\p{Nd}] // Any Unicode Digit
  ;

LBRACKET : '[' -> pushMode(HEADER_MODE);
RBRACKET : ']';
ASSIGN   : '=' -> pushMode(ASSIGN_MODE);
DQUOTE   : '"';
SQUOTE   : '\'';
EOL      : '\n';
ESCAPE   : '\\';

WHITESPACE
  : [ \t]+
  -> skip
  ;

IDENTIFIER
  : LETTER (LETTER | DIGIT)*
  ;

COMMENT
  : ('#' | ';')
    ~[\r\n]* EOL
  -> channel(HIDDEN)
  ;

mode HEADER_MODE;

ID_IN_HEADER
  : LETTER
    (LETTER | DIGIT | ' ')*
    (LETTER | DIGIT) // Cannot end with a space.
  -> type(IDENTIFIER)
  ;

RBRACKET_IN_HEADER
  : ']'
  -> type(RBRACKET), popMode
  ;

mode ASSIGN_MODE;

COMMENT_IN_ASSIGN
  : ('#' | ';')
    ~[\n]* EOL
  -> type(COMMENT), channel(HIDDEN)
  ;

CONTINUATION
  : '\\\n'
  ;

DQUOTE_IN_ASSIGN
  : '"'
  -> type(DQUOTE), pushMode(DQUOTE_STRING_MODE)
  ;

EOL_IN_ASSIGN
  : EOL    
  -> type(EOL), popMode
  ;

EOF_IN_ASSIGN
  : EOF
  -> type(EOF), popMode
  ;

SQUOTE_IN_ASSIGN
  : '\''
  -> type(SQUOTE), pushMode(SQUOTE_STRING_MODE)
  ;

VALUE_IN_ASSIGN
  : ~['" \t\n]+
  -> type(VALUE), pushMode(PLAIN_MODE)
  ;

WHITESPACE_IN_ASSIGN
  : [ \t]+
  -> skip
  ;

mode PLAIN_MODE;

COMMENT_IN_PLAIN
  : ('#' | ';')
    ~[\n]* EOL
  -> type(COMMENT), channel(HIDDEN)
  ;

CONTINUATION_IN_PLAIN
  : '\\\n'
  -> type(VALUE)
  ;

EOL_IN_PLAIN
  : EOL    
  -> type(EOL), mode(DEFAULT_MODE)
  ;

EOF_IN_PLAIN
  : EOF
  -> type(EOF), mode(DEFAULT_MODE)
  ;

ESCAPE_IN_PLAIN
  : '\\'
  -> type(ESCAPE)
  ;

VALUE
  : ~[#;\\\n]+
  ;

mode DQUOTE_STRING_MODE;

END_DQUOTE
  : '"'
  -> type(DQUOTE), popMode
  ;

ESCAPE_IN_DQUOTE
  : ESCAPE
  -> type(ESCAPE)
  ;

VALUE_IN_DQUOTE
  : ~["\\\n]+
  -> type(VALUE)
  ;

mode SQUOTE_STRING_MODE;

END_SQUOTE
  : '\''
  -> type(SQUOTE), popMode
  ;

VALUE_IN_SQUOTE
  : ~['\n]+
  -> type(VALUE)
  ;

ESCAPE_IN_SQUOTE
  : ESCAPE
  -> type(ESCAPE)
  ;
