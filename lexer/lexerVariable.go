package lexer

import (
	"errors"
	"fmt"
	"strconv"
)

var Variables = map[string]any{}

func RuneToStr(ch rune) string {
	return fmt.Sprintf(`%q`, string(ch))
}
func ParseVariable(curToken Token, lexer *Lexer) (key string, value interface{}, err error) {
	tok := curToken //
	keyT := lexer.NextToken()
	fmt.Println(tok)
	if tok.Type == 0 {
		return "", nil, nil
	}
	if tok.Type != VAR {
		return "", nil, errors.New("var error unkown")
	}
	if keyT.Type != IDENTIFIER {
		return "", nil, errors.New(fmt.Sprintf("'%s' must be an identifier", keyT.Value))
	}
	Eq := lexer.NextToken()
	fmt.Println(Eq)
	if Eq.Type != ASSIGN {
		return "", nil, errors.New(fmt.Sprintf("'=' sign is expected after the '%s'", keyT.Value))
	}
	valT := lexer.NextToken()
	parsedVal, err := parseVariableValue(valT)
	if err != nil {
		return "", nil, err
	}
	Variables[keyT.Value] = parsedVal
	return keyT.Value, parsedVal, nil

}

func parseVariableValue(token Token) (interface{}, error) {
	switch token.Type {
	case 0:
		return nil, nil
	case NUMBER:
		return strconv.Atoi(token.Value)
	case STRING:
		return token.Value[1 : len(token.Value)-1], nil
	case BOOL:
		if token.Value == "true" {
			return true, nil
		}
		return false, nil

	case IDENTIFIER:
		return parseIdentifier(token)
	default:
		/*
		   TODO: {
		   1. check if identifier exists and if does return identifier value
		   2. check if function then take function output if function returns nothing throw error

		   }




		*/
		return nil, nil

	}

}
func parseIdentifier(token Token) (interface{}, error) {
	if val, ok := Variables[token.Value]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("Undefined identifier: %s", token.Value)
}
