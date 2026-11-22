package engine

import "context"

type FtpCode int

const (
	FTPCodeMarkerResumption          = 110 //Ответ с маркером возобновления.
	FTPCodeServiceWillBeReadyIn      = 120 //Служба будет готова через nnn минут.
	FTPCodeDataConnectionAlreadyOpen = 125 //Соединение данных уже открыто; передача начинается.
	FTPCodeFileStatusOk              = 150 //Состояние файла в порядке; сейчас будет открыто соединение данных.

	FTPCodeCommandCompleted             = 200 //Команда выполнена успешно.
	FTPCodeCommandNotImplementedForNode = 202 //Команда не реализована, на этом узле она излишня.
	FTPCodeSystemStatus                 = 211 //Состояние системы или ответ со справкой по системе.
	FTPCodeDirectoryStatus              = 212 //Состояние каталога.
	FTPCodeFileStatus                   = 213 //Состояние файла.
	FTPCodeHelpMessage                  = 214 //Справочное сообщение. О том, как пользоваться сервером, или о значении определённой нестандартной команды. Этот ответ полезен только человеку-пользователю.
	FTPCodeSystemType                   = 215 //NAME — тип системы. Здесь NAME — официальное имя системы из списка в документе «Назначенные номера» (Assigned Numbers).
	FTPCodeServiceReady                 = 220 //Служба готова для нового пользователя.
	FTPCodeServiceClosing               = 221 //Служба закрывает управляющее соединение. При необходимости выполнен выход из системы.
	FTPCodeDataConnectionOpen           = 225 //Соединение данных открыто; передача не выполняется.
	FTPCodeDataConnectionClosed         = 226 //Соединение данных закрывается.
	FTPCodePassiveMode                  = 227 //Переход в пассивный режим (h1,h2,h3,h4,p1,p2).
	FTPCodeUserLoggedIn                 = 230 //Пользователь вошёл в систему, продолжайте.
	FTPCodeFileActionCompleted          = 250 //Запрошенное действие с файлом успешно выполнено и завершено.
	FTPCodeDirectoryCreated             = 257 //\"PATHNAME\" создан.

	FTPCodeNameUserAccepted  = 331 //Имя пользователя принято, требуется пароль.
	FTPCodeAccountRequired   = 332 //Для входа в систему требуется счёт.
	FTPCodeFileActionPending = 350 //Запрошенное действие с файлом ожидает дополнительной информации.

	FTPCodeServiceUnavailable           = 421 //Служба недоступна, управляющее соединение закрывается.
	FTPCodeDataConnectionUnavailable    = 425 //Невозможно открыть соединение данных.
	FTPCodeConnectionClosed             = 426 //Соединение закрыто; передача прервана.
	FTPCodeFileActionNotCompleted       = 450 //Запрошенное действие с файлом не выполнено.
	FTPCodeFileUnavailableLocal         = 451 //Запрошенное действие прервано: локальная ошибка при обработке.
	FTPCodeActionNotCompleted           = 452 //Запрошенное действие не выполнено.
	FTPCodeTemporaryInsufficientStorage = 455 //В системе недостаточно места для хранения.

	FTPCodeSyntaxErrorCommandNotRecognized   = 500 //Синтаксическая ошибка, команда не распознана.
	FTPCodeSyntaxErrorInParameters           = 501 //Синтаксическая ошибка в параметрах или аргументах.
	FTPCodeCommandNotImplemented             = 502 //Команда не реализована.
	FTPCodeIncorrectSequenceOfCommands       = 503 //Неправильная последовательность команд.
	FTPCodeCommandNotImplementedForParameter = 504 //Команда не реализована для этого параметра.
	FTPCodeLoginIncorrect                    = 530 //Вход в систему не выполнен.
	FTPCodeFileAccountRequired               = 532 //Для сохранения файлов требуется счёт.
	FTPCodeFileUnavailable                   = 550 //Запрошенное действие не выполнено.
	FTPCodeFileUnavailableMessage            = 551 //Запрошенное действие прервано: тип страницы неизвестен.
	FTPCodeFileActionAborted                 = 552 //Запрошенное действие с файлом прервано.
	FTPCodeInsufficientStorage               = 553 //В системе недостаточно места для хранения.
)

type Reply struct {
	Code     FtpCode
	Messages []string
	Body     []string
}

type FtpHandler func(ctx context.Context, args []string) Reply
