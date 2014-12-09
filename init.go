package core

func init() {
	ruleHandler["rename_col"] = &renameCol{}
	ruleHandler["find_replace"] = &findReplace{}
}
