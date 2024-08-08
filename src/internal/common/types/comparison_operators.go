package common_types

const (
	ComparisonOperatorsEqual              = "eq"
	ComparisonOperatorsNotEqual           = "ne"
	ComparisonOperatorsGreater            = "gt"
	ComparisonOperatorsLess               = "lt"
	ComparisonOperatorsGreaterThanOrEqual = "ge"
	ComparisonOperatorsLessThanOrEqual    = "le"
)

// ComparisonOperators - операторы сравнения.
type ComparisonOperators string

// ParseComparisonOperators - парсинг операторов сравнения.
func ParseComparisonOperators(operator string) (result ComparisonOperators) {
	switch operator {
	case "eq":
		{
			result = ComparisonOperatorsEqual
		}
	case "ne":
		{
			result = ComparisonOperatorsNotEqual
		}
	case "gt":
		{
			result = ComparisonOperatorsGreater
		}
	case "lt":
		{
			result = ComparisonOperatorsLess
		}
	case "ge":
		{
			result = ComparisonOperatorsGreaterThanOrEqual
		}
	case "le":
		{
			result = ComparisonOperatorsLessThanOrEqual
		}
	}

	return
}

// TranslateToSign - получение знака оператора.
func (operator ComparisonOperators) TranslateToSign() (result string) {
	switch operator {
	case ComparisonOperatorsEqual:
		{
			result = "="
		}
	case ComparisonOperatorsNotEqual:
		{
			result = "!="
		}
	case ComparisonOperatorsGreater:
		{
			result = ">"
		}
	case ComparisonOperatorsLess:
		{
			result = "<"
		}
	case ComparisonOperatorsGreaterThanOrEqual:
		{
			result = ">="
		}
	case ComparisonOperatorsLessThanOrEqual:
		{
			result = "<="
		}
	}

	return
}
