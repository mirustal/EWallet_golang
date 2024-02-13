Для развертывания докера: 
docker compose up --build

Приложение EWallet реализующее систему обработки транзакций платёжной системы

Создание кошелька: POST /api/v1/wallet
Перевод средств с одного кошелька на другой: POST /api/v1/wallet/{walletId}/send
Получение историй входящих и исходящих транзакций: GET /api/v1/wallet/{walletId}/history
Получение текущего состояния кошелька: GET /api/v1/wallet/{walletId}

Аутенфикатор доступа Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkFkbWluIiwiaWF0IjoiQWRtaW4ifQ.8NqkbR4i2NTzeMA9J8Qnn23yx3nzlO4E8YxJF1XspOU, создание токена закомичена в main
