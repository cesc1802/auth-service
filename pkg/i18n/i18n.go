package i18n

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const (
	i18nMessage       = "message"
	messageFolderName = "messages"
	defaultLanguage   = "en"
)

type Config struct {
	Languages []string
}

func NewI18nConfig(languages []string) *Config {
	return &Config{
		Languages: languages,
	}
}

type AppI18n struct {
	Bundle    *i18n.Bundle
	Localizer map[string]*i18n.Localizer
}

func NewI18n(c *Config) (*AppI18n, error) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	mapLocalizer := make(map[string]*i18n.Localizer)
	for _, lang := range c.Languages {
		bundle.MustLoadMessageFile(fmt.Sprintf("./%s/%v.%v.json", messageFolderName, i18nMessage, lang))
		mapLocalizer[lang] = i18n.NewLocalizer(bundle, lang)
	}

	return &AppI18n{
		Bundle:    bundle,
		Localizer: mapLocalizer,
	}, nil
}

func (r *AppI18n) MustLocalize(lang string, msgId string, templateData map[string]string) string {
	var localizePtr *i18n.Localizer
	if _, ok := r.Localizer[lang]; !ok {
		localizePtr = r.Localizer[defaultLanguage]
	}
	return localizePtr.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    msgId,
		TemplateData: templateData,
	})
}
