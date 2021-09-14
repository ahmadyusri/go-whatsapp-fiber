package utils

import (
	"fmt"
	"os"

	"github.com/Rhymen/go-whatsapp"
)

type WhatsappHandler struct{}

func (WhatsappHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "%v", err)
}

func (WhatsappHandler) HandleTextMessage(message whatsapp.TextMessage) {
	fmt.Println("------------------------------")
	fmt.Println("RemoteJid: ", message.Info.RemoteJid)
	fmt.Println("Text: ", message.Text)
	fmt.Println("------------------------------")
}

func (WhatsappHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	//fmt.Println(message)
}

func (WhatsappHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {
	//fmt.Println(message)
}

func (WhatsappHandler) HandleVideoMessage(message whatsapp.VideoMessage) {
	//fmt.Println(message)
}

func (WhatsappHandler) HandleAudioMessage(message whatsapp.AudioMessage) {
	//fmt.Println(message)
}

func (WhatsappHandler) HandleJsonMessage(message string) {
	//fmt.Println(message)
}

func (WhatsappHandler) HandleContactMessage(message whatsapp.ContactMessage) {
	//fmt.Println(message)
}

func (WhatsappHandler) HandleBatteryMessage(message whatsapp.BatteryMessage) {
	//fmt.Println(message)
}

func (WhatsappHandler) HandleNewContact(contact whatsapp.Contact) {
	//fmt.Println(contact)
}
