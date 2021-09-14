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
	// fmt.Println("------------------------------")
	// fmt.Println("Type: TextMessage")
	// fmt.Println("RemoteJid: ", message.Info.RemoteJid)
	// fmt.Println("Text: ", message.Text)
	// fmt.Println("------------------------------")
}

func (WhatsappHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	// fmt.Println("------------------------------")
	// fmt.Println("Type: ImageMessage")
	// fmt.Println("RemoteJid: ", message.Info.RemoteJid)
	// fmt.Println("Caption: ", message.Caption)
	// fmt.Println("Type: ", message.Type)
	// fmt.Println("------------------------------")
}

func (WhatsappHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {
	// fmt.Println("------------------------------")
	// fmt.Println("Type: DocumentMessage")
	// fmt.Println("RemoteJid: ", message.Info.RemoteJid)
	// fmt.Println("Type: ", message.Type)
	// fmt.Println("------------------------------")
}

func (WhatsappHandler) HandleVideoMessage(message whatsapp.VideoMessage) {
	// fmt.Println("------------------------------")
	// fmt.Println("Type: VideoMessage")
	// fmt.Println("RemoteJid: ", message.Info.RemoteJid)
	// fmt.Println("Caption: ", message.Caption)
	// fmt.Println("Type: ", message.Type)
	// fmt.Println("------------------------------")
}

func (WhatsappHandler) HandleAudioMessage(message whatsapp.AudioMessage) {
	// fmt.Println("------------------------------")
	// fmt.Println("Type: AudioMessage")
	// fmt.Println("RemoteJid: ", message.Info.RemoteJid)
	// fmt.Println("Type: ", message.Type)
	// fmt.Println("------------------------------")
}

func (WhatsappHandler) HandleJsonMessage(message string) {
	//fmt.Println(message)
}

func (WhatsappHandler) HandleContactMessage(message whatsapp.ContactMessage) {
	// fmt.Println(message)
}

func (WhatsappHandler) HandleBatteryMessage(message whatsapp.BatteryMessage) {
	// fmt.Println("------------------------------")
	// fmt.Println("Plugged: ", message.Plugged)
	// fmt.Println("Powersave: ", message.Powersave)
	// fmt.Println("Percentage: ", message.Percentage)
	// fmt.Println("------------------------------")
}

func (WhatsappHandler) HandleNewContact(contact whatsapp.Contact) {
	// fmt.Println(contact)
}
