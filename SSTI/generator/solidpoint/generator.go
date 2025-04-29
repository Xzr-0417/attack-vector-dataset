package engineconfig

import (
	"fmt"
	"math/rand"
	"strconv"
)

func generateRenderPayload(rawPayload string) (string, string, error) {
	a := rand.Intn(1000) + 5
	b := rand.Intn(1000) + 5

	data := map[string]string{
		firstNumPlaceholder:  strconv.Itoa(a),
		secondNumPlaceholder: strconv.Itoa(b),
	}

	payload, err := executeTemplate(rawPayload, data)
	if err != nil {
		return "", "", err
	}

	result := mathExprResult(a, b, payload)

	return payload, result, nil
}

func generateExecPayload(rawPayload string) (string, string, error) {
	val := strconv.Itoa(rand.Intn(9) + 1)
	pldValue := fmt.Sprintf("\\\\x3%s", val)
	result := val

	pldLength := 10
	for i := 0; i < pldLength; i++ {
		val := strconv.Itoa(rand.Intn(10))
		pldValue += fmt.Sprintf("\\\\x3%s", val)
		result = result + val
	}

	data := map[string]string{
		formattedNumPlaceholder: pldValue,
	}

	payload, err := executeTemplate(rawPayload, data)
	if err != nil {
		return "", "", err
	}

	return payload, result, nil
}

func generateBlindSleepPayload(rawPayload string, waitTime int) (string, error) {
	data := map[string]string{
		waitTimePlaceholder: strconv.Itoa(waitTime),
	}

	payload, err := executeTemplate(rawPayload, data)
	if err != nil {
		return "", err
	}

	return payload, nil
}

func GenerateCollaboratorPayload(rawPayload string, collaboratorDomain string) (string, error) {
	data := map[string]string{
		pingDomainPlaceholder: collaboratorDomain,
	}

	payload, err := executeTemplate(rawPayload, data)
	if err != nil {
		return "", err
	}

	return payload, nil
}
