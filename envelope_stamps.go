package hermessenger

type EnvelopeStamps struct {
	template string
}

func getDefaultEnvelopeStamps() *EnvelopeStamps {
	return &EnvelopeStamps{}
}

func EnvelopeStampWithTemplate(templateName string) func(*EnvelopeStamps) {
	return func(stamps *EnvelopeStamps) {
		stamps.template = templateName
	}
}
