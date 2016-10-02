package telegram

type User struct {
	// Unique identifier for this user or bot
	Id int64 `json:"id,omitempty"`
	// User‘s or bot’s first name
	FirstName string `json:"first_name,omitempty"`
	// Optional. User‘s or bot’s last name
	LastName string `json:"last_name,omitempty"`
	// Optional. User‘s or bot’s username
	Username string `json:"username,omitempty"`
}

type Chat struct {
	// Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	Id int64 `json:"id,omitempty"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type,omitempty"`
	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`
	// Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username,omitempty"`
	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`
	// Optional. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`
}

type Message struct {
	// Unique message identifier
	MessageId int64 `json:"message_id,omitempty"`
	// Optional. Sender, can be empty for messages sent to channels
	From *User `json:"from,omitempty"`
	// Date the message was sent in Unix time
	Date int64 `json:"date,omitempty"`
	// Conversation the message belongs to
	Chat *Chat `json:"chat,omitempty"`
	// Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`
	// Optional. For messages forwarded from a channel, information about the original channel
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
	// Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardDate int64 `json:"forward_date,omitempty"`
	// Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// Optional. Date the message was last edited in Unix time
	EditDate int64 `json:"edit_date,omitempty"`
	// Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Text string `json:"text,omitempty"`
	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []*MessageEntity `json:"entities,omitempty"`
	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`
	// Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`
	// Optional. Message is a photo, available sizes of the photo
	Photo []*PhotoSize `json:"photo,omitempty"`
	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`
	// Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`
	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`
	// Optional. Caption for the document, photo or video, 0-200 characters
	Caption string `json:"caption,omitempty"`
	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`
	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`
	// Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue,omitempty"`
	// Optional. A new member was added to the group, information about them (this member may be the bot itself)
	NewChatMember *User `json:"new_chat_member,omitempty"`
	// Optional. A member was removed from the group, information about them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`
	// Optional. A chat photo was change to this value
	NewChatPhoto []*PhotoSize `json:"new_chat_photo,omitempty"`
	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`
	// Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
	// Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
	// Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	// Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatId int64 `json:"migrate_from_chat_id,omitempty"`
	// Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
}

type MessageEntity struct {
	// Type of the entity. Can be mention (@username), hashtag, bot_command, url, email, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
	Type string `json:"type,omitempty"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int64 `json:"offset,omitempty"`
	// Length of the entity in UTF-16 code units
	Length int64 `json:"length,omitempty"`
	// Optional. For “text_link” only, url that will be opened after user taps on the text
	Url string `json:"url,omitempty"`
	// Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`
}

type PhotoSize struct {
	// Unique identifier for this file
	FileId string `json:"file_id,omitempty"`
	// Photo width
	Width int64 `json:"width,omitempty"`
	// Photo height
	Height int64 `json:"height,omitempty"`
	// Optional. File size
	FileSize int64 `json:"file_size,omitempty"`
}

type Audio struct {
	// Unique identifier for this file
	FileId string `json:"file_id,omitempty"`
	// Duration of the audio in seconds as defined by sender
	Duration int64 `json:"duration,omitempty"`
	// Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer,omitempty"`
	// Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int64 `json:"file_size,omitempty"`
}

type Document struct {
	// Unique file identifier
	FileId string `json:"file_id,omitempty"`
	// Optional. Document thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int64 `json:"file_size,omitempty"`
}

type Sticker struct {
	// Unique identifier for this file
	FileId string `json:"file_id,omitempty"`
	// Sticker width
	Width int64 `json:"width,omitempty"`
	// Sticker height
	Height int64 `json:"height,omitempty"`
	// Optional. Sticker thumbnail in .webp or .jpg format
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`
	// Optional. File size
	FileSize int64 `json:"file_size,omitempty"`
}

type Video struct {
	// Unique identifier for this file
	FileId string `json:"file_id,omitempty"`
	// Video width as defined by sender
	Width int64 `json:"width,omitempty"`
	// Video height as defined by sender
	Height int64 `json:"height,omitempty"`
	// Duration of the video in seconds as defined by sender
	Duration int64 `json:"duration,omitempty"`
	// Optional. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Mime type of a file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int64 `json:"file_size,omitempty"`
}

type Voice struct {
	// Unique identifier for this file
	FileId string `json:"file_id,omitempty"`
	// Duration of the audio in seconds as defined by sender
	Duration int64 `json:"duration,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size
	FileSize int64 `json:"file_size,omitempty"`
}

type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number,omitempty"`
	// Contact's first name
	FirstName string `json:"first_name,omitempty"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Contact's user identifier in Telegram
	UserId int64 `json:"user_id,omitempty"`
}

type Location struct {
	// Longitude as defined by sender
	Longitude float64 `json:"longitude,omitempty"`
	// Latitude as defined by sender
	Latitude float64 `json:"latitude,omitempty"`
}

type Venue struct {
	// Venue location
	Location *Location `json:"location,omitempty"`
	// Name of the venue
	Title string `json:"title,omitempty"`
	// Address of the venue
	Address string `json:"address,omitempty"`
	// Optional. Foursquare identifier of the venue
	FoursquareId string `json:"foursquare_id,omitempty"`
}

type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int64 `json:"total_count,omitempty"`
	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]*PhotoSize `json:"photos,omitempty"`
}

type File struct {
	// Unique identifier for this file
	FileId string `json:"file_id,omitempty"`
	// Optional. File size, if known
	FileSize int64 `json:"file_size,omitempty"`
	// Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
	FilePath string `json:"file_path,omitempty"`
}

type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]*KeyboardButton `json:"keyboard,omitempty"`
	// Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	// Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	// Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used, it will be sent to the bot as a message when the button is pressed
	Text string `json:"text,omitempty"`
	// Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestContact bool `json:"request_contact,omitempty"`
	// Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
	RequestLocation bool `json:"request_location,omitempty"`
}

type ReplyKeyboardHide struct {
	// Requests clients to hide the custom keyboard
	HideKeyboard bool `json:"hide_keyboard,omitempty"`
	// Optional. Use this parameter if you want to hide keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text,omitempty"`
	// Optional. HTTP url to be opened when button is pressed
	Url string `json:"url,omitempty"`
	// Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`
	// Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`
}

type CallbackQuery struct {
	// Unique identifier for this query
	Id string `json:"id,omitempty"`
	// Sender
	From *User `json:"from,omitempty"`
	// Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	Message *Message `json:"message,omitempty"`
	// Optional. Identifier of the message sent via the bot in inline mode, that originated the query
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field
	Data string `json:"data,omitempty"`
}

type ForceReply struct {
	// Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	ForceReply bool `json:"force_reply,omitempty"`
	// Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

type ChatMember struct {
	// Information about the user
	User *User `json:"user,omitempty"`
	// The member's status in the chat. Can be “creator”, “administrator”, “member”, “left” or “kicked”
	Status string `json:"status,omitempty"`
}

type Update struct {
	// The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order.
	UpdateId int64 `json:"update_id,omitempty"`
	// Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`
	// Optional. New version of a message that is known to the bot and was edited
	EditedMessage *Message `json:"edited_message,omitempty"`
	// Optional. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	// Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	// Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}

type InlineQuery struct {
	// Unique identifier for this query
	Id string `json:"id,omitempty"`
	// Sender
	From *User `json:"from,omitempty"`
	// Optional. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
	// Text of the query (up to 512 characters)
	Query string `json:"query,omitempty"`
	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset,omitempty"`
}

type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ResultId string `json:"result_id,omitempty"`
	// The user that chose the result
	From *User `json:"from,omitempty"`
	// Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`
	// Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	InlineMessageId string `json:"inline_message_id,omitempty"`
	// The query that was used to obtain the result
	Query string `json:"query,omitempty"`
}
