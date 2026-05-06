# pratice-ascii
 None selected

Skip to content
Using Gmail with screen readers
in:draft
1 of 11
Fwd:
Inbox

Okoh Christopher
Mon, Apr 27, 3:08 PM (9 days ago)
to me



---------- Forwarded message ---------
From: Okoh Christopher <etzkristokency2@gmail.com>
Date: Mon, Apr 27, 2026 at 1:56 PM
Subject:
To: joyceejeks4@gmail.com <joyceejeks4@gmail.com>



ASCII
Write a function that loads a banner file and returns its characters as a map

The banner files (standard, shadow, thinkertoy) store ASCII characters separated by newlines. Each character occupies exactly 8 lines. Write a function that reads the file content and returns a map where each key is a rune and each value is a slice of 8 strings one per line of that character's art.

func LoadBanner(filename string) (map[rune][]string, error)

what it must handle:

LoadBanner("standard.txt")→map with 95 printable ASCII chars, each having exactly 8 lines

LoadBanner("notfound.txt")→returns a non-nil error

result[' ']→slice of 8 empty/space strings (space is a valid char)

Characters in the file start from ASCII 32 (space) and go to 126 (~). The file has a blank line separating each character your parser must account for that extra line.


hint:

Split file content by "\n", then chunk into groups of 9 (8 lines + 1 separator). The character index maps to ASCII code: index 0 = space (32), index 1 = '!' (33), etc.






ascii-art/

├── main.go

├── banner.go          // contains LoadBanner

├── standard.txt

├──  banner_test.go  // must have an empty

├── shadow.txt

└── thinkertoy.txt





igbokennedy@gmail.com

---------- Forwarded message ---------
From: Okoh Christopher <etzkristokency2@gmail.com>
Date: Mon, Apr 27, 2026, 3:08 PM
Subject: Fwd:
To: adewahjames@gmail.com <adewahjames@gmail.com>




---------- Forwarded message ---------
From: Okoh Christopher <etzkristokency2@gmail.com>
Date: Mon, Apr 27, 2026 at 1:56 PM
Subject:
To: joyceejeks4@gmail.com <joyceejeks4@gmail.com>



ASCII
Write a function that loads a banner file and returns its characters as a map

The banner files (standard, shadow, thinkertoy) store ASCII characters separated by newlines. Each character occupies exactly 8 lines. Write a function that reads the file content and returns a map where each key is a rune and each value is a slice of 8 strings one per line of that character's art.

func LoadBanner(filename string) (map[rune][]string, error)

what it must handle:

LoadBanner("standard.txt")→map with 95 printable ASCII chars, each having exactly 8 lines

LoadBanner("notfound.txt")→returns a non-nil error

result[' ']→slice of 8 empty/space strings (space is a valid char)

Characters in the file start from ASCII 32 (space) and go to 126 (~). The file has a blank line separating each character your parser must account for that extra line.


hint:

Split file content by "\n", then chunk into groups of 9 (8 lines + 1 separator). The character index maps to ASCII code: index 0 = space (32), index 1 = '!' (33), etc.






ascii-art/

├── main.go

├── banner.go          // contains LoadBanner

├── standard.txt

├──  banner_test.go  // must have an empty

├── shadow.txt

└── thinkertoy.txt





Adewah James <adewahjames@gmail.com>
Mon, Apr 27, 3:31 PM (9 days ago)
to igbokennedy

