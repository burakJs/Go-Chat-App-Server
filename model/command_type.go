package model

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type ClientStatusType interface {
	getMessage(client Client) string
	checkStatus(statusCode string) bool
}

type ClientExitType struct {
}

type ClientEnterType struct {
}

type ClientSendDataType struct {
}

func (t ClientExitType) getMessage(client Client) string {
	coloredLeftText := color.New(color.FgRed).Sprint("left")
	boldNameText := color.New(color.FgWhite, color.Bold).Sprint(client.Name)
	return fmt.Sprintf("%s %s the room", boldNameText, coloredLeftText)
}

func (t ClientExitType) checkStatus(statusCode string) bool {
	return strings.Compare(statusCode, "/exit") == 0
}

func (t ClientEnterType) getMessage(client Client) string {
	coloredJoinedText := color.New(color.FgGreen).Sprint("joined")
	boldNameText := color.New(color.FgWhite, color.Bold).Sprint(client.Name)
	return fmt.Sprintf("%s %s the room", boldNameText, coloredJoinedText)
}

func (t ClientEnterType) checkStatus(statusCode string) bool {
	return strings.Compare(statusCode, "/enter") == 0
}

func (c ClientSendDataType) getMessage(client Client) string {
	boldNameText := color.New(color.FgWhite, color.Bold).Sprint(client.Name)
	return fmt.Sprintf("%s: %s", boldNameText, client.Message)
}

func (c ClientSendDataType) checkStatus(statusCode string) bool {
	return string(statusCode[0]) != "/"
}
