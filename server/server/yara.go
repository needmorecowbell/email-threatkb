package server

import "github.com/hillu/go-yara/v4"

// InitYARA initializes the YARA rules for the EMLServer.
// It compiles a YARA rule that detects malicious patterns in emails.
// The compiled rules are stored in the EMLServer's rules field.
// Returns an error if there was a problem compiling the rules.
func (s *EMLServer) InitYARA() error {
	yaraCompiler, err := yara.NewCompiler()
	if err != nil {
		return err
	}
	err = yaraCompiler.AddString(`rule DetectMalicious {
		strings:
			$malicious_string = "malicious_pattern"
		condition:
			$malicious_string
	}`, "rules")

	if err != nil {
		return err
	}
	yaraRules, err := yaraCompiler.GetRules()
	if err != nil {
		return err
	}
	s.rules = yaraRules
	return nil
}
