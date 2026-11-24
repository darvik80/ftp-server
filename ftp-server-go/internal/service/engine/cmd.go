package engine

import "context"

type FtpCode int

const (
	FTPCodeMarkerResumption          FtpCode = 110 // Restart marker reply.
	FTPCodeServiceWillBeReadyIn      FtpCode = 120 // Service ready in nnn minutes.
	FTPCodeDataConnectionAlreadyOpen FtpCode = 125 // Data connection already open; transfer starting.
	FTPCodeFileStatusOk              FtpCode = 150 // File status okay; about to open data connection.

	FTPCodeCommandCompleted             FtpCode = 200 // Command okay.
	FTPCodeCommandNotImplementedForNode FtpCode = 202 // Command not implemented, superfluous at this site.
	FTPCodeSystemStatus                 FtpCode = 211 // System status, or system help reply.
	FTPCodeDirectoryStatus              FtpCode = 212 // Directory status.
	FTPCodeFileStatus                   FtpCode = 213 // File status.
	FTPCodeHelpMessage                  FtpCode = 214 // Help message. How to use the server, or the meaning of a particular non-standard command. This reply is useful only to the human user.
	FTPCodeSystemType                   FtpCode = 215 // NAME system type. Where NAME is an official system name from the list in the Assigned Numbers document.
	FTPCodeServiceReady                 FtpCode = 220 // Service ready for new user.
	FTPCodeServiceClosing               FtpCode = 221 // Service closing control connection. Logged out if appropriate.
	FTPCodeDataConnectionOpen           FtpCode = 225 // Data connection open; no transfer in progress.
	FTPCodeDataConnectionClosed         FtpCode = 226 // Closing data connection. Requested file action successful.
	FTPCodePassiveMode                  FtpCode = 227 // Entering Passive Mode (h1,h2,h3,h4,p1,p2).
	FTPCodeUserLoggedIn                 FtpCode = 230 // User logged in, proceed.
	FTPCodeFileActionCompleted          FtpCode = 250 // Requested file action okay, completed.
	FTPCodeDirectoryCreated             FtpCode = 257 // "PATHNAME" created.

	FTPCodeNameUserAccepted  FtpCode = 331 // User name okay, need password.
	FTPCodeAccountRequired   FtpCode = 332 // Need account for login.
	FTPCodeFileActionPending FtpCode = 350 // Requested file action pending further information.

	FTPCodeServiceUnavailable           FtpCode = 421 // Service not available, closing control connection.
	FTPCodeDataConnectionUnavailable    FtpCode = 425 // Can't open data connection.
	FTPCodeConnectionClosed             FtpCode = 426 // Connection closed; transfer aborted.
	FTPCodeFileActionNotCompleted       FtpCode = 450 // Requested file action not taken. File unavailable.
	FTPCodeFileUnavailableLocal         FtpCode = 451 // Requested action aborted: local error in processing.
	FTPCodeActionNotCompleted           FtpCode = 452 // Requested action not taken. Insufficient storage space in system.
	FTPCodeTemporaryInsufficientStorage FtpCode = 455 // Requested action aborted. Insufficient storage space in system.

	FTPCodeSyntaxErrorCommandNotRecognized   FtpCode = 500 // Syntax error, command unrecognized.
	FTPCodeSyntaxErrorInParameters           FtpCode = 501 // Syntax error in parameters or arguments.
	FTPCodeCommandNotImplemented             FtpCode = 502 // Command not implemented.
	FTPCodeIncorrectSequenceOfCommands       FtpCode = 503 // Bad sequence of commands.
	FTPCodeCommandNotImplementedForParameter FtpCode = 504 // Command not implemented for that parameter.
	FTPCodeLoginIncorrect                    FtpCode = 530 // Not logged in.
	FTPCodeFileAccountRequired               FtpCode = 532 // Need account for storing files.
	FTPCodeFileUnavailable                   FtpCode = 550 // Requested action not taken. File unavailable.
	FTPCodeFileUnavailableMessage            FtpCode = 551 // Requested action aborted: page type unknown.
	FTPCodeFileActionAborted                 FtpCode = 552 // Requested file action aborted. Exceeded storage allocation.
	FTPCodeInsufficientStorage               FtpCode = 553 // Requested action not taken. File name not allowed.
)

// MsgByCode returns the standard RFC 959 reply text for a control-connection reply code.
func MsgByCode(code FtpCode) string {
	switch code {
	case FTPCodeMarkerResumption:
		return "Restart marker reply."
	case FTPCodeServiceWillBeReadyIn:
		return "Service ready in nnn minutes."
	case FTPCodeDataConnectionAlreadyOpen:
		return "Data connection already open; transfer starting."
	case FTPCodeFileStatusOk:
		return "File status okay; about to open data connection."
	case FTPCodeCommandCompleted:
		return "Command okay."
	case FTPCodeCommandNotImplementedForNode:
		return "Command not implemented, superfluous at this site."
	case FTPCodeSystemStatus:
		return "System status, or system help reply."
	case FTPCodeDirectoryStatus:
		return "Directory status."
	case FTPCodeFileStatus:
		return "File status."
	case FTPCodeHelpMessage:
		return "Help message."
	case FTPCodeSystemType:
		return "NAME system type."
	case FTPCodeServiceReady:
		return "Service ready for new user."
	case FTPCodeServiceClosing:
		return "Service closing control connection."
	case FTPCodeDataConnectionOpen:
		return "Data connection open; no transfer in progress."
	case FTPCodeDataConnectionClosed:
		return "Closing data connection."
	case FTPCodePassiveMode:
		return "Entering Passive Mode (h1,h2,h3,h4,p1,p2)."
	case FTPCodeUserLoggedIn:
		return "User logged in, proceed."
	case FTPCodeFileActionCompleted:
		return "Requested file action okay, completed."
	case FTPCodeDirectoryCreated:
		return "\"PATHNAME\" created."

	case FTPCodeNameUserAccepted:
		return "User name okay, need password."
	case FTPCodeAccountRequired:
		return "Need account for login."
	case FTPCodeFileActionPending:
		return "Requested file action pending further information."

	case FTPCodeServiceUnavailable:
		return "Service not available, closing control connection."
	case FTPCodeDataConnectionUnavailable:
		return "Can't open data connection."
	case FTPCodeConnectionClosed:
		return "Connection closed; transfer aborted."
	case FTPCodeFileActionNotCompleted:
		return "Requested file action not taken. File unavailable."
	case FTPCodeFileUnavailableLocal:
		return "Requested action aborted: local error in processing."
	case FTPCodeActionNotCompleted:
		return "Requested action not taken. Insufficient storage space in system."
	case FTPCodeTemporaryInsufficientStorage:
		return "Requested action aborted. Insufficient storage space in system."

	case FTPCodeSyntaxErrorCommandNotRecognized:
		return "Syntax error, command unrecognized."
	case FTPCodeSyntaxErrorInParameters:
		return "Syntax error in parameters or arguments."
	case FTPCodeCommandNotImplemented:
		return "Command not implemented."
	case FTPCodeIncorrectSequenceOfCommands:
		return "Bad sequence of commands."
	case FTPCodeCommandNotImplementedForParameter:
		return "Command not implemented for that parameter."
	case FTPCodeLoginIncorrect:
		return "Not logged in."
	case FTPCodeFileAccountRequired:
		return "Need account for storing files."
	case FTPCodeFileUnavailable:
		return "Requested action not taken. File unavailable."
	case FTPCodeFileUnavailableMessage:
		return "Requested action aborted: page type unknown."
	case FTPCodeFileActionAborted:
		return "Requested file action aborted. Exceeded storage allocation."
	case FTPCodeInsufficientStorage:
		return "Requested action not taken. File name not allowed."
	default:
		return "Unknown status code."
	}
}

type Reply struct {
	Code     FtpCode
	Messages []string
	Body     []string
}

func ReplyByCode(code FtpCode) Reply {
	return Reply{
		Code:     code,
		Messages: []string{MsgByCode(code)},
	}
}

type FtpHandler func(ctx context.Context, s *Session, args []string) Reply
