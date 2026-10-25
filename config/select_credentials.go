package config

import (
	"fmt"
	"strings"

	exerrors "github.com/exsited/exsited-sdk-go/v4/errors"
)

// One spelling only, so a document cannot name the instance twice with no rule for which wins.
const (
	clientNameKeyCanon  = "client_name"
	unnamedIndexPattern = "<index %d>"
)

// ReadClientName returns one credentials entry's client name, whatever shape the entry has.
func ReadClientName(block any) string {
	switch typed := block.(type) {
	case SdkConfig:
		return typed.Auth.ClientName
	case *SdkConfig:
		if typed == nil {
			return ""
		}
		return typed.Auth.ClientName
	case AuthConfig:
		return typed.ClientName
	case *AuthConfig:
		if typed == nil {
			return ""
		}
		return typed.ClientName
	case map[string]any:
		for _, key := range []string{clientNameKeyCanon} {
			if value, ok := typed[key]; ok && value != nil {
				if name := fmt.Sprint(value); name != "" {
					return name
				}
			}
		}
		return ""
	case map[string]string:
		for _, key := range []string{clientNameKeyCanon} {
			if name, ok := typed[key]; ok && name != "" {
				return name
			}
		}
		return ""
	default:
		return ""
	}
}

// LabelCredentials names every entry for an error message, position-labelling the ones that name no instance.
func LabelCredentials(credentials []any) []string {
	labels := make([]string, 0, len(credentials))
	for index, block := range credentials {
		label := ReadClientName(block)
		if label == "" {
			label = fmt.Sprintf(unnamedIndexPattern, index)
		}
		labels = append(labels, label)
	}
	return labels
}

// SelectCredentials returns one entry: nil is the first, an int a position (negative from the end), a string a name.
func SelectCredentials(credentials []any, instance any) (any, error) {
	if len(credentials) == 0 {
		return nil, exerrors.NewUsageError("no credentials were supplied")
	}
	if instance == nil {
		return credentials[0], nil
	}
	switch typed := instance.(type) {
	case int:
		size := len(credentials)
		if typed >= -size && typed < size {
			return credentials[((typed%size)+size)%size], nil
		}
		return nil, exerrors.NewUsageError(fmt.Sprintf(
			"no credentials at index %d; %d supplied (0-%d)", typed, size, size-1))
	case string:
		return selectByName(credentials, typed)
	case bool:
		// Not an index: a YAML `instance: yes` must fail by name rather than select position 1.
		if typed {
			return selectByName(credentials, "True")
		}
		return selectByName(credentials, "False")
	default:
		return nil, exerrors.NewUsageError(
			"instance must be nil, an index or a clientName")
	}
}

// selectByName refuses a name two entries carry, which would send one tenant's credentials to another.
func selectByName(credentials []any, name string) (any, error) {
	var matched []int
	for index, block := range credentials {
		if ReadClientName(block) == name {
			matched = append(matched, index)
		}
	}
	if len(matched) > 1 {
		positions := make([]string, 0, len(matched))
		for _, index := range matched {
			positions = append(positions, fmt.Sprintf("%d", index))
		}
		return nil, exerrors.NewUsageError(fmt.Sprintf(
			"instance '%s' is defined more than once, at index %s; clientName has to identify "+
				"one entry", name, strings.Join(positions, " and ")))
	}
	if len(matched) == 1 {
		return credentials[matched[0]], nil
	}
	return nil, exerrors.NewUsageError(fmt.Sprintf(
		"unknown instance '%s'; available: %s", name, strings.Join(LabelCredentials(credentials), ", ")))
}
