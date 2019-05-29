package utils

// GetMessageByKey : A method to get all locales messages for response messages by key
func GetMessageByKey(key string) MessageItem {
	messages := map[string]MessageItem{
		NotFoundErrorMessageKey: MessageItem{Fa: "یافت نشد!", En: "Ooops! Not found..."},
		InputErrorMessageKey:       MessageItem{Fa: "لطفا ورودی ها را چک کنید", En: "Check inputs!"},
	}
	suitableMessage := messages[key]
	if (suitableMessage == MessageItem{}) {
		return MessageItem{Fa: key, En: key}
	}
	return suitableMessage
}
