package types

type LogType struct {
	Time        string `json:"time"`
	Date        string `json:"date"`
	ActionType  string `json:"action_type"`
	Description string `json:"description"`
}
