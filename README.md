Приложение EWallet реализующее систему обработки транзакций платёжной системы
# Для развертывания докера: 
	docker compose up --build



# Описание структуры программы 

	./app Здесь хранится все связанное с бизнесс логикой:
	
	./models модели структур данных
	./queries интерфейсы для связи с платформой\
	./routes все что связано с маршрутами 
	
	
	./pkg Здесь хранится функционал проекта. конфигурации, 
	./config обработка конфигурации 

	./platform Здесь хранится логика уровня платформы, настройка базы данных, 
 	./database  инициализация и функции с обработкой запросов в базу данных
  
# Описание запросов:
	Создание кошелька: 
		POST /api/v1/wallet
	Перевод средств с одного кошелька на другой:
		POST /api/v1/wallet/{walletId}/send
	Получение историй входящих и исходящих транзакций: 
		GET /api/v1/wallet/{walletId}/history
	Получение текущего состояния кошелька: 
		GET /api/v1/wallet/{walletId}



# Аутенфикатор доступа Bearer
	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkFkbWluIiwiaWF0IjoiQWRtaW4ifQ.8NqkbR4i2NTzeMA9J8Qnn23yx3nzlO4E8YxJF1XspOU
	
