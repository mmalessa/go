package envelope

type EnvelopeStamps map[string]interface{}

func StampWithDelay(delaySec int) func(*EnvelopeStamps) {
	return func(stamps *EnvelopeStamps) {
		(*stamps)["delay"] = delaySec
	}
}

func StampWithMessageType(templateName string) func(*EnvelopeStamps) {
	return func(stamps *EnvelopeStamps) {
		(*stamps)["messageType"] = templateName
	}
}
