package server

import "github.com/hillu/go-yara/v4"

func initYARACompiler() (*yara.Compiler, error) {
	yaraCompiler, err := yara.NewCompiler()
	if err != nil {
		return nil, err
	}
	err = yaraCompiler.AddString(`rule DetectMalicious {
		strings:
			$malicious_string = "malicious_pattern"
		condition:
			$malicious_string
	}`, "rules")

	if err != nil {
		//c.String(http.StatusInternalServerError, "Error retrieving compiled rules")
		return nil, err
	}
	return yaraCompiler, nil
}
