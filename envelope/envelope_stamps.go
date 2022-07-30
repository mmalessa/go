package envelope

type EnvelopeStamps map[string]interface{}

func StampWithDelay(delaySec int) func(*EnvelopeStamps) {
	return func(stamps *EnvelopeStamps) {
		(*stamps)["delaySec"] = delaySec
	}
}

func StampWithMessageTemplate(templateName string) func(*EnvelopeStamps) {
	return func(stamps *EnvelopeStamps) {
		(*stamps)["messageTemplate"] = templateName
	}
}
