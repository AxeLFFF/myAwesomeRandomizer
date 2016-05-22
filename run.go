// run
package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	theBot, err := tgbotapi.NewBotAPI("216709974:AAFRd_MSJF08IgOOZzdLNilyCRrviopIFqw")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Logged as ", theBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := theBot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	for update := range updates {
		resp_msg, ok := processBotCommand(update)
		if ok == true {
			log.Println("Response sent: ", resp_msg.Text)
			theBot.Send(resp_msg)
		}
	}
}

func processBotCommand(received_msg tgbotapi.Update) (tgbotapi.MessageConfig, bool) {
	varList := map[string]func() string{
		"/воровач": vorovach,
		"/дворяч":  dvoryach,
		"/шар":     ball,
	}
	var resp_msg tgbotapi.MessageConfig
	for k := range varList {
		if strings.Index(received_msg.Message.Text, k) == 0 {
			log.Println("Cmd received: ", k)
			response := varList[k]()
			resp_msg := tgbotapi.NewMessage(received_msg.Message.Chat.ID, response)
			resp_msg.ReplyToMessageID = received_msg.Message.MessageID
			return resp_msg, true
		}
	}
	return resp_msg, false
}

func getRandomItem(list []string) string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	n := r.Intn(len(list))
	if n > 0 {
		return list[n-1]
	} else {
		return list[0]
	}
}

func vorovach() string {
	vars := []string{"Петух",
		"Чухан",
		"Малолетка",
		"Мужик",
		"Стремящийся",
		"Шестерка",
		"Блатной",
		"Вор",
		"Вор в законе",
		"Черт",
		"Вафлер",
		"Барсук",
		"Главпетух",
		"Король чуханов"}
	return "Решаем за твой положняк...\nПоздравляем, ты " + getRandomItem(vars)
}

func dvoryach() string {
	vars := []string{"Лакеi",
		"Холопъ",
		"Крепостноi",
		"Ля мюжик",
		"Статсъ-секретарь",
		"Сударь",
		"Дворянiнъ",
		"Государь Императоръ",
		"Разночiнецъ",
		"Полiцмейстеръ",
		"Ямщiкъ",
		"Тiтулярный советникъ",
		"Генерал-губернаторъ",
		"Сенатскiй регистраторъ",
		"Юнкер",
		"Поручiк",
		"Есаулъ"}
	return "Листаем Бархатную книгу...\nИтакъ, Вы " + getRandomItem(vars)
}

func ball() string {
	vars := []string{"Да",
		"Не сейчас",
		"Ни в коем случае",
		"Уверенное да",
		"Кто знает",
		"Быть может",
		"Забудь об этом",
		"Я не уверен",
		"Слишком рано",
		"Это возможно",
		"Вперед!",
		"Это неплохо",
		"Ты шутишь?",
		"Конечно, да",
		"Не надейся на это",
		"Да, но позднее",
		"Я думаю, хорошо",
		"Даже не думай",
		"Не делай этого",
		"Думаю, не стоит"}
	return "Магический шар ищет ответ в глубинах мироздания...\nОтвет: " + getRandomItem(vars)
}
